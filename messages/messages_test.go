package messages

import (
	"fmt"
	"github.com/nate-droid/go-orchestra/messages/fake"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type TestMessage struct {
	Message string
}

func TestThing(t *testing.T) {
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()

	natsURI := os.Getenv("NATS_URI")
	if natsURI == "" {
		natsURI = nats.DefaultURL
	}
	nc, err := nats.Connect(natsURI)
	assert.NoError(t, err)
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	assert.NoError(t, err)

	err = ec.Publish("ORDERS", TestMessage{Message: "h"})

	_, err = ec.Subscribe("ORDERS", func(x TestMessage) {

		fmt.Println("Received a message: ", x)
	})
	assert.NoError(t, err)

	assert.NoError(t, err)

}
