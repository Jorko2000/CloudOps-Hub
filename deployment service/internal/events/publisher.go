package events

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type DeploymentEvent struct {

	AppName string `json:"app_name"`

	Image string `json:"image"`

	Replicas int32 `json:"replicas"`

	Status string `json:"status"`
}

type Publisher struct {
	Conn *nats.Conn
}

func NewPublisher() (
	*Publisher,
	error,
) {

	conn, err := nats.Connect(
		nats.DefaultURL,
	)

	if err != nil {
		return nil, err
	}

	return &Publisher{
		Conn: conn,
	}, nil
}

func (p *Publisher) Publish(
	event DeploymentEvent,
) error {

	payload, err := json.Marshal(
		event,
	)

	if err != nil {
		return err
	}

	return p.Conn.Publish(
		"deployment.created",
		payload,
	)
}
