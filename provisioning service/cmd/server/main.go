package main

import (
	"github.com/gin-gonic/gin"

	"provisioning-service/internal/handler"
	"provisioning-service/internal/service"
)

func main() {

	router := gin.Default()

	provisionService := &service.ProvisionService{}

	provisionHandler := &handler.ProvisionHandler{
		Service: provisionService,
	}

	router.POST(
		"/infrastructure/provision",
		provisionHandler.Provision,
	)

	router.Run(":8082")
}
