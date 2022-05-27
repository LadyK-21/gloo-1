// Code generated by solo-kit. DO NOT EDIT.

//Source: pkg/code-generator/codegen/templates/resource_client_template.go
package v1

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type RouteOptionWatcher interface {
	// watch namespace-scoped RouteOptions
	Watch(namespace string, opts clients.WatchOpts) (<-chan RouteOptionList, <-chan error, error)
}

type RouteOptionClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*RouteOption, error)
	Write(resource *RouteOption, opts clients.WriteOpts) (*RouteOption, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (RouteOptionList, error)
	RouteOptionWatcher
}

type routeOptionClient struct {
	rc clients.ResourceClient
}

func NewRouteOptionClient(ctx context.Context, rcFactory factory.ResourceClientFactory) (RouteOptionClient, error) {
	return NewRouteOptionClientWithToken(ctx, rcFactory, "")
}

func NewRouteOptionClientWithToken(ctx context.Context, rcFactory factory.ResourceClientFactory, token string) (RouteOptionClient, error) {
	rc, err := rcFactory.NewResourceClient(ctx, factory.NewResourceClientParams{
		ResourceType: &RouteOption{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base RouteOption resource client")
	}
	return NewRouteOptionClientWithBase(rc), nil
}

func NewRouteOptionClientWithBase(rc clients.ResourceClient) RouteOptionClient {
	return &routeOptionClient{
		rc: rc,
	}
}

func (client *routeOptionClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *routeOptionClient) Register() error {
	return client.rc.Register()
}

func (client *routeOptionClient) Read(namespace, name string, opts clients.ReadOpts) (*RouteOption, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*RouteOption), nil
}

func (client *routeOptionClient) Write(routeOption *RouteOption, opts clients.WriteOpts) (*RouteOption, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(routeOption, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*RouteOption), nil
}

func (client *routeOptionClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *routeOptionClient) List(namespace string, opts clients.ListOpts) (RouteOptionList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToRouteOption(resourceList), nil
}

func (client *routeOptionClient) Watch(namespace string, opts clients.WatchOpts) (<-chan RouteOptionList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	routeOptionsChan := make(chan RouteOptionList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				select {
				case routeOptionsChan <- convertToRouteOption(resourceList):
				case <-opts.Ctx.Done():
					close(routeOptionsChan)
					return
				}
			case <-opts.Ctx.Done():
				close(routeOptionsChan)
				return
			}
		}
	}()
	return routeOptionsChan, errs, nil
}

func convertToRouteOption(resources resources.ResourceList) RouteOptionList {
	var routeOptionList RouteOptionList
	for _, resource := range resources {
		routeOptionList = append(routeOptionList, resource.(*RouteOption))
	}
	return routeOptionList
}
