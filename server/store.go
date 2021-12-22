package server

import (
	"context"
	"fmt"
	"github.com/nate-droid/go-orchestra/core/symphony"
	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
	"os"
	"sync"
)

var ReceiveSongSubject = "sendSong"

type Mem struct {
	ReceiveSong chan *symphony.Song
	Songs       map[string]symphony.Song
	Symphonies  map[string]SymphonyEntry

	mu sync.Mutex
}

type SymphonyEntry struct {
	Symphony      symphony.Symphony
	MusicianCount int
	TimesPlayed   int
}

func newMem() (*Mem, error) {
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

	return &Mem{
		Songs:       map[string]symphony.Song{},
		Symphonies:  map[string]SymphonyEntry{},
		ReceiveSong: recvCh,
	}, nil
}

func (m *Mem) run(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		for {
			select {
			case song := <-m.ReceiveSong:
				err := m.StoreSong(song)
				if err != nil {
					return err
				}
			}
		}
	})

	return errs.Wait()
}

func (m *Mem) StoreSong(song *symphony.Song) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.Songs) >= 10 {
		fmt.Println("storage capacity for songs has been reached. Skipped")
		return nil
	}

	m.Songs[song.SymphonyID] = *song

	fmt.Printf("storing song with SymphonyID: %s\n", song.SymphonyID)

	return nil
}

func (m *Mem) StoreSymphony(s symphony.Symphony) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.Symphonies) >= 10 {
		fmt.Println("storage capacity for songs has been reached. Skipped")
		return nil
	}

	m.Symphonies[s.ID] = SymphonyEntry{
		Symphony:      s,
		MusicianCount: s.GroupSize,
		TimesPlayed:   0,
	}

	fmt.Printf("storing song with SymphonyID: %s\n", s.ID)

	return nil
}

func (m *Mem) FetchSymphonyByID(id string) (*symphony.Symphony, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	s, ok := m.Symphonies[id]
	if !ok {
		return nil, fmt.Errorf("could not find Symphony with ID: %s", id)
	}
	if s.TimesPlayed >= s.MusicianCount {
		return nil, fmt.Errorf("symphony with ID: %s has alread been played the maximum amount")
	}

	return &s.Symphony, nil
}

func (m *Mem) UpdateSymphony(id string) error {
	return nil
}

func (m *Mem) FetchAllSongs() (map[string]symphony.Song, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	songs := m.Songs

	return songs, nil
}
