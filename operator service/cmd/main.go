package main

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/georgigeorgiev/cloudops/operator-service/controllers"
)

func main() {

	mgr, err := ctrl.NewManager(
		ctrl.GetConfigOrDie(),
		ctrl.Options{},
	)

	if err != nil {
		panic(err)
	}

	if err := (&controllers.ApplicationReconciler{}).
		SetupWithManager(mgr); err != nil {

		panic(err)
	}

	if err := mgr.Start(
		ctrl.SetupSignalHandler(),
	); err != nil {

		panic(err)
	}
}
