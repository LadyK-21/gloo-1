// Code generated by solo-kit. DO NOT EDIT.

//Source: pkg/code-generator/codegen/templates/resource_reconciler_template.go
package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionSecretFunc func(original, desired *Secret) (bool, error)

type SecretReconciler interface {
	Reconcile(namespace string, desiredResources SecretList, transition TransitionSecretFunc, opts clients.ListOpts) error
}

func secretsToResources(list SecretList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, secret := range list {
		resourceList = append(resourceList, secret)
	}
	return resourceList
}

func NewSecretReconciler(client SecretClient, statusSetter resources.StatusSetter) SecretReconciler {
	return &secretReconciler{
		base: reconcile.NewReconciler(client.BaseClient(), statusSetter),
	}
}

type secretReconciler struct {
	base reconcile.Reconciler
}

func (r *secretReconciler) Reconcile(namespace string, desiredResources SecretList, transition TransitionSecretFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "secret_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*Secret), desired.(*Secret))
		}
	}
	return r.base.Reconcile(namespace, secretsToResources(desiredResources), transitionResources, opts)
}
