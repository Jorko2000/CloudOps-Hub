package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"
)

func BuildDeployment(
	name string,
	image string,
	replicas int32,
) *appsv1.Deployment {

	return &appsv1.Deployment{}
}
