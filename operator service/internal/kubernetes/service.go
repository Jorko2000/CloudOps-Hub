package kubernetes

import corev1 "k8s.io/api/core/v1"

func BuildService(
	name string,
) *corev1.Service {

	return &corev1.Service{}
}
