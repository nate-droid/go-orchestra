package messages

import (
	"fmt"
	"github.com/nats-io/jsm.go"
	"github.com/nats-io/nats.go"
	"os"
	"time"
)

// NewEncodedNatsCon will return a new connection to a nats instance
func NewEncodedNatsCon() (*nats.EncodedConn, error) {
	// TODO make nats stuff generic
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

func NewStream(subject string) (*jsm.Stream, error) {
	natsURI := os.Getenv("NATS_URI")
	if natsURI == "" {
		natsURI = nats.DefaultURL
	}
	nc, err := nats.Connect(natsURI)
	if err != nil {
		return nil, err
	}
	_, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}

	mgr, _ := jsm.New(nc, jsm.WithTimeout(10*time.Second))
	// stream, _ := mgr.NewStream("ORDERS", jsm.Subjects("ORDERS.*"), jsm.MaxAge(24*365*time.Hour), jsm.FileStorage())
	pattern := subject + ".*"
	stream, err := mgr.NewStream(subject, jsm.Subjects(pattern), jsm.MaxAge(24*365*time.Hour), jsm.FileStorage())
	return stream, err
}

func NewConsumer(nc *nats.Conn, subject string) (*jsm.Consumer, error) {
	// sub, err := nc.Subscribe(subject, func(msg *nats.Msg) error {
	//	meta, _ := msg.Metadata()
	//	fmt.Println("Received a message: ", string(msg.Data))
	//	return msg.Ack()
	// })
	mgr, _ := jsm.New(nc, jsm.WithTimeout(10*time.Second))
	ib := nats.NewInbox()

	nc.Publish(subject, []byte("hello"))
	_, err := nc.Subscribe(ib, func(m *nats.Msg) {
		meta, _ := m.Metadata()
		fmt.Println("Received a message: ", string(m.Data), meta)
		m.Ack()
	})
	if err != nil {
		return nil, err
	}
	consumer, _ := mgr.NewConsumer("ORDERS", nil, jsm.FilterStreamBySubject("ORDERS.received"), jsm.SamplePercent(100), jsm.DeliverySubject(ib))

	return consumer, nil
}
