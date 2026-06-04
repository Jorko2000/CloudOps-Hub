package kubernetes

func CreateDeployment(
	client *kubernetes.Clientset,
	appName string,
	image string,
	replicas int32,
) error {

	// create k8s deployment

	return nil
}
