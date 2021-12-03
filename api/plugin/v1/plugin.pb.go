//
//Copyright 2021 The tKeel Authors.
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/plugin/v1/plugin.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v1 "github.com/tkeel-io/tkeel-interface/openapi/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//*
// Register Addons.
type RegisterAddons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addons   string `protobuf:"bytes,1,opt,name=addons,proto3" json:"addons,omitempty"`     // addons name.
	Upstream string `protobuf:"bytes,2,opt,name=upstream,proto3" json:"upstream,omitempty"` // upstream path.
}

func (x *RegisterAddons) Reset() {
	*x = RegisterAddons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAddons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAddons) ProtoMessage() {}

func (x *RegisterAddons) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAddons.ProtoReflect.Descriptor instead.
func (*RegisterAddons) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterAddons) GetAddons() string {
	if x != nil {
		return x.Addons
	}
	return ""
}

func (x *RegisterAddons) GetUpstream() string {
	if x != nil {
		return x.Upstream
	}
	return ""
}

//*
// plugin object.
type PluginObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                         // plugin id.
	PluginVersion     string                  `protobuf:"bytes,2,opt,name=plugin_version,json=pluginVersion,proto3" json:"plugin_version,omitempty"`              // plugin version.
	TkeelVersion      string                  `protobuf:"bytes,3,opt,name=tkeel_version,json=tkeelVersion,proto3" json:"tkeel_version,omitempty"`                 // plugin depend tkeel version.
	AddonsPoint       []*v1.AddonsPoint       `protobuf:"bytes,4,rep,name=addons_point,json=addonsPoint,proto3" json:"addons_point,omitempty"`                    // plugin declares addons.
	ImplementedPlugin []*v1.ImplementedPlugin `protobuf:"bytes,5,rep,name=implemented_plugin,json=implementedPlugin,proto3" json:"implemented_plugin,omitempty"`  // plugin implemented plugin list.
	Secret            string                  `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`                                                 // plugin registered secret.
	RegisterTimestamp int64                   `protobuf:"varint,7,opt,name=register_timestamp,json=registerTimestamp,proto3" json:"register_timestamp,omitempty"` // register timestamp.
	ActiveTenantes    []string                `protobuf:"bytes,8,rep,name=active_tenantes,json=activeTenantes,proto3" json:"active_tenantes,omitempty"`           // active tenant's id list.
	RegisterAddons    []*RegisterAddons       `protobuf:"bytes,9,rep,name=register_addons,json=registerAddons,proto3" json:"register_addons,omitempty"`           // register addons router.
	Status            v1.PluginStatus         `protobuf:"varint,10,opt,name=status,proto3,enum=openapi.v1.PluginStatus" json:"status,omitempty"`                  // register plugin status.
}

func (x *PluginObject) Reset() {
	*x = PluginObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginObject) ProtoMessage() {}

func (x *PluginObject) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginObject.ProtoReflect.Descriptor instead.
func (*PluginObject) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PluginObject) GetPluginVersion() string {
	if x != nil {
		return x.PluginVersion
	}
	return ""
}

func (x *PluginObject) GetTkeelVersion() string {
	if x != nil {
		return x.TkeelVersion
	}
	return ""
}

func (x *PluginObject) GetAddonsPoint() []*v1.AddonsPoint {
	if x != nil {
		return x.AddonsPoint
	}
	return nil
}

func (x *PluginObject) GetImplementedPlugin() []*v1.ImplementedPlugin {
	if x != nil {
		return x.ImplementedPlugin
	}
	return nil
}

func (x *PluginObject) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *PluginObject) GetRegisterTimestamp() int64 {
	if x != nil {
		return x.RegisterTimestamp
	}
	return 0
}

func (x *PluginObject) GetActiveTenantes() []string {
	if x != nil {
		return x.ActiveTenantes
	}
	return nil
}

func (x *PluginObject) GetRegisterAddons() []*RegisterAddons {
	if x != nil {
		return x.RegisterAddons
	}
	return nil
}

func (x *PluginObject) GetStatus() v1.PluginStatus {
	if x != nil {
		return x.Status
	}
	return v1.PluginStatus(0)
}

type RegisterPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *RegisterPluginRequest) Reset() {
	*x = RegisterPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPluginRequest) ProtoMessage() {}

func (x *RegisterPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPluginRequest.ProtoReflect.Descriptor instead.
func (*RegisterPluginRequest) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterPluginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterPluginRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type DeletePluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePluginRequest) Reset() {
	*x = DeletePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePluginRequest) ProtoMessage() {}

func (x *DeletePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePluginRequest.ProtoReflect.Descriptor instead.
func (*DeletePluginRequest) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePluginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin *PluginObject `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *DeletePluginResponse) Reset() {
	*x = DeletePluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePluginResponse) ProtoMessage() {}

func (x *DeletePluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePluginResponse.ProtoReflect.Descriptor instead.
func (*DeletePluginResponse) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePluginResponse) GetPlugin() *PluginObject {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type GetPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPluginRequest) Reset() {
	*x = GetPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginRequest) ProtoMessage() {}

func (x *GetPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginRequest.ProtoReflect.Descriptor instead.
func (*GetPluginRequest) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *GetPluginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin *PluginObject `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *GetPluginResponse) Reset() {
	*x = GetPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginResponse) ProtoMessage() {}

