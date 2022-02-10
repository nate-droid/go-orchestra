package musician

import (
	"context"
	"fmt"
	"github.com/nate-droid/go-orchestra/conductor"
	"github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nate-droid/go-orchestra/core/symphony"
	"os"

	"golang.org/x/sync/errgroup"
	"time"

	"github.com/nats-io/nats.go"
)

type Musician struct {
	ReceiveSongStructure chan *symphony.SongStructure
	SendSong             chan *symphony.Song

	EncodedConn *nats.EncodedConn
}

var SendSongSubject = "sendSong"
var SendSongToSectionSubject = "sendToSection"

func (m *Musician) playSong(song *symphony.SongStructure) (*symphony.Song, error) {
	// TODO here we could add some extensions based on an adjective like "jazzy" or "normal"
	progression, err := chords.Progression(song.Mode, song.ChordProgression, song.Key)
	return &symphony.Song{
		Progression:      song.ChordProgression,
		ChordProgression: progression,
		Section:          song.Section,
		SymphonyID:       song.SymphonyID,
	}, err
}

func NewMusician() (*Musician, error) {
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

	sendCh := make(chan *symphony.Song)
	err = ec.BindSendChan(SendSongSubject, sendCh)
	if err != nil {
		return nil, err
	}

	recvCh := make(chan *symphony.SongStructure)
	_, err = ec.BindRecvChan(SendSongToSectionSubject, recvCh)
	if err != nil {
		return nil, err
	}

	m := &Musician{
		ReceiveSongStructure: recvCh,
		SendSong:             sendCh,
		EncodedConn:          ec,
	}

	return m, nil
}

func main() {
	time.Sleep(3 * time.Second)
	fmt.Println("musician started")
	m, err := NewMusician()
	if err != nil {
		panic(err)
	}
	err = m.Run(context.Background())
	if err != nil {
		panic(err)
	}
}

func (m *Musician) Run(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		for {
			//select {
			//case songStructure := <-m.ReceiveSongStructure:
			//	song, err := m.playSong(songStructure)
			//	if err != nil {
			//		return err
			//	}
			//	fmt.Println("Played Song: ", song)
			//	m.SendSong <- song
			//
			//}

			// TODO This breaks
			var song symphony.SongStructure

			fmt.Println("Waiting for song!!")
			_, err := m.EncodedConn.Subscribe(conductor.SendSongToSectionSubject, func(s *symphony.SongStructure) {
				fmt.Println("Received Song: ", s)

			})
			if err != nil {
				return err
			}
			fmt.Println("Received Song: ", song)
			if os.Getenv("SINGLE_RUN") == "true" {
				return nil
			}
		}
	})

	return errs.Wait()
}
