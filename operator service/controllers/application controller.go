package controllers

import (
	"context"

	appv1 "github.com/georgigeorgiev/cloudops/operator-service/api/v1"

	ctrl "sigs.k8s.io/controller-runtime"
)

type ApplicationReconciler struct {
}

func (r *ApplicationReconciler) Reconcile(
	ctx context.Context,
	req ctrl.Request,
) (ctrl.Result, error) {

	/*
		Reconciliation Loop

		1. Read Application CR
		2. Ensure Deployment exists
		3. Ensure Service exists
		4. Ensure Monitoring exists
		5. Update Status
	*/

	return ctrl.Result{}, nil
}

func (r *ApplicationReconciler) SetupWithManager(
	mgr ctrl.Manager,
) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.Application{}).
		Complete(r)
}