func (x *GetPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginResponse.ProtoReflect.Descriptor instead.
func (*GetPluginResponse) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *GetPluginResponse) GetPlugin() *PluginObject {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type ListPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginList []*PluginObject `protobuf:"bytes,1,rep,name=plugin_list,json=pluginList,proto3" json:"plugin_list,omitempty"`
}

func (x *ListPluginResponse) Reset() {
	*x = ListPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_plugin_v1_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginResponse) ProtoMessage() {}

func (x *ListPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_plugin_v1_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginResponse.ProtoReflect.Descriptor instead.
func (*ListPluginResponse) Descriptor() ([]byte, []int) {
	return file_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *ListPluginResponse) GetPluginList() []*PluginObject {
	if x != nil {
		return x.PluginList
	}
	return nil
}

var File_api_plugin_v1_plugin_proto protoreflect.FileDescriptor

var file_api_plugin_v1_plugin_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x44, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x6f,
	0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0xf7, 0x05, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x20, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x13, 0x92, 0x41, 0x10, 0x32, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0d, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0x92, 0x41, 0x1d, 0x32,
	0x1b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x20, 0x74,
	0x6b, 0x65, 0x65, 0x6c, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x6b,
	0x65, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x0c, 0x61, 0x64,
	0x64, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x6f, 0x6e, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x32, 0x16,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x73, 0x20,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x71, 0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x64, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x42, 0x23,
	0x92, 0x41, 0x20, 0x32, 0x1e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x11, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0x92, 0x41, 0x1a, 0x32, 0x18, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x46, 0x0a,
	0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x12,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x45, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x1c,
	0x92, 0x41, 0x19, 0x32, 0x17, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x20, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x27, 0x73, 0x20, 0x69, 0x64, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x12, 0x63, 0x0a, 0x0f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x6f, 0x6e, 0x73, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x32, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x20, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x20, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x6f, 0x6e,
	0x73, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1b, 0x92, 0x41, 0x18,
	0x32, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x20, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x6e, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x20, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0x92, 0x41, 0x1a, 0x32, 0x18,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x20, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x22, 0x35, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x20, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x12, 0x92, 0x41,
	0x0f, 0x32, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x32, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x12,
	0x92, 0x41, 0x0f, 0x32, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x6b, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x20,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xb2, 0x08, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x12, 0xc2, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0xf1, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0xda, 0x01, 0x0a, 0x06, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0xe6, 0xb3, 0xa8, 0xe5, 0x86, 0x8c, 0xe6, 0x8f, 0x92,
	0xe4, 0xbb, 0xb6, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x2a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4a, 0x1c, 0x0a, 0x03, 0x32, 0x30, 0x34,
	0x12, 0x15, 0x0a, 0x13, 0x53, 0x55, 0x43, 0x43, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4e, 0x4f, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x4a, 0x19, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x12,
	0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45,
	0x4e, 0x54, 0x4a, 0x19, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x12, 0x0a, 0x10, 0x50, 0x4c, 0x55,
	0x47, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x4a, 0x17, 0x0a,
	0x03, 0x34, 0x30, 0x39, 0x12, 0x10, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x4a, 0x26, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x1f, 0x0a,
	0x1d, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f,
	0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x41, 0x50, 0x49, 0x4a, 0x17,
	0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x10, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x12, 0xa3, 0x02, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xc9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0xb0, 0x01, 0x0a, 0x06, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe6, 0x8f, 0x92,
	0xe4, 0xbb, 0xb6, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x2a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04,
	0x0a, 0x02, 0x4f, 0x4b, 0x4a, 0x19, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x12, 0x0a, 0x10, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x4a,
	0x19, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x12, 0x0a, 0x10, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x4a, 0x28, 0x0a, 0x03, 0x35, 0x30,
	0x30, 0x12, 0x21, 0x0a, 0x1f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x50, 0x4c, 0x55, 0x47,
	0x49, 0x4e, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x42, 0x45, 0x45, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x45,
	0x4e, 0x44, 0x45, 0x44, 0x4a, 0x17, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x10, 0x0a, 0x0e, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0xed, 0x01,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9c,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x83, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x12, 0x12, 0xe6, 0x9f, 0xa5, 0xe8, 0xaf, 0xa2, 0xe6, 0x8f, 0x92, 0xe4, 0xbb, 0xb6,
	0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x2a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x4a, 0x19,
	0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x12, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x4a, 0x19, 0x0a, 0x03, 0x34, 0x30, 0x34,
	0x12, 0x12, 0x0a, 0x10, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x4a, 0x17, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x10, 0x0a, 0x0e, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0xcc, 0x01,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x82, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x92, 0x41, 0x6f, 0x0a, 0x06, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x18, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe6, 0x8f, 0x92,
	0xe4, 0xbb, 0xb6, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x2a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4a, 0x0b, 0x0a, 0x03, 0x32,
	0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x4a, 0x19, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12,
	0x12, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x4a, 0x17, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x10, 0x0a, 0x0e, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x42, 0x3d, 0x0a, 0x0d,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6b, 0x65, 0x65,
	0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_plugin_v1_plugin_proto_rawDescOnce sync.Once
	file_api_plugin_v1_plugin_proto_rawDescData = file_api_plugin_v1_plugin_proto_rawDesc
)

