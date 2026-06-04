package model

import "time"

type Deployment struct {
	ID        int64
	AppName   string
	Image     string
	Replicas  int32
	Status    string
	CreatedAt time.Time
}
