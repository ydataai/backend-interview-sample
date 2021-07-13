package reconcilers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

// ApplicationReconciler reconciles a Application object
type ApplicationReconciler struct {
	ctx    context.Context
	scheme *runtime.Scheme
}

// NewApplicationReconciler initializes application reconciler structure
func NewApplicationReconciler(
	ctx context.Context,
	scheme *runtime.Scheme,
) ApplicationReconciler {
	return ApplicationReconciler{
		ctx:    ctx,
		scheme: scheme,
	}
}

// Reconcile is a required function responsible for listening and making changes in a specific resource
func (r ApplicationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	// TODO: Complete

	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=ydata.ai,resources=applications,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ydata.ai,resources=applications/status,verbs=get;update;patch
