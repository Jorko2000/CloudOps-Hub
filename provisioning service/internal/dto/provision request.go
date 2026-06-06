package dto

type ProvisionRequest struct {
	Namespace string `json:"namespace"`
	Postgres  string `json:"postgres"`
	Redis     string `json:"redis"`
}
