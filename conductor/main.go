package main

import (
	"context"
	"fmt"
	"github.com/nate-droid/core/chords"
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
	"github.com/nate-droid/core/symphony"
	uuid "github.com/nu7hatch/gouuid"
	"math/rand"
	"os"

	"time"

	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
)

type Conductor struct {
	SendSymphony  chan *symphony.Symphony
	SymphonyReady chan bool

	SendSongStructure chan *symphony.SongStructure
}

type SectionType string

var SendSongToSectionSubject = "sendToSection"

const (
	StringsSection    SectionType = "strings"
	WoodwindSection   SectionType = "woodwind"
	BrassSection      SectionType = "brass"
	PercussionSection SectionType = "percussion" // xylophones, marimbas, bells for percussion
)

func (cond *Conductor) sendSymphony(symphony *symphony.Symphony) error {
	cond.SendSymphony <- symphony
	return nil
}

func (cond *Conductor) sendSongStructure(songStructure *symphony.SongStructure) error {
	cond.SendSongStructure <- songStructure
	return nil
}

func (cond *Conductor) Run(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		for {
			fmt.Println("creating symphony")
			s := newSymphony()
			fmt.Println("sending: ", s)
			for _, section := range s.Sections {
				for i := 0; i < section.GroupSize; i++ {
					fmt.Println("sending song to: ", section)
					err := cond.sendSongStructure(&s.SongStructure)
					if err != nil {
						return err
					}
				}
			}
			time.Sleep(time.Second * 10)
		}
	})

	return errs.Wait()
}

// newSymphony will randomly generate a symphony
func newSymphony() *symphony.Symphony {
	rand.Seed(time.Now().Unix())

	prog := chords.CommonProgressions[rand.Intn(len(chords.CommonProgressions))]
	key := notes.GetAllNotes()[rand.Intn(len(notes.GetAllNotes()))]

	u, _ := uuid.NewV4()

	return &symphony.Symphony{
		SongStructure: symphony.SongStructure{
			Key:              key,
			ChordProgression: prog,
			Mode:             scales.Major,
			SymphonyID:       u.String(),
			Section:          symphony.Section{
				Type: "strings",
				GroupSize: 1,
			},
		},
		Sections: []symphony.Section{*newSection()},
		ID: u.String(),
	}
}

func newSection() *symphony.Section {
	return &symphony.Section{
		GroupSize: 1,
		Type:      "strings",
	}
}

func newConductor() (*Conductor, error) {
	ec, err := newEncodedNatsCon()
	if err != nil {

	}

	sendCh := make(chan *symphony.Symphony)
	err = ec.BindSendChan("createSymphony", sendCh)
	if err != nil {
		return nil, err
	}

	sendSongStructureCh := make(chan *symphony.SongStructure)
	err = ec.BindSendChan(SendSongToSectionSubject, sendSongStructureCh)
	if err != nil {
		return nil, err
	}

	recvCh := make(chan bool)
	_, err = ec.BindRecvChan("symphonyReady", recvCh)
	if err != nil {
		return nil, err
	}

	cond := &Conductor{
		SendSymphony:      sendCh,
		SymphonyReady:     recvCh,
		SendSongStructure: sendSongStructureCh,
	}

	return cond, nil
}

// newEncodedNatsCon will return a new connection to a nats instance
func newEncodedNatsCon() (*nats.EncodedConn, error){
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
	return ec, nil
}

func main() {
	time.Sleep(3 * time.Second)
	fmt.Println("conductor started")
	conductor, err := newConductor()
	if err != nil {
		panic(err)
	}
	fmt.Println("created conductor")
	err = conductor.Run(context.Background())
	if err != nil {
		panic(err)
	}

}