func file_api_plugin_v1_plugin_proto_rawDescGZIP() []byte {
	file_api_plugin_v1_plugin_proto_rawDescOnce.Do(func() {
		file_api_plugin_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_plugin_v1_plugin_proto_rawDescData)
	})
	return file_api_plugin_v1_plugin_proto_rawDescData
}

var file_api_plugin_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_plugin_v1_plugin_proto_goTypes = []interface{}{
	(*RegisterAddons)(nil),        // 0: api.plugin.v1.RegisterAddons
	(*PluginObject)(nil),          // 1: api.plugin.v1.PluginObject
	(*RegisterPluginRequest)(nil), // 2: api.plugin.v1.RegisterPluginRequest
	(*DeletePluginRequest)(nil),   // 3: api.plugin.v1.DeletePluginRequest
	(*DeletePluginResponse)(nil),  // 4: api.plugin.v1.DeletePluginResponse
	(*GetPluginRequest)(nil),      // 5: api.plugin.v1.GetPluginRequest
	(*GetPluginResponse)(nil),     // 6: api.plugin.v1.GetPluginResponse
	(*ListPluginResponse)(nil),    // 7: api.plugin.v1.ListPluginResponse
	(*v1.AddonsPoint)(nil),        // 8: openapi.v1.AddonsPoint
	(*v1.ImplementedPlugin)(nil),  // 9: openapi.v1.ImplementedPlugin
	(v1.PluginStatus)(0),          // 10: openapi.v1.PluginStatus
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_api_plugin_v1_plugin_proto_depIdxs = []int32{
	8,  // 0: api.plugin.v1.PluginObject.addons_point:type_name -> openapi.v1.AddonsPoint
	9,  // 1: api.plugin.v1.PluginObject.implemented_plugin:type_name -> openapi.v1.ImplementedPlugin
	0,  // 2: api.plugin.v1.PluginObject.register_addons:type_name -> api.plugin.v1.RegisterAddons
	10, // 3: api.plugin.v1.PluginObject.status:type_name -> openapi.v1.PluginStatus
	1,  // 4: api.plugin.v1.DeletePluginResponse.plugin:type_name -> api.plugin.v1.PluginObject
	1,  // 5: api.plugin.v1.GetPluginResponse.plugin:type_name -> api.plugin.v1.PluginObject
	1,  // 6: api.plugin.v1.ListPluginResponse.plugin_list:type_name -> api.plugin.v1.PluginObject
	2,  // 7: api.plugin.v1.Plugin.RegisterPlugin:input_type -> api.plugin.v1.RegisterPluginRequest
	3,  // 8: api.plugin.v1.Plugin.DeletePlugin:input_type -> api.plugin.v1.DeletePluginRequest
	5,  // 9: api.plugin.v1.Plugin.GetPlugin:input_type -> api.plugin.v1.GetPluginRequest
	11, // 10: api.plugin.v1.Plugin.ListPlugin:input_type -> google.protobuf.Empty
	11, // 11: api.plugin.v1.Plugin.RegisterPlugin:output_type -> google.protobuf.Empty
	4,  // 12: api.plugin.v1.Plugin.DeletePlugin:output_type -> api.plugin.v1.DeletePluginResponse
	6,  // 13: api.plugin.v1.Plugin.GetPlugin:output_type -> api.plugin.v1.GetPluginResponse
	7,  // 14: api.plugin.v1.Plugin.ListPlugin:output_type -> api.plugin.v1.ListPluginResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_plugin_v1_plugin_proto_init() }
func file_api_plugin_v1_plugin_proto_init() {
	if File_api_plugin_v1_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_plugin_v1_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAddons); i {
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
		file_api_plugin_v1_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginObject); i {
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
		file_api_plugin_v1_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPluginRequest); i {
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
		file_api_plugin_v1_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePluginRequest); i {
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
		file_api_plugin_v1_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePluginResponse); i {
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
		file_api_plugin_v1_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginRequest); i {
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
		file_api_plugin_v1_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginResponse); i {
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
		file_api_plugin_v1_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_plugin_v1_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_plugin_v1_plugin_proto_goTypes,
		DependencyIndexes: file_api_plugin_v1_plugin_proto_depIdxs,
		MessageInfos:      file_api_plugin_v1_plugin_proto_msgTypes,
	}.Build()
	File_api_plugin_v1_plugin_proto = out.File
	file_api_plugin_v1_plugin_proto_rawDesc = nil
	file_api_plugin_v1_plugin_proto_goTypes = nil
	file_api_plugin_v1_plugin_proto_depIdxs = nil
}
