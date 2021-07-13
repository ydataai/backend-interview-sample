package main

import (
	"github.com/ydataai/backend-interview-sample/pkg/models/kube/v1alpha1"

	v1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientGoScheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	// Scheme ...
	Scheme = runtime.NewScheme()
)

// AddToScheme ...
func AddToScheme(scheme *runtime.Scheme) {
	v1alpha1.AddToScheme(scheme)

	v1.AddToScheme(scheme)
	coreV1.AddToScheme(scheme)
	clientGoScheme.AddToScheme(scheme)
}
