// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/wasm/wasm.proto

package wasm

import (
	reflect "reflect"
	sync "sync"

	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// represents the different types of WASM VMs available with which envoy can run
// the WASM filter module
type WasmFilter_VmType int32

const (
	WasmFilter_V8   WasmFilter_VmType = 0
	WasmFilter_WAVM WasmFilter_VmType = 1
)

// Enum value maps for WasmFilter_VmType.
var (
	WasmFilter_VmType_name = map[int32]string{
		0: "V8",
		1: "WAVM",
	}
	WasmFilter_VmType_value = map[string]int32{
		"V8":   0,
		"WAVM": 1,
	}
)

func (x WasmFilter_VmType) Enum() *WasmFilter_VmType {
	p := new(WasmFilter_VmType)
	*p = x
	return p
}

func (x WasmFilter_VmType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WasmFilter_VmType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes[0].Descriptor()
}

func (WasmFilter_VmType) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes[0]
}

func (x WasmFilter_VmType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WasmFilter_VmType.Descriptor instead.
func (WasmFilter_VmType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescGZIP(), []int{1, 0}
}

// list of filter stages which can be selected for a WASM filter
type FilterStage_Stage int32

const (
	FilterStage_FaultStage     FilterStage_Stage = 0
	FilterStage_CorsStage      FilterStage_Stage = 1
	FilterStage_WafStage       FilterStage_Stage = 2
	FilterStage_AuthNStage     FilterStage_Stage = 3
	FilterStage_AuthZStage     FilterStage_Stage = 4
	FilterStage_RateLimitStage FilterStage_Stage = 5
	FilterStage_AcceptedStage  FilterStage_Stage = 6
	FilterStage_OutAuthStage   FilterStage_Stage = 7
	FilterStage_RouteStage     FilterStage_Stage = 8
)

// Enum value maps for FilterStage_Stage.
var (
	FilterStage_Stage_name = map[int32]string{
		0: "FaultStage",
		1: "CorsStage",
		2: "WafStage",
		3: "AuthNStage",
		4: "AuthZStage",
		5: "RateLimitStage",
		6: "AcceptedStage",
		7: "OutAuthStage",
		8: "RouteStage",
	}
	FilterStage_Stage_value = map[string]int32{
		"FaultStage":     0,
		"CorsStage":      1,
		"WafStage":       2,
		"AuthNStage":     3,
		"AuthZStage":     4,
		"RateLimitStage": 5,
		"AcceptedStage":  6,
		"OutAuthStage":   7,
		"RouteStage":     8,
	}
)

func (x FilterStage_Stage) Enum() *FilterStage_Stage {
	p := new(FilterStage_Stage)
	*p = x
	return p
}

func (x FilterStage_Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterStage_Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes[1].Descriptor()
}

func (FilterStage_Stage) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes[1]
}

func (x FilterStage_Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterStage_Stage.Descriptor instead.
func (FilterStage_Stage) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescGZIP(), []int{2, 0}
}

// During is the 0th member so that it is the default, even though
// the reading order can be a little confusing
type FilterStage_Predicate int32

const (
	FilterStage_During FilterStage_Predicate = 0
	FilterStage_Before FilterStage_Predicate = 1
	FilterStage_After  FilterStage_Predicate = 2
)

// Enum value maps for FilterStage_Predicate.
var (
	FilterStage_Predicate_name = map[int32]string{
		0: "During",
		1: "Before",
		2: "After",
	}
	FilterStage_Predicate_value = map[string]int32{
		"During": 0,
		"Before": 1,
		"After":  2,
	}
)

func (x FilterStage_Predicate) Enum() *FilterStage_Predicate {
	p := new(FilterStage_Predicate)
	*p = x
	return p
}

func (x FilterStage_Predicate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterStage_Predicate) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes[2].Descriptor()
}

func (FilterStage_Predicate) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes[2]
}

func (x FilterStage_Predicate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterStage_Predicate.Descriptor instead.
func (FilterStage_Predicate) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescGZIP(), []int{2, 1}
}

//
//Options config for WASM filters
type PluginSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of WASM filters to be added into the filter chain
	Filters []*WasmFilter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *PluginSource) Reset() {
	*x = PluginSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginSource) ProtoMessage() {}

func (x *PluginSource) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginSource.ProtoReflect.Descriptor instead.
func (*PluginSource) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *PluginSource) GetFilters() []*WasmFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

