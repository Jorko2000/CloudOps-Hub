package service

import (
	"deployment-service/internal/events"
	"deployment-service/internal/model"
	"deployment-service/internal/repository"
)

type DeploymentService struct {
	Repo      repository.DeploymentRepository
	Publisher *events.Publisher
}

func (s *DeploymentService) Deploy(
	appName string,
	image string,
	replicas int32,
) error {

	// Create deployment record
	record := &model.Deployment{
		AppName:  appName,
		Image:    image,
		Replicas: replicas,
		Status:   "DEPLOYED",
	}

	// Save deployment in PostgreSQL
	if err := s.Repo.Save(record); err != nil {
		return err
	}

	// Publish deployment event to NATS
	err := s.Publisher.Publish(
		events.DeploymentEvent{
			AppName:  appName,
			Image:    image,
			Replicas: replicas,
			Status:   "DEPLOYED",
		},
	)

	if err != nil {
		return err
	}

	return nil
}
