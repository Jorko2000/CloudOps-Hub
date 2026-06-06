package events

import (
	"event-service/internal/processor"

	"github.com/nats-io/nats.go"
)

func StartSubscriber() error {

	conn, err := nats.Connect(
		nats.DefaultURL,
	)

	if err != nil {
		return err
	}

	_, err = conn.Subscribe(
		"deployment.created",
		func(msg *nats.Msg) {

			processor.ProcessDeployment(
				msg.Data,
			)
		},
	)

	if err != nil {
		return err
	}

	select {}
}
