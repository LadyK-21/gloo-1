// Code generated by solo-kit. DO NOT EDIT.

//Source: pkg/code-generator/codegen/templates/resource_template.go
package v1

import (
	"log"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewSecret(namespace, name string) *Secret {
	secret := &Secret{}
	secret.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return secret
}

func (r *Secret) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

func (r *Secret) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *Secret) GroupVersionKind() schema.GroupVersionKind {
	return SecretGVK
}

type SecretList []*Secret

func (list SecretList) Find(namespace, name string) (*Secret, error) {
	for _, secret := range list {
		if secret.GetMetadata().Name == name && secret.GetMetadata().Namespace == namespace {
			return secret, nil
		}
	}
	return nil, errors.Errorf("list did not find secret %v.%v", namespace, name)
}

func (list SecretList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, secret := range list {
		ress = append(ress, secret)
	}
	return ress
}

func (list SecretList) Names() []string {
	var names []string
	for _, secret := range list {
		names = append(names, secret.GetMetadata().Name)
	}
	return names
}

func (list SecretList) NamespacesDotNames() []string {
	var names []string
	for _, secret := range list {
		names = append(names, secret.GetMetadata().Namespace+"."+secret.GetMetadata().Name)
	}
	return names
}

func (list SecretList) Sort() SecretList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list SecretList) Clone() SecretList {
	var secretList SecretList
	for _, secret := range list {
		secretList = append(secretList, resources.Clone(secret).(*Secret))
	}
	return secretList
}

func (list SecretList) Each(f func(element *Secret)) {
	for _, secret := range list {
		f(secret)
	}
}

func (list SecretList) EachResource(f func(element resources.Resource)) {
	for _, secret := range list {
		f(secret)
	}
}

func (list SecretList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Secret) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Secret

func (o *Secret) GetObjectKind() schema.ObjectKind {
	t := SecretCrd.TypeMeta()
	return &t
}

func (o *Secret) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Secret)
}

func (o *Secret) DeepCopyInto(out *Secret) {
	clone := resources.Clone(o).(*Secret)
	*out = *clone
}

var (
	SecretCrd = crd.NewCrd(
		"secrets",
		SecretGVK.Group,
		SecretGVK.Version,
		SecretGVK.Kind,
		"sec",
		false,
		&Secret{})
)

var (
	SecretGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gloo.solo.io",
		Kind:    "Secret",
	}
)
