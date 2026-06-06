package kubernetes

import netv1 "k8s.io/api/networking/v1"

func BuildIngress(
	name string,
) *netv1.Ingress {

	return &netv1.Ingress{}
}
