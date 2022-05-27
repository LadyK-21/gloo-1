// Code generated by solo-kit. DO NOT EDIT.

//Source: pkg/code-generator/codegen/templates/snapshot_emitter_template.go
package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/zap"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
)

var (
	// Deprecated. See mEnterpriseResourcesIn
	mEnterpriseSnapshotIn = stats.Int64("enterprise.gloo.solo.io/emitter/snap_in", "Deprecated. Use enterprise.gloo.solo.io/emitter/resources_in. The number of snapshots in", "1")

	// metrics for emitter
	mEnterpriseResourcesIn    = stats.Int64("enterprise.gloo.solo.io/emitter/resources_in", "The number of resource lists received on open watch channels", "1")
	mEnterpriseSnapshotOut    = stats.Int64("enterprise.gloo.solo.io/emitter/snap_out", "The number of snapshots out", "1")
	mEnterpriseSnapshotMissed = stats.Int64("enterprise.gloo.solo.io/emitter/snap_missed", "The number of snapshots missed", "1")

	// views for emitter
	// deprecated: see enterpriseResourcesInView
	enterprisesnapshotInView = &view.View{
		Name:        "enterprise.gloo.solo.io/emitter/snap_in",
		Measure:     mEnterpriseSnapshotIn,
		Description: "Deprecated. Use enterprise.gloo.solo.io/emitter/resources_in. The number of snapshots updates coming in.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}

	enterpriseResourcesInView = &view.View{
		Name:        "enterprise.gloo.solo.io/emitter/resources_in",
		Measure:     mEnterpriseResourcesIn,
		Description: "The number of resource lists received on open watch channels",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			skstats.NamespaceKey,
			skstats.ResourceKey,
		},
	}
	enterprisesnapshotOutView = &view.View{
		Name:        "enterprise.gloo.solo.io/emitter/snap_out",
		Measure:     mEnterpriseSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	enterprisesnapshotMissedView = &view.View{
		Name:        "enterprise.gloo.solo.io/emitter/snap_missed",
		Measure:     mEnterpriseSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(
		enterprisesnapshotInView,
		enterprisesnapshotOutView,
		enterprisesnapshotMissedView,
		enterpriseResourcesInView,
	)
}

type EnterpriseSnapshotEmitter interface {
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *EnterpriseSnapshot, <-chan error, error)
}

type EnterpriseEmitter interface {
	EnterpriseSnapshotEmitter
	Register() error
	AuthConfig() AuthConfigClient
}

func NewEnterpriseEmitter(authConfigClient AuthConfigClient) EnterpriseEmitter {
	return NewEnterpriseEmitterWithEmit(authConfigClient, make(chan struct{}))
}

func NewEnterpriseEmitterWithEmit(authConfigClient AuthConfigClient, emit <-chan struct{}) EnterpriseEmitter {
	return &enterpriseEmitter{
		authConfig: authConfigClient,
		forceEmit:  emit,
	}
}

type enterpriseEmitter struct {
	forceEmit  <-chan struct{}
	authConfig AuthConfigClient
}

func (c *enterpriseEmitter) Register() error {
	if err := c.authConfig.Register(); err != nil {
		return err
	}
	return nil
}

func (c *enterpriseEmitter) AuthConfig() AuthConfigClient {
	return c.authConfig
}

func (c *enterpriseEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *EnterpriseSnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for AuthConfig */
	type authConfigListWithNamespace struct {
		list      AuthConfigList
		namespace string
	}
	authConfigChan := make(chan authConfigListWithNamespace)

	var initialAuthConfigList AuthConfigList

	currentSnapshot := EnterpriseSnapshot{}

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for AuthConfig */
		{
			authConfigs, err := c.authConfig.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial AuthConfig list")
			}
			initialAuthConfigList = append(initialAuthConfigList, authConfigs...)
		}
		authConfigNamespacesChan, authConfigErrs, err := c.authConfig.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting AuthConfig watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, authConfigErrs, namespace+"-authConfigs")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case authConfigList, ok := <-authConfigNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case authConfigChan <- authConfigListWithNamespace{list: authConfigList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for AuthConfigs */
	currentSnapshot.AuthConfigs = initialAuthConfigList.Sort()

	snapshots := make(chan *EnterpriseSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		previousHash, err := currentSnapshot.Hash(nil)
		if err != nil {
			contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
		}
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			// this should never happen, so panic if it does
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mEnterpriseSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mEnterpriseSnapshotMissed.M(1))
			}
		}
		authConfigsByNamespace := make(map[string]AuthConfigList)
		defer func() {
			close(snapshots)
			// we must wait for done before closing the error chan,
			// to avoid sending on close channel.
			done.Wait()
			close(errs)
		}()
		for {
			record := func() { stats.Record(ctx, mEnterpriseSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case authConfigNamespacedList, ok := <-authConfigChan:
				if !ok {
					return
				}
				record()

				namespace := authConfigNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"auth_config",
					mEnterpriseResourcesIn,
				)

				// merge lists by namespace
				authConfigsByNamespace[namespace] = authConfigNamespacedList.list
				var authConfigList AuthConfigList
				for _, authConfigs := range authConfigsByNamespace {
					authConfigList = append(authConfigList, authConfigs...)
				}
				currentSnapshot.AuthConfigs = authConfigList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}
