package models

import "time"

type DeploymentEvent struct {

	AppName string `json:"app_name"`

	Image string `json:"image"`

	Replicas int32 `json:"replicas"`

	Status string `json:"status"`

	Timestamp time.Time `json:"timestamp"`
}
