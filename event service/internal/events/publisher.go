package events

import (
	"encoding/json"

	"github.com/nats-io/nats.go"

	"event-service/internal/models"
)

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

func (p *Publisher) PublishDeployment(
	event models.DeploymentEvent,
) error {

	payload, err := json.Marshal(event)

	if err != nil {
		return err
	}

	return p.Conn.Publish(
		"deployment.created",
		payload,
	)
}
