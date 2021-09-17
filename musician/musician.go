package musician

import (
	"fmt"
	"github.com/nate-droid/go-orchestra/conductor"
	"github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nats-io/nats.go"
)

type Musician struct {
	ReceiveSongStructure chan *conductor.SongStructure
	SendSong             chan *conductor.Song
}

var SendSongSubject = "sendSong"
var ReceiveSongSubject = "receiveSong"

func (m *Musician) playSong(song conductor.SongStructure) (*conductor.Song, error) {
	progression, err := chords.Progression(song.Mode, song.ChordProgression, song.Key)
	return  &conductor.Song{
		ChordProgression: progression,
	}, err
}

func newMusician() (*Musician, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}

	sendCh := make(chan *conductor.Song)
	err = ec.BindSendChan(SendSongSubject, sendCh)
	if err != nil {
		return nil, err
	}

	recvCh := make(chan *conductor.SongStructure)
	_, err = ec.BindRecvChan(ReceiveSongSubject, recvCh)
	if err != nil {
		return nil, err
	}

	m := &Musician{
		ReceiveSongStructure: recvCh,
		SendSong:             sendCh,
	}
	return m, nil
}

func almostMain() {
	m, err := newMusician()
	if err != nil {
		panic(err)
	}
	songStructure := <- m.ReceiveSongStructure
	song, err := m.playSong(*songStructure)
	fmt.Println(song)
}