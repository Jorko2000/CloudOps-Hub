package service

import (
	tf "provisioning-service/internal/terraform"
)

type ProvisionService struct{}

func (s *ProvisionService) Provision() error {

	return tf.Apply()
}
