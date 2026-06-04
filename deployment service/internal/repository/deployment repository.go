package repository

import "deployment-service/internal/model"

type DeploymentRepository interface {

	Save(
		deployment *model.Deployment,
	) error
}
