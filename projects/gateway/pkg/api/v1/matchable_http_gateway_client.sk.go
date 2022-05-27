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

type MatchableHttpGatewayWatcher interface {
	// watch namespace-scoped HttpGateways
	Watch(namespace string, opts clients.WatchOpts) (<-chan MatchableHttpGatewayList, <-chan error, error)
}

type MatchableHttpGatewayClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*MatchableHttpGateway, error)
	Write(resource *MatchableHttpGateway, opts clients.WriteOpts) (*MatchableHttpGateway, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (MatchableHttpGatewayList, error)
	MatchableHttpGatewayWatcher
}

type matchableHttpGatewayClient struct {
	rc clients.ResourceClient
}

func NewMatchableHttpGatewayClient(ctx context.Context, rcFactory factory.ResourceClientFactory) (MatchableHttpGatewayClient, error) {
	return NewMatchableHttpGatewayClientWithToken(ctx, rcFactory, "")
}

func NewMatchableHttpGatewayClientWithToken(ctx context.Context, rcFactory factory.ResourceClientFactory, token string) (MatchableHttpGatewayClient, error) {
	rc, err := rcFactory.NewResourceClient(ctx, factory.NewResourceClientParams{
		ResourceType: &MatchableHttpGateway{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base MatchableHttpGateway resource client")
	}
	return NewMatchableHttpGatewayClientWithBase(rc), nil
}

func NewMatchableHttpGatewayClientWithBase(rc clients.ResourceClient) MatchableHttpGatewayClient {
	return &matchableHttpGatewayClient{
		rc: rc,
	}
}

func (client *matchableHttpGatewayClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *matchableHttpGatewayClient) Register() error {
	return client.rc.Register()
}

func (client *matchableHttpGatewayClient) Read(namespace, name string, opts clients.ReadOpts) (*MatchableHttpGateway, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*MatchableHttpGateway), nil
}

func (client *matchableHttpGatewayClient) Write(matchableHttpGateway *MatchableHttpGateway, opts clients.WriteOpts) (*MatchableHttpGateway, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(matchableHttpGateway, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*MatchableHttpGateway), nil
}

func (client *matchableHttpGatewayClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *matchableHttpGatewayClient) List(namespace string, opts clients.ListOpts) (MatchableHttpGatewayList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToMatchableHttpGateway(resourceList), nil
}

func (client *matchableHttpGatewayClient) Watch(namespace string, opts clients.WatchOpts) (<-chan MatchableHttpGatewayList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	httpGatewaysChan := make(chan MatchableHttpGatewayList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				select {
				case httpGatewaysChan <- convertToMatchableHttpGateway(resourceList):
				case <-opts.Ctx.Done():
					close(httpGatewaysChan)
					return
				}
			case <-opts.Ctx.Done():
				close(httpGatewaysChan)
				return
			}
		}
	}()
	return httpGatewaysChan, errs, nil
}

func convertToMatchableHttpGateway(resources resources.ResourceList) MatchableHttpGatewayList {
	var matchableHttpGatewayList MatchableHttpGatewayList
	for _, resource := range resources {
		matchableHttpGatewayList = append(matchableHttpGatewayList, resource.(*MatchableHttpGateway))
	}
	return matchableHttpGatewayList
}
