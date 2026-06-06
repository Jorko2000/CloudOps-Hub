package main

import (
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/georgigeorgiev/cloudops/observability-service/internal/metrics"
)

func main() {

	metrics.Register()

	router := gin.Default()

	router.GET(
		"/metrics",
		gin.WrapH(promhttp.Handler()),
	)

	router.Run(":8083")
}
