package controllers

import (
	"context"

	kubeModel "github.com/ydataai/backend-interview-sample/pkg/models/kube/v1alpha1"
	"github.com/ydataai/backend-interview-sample/pkg/reconcilers"

	ctrl "sigs.k8s.io/controller-runtime"
)

// ReconcilerController defines the a reconciler
type ReconcilerController struct {
	ctx           context.Context
	mngr          ctrl.Manager
	appReconciler reconcilers.ApplicationReconciler
}

// NewReconcilerController initializes reconciler controller
func NewReconcilerController(
	ctx context.Context,
	mngr ctrl.Manager,
	appReconciler reconcilers.ApplicationReconciler,
) ReconcilerController {

	return ReconcilerController{
		mngr:          mngr,
		ctx:           ctx,
		appReconciler: appReconciler,
	}
}

// Boot ...
func (r ReconcilerController) Boot() error {
	managedController := ctrl.NewControllerManagedBy(r.mngr)

	if err := managedController.For(&kubeModel.Application{}).Complete(r.appReconciler); err != nil {

		return err
	}

	return nil
}
