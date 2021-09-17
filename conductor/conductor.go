package conductor

import (
	"fmt"
	chords2 "github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nate-droid/go-orchestra/core/notes"
	"github.com/nate-droid/go-orchestra/core/scales"
	"github.com/nats-io/nats.go"
)

// Responsibilities
// Create song structure
// Choose Key
// Choose how many in each group

type Conductor struct {
	SendSymphony  chan *Symphony
	SymphonyReady chan bool
}

type Symphony struct {
	Key              string
	ChordProgression []string // for now just I IV and so on
	Sections         []Section
}

type Song struct {
	ChordProgression []chords2.Chord
}

type SongStructure struct {
	Key              notes.Note
	ChordProgression []chords2.ChordInterval
	Mode             scales.ModeName
}

type Section struct {
	Type      string
	GroupSize int
}

type Chord []string

// strings, woodwind, brass, percussion

// xylophones, marimbas, bells for percussion

func (cond *Conductor) sendSymphony(symphony Symphony) error {
	cond.SendSymphony <- &symphony
	return nil
}

func (cond *Conductor) Run() error {
	// TODO this will be the loop that runs ever on
	return nil
}

func newSymphony() *Symphony {
	// TODO randomize

	return &Symphony{
		Key:              "C",
		ChordProgression: []string{"I", "IV", "V"},
		Sections:         []Section{*newSection()},
	}
}

func newSection() *Section {
	return &Section{
		GroupSize: 1,
		Type:      "strings",
	}
}

func newConductor() (*Conductor, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	sendCh := make(chan *Symphony)
	err = ec.BindSendChan("createSymphony", sendCh)
	if err != nil {
		return nil, err
	}

	recvCh := make(chan bool)
	_, err = ec.BindRecvChan("symphonyReady", recvCh)
	if err != nil {
		return nil, err
	}

	cond := &Conductor{
		SendSymphony: sendCh,
		SymphonyReady: recvCh,
	}
	return cond, nil
}

func main() {
	fmt.Println("something!")
}