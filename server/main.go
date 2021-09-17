package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nate-droid/core/symphony"
	"github.com/nats-io/nats.go"
	uuid "github.com/nu7hatch/gouuid"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
)

type Store struct {
	db *bolt.DB
	ReceiveSong chan *symphony.Song
}

var path = "store.db"
var ReceiveSongSubject = "sendSong"

func (store *Store) SaveSong(song *symphony.Song) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(song.SymphonyID))
		if err != nil {
			return err
		}
		encodedSong, err := encodeSong(song)
		if err != nil {
			return err
		}

		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		err = b.Put([]byte(id.String()), encodedSong)
		fmt.Println("saved song!")
		return nil
	})
	if err != nil {
		return err
	}
	err = store.SaveSymphonyID(song.SymphonyID)
	if err != nil {
		return err
	}
	return nil
}

func (store *Store) SaveSymphonyID(id string) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("symphonies"))
		if err != nil {
			return err
		}
		err = b.Put([]byte(id), []byte(""))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *Store) FetchSong(id string) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("songs"))
		v := b.Get([]byte("test"))
		fmt.Println("results: ", v)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *Store) FetchLatestSymphonies() ([][]symphony.Song, error) {
	var all [][]symphony.Song
	ids, err := store.FetchSymphoniesIDs()
	if err != nil {
		return nil, err
	}
	for _, i := range ids {
		allSongs, err := store.FetchSongsByBucket(i)
		if err != nil {
			return nil, err
		}
		all = append(all, allSongs)
	}

	return all,	nil
}

func (store *Store) FetchSymphoniesIDs() ([]string, error) {
	var symphonies []string
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("symphonies"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			symphonies = append(symphonies, string(k))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return symphonies, nil
}

func (store *Store) run(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		for {
			select {
				case song := <- store.ReceiveSong:
					err := store.SaveSong(song)
					if err != nil {
						return err
					}
			}
		}
	})

	return errs.Wait()
}

func (store *Store) FetchSongsByBucket(bucket string) ([]symphony.Song, error) {
	songs := []symphony.Song{}
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			temp, err := decodeSong(v)
			if err != nil {
				return err
			}
			fmt.Println("cool? ", temp)
			songs = append(songs, *temp)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func newStore() (*Store, error) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("nats: ", os.Getenv("NATS_URI"))
	natsURI := os.Getenv("NATS_URI")
	if natsURI == "" {
		natsURI = nats.DefaultURL
	}
	fmt.Println("new Nats: ", natsURI)
	nc, err := nats.Connect(natsURI)
	if err != nil {
		return nil, err
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}


	recvCh := make(chan *symphony.Song)
	_, err = ec.BindRecvChan(ReceiveSongSubject, recvCh)
	if err != nil {
		return nil, err
	}

	return &Store{
		db,
		recvCh,
	}, nil
}

type Results struct {
	ResultsJson []symphony.Song `json:"results"`
}

func main() {
	errs, ctx := errgroup.WithContext(context.Background())
	store, err := newStore()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/symphonies", func(c echo.Context) error {
		res, err := store.FetchLatestSymphonies()
		var x []symphony.Song
		for _, song := range res {
			x = append(x, song[0])
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, "oops")
		}
		return c.JSON(http.StatusOK, Results{x})
	})

	errs.Go(func() error {
		return store.run(ctx)
	})
	errs.Go(func() error {
		return e.Start("0.0.0.0:1323")
	})

	err = errs.Wait()
	if err != nil {
		panic(err)
	}
}

func encodeSong(song *symphony.Song) ([]byte, error){
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(song)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func decodeSong(song []byte) (*symphony.Song, error) {
	r := bytes.NewReader(song)
	dec := gob.NewDecoder(r)
	var n *symphony.Song
	err := dec.Decode(&n)
	if err != nil {
		return nil, err
	}

	return n, nil
}
