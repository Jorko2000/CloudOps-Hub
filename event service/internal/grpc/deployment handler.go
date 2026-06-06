package grpc

import (
	"context"

	pb "event-service/proto"
)

type Server struct {
	pb.UnimplementedDeploymentServiceServer
}

func (s *Server) DeployApplication(
	ctx context.Context,
	req *pb.DeploymentRequest,
)(
	*pb.DeploymentResponse,
	error,
){

	return &pb.DeploymentResponse{

		Status: "SUCCESS",

		Message: "Deployment event received",

	}, nil
}
