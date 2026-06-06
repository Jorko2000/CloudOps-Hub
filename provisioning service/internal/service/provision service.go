package service

import (
	tf "provisioning-service/internal/terraform"
)

type ProvisionService struct{}

func (s *ProvisionService) Provision(
	namespace string,
	postgres string,
	redis string,
) error {

	err := tf.WriteTFVars(
		tf.Variables{
			Namespace: namespace,
			Postgres:  postgres,
			Redis:     redis,
		},
	)

	if err != nil {
		return err
	}

	return tf.Apply()
}
