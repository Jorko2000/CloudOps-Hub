package service

import (
	"deployment-service/internal/model"
	"deployment-service/internal/repository"
)

type DeploymentService struct {
	Repo repository.DeploymentRepository
}

func (s *DeploymentService) Deploy(
	appName string,
	image string,
	replicas int32,
) error {

	record := &model.Deployment{
		AppName: appName,
		Image: image,
		Replicas: replicas,
		Status: "DEPLOYED",
	}

	return s.Repo.Save(
		record,
	)
}
