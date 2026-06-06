package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var DeploymentsCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "cloudops_deployments_total",
		Help: "Total Deployments",
	},
)

func Register() {

	prometheus.MustRegister(
		DeploymentsCounter,
	)
}
