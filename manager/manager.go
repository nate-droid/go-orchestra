package manager

import (
	"fmt"
	// "github.com/nate-droid/go-orchestra/core"
	"github.com/nats-io/nats.go"
	"time"
)

type Manager struct {
	WaitForSymphony chan *core.Symphony
	SymphonyReady   chan bool
}

func newManager() (*Manager, error) {
	nc, _ := nats.Connect(nats.DefaultURL)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	recvCh := make(chan *core.Symphony)
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
		SymphonyReady:   sendCh,
	}

	return man, nil
}

func (m *Manager) hireOrchestra(symphony *core.Symphony) {
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

func (m *Manager) hireMusician(song *core.SongStructure) bool {
	// create a container here!
	return true
}
