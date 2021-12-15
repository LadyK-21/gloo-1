
---
title: "graphql.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `graphql.gloo.solo.io` 
#### Types:


- [RequestTemplate](#requesttemplate)
- [ResponseTemplate](#responsetemplate)
- [RESTResolver](#restresolver)
- [QueryMatcher](#querymatcher)
- [FieldMatcher](#fieldmatcher)
- [Resolution](#resolution)
- [GraphQLSchema](#graphqlschema) **Top-Level Resource**
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto)





---
### RequestTemplate

 
Defines a configuration for generating outgoing requests for a resolver.

```yaml
"headers": map<string, string>
"queryParams": map<string, string>
"body": .google.protobuf.Value

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `headers` | `map<string, string>` | Use this attribute to set request headers to your REST service. It consists of a map of strings to value providers. The string key determines the name of the resulting header, the value provided will be the value. at least need ":method" and ":path". |
| `queryParams` | `map<string, string>` | Use this attribute to set query parameters to your REST service. It consists of a map of strings to value providers. The string key determines the name of the query param, the provided value will be the value. This value is appended to any value set to the :path header in `headers`. Interpolation is done in envoy rather than the control plane to prevent escaped character issues. Additionally, we may be providing values not known until the request is being executed (e.g., graphql parent info). |
| `body` | [.google.protobuf.Value](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/value) | Used to construct the outgoing body to the upstream from the graphql value providers. |




---
### ResponseTemplate



```yaml
"resultRoot": string
"setters": map<string, string>

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `resultRoot` | `string` |  |
| `setters` | `map<string, string>` |  |




---
### RESTResolver

 
control-plane API

```yaml
"upstreamRef": .core.solo.io.ResourceRef
"request": .graphql.gloo.solo.io.RequestTemplate
"response": .graphql.gloo.solo.io.ResponseTemplate
"spanName": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `upstreamRef` | [.core.solo.io.ResourceRef](../../../../../../../../../../solo-kit/api/v1/ref.proto.sk/#resourceref) |  |
| `request` | [.graphql.gloo.solo.io.RequestTemplate](../graphql.proto.sk/#requesttemplate) | configuration used to compose the outgoing request to a REST API. |
| `response` | [.graphql.gloo.solo.io.ResponseTemplate](../graphql.proto.sk/#responsetemplate) |  |
| `spanName` | `string` |  |




---
### QueryMatcher



```yaml
"fieldMatcher": .graphql.gloo.solo.io.QueryMatcher.FieldMatcher

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `fieldMatcher` | [.graphql.gloo.solo.io.QueryMatcher.FieldMatcher](../graphql.proto.sk/#fieldmatcher) |  |




---
### FieldMatcher



```yaml
"type": string
"field": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `type` | `string` | Object type. For example, Query. |
| `field` | `string` | Field with in the object. |




---
### Resolution

 
This is the resolver map for the schema.
For each Type.Field, we can define a resolver.
if a field does not have resolver, the default resolver will be used.
the default resolver takes the field with the same name from the parent, and uses that value
to resolve the field.
if a field with the same name does not exist in the parent, null will be used.

```yaml
"matcher": .graphql.gloo.solo.io.QueryMatcher
"restResolver": .graphql.gloo.solo.io.RESTResolver

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `matcher` | [.graphql.gloo.solo.io.QueryMatcher](../graphql.proto.sk/#querymatcher) | Match an object type and field. |
| `restResolver` | [.graphql.gloo.solo.io.RESTResolver](../graphql.proto.sk/#restresolver) |  |




---
### GraphQLSchema

 
Enterprise-Only: THIS FEATURE IS IN TECH PREVIEW. APIs are versioned as alpha and subject to change.
User-facing CR config for resolving client requests to graphql schemas.
Routes that have this config will execute graphql queries, and will not make it to the router filter. i.e. this
filter will terminate the request for these routes.
Note: while users can provide this configuration manually, the eventual UX will
be to generate the Executable Schema CRs from other sources and just have users
configure the routes to point to these schema CRs.

```yaml
"namespacedStatuses": .core.solo.io.NamespacedStatuses
"metadata": .core.solo.io.Metadata
"schema": string
"enableIntrospection": bool
"resolutions": []graphql.gloo.solo.io.Resolution

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `namespacedStatuses` | [.core.solo.io.NamespacedStatuses](../../../../../../../../../../solo-kit/api/v1/status.proto.sk/#namespacedstatuses) | NamespacedStatuses indicates the validation status of this resource. NamespacedStatuses is read-only by clients, and set by gloo during validation. |
| `metadata` | [.core.solo.io.Metadata](../../../../../../../../../../solo-kit/api/v1/metadata.proto.sk/#metadata) | Metadata contains the object metadata for this resource. |
| `schema` | `string` | Schema to use in string format. |
| `enableIntrospection` | `bool` | Do we enable introspection for the schema? general recommendation is to disable this for production and hence it defaults to false. |
| `resolutions` | [[]graphql.gloo.solo.io.Resolution](../graphql.proto.sk/#resolution) | The resolver map to use to resolve the schema. Omitted fields will use the default resolver, which looks for a field with that name in the parent's object, and errors if the field cannot be found. |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->