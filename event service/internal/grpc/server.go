package grpc

import (
	"net"

	pb "event-service/proto"

	"google.golang.org/grpc"
)

func Start() error {

	lis, err := net.Listen(
		"tcp",
		":50051",
	)

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	pb.RegisterDeploymentServiceServer(
		server,
		&Server{},
	)

	return server.Serve(lis)
}
