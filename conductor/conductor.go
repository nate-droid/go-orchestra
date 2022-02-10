package conductor

import (
	"context"
	"fmt"
	"github.com/nate-droid/go-orchestra/core/symphony"
	"github.com/nate-droid/go-orchestra/messages"
	"github.com/nats-io/nats.go"

	"time"

	"golang.org/x/sync/errgroup"
)

type Conductor struct {
	SendSymphony  chan *symphony.Symphony
	SymphonyReady chan bool

	SendSongStructure chan *symphony.SongStructure
	EncodedConn       *nats.EncodedConn
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
	err := cond.EncodedConn.Publish(SendSongToSectionSubject, songStructure)
	if err != nil {
		return err
	}
	return nil
}

func (cond *Conductor) Run(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		for {
			fmt.Println("creating symphony")
			s := symphony.NewSymphony()
			fmt.Println("sending: ", s)
			for _, section := range s.Sections {
				for i := 0; i < section.GroupSize; i++ {
					fmt.Println("sending song to: ", section)
					err := cond.sendSongStructure(s.SongStructure)
					if err != nil {
						return err
					}
					err = cond.sendSymphony(s)
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

func NewConductor() (*Conductor, error) {
	ec, err := messages.NewEncodedNatsCon()
	if err != nil {
		return nil, err
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
		EncodedConn:       ec,
	}

	return cond, nil
}
