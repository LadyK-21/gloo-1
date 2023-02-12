---
title: Proxy CRD Spec Compression
description: This document explains how to use a compressed spec for the Proxy CRD.
weight: 20
---

Gloo Edge's 2-layer API aggregates all the routes in to a single Proxy CRD. This is helpful for usability and debugging, and allows rapid support in various user level APIs.

When using many routes, the Proxy CRDs can get quite big, and may surpass the storage limits set by etcd (you will see a message in the logs: `etcdserver: request is too large`).

To workaround that (as of Gloo Edge 1.5.0+), you can enable compression for the Proxy CRD spec, so that the data stored in etcd is reduced significantly (by a factor of 10 or so).


 {{% notice note %}} Note that this is a **workaround** and should only be used when necessary. {{% /notice %}}


To enable proxy spec compression, edit the gloo the settings, like so:
```
kubectl patch settings -n gloo-system default --type='merge' -p '{"spec":{"gateway":{"compressedProxySpec":true}}}'
```

This will result in settings looking like this:
```yaml
apiVersion: gloo.solo.io/v1
kind: Settings
metadata:
  labels:
    app: gloo
  name: default
  namespace: gloo-system
spec:
  gateway:
    compressedProxySpec: true
    readGatewaysFromAllNamespaces: false
    validation:
      allowWarnings: true
      alwaysAccept: true
      proxyValidationServerAddr: gloo:9988
  …
  …
  …
```

Once set, the Proxies generated by the gateway proxy pod will contain a compressed spec. For example:

```yaml
apiVersion: gloo.solo.io/v1
kind: Proxy
metadata:
  annotations:
    # do NOT remove this annotation!
    gloo.solo.io/compress: "true"
  labels:
    created_by: gateway
  name: gateway-proxy
  namespace: gloo-system
spec:
  compressedSpec: eJzEUj1rwzA…SfOKY=
status:
  reportedBy: gloo
  state: 1
```
 {{% notice warning %}} The proxy object will have an extra annotation: `gloo.solo.io/compress: "true"`. Do not remove it! {{% /notice %}}


To view the contents of the spec in an uncompressed form, you can use one of the following commands:

{{< tabs >}}
{{< tab name="Linux Shell" codelang="shell">}}
kubectl get proxy gateway-proxy -n gloo-system -o 'jsonpath={@.spec.compressedSpec}'|base64 -d | openssl zlib -d|jq
{{< /tab >}}
{{< tab name="Mac OS" codelang="shell">}}
kubectl get proxy gateway-proxy -n gloo-system -o 'jsonpath={@.spec.compressedSpec}'|base64 -d | perl -e 'use Compress::Raw::Zlib;my $d=new Compress::Raw::Zlib::Inflate();my $o;undef $/;$d->inflate(<>,$o);print $o;'|jq
{{< /tab >}}
{{< tab name="Python" codelang="shell">}}
kubectl get proxy gateway-proxy -n gloo-system -o 'jsonpath={@.spec.compressedSpec}' | python3 -c 'import base64,zlib,sys;sys.stdout.write(zlib.decompress(base64.decodebytes(sys.stdin.read().encode("utf8"))).decode("utf8"))'|jq
{{< /tab >}}
{{< tab name="glooctl" codelang="shell">}}
glooctl get proxy -o yaml gateway-proxy
{{< /tab >}}
{{< /tabs >}}
