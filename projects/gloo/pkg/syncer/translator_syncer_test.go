package syncer_test

import (
	"context"

	"github.com/solo-io/gloo/projects/gloo/pkg/xds"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/utils/statusutils"
	"github.com/solo-io/gloo/projects/gateway/pkg/utils/metrics"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	. "github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	"github.com/solo-io/solo-kit/pkg/errors"
)

var _ = Describe("Translate Proxy", func() {

	var xdsCache *MockXdsCache
	var sanitizer *MockXdsSanitizer
	var syncer v1snap.ApiSyncer
	var snap *v1snap.ApiSnapshot
	var settings *v1.Settings
	var upstreamClient clients.ResourceClient
	var proxyClient v1.ProxyClient
	var ctx context.Context
	var cancel context.CancelFunc
	var proxyName = "proxy-name"
	var ns = "any-ns"
	var ref = "syncer-test"
	var statusClient resources.StatusClient
	var statusMetrics metrics.ConfigStatusMetrics

	BeforeEach(func() {
		var err error
		xdsCache = &MockXdsCache{}
		sanitizer = &MockXdsSanitizer{}
		ctx, cancel = context.WithCancel(context.Background())

		resourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		proxyClient, _ = v1.NewProxyClient(ctx, resourceClientFactory)

		upstreamClient, err = resourceClientFactory.NewResourceClient(ctx, factory.NewResourceClientParams{ResourceType: &v1.Upstream{}})
		Expect(err).NotTo(HaveOccurred())

		proxy := &v1.Proxy{
			Metadata: &core.Metadata{
				Namespace: ns,
				Name:      proxyName,
			},
		}

		settings = &v1.Settings{}

		statusClient = statusutils.GetStatusClientFromEnvOrDefault(ns)
		statusMetrics, err = metrics.NewConfigStatusMetrics(metrics.GetDefaultConfigStatusOptions())
		Expect(err).NotTo(HaveOccurred())

		rep := reporter.NewReporter(ref, statusClient, proxyClient.BaseClient(), upstreamClient)

		xdsHasher := xds.NewNodeRoleHasher()
		syncer = NewTranslatorSyncer(&mockTranslator{true, false, nil}, xdsCache, xdsHasher, sanitizer, rep, false, nil, settings, statusMetrics, nil, proxyClient, "")
		snap = &v1snap.ApiSnapshot{
			Proxies: v1.ProxyList{
				proxy,
			},
		}
		_, err = proxyClient.Write(proxy, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		err = syncer.Sync(context.Background(), snap)
		Expect(err).NotTo(HaveOccurred())

		proxies, err := proxyClient.List(proxy.GetMetadata().Namespace, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		Expect(proxies).To(HaveLen(1))
		Expect(proxies[0]).To(BeAssignableToTypeOf(&v1.Proxy{}))
		Expect(statusClient.GetStatus(proxies[0])).To(Equal(&core.Status{
			State:      2,
			Reason:     "1 error occurred:\n\t* hi, how ya doin'?\n\n",
			ReportedBy: ref,
		}))

		// NilSnapshot is always consistent, so snapshot will always be set as part of endpoints update
		Expect(xdsCache.Called).To(BeTrue())

		// update rv for proxy
		p1, err := proxyClient.Read(proxy.Metadata.Namespace, proxy.Metadata.Name, clients.ReadOpts{})
		Expect(err).NotTo(HaveOccurred())
		snap.Proxies[0] = p1

		syncer = NewTranslatorSyncer(&mockTranslator{false, false, nil}, xdsCache, xdsHasher, sanitizer, rep, false, nil, settings, statusMetrics, nil, proxyClient, "")

		err = syncer.Sync(context.Background(), snap)
		Expect(err).NotTo(HaveOccurred())

	})

	AfterEach(func() { cancel() })

	It("writes the reports the translator spits out and calls SetSnapshot on the cache", func() {
		proxies, err := proxyClient.List(ns, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		Expect(proxies).To(HaveLen(1))
		Expect(proxies[0]).To(BeAssignableToTypeOf(&v1.Proxy{}))
		Expect(statusClient.GetStatus(proxies[0])).To(Equal(&core.Status{
			State:      1,
			ReportedBy: ref,
		}))

		Expect(xdsCache.Called).To(BeTrue())
	})

	It("updates the cache with the sanitized snapshot", func() {
		sanitizer.Snap = envoycache.NewEasyGenericSnapshot("easy")
		err := syncer.Sync(context.Background(), snap)
		Expect(err).NotTo(HaveOccurred())

		Expect(sanitizer.Called).To(BeTrue())
		Expect(xdsCache.SetSnap).To(BeEquivalentTo(sanitizer.Snap))
	})
})

var _ = Describe("Translate multiple proxies with errors", func() {

	var (
		xdsCache       *MockXdsCache
		sanitizer      *MockXdsSanitizer
		syncer         v1snap.ApiSyncer
		snap           *v1snap.ApiSnapshot
		settings       *v1.Settings
		proxyClient    v1.ProxyClient
		upstreamClient v1.UpstreamClient
		proxyName      = "proxy-name"
		upstreamName   = "upstream-name"
		ns             = "any-ns"
		ref            = "syncer-test"
		statusClient   resources.StatusClient
		statusMetrics  metrics.ConfigStatusMetrics
	)

	proxiesShouldHaveErrors := func(proxies v1.ProxyList, numProxies int) {
		Expect(proxies).To(HaveLen(numProxies))
		for _, proxy := range proxies {
			Expect(proxy).To(BeAssignableToTypeOf(&v1.Proxy{}))
			Expect(statusClient.GetStatus(proxy)).To(Equal(&core.Status{
				State:      2,
				Reason:     "1 error occurred:\n\t* hi, how ya doin'?\n\n",
				ReportedBy: ref,
			}))

		}

	}
	writeUniqueErrsToUpstreams := func() {
		// Re-writes existing upstream to have an annotation
		// which triggers a unique error to be written from each proxy's mockTranslator
		upstreams, err := upstreamClient.List(ns, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		Expect(upstreams).To(HaveLen(1))

		us := upstreams[0]
		// This annotation causes the translator mock to generate a unique error per proxy on each upstream
		us.Metadata.Annotations = map[string]string{"uniqueErrPerProxy": "true"}
		_, err = upstreamClient.Write(us, clients.WriteOpts{OverwriteExisting: true})
		Expect(err).NotTo(HaveOccurred())
		snap.Upstreams = upstreams
		err = syncer.Sync(context.Background(), snap)
		Expect(err).NotTo(HaveOccurred())
	}

	BeforeEach(func() {
		var err error
		xdsCache = &MockXdsCache{}
		sanitizer = &MockXdsSanitizer{}

		resourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		proxyClient, _ = v1.NewProxyClient(context.Background(), resourceClientFactory)

		usClient, err := resourceClientFactory.NewResourceClient(context.Background(), factory.NewResourceClientParams{ResourceType: &v1.Upstream{}})
		Expect(err).NotTo(HaveOccurred())

		proxy1 := &v1.Proxy{
			Metadata: &core.Metadata{
				Namespace: ns,
				Name:      proxyName + "1",
			},
		}
		proxy2 := &v1.Proxy{
			Metadata: &core.Metadata{
				Namespace: ns,
				Name:      proxyName + "2",
			},
		}

		us := &v1.Upstream{
			Metadata: &core.Metadata{
				Name:      upstreamName,
				Namespace: ns,
			},
		}

		settings = &v1.Settings{}

		statusClient = statusutils.GetStatusClientFromEnvOrDefault(ns)
		statusMetrics, err = metrics.NewConfigStatusMetrics(metrics.GetDefaultConfigStatusOptions())
		Expect(err).NotTo(HaveOccurred())

		rep := reporter.NewReporter(ref, statusClient, proxyClient.BaseClient(), usClient)

		xdsHasher := xds.NewNodeRoleHasher()
		syncer = NewTranslatorSyncer(&mockTranslator{true, true, nil}, xdsCache, xdsHasher, sanitizer, rep, false, nil, settings, statusMetrics, nil, proxyClient, "")
		snap = &v1snap.ApiSnapshot{
			Proxies: v1.ProxyList{
				proxy1,
				proxy2,
			},
			Upstreams: v1.UpstreamList{
				us,
			},
		}

		_, err = usClient.Write(us, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = proxyClient.Write(proxy1, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = proxyClient.Write(proxy2, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		err = syncer.Sync(context.Background(), snap)
		Expect(err).NotTo(HaveOccurred())

		proxies, err := proxyClient.List(proxy1.GetMetadata().Namespace, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		Expect(proxies).To(HaveLen(2))
		Expect(proxies[0]).To(BeAssignableToTypeOf(&v1.Proxy{}))
		Expect(statusClient.GetStatus(proxies[0])).To(Equal(&core.Status{
			State:      2,
			Reason:     "1 error occurred:\n\t* hi, how ya doin'?\n\n",
			ReportedBy: ref,
		}))

		// NilSnapshot is always consistent, so snapshot will always be set as part of endpoints update
		Expect(xdsCache.Called).To(BeTrue())

		upstreamClient, err = v1.NewUpstreamClient(context.Background(), resourceClientFactory)
		Expect(err).NotTo(HaveOccurred())
	})

	It("handles reporting errors on multiple proxies sharing an upstream reporting 2 different errors", func() {
		// Testing the scenario where we have multiple proxies,
		// each of which should report a different unique error on an upstream.
		proxies, err := proxyClient.List(ns, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		proxiesShouldHaveErrors(proxies, 2)

		writeUniqueErrsToUpstreams()

		upstreams, err := upstreamClient.List(ns, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())

		Expect(statusClient.GetStatus(upstreams[0])).To(Equal(&core.Status{
			State:      2,
			Reason:     "2 errors occurred:\n\t* upstream is bad - determined by proxy-name1\n\t* upstream is bad - determined by proxy-name2\n\n",
			ReportedBy: ref,
		}))

		Expect(xdsCache.Called).To(BeTrue())
	})

	It("handles reporting errors on multiple proxies sharing an upstream, each reporting the same upstream error", func() {
		// Testing the scenario where we have multiple proxies,
		// each of which should report the same error on an upstream.
		proxies, err := proxyClient.List(ns, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		proxiesShouldHaveErrors(proxies, 2)

		upstreams, err := upstreamClient.List(ns, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		Expect(upstreams).To(HaveLen(1))
		Expect(statusClient.GetStatus(upstreams[0])).To(Equal(&core.Status{
			State:      2,
			Reason:     "1 error occurred:\n\t* generic upstream error\n\n",
			ReportedBy: ref,
		}))

		Expect(xdsCache.Called).To(BeTrue())
	})
})

type mockTranslator struct {
	reportErrs         bool
	reportUpstreamErrs bool // Adds an error to every upstream in the snapshot
	currentSnapshot    envoycache.Snapshot
}

func (t *mockTranslator) Translate(params plugins.Params, proxy *v1.Proxy) (envoycache.Snapshot, reporter.ResourceReports, *validation.ProxyReport, error) {
	if t.reportErrs {
		rpts := reporter.ResourceReports{}
		rpts.AddError(proxy, errors.Errorf("hi, how ya doin'?"))
		rpts.AddMessages(proxy, "This is a message")
		rpts.AddMessages(proxy, "And also this")
		if t.reportUpstreamErrs {
			for _, upstream := range params.Snapshot.Upstreams {
				if upstream.Metadata.Annotations["uniqueErrPerProxy"] == "true" {
					rpts.AddError(upstream, errors.Errorf("upstream is bad - determined by %s", proxy.Metadata.Name))
				} else {
					rpts.AddError(upstream, errors.Errorf("generic upstream error"))
				}
			}
		}
		if t.currentSnapshot != nil {
			return t.currentSnapshot, rpts, &validation.ProxyReport{}, nil
		}
		return envoycache.NilSnapshot{}, rpts, &validation.ProxyReport{}, nil
	}
	if t.currentSnapshot != nil {
		return t.currentSnapshot, nil, &validation.ProxyReport{}, nil
	}
	return envoycache.NilSnapshot{}, nil, &validation.ProxyReport{}, nil
}

var _ envoycache.SnapshotCache = &MockXdsCache{}
