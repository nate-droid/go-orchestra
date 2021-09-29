package main

import (
	"context"
	"fmt"
	"github.com/nate-droid/core/symphony"
	"github.com/nats-io/nats.go"
	uuid "github.com/nu7hatch/gouuid"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/sync/errgroup"
	"os"
)

var path = "store.db"
var ReceiveSongSubject = "sendSong"

type Store struct {
	db *bolt.DB
	ReceiveSong chan *symphony.Song
}

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

