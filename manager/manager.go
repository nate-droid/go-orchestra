package manager

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nate-droid/go-orchestra/conductor"
	"time"
)

type Manager struct {
	WaitForSymphony chan *conductor.Symphony
	SymphonyReady   chan bool
}

func newManager() (*Manager, error) {
	nc, _ := nats.Connect(nats.DefaultURL)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	recvCh := make(chan *conductor.Symphony)
	_, err := ec.BindRecvChan("createSymphony", recvCh)
	if err != nil {
		return nil, err
	}
	sendCh := make(chan bool)
	err = ec.BindSendChan("symphonyReady", sendCh)
	if err != nil {
		return nil, err
	}

	man := &Manager{
		WaitForSymphony: recvCh,
		SymphonyReady: sendCh,
	}

	return man, nil
}

func (m *Manager) hireOrchestra(symphony *conductor.Symphony) {
	// for each section, we need to hire the musicians
	for _, section := range symphony.Sections {
		for i := 0; i < section.GroupSize; i++ {
			// create Job / Container
			fmt.Println("Hired!")
		}
	}
	// TODO send orchestra ready
	// sleep for 2 seconds then send
	time.Sleep(3 * time.Second)
	m.SymphonyReady <- true
}

func (m *Manager) hireMusician(song *conductor.SongStructure) bool {
	// create a container here!
	return true
}
