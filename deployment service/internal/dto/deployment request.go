package dto

type DeploymentRequest struct {
	AppName string `json:"appName"`
	Image string `json:"image"`
	Replicas int32 `json:"replicas"`
}
