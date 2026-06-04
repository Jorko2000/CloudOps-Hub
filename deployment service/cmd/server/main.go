package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST(
		"/deployments",
		func(c *gin.Context) {},
	)

	router.Run(":8081")
}