//
//This message defines a single Envoy WASM filter to be placed into the filter chain
type WasmFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Src:
	//	*WasmFilter_Image
	//	*WasmFilter_FilePath
	Src isWasmFilter_Src `protobuf_oneof:"src"`
	// Filter/service configuration used to configure or reconfigure a plugin
	// (proxy_on_configuration).
	// `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Config *any.Any `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	// the stage in the filter chain where this filter should be placed
	FilterStage *FilterStage `protobuf:"bytes,4,opt,name=filter_stage,json=filterStage,proto3" json:"filter_stage,omitempty"`
	// the name of the filter, used for logging
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// the root_id of the filter which should be run, if this value is incorrect, or
	// empty the filter will crash
	RootId string `protobuf:"bytes,6,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	// selected VM type
	VmType WasmFilter_VmType `protobuf:"varint,7,opt,name=vm_type,json=vmType,proto3,enum=wasm.options.gloo.solo.io.WasmFilter_VmType" json:"vm_type,omitempty"`
	// when set to true if there is a fatal error on the VM the filter will be bypassed
	FailOpen bool `protobuf:"varint,9,opt,name=fail_open,json=failOpen,proto3" json:"fail_open,omitempty"`
}

func (x *WasmFilter) Reset() {
	*x = WasmFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasmFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasmFilter) ProtoMessage() {}

func (x *WasmFilter) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasmFilter.ProtoReflect.Descriptor instead.
func (*WasmFilter) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescGZIP(), []int{1}
}

func (m *WasmFilter) GetSrc() isWasmFilter_Src {
	if m != nil {
		return m.Src
	}
	return nil
}

func (x *WasmFilter) GetImage() string {
	if x, ok := x.GetSrc().(*WasmFilter_Image); ok {
		return x.Image
	}
	return ""
}

func (x *WasmFilter) GetFilePath() string {
	if x, ok := x.GetSrc().(*WasmFilter_FilePath); ok {
		return x.FilePath
	}
	return ""
}

func (x *WasmFilter) GetConfig() *any.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *WasmFilter) GetFilterStage() *FilterStage {
	if x != nil {
		return x.FilterStage
	}
	return nil
}

func (x *WasmFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WasmFilter) GetRootId() string {
	if x != nil {
		return x.RootId
	}
	return ""
}

func (x *WasmFilter) GetVmType() WasmFilter_VmType {
	if x != nil {
		return x.VmType
	}
	return WasmFilter_V8
}

func (x *WasmFilter) GetFailOpen() bool {
	if x != nil {
		return x.FailOpen
	}
	return false
}

type isWasmFilter_Src interface {
	isWasmFilter_Src()
}

type WasmFilter_Image struct {
	// name of image which houses the compiled wasm filter
	Image string `protobuf:"bytes,2,opt,name=image,proto3,oneof"`
}

type WasmFilter_FilePath struct {
	// path from which to load wasm filter from disk
	FilePath string `protobuf:"bytes,8,opt,name=file_path,json=filePath,proto3,oneof"`
}

func (*WasmFilter_Image) isWasmFilter_Src() {}

func (*WasmFilter_FilePath) isWasmFilter_Src() {}

type FilterStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stage of the filter chain in which the selected filter should be added
	Stage FilterStage_Stage `protobuf:"varint,1,opt,name=stage,proto3,enum=wasm.options.gloo.solo.io.FilterStage_Stage" json:"stage,omitempty"`
	// How this filter should be placed relative to the stage
	Predicate FilterStage_Predicate `protobuf:"varint,2,opt,name=predicate,proto3,enum=wasm.options.gloo.solo.io.FilterStage_Predicate" json:"predicate,omitempty"`
}

func (x *FilterStage) Reset() {
	*x = FilterStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterStage) ProtoMessage() {}

func (x *FilterStage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterStage.ProtoReflect.Descriptor instead.
func (*FilterStage) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescGZIP(), []int{2}
}

func (x *FilterStage) GetStage() FilterStage_Stage {
	if x != nil {
		return x.Stage
	}
	return FilterStage_FaultStage
}

func (x *FilterStage) GetPredicate() FilterStage_Predicate {
	if x != nil {
		return x.Predicate
	}
	return FilterStage_During
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2f, 0x77, 0x61, 0x73, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4f, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x61,
	0x73, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x22, 0xf0, 0x02, 0x0a, 0x0a, 0x57, 0x61, 0x73, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x61,
	0x73, 0x6d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a,
	0x07, 0x76, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x61, 0x73, 0x6d, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x76, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x6f, 0x70, 0x65,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x65,
	0x6e, 0x22, 0x1a, 0x0a, 0x06, 0x56, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x56,
	0x38, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x56, 0x4d, 0x10, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x73, 0x72, 0x63, 0x22, 0xf1, 0x02, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x77, 0x61,
	0x73, 0x6d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x72, 0x73, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x61, 0x66, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x5a, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x04, 0x12,
	0x12, 0x0a, 0x0e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x08, 0x22, 0x2e, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x66, 0x74, 0x65, 0x72, 0x10, 0x02, 0x42, 0x4b, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04,
	0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_goTypes = []interface{}{
	(WasmFilter_VmType)(0),     // 0: wasm.options.gloo.solo.io.WasmFilter.VmType
	(FilterStage_Stage)(0),     // 1: wasm.options.gloo.solo.io.FilterStage.Stage
	(FilterStage_Predicate)(0), // 2: wasm.options.gloo.solo.io.FilterStage.Predicate
	(*PluginSource)(nil),       // 3: wasm.options.gloo.solo.io.PluginSource
	(*WasmFilter)(nil),         // 4: wasm.options.gloo.solo.io.WasmFilter
	(*FilterStage)(nil),        // 5: wasm.options.gloo.solo.io.FilterStage
	(*any.Any)(nil),            // 6: google.protobuf.Any
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_depIdxs = []int32{
	4, // 0: wasm.options.gloo.solo.io.PluginSource.filters:type_name -> wasm.options.gloo.solo.io.WasmFilter
	6, // 1: wasm.options.gloo.solo.io.WasmFilter.config:type_name -> google.protobuf.Any
	5, // 2: wasm.options.gloo.solo.io.WasmFilter.filter_stage:type_name -> wasm.options.gloo.solo.io.FilterStage
	0, // 3: wasm.options.gloo.solo.io.WasmFilter.vm_type:type_name -> wasm.options.gloo.solo.io.WasmFilter.VmType
	1, // 4: wasm.options.gloo.solo.io.FilterStage.stage:type_name -> wasm.options.gloo.solo.io.FilterStage.Stage
	2, // 5: wasm.options.gloo.solo.io.FilterStage.predicate:type_name -> wasm.options.gloo.solo.io.FilterStage.Predicate
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginSource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WasmFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterStage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*WasmFilter_Image)(nil),
		(*WasmFilter_FilePath)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_wasm_wasm_proto_depIdxs = nil
}
