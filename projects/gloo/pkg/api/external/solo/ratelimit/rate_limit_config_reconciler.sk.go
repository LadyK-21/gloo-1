// Code generated by solo-kit. DO NOT EDIT.

//Source: pkg/code-generator/codegen/templates/resource_reconciler_template.go
package v1alpha1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionRateLimitConfigFunc func(original, desired *RateLimitConfig) (bool, error)

type RateLimitConfigReconciler interface {
	Reconcile(namespace string, desiredResources RateLimitConfigList, transition TransitionRateLimitConfigFunc, opts clients.ListOpts) error
}

func rateLimitConfigsToResources(list RateLimitConfigList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, rateLimitConfig := range list {
		resourceList = append(resourceList, rateLimitConfig)
	}
	return resourceList
}

func NewRateLimitConfigReconciler(client RateLimitConfigClient, statusSetter resources.StatusSetter) RateLimitConfigReconciler {
	return &rateLimitConfigReconciler{
		base: reconcile.NewReconciler(client.BaseClient(), statusSetter),
	}
}

type rateLimitConfigReconciler struct {
	base reconcile.Reconciler
}

func (r *rateLimitConfigReconciler) Reconcile(namespace string, desiredResources RateLimitConfigList, transition TransitionRateLimitConfigFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "rateLimitConfig_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*RateLimitConfig), desired.(*RateLimitConfig))
		}
	}
	return r.base.Reconcile(namespace, rateLimitConfigsToResources(desiredResources), transitionResources, opts)
}
