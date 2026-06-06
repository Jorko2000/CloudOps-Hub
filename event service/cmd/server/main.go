package main

import (
	"log"

	grpcserver "event-service/internal/grpc"
	"event-service/internal/events"
)

func main() {

	go func() {

		err := events.StartSubscriber()

		if err != nil {

			log.Fatal(err)
		}
	}()

	err := grpcserver.Start()

	if err != nil {

		log.Fatal(err)
	}
}
