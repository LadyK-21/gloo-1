// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/connection.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *ConnectionConfig) Clone() proto.Message {
	var target *ConnectionConfig
	if m == nil {
		return target
	}
	target = &ConnectionConfig{}

	target.MaxRequestsPerConnection = m.GetMaxRequestsPerConnection()

	if h, ok := interface{}(m.GetConnectTimeout()).(clone.Cloner); ok {
		target.ConnectTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.ConnectTimeout = proto.Clone(m.GetConnectTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetTcpKeepalive()).(clone.Cloner); ok {
		target.TcpKeepalive = h.Clone().(*ConnectionConfig_TcpKeepAlive)
	} else {
		target.TcpKeepalive = proto.Clone(m.GetTcpKeepalive()).(*ConnectionConfig_TcpKeepAlive)
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(clone.Cloner); ok {
		target.PerConnectionBufferLimitBytes = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.PerConnectionBufferLimitBytes = proto.Clone(m.GetPerConnectionBufferLimitBytes()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetCommonHttpProtocolOptions()).(clone.Cloner); ok {
		target.CommonHttpProtocolOptions = h.Clone().(*ConnectionConfig_HttpProtocolOptions)
	} else {
		target.CommonHttpProtocolOptions = proto.Clone(m.GetCommonHttpProtocolOptions()).(*ConnectionConfig_HttpProtocolOptions)
	}

	if h, ok := interface{}(m.GetHttp1ProtocolOptions()).(clone.Cloner); ok {
		target.Http1ProtocolOptions = h.Clone().(*ConnectionConfig_Http1ProtocolOptions)
	} else {
		target.Http1ProtocolOptions = proto.Clone(m.GetHttp1ProtocolOptions()).(*ConnectionConfig_Http1ProtocolOptions)
	}

	return target
}

// Clone function
func (m *ConnectionConfig_TcpKeepAlive) Clone() proto.Message {
	var target *ConnectionConfig_TcpKeepAlive
	if m == nil {
		return target
	}
	target = &ConnectionConfig_TcpKeepAlive{}

	target.KeepaliveProbes = m.GetKeepaliveProbes()

	if h, ok := interface{}(m.GetKeepaliveTime()).(clone.Cloner); ok {
		target.KeepaliveTime = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.KeepaliveTime = proto.Clone(m.GetKeepaliveTime()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetKeepaliveInterval()).(clone.Cloner); ok {
		target.KeepaliveInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.KeepaliveInterval = proto.Clone(m.GetKeepaliveInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *ConnectionConfig_HttpProtocolOptions) Clone() proto.Message {
	var target *ConnectionConfig_HttpProtocolOptions
	if m == nil {
		return target
	}
	target = &ConnectionConfig_HttpProtocolOptions{}

	if h, ok := interface{}(m.GetIdleTimeout()).(clone.Cloner); ok {
		target.IdleTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.IdleTimeout = proto.Clone(m.GetIdleTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.MaxHeadersCount = m.GetMaxHeadersCount()

	if h, ok := interface{}(m.GetMaxStreamDuration()).(clone.Cloner); ok {
		target.MaxStreamDuration = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxStreamDuration = proto.Clone(m.GetMaxStreamDuration()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	target.HeadersWithUnderscoresAction = m.GetHeadersWithUnderscoresAction()

	target.OverrideStreamErrorOnInvalidHttpMessage = m.GetOverrideStreamErrorOnInvalidHttpMessage()

	return target
}

// Clone function
func (m *ConnectionConfig_Http1ProtocolOptions) Clone() proto.Message {
	var target *ConnectionConfig_Http1ProtocolOptions
	if m == nil {
		return target
	}
	target = &ConnectionConfig_Http1ProtocolOptions{}

	target.EnableTrailers = m.GetEnableTrailers()

	switch m.HeaderFormat.(type) {

	case *ConnectionConfig_Http1ProtocolOptions_ProperCaseHeaderKeyFormat:

		target.HeaderFormat = &ConnectionConfig_Http1ProtocolOptions_ProperCaseHeaderKeyFormat{
			ProperCaseHeaderKeyFormat: m.GetProperCaseHeaderKeyFormat(),
		}

	case *ConnectionConfig_Http1ProtocolOptions_PreserveCaseHeaderKeyFormat:

		target.HeaderFormat = &ConnectionConfig_Http1ProtocolOptions_PreserveCaseHeaderKeyFormat{
			PreserveCaseHeaderKeyFormat: m.GetPreserveCaseHeaderKeyFormat(),
		}

	}

	return target
}
