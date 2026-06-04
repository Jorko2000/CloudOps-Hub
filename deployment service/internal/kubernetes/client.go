package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient() (*kubernetes.Clientset,error) {

	config, err :=
		clientcmd.BuildConfigFromFlags(
			"",
			"./kubeconfig",
		)

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(
		config,
	)
}
