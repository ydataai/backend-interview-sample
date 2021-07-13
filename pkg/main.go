package main

import (
	"context"
	"fmt"

	"github.com/ydataai/backend-interview-sample/pkg/controllers"
	"github.com/ydataai/backend-interview-sample/pkg/reconcilers"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	errChan = make(chan error)
)

func init() {
	AddToScheme(Scheme)
}

func main() {
	appConfiguration := Configuration{}

	if err := appConfiguration.LoadEnvVars(); err != nil {
		panic(fmt.Errorf("could not set configurationuration variables. Err: %v", err))
	}

	// log.SetLevel(appConfiguration.logLevel)

	mngr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             Scheme,
		MetricsBindAddress: fmt.Sprintf(":%s", appConfiguration.metricsPort),
		LeaderElection:     appConfiguration.enableLeaderElection,
	})
	if err != nil {
		panic(err)
	}

	reconcilerContext := context.Background()
	appReconciler := reconcilers.NewApplicationReconciler(reconcilerContext, mngr.GetScheme())
	reconcilerController := controllers.NewReconcilerController(
		reconcilerContext,
		mngr,
		appReconciler,
	)

	if err = reconcilerController.Boot(); err != nil {
		panic(err)
	}

	go func() {
		if err = mngr.Start(ctrl.SetupSignalHandler()); err != nil {
			errChan <- err
		}
	}()

	for err := range errChan {
		panic(err)
	}
}
