package manager

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/nate-droid/core/symphony"
	"github.com/nats-io/nats.go"
	"time"
)

type Manager struct {
	WaitForSymphony chan *symphony.Symphony
	SymphonyReady   chan bool
}

func newManager() (*Manager, error) {
	nc, _ := nats.Connect(nats.DefaultURL)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	recvCh := make(chan *symphony.Symphony)
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

func (m *Manager) hireOrchestra(symphony *symphony.Symphony) {
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

func (m *Manager) hireMusician(song *symphony.SongStructure) bool {
	// create a container here!
	return true
}

func dockStuff() error {
	ctx := context.Background()
	c, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	l, err := c.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		return err
	}
	fmt.Println(l)
	serviceList, err := c.ServiceList(ctx, types.ServiceListOptions{Status: true})
	if err != nil {
		return err
	}
	fmt.Println("serviceList: ", serviceList)
	//spec := swarm.ServiceSpec{
	//	Annotations: swarm.Annotations{
	//		Name: "superdelete",
	//	},
	//	TaskTemplate: swarm.TaskSpec{
	//		ContainerSpec: &swarm.ContainerSpec{Image: "conductor"},
	//	},
	//}
	//_, err = c.ServiceCreate(ctx, spec, types.ServiceCreateOptions{})
	if err != nil {
		return err
	}
	for _, service := range serviceList {
		if service.Spec.Annotations.Name == "superdelete" {
			err = c.ServiceRemove(ctx, service.ID)
			if err != nil {
				return err
			}
			fmt.Println("deleted service")
		}
	}
	//r, err := c.ServiceCreate()
	//if err != nil {
	//
	//}
	//c.ServiceCreate()
	//c.SwarmInit()
	return nil
}
