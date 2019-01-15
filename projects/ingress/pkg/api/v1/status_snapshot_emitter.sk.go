// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	mStatusSnapshotIn  = stats.Int64("status.ingress.solo.io/snap_emitter/snap_in", "The number of snapshots in", "1")
	mStatusSnapshotOut = stats.Int64("status.ingress.solo.io/snap_emitter/snap_out", "The number of snapshots out", "1")

	statussnapshotInView = &view.View{
		Name:        "status.ingress.solo.io_snap_emitter/snap_in",
		Measure:     mStatusSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	statussnapshotOutView = &view.View{
		Name:        "status.ingress.solo.io/snap_emitter/snap_out",
		Measure:     mStatusSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(statussnapshotInView, statussnapshotOutView)
}

type StatusEmitter interface {
	Register() error
	KubeService() KubeServiceClient
	Ingress() IngressClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *StatusSnapshot, <-chan error, error)
}

func NewStatusEmitter(kubeServiceClient KubeServiceClient, ingressClient IngressClient) StatusEmitter {
	return NewStatusEmitterWithEmit(kubeServiceClient, ingressClient, make(chan struct{}))
}

func NewStatusEmitterWithEmit(kubeServiceClient KubeServiceClient, ingressClient IngressClient, emit <-chan struct{}) StatusEmitter {
	return &statusEmitter{
		kubeService: kubeServiceClient,
		ingress:     ingressClient,
		forceEmit:   emit,
	}
}

type statusEmitter struct {
	forceEmit   <-chan struct{}
	kubeService KubeServiceClient
	ingress     IngressClient
}

func (c *statusEmitter) Register() error {
	if err := c.kubeService.Register(); err != nil {
		return err
	}
	if err := c.ingress.Register(); err != nil {
		return err
	}
	return nil
}

func (c *statusEmitter) KubeService() KubeServiceClient {
	return c.kubeService
}

func (c *statusEmitter) Ingress() IngressClient {
	return c.ingress
}

func (c *statusEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *StatusSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for KubeService */
	type kubeServiceListWithNamespace struct {
		list      KubeServiceList
		namespace string
	}
	kubeServiceChan := make(chan kubeServiceListWithNamespace)
	/* Create channel for Ingress */
	type ingressListWithNamespace struct {
		list      IngressList
		namespace string
	}
	ingressChan := make(chan ingressListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for KubeService */
		kubeServiceNamespacesChan, kubeServiceErrs, err := c.kubeService.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting KubeService watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, kubeServiceErrs, namespace+"-services")
		}(namespace)
		/* Setup namespaced watch for Ingress */
		ingressNamespacesChan, ingressErrs, err := c.ingress.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Ingress watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, ingressErrs, namespace+"-ingresses")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case kubeServiceList := <-kubeServiceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case kubeServiceChan <- kubeServiceListWithNamespace{list: kubeServiceList, namespace: namespace}:
					}
				case ingressList := <-ingressNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case ingressChan <- ingressListWithNamespace{list: ingressList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *StatusSnapshot)
	go func() {
		originalSnapshot := StatusSnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 1)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}

			stats.Record(ctx, mStatusSnapshotOut.M(1))
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		/* TODO (yuval-k): figure out how to make this work to avoid a stale snapshot.
		   		// construct the first snapshot from all the configs that are currently there
		   		// that guarantees that the first snapshot contains all the data.
		   		for range watchNamespaces {
		      kubeServiceNamespacedList := <- kubeServiceChan
		      currentSnapshot.Services.Clear(kubeServiceNamespacedList.namespace)
		      kubeServiceList := kubeServiceNamespacedList.list
		   	currentSnapshot.Services.Add(kubeServiceList...)
		      ingressNamespacedList := <- ingressChan
		      currentSnapshot.Ingresses.Clear(ingressNamespacedList.namespace)
		      ingressList := ingressNamespacedList.list
		   	currentSnapshot.Ingresses.Add(ingressList...)
		   		}
		*/

		for {
			record := func() { stats.Record(ctx, mStatusSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case kubeServiceNamespacedList := <-kubeServiceChan:
				record()

				namespace := kubeServiceNamespacedList.namespace
				kubeServiceList := kubeServiceNamespacedList.list

				currentSnapshot.Services.Clear(namespace)
				currentSnapshot.Services.Add(kubeServiceList...)
			case ingressNamespacedList := <-ingressChan:
				record()

				namespace := ingressNamespacedList.namespace
				ingressList := ingressNamespacedList.list

				currentSnapshot.Ingresses.Clear(namespace)
				currentSnapshot.Ingresses.Add(ingressList...)
			}
		}
	}()
	return snapshots, errs, nil
}
