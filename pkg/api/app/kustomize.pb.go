// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/app/kustomize.proto

package apps

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Kustomization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path      string     `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	SourceRef *SourceRef `protobuf:"bytes,4,opt,name=sourceRef,proto3" json:"sourceRef,omitempty"`
	Interval  *Interval  `protobuf:"bytes,5,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *Kustomization) Reset() {
	*x = Kustomization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kustomization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kustomization) ProtoMessage() {}

func (x *Kustomization) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kustomization.ProtoReflect.Descriptor instead.
func (*Kustomization) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{0}
}

func (x *Kustomization) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Kustomization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Kustomization) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Kustomization) GetSourceRef() *SourceRef {
	if x != nil {
		return x.SourceRef
	}
	return nil
}

func (x *Kustomization) GetInterval() *Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

type AddKustomizationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	AppName   string     `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Name      string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Path      string     `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	SourceRef *SourceRef `protobuf:"bytes,5,opt,name=sourceRef,proto3" json:"sourceRef,omitempty"`
	Interval  *Interval  `protobuf:"bytes,6,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *AddKustomizationReq) Reset() {
	*x = AddKustomizationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKustomizationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKustomizationReq) ProtoMessage() {}

func (x *AddKustomizationReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKustomizationReq.ProtoReflect.Descriptor instead.
func (*AddKustomizationReq) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{1}
}

func (x *AddKustomizationReq) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AddKustomizationReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AddKustomizationReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddKustomizationReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AddKustomizationReq) GetSourceRef() *SourceRef {
	if x != nil {
		return x.SourceRef
	}
	return nil
}

func (x *AddKustomizationReq) GetInterval() *Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

type AddKustomizationRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success       bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Kustomization *Kustomization `protobuf:"bytes,2,opt,name=kustomization,proto3" json:"kustomization,omitempty"`
}

func (x *AddKustomizationRes) Reset() {
	*x = AddKustomizationRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKustomizationRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKustomizationRes) ProtoMessage() {}

func (x *AddKustomizationRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKustomizationRes.ProtoReflect.Descriptor instead.
func (*AddKustomizationRes) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{2}
}

func (x *AddKustomizationRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddKustomizationRes) GetKustomization() *Kustomization {
	if x != nil {
		return x.Kustomization
	}
	return nil
}

type ListKustomizationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	AppName   string `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

func (x *ListKustomizationsReq) Reset() {
	*x = ListKustomizationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKustomizationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKustomizationsReq) ProtoMessage() {}

func (x *ListKustomizationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKustomizationsReq.ProtoReflect.Descriptor instead.
func (*ListKustomizationsReq) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{3}
}

func (x *ListKustomizationsReq) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListKustomizationsReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type ListKustomizationsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kustomizations []*Kustomization `protobuf:"bytes,1,rep,name=kustomizations,proto3" json:"kustomizations,omitempty"`
}

func (x *ListKustomizationsRes) Reset() {
	*x = ListKustomizationsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKustomizationsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKustomizationsRes) ProtoMessage() {}

func (x *ListKustomizationsRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKustomizationsRes.ProtoReflect.Descriptor instead.
func (*ListKustomizationsRes) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{4}
}

func (x *ListKustomizationsRes) GetKustomizations() []*Kustomization {
	if x != nil {
		return x.Kustomizations
	}
	return nil
}

type RemoveKustomizationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace         string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	AppName           string `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	KustomizationName string `protobuf:"bytes,3,opt,name=kustomization_name,json=kustomizationName,proto3" json:"kustomization_name,omitempty"`
}

func (x *RemoveKustomizationReq) Reset() {
	*x = RemoveKustomizationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKustomizationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKustomizationReq) ProtoMessage() {}

func (x *RemoveKustomizationReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKustomizationReq.ProtoReflect.Descriptor instead.
func (*RemoveKustomizationReq) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveKustomizationReq) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RemoveKustomizationReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RemoveKustomizationReq) GetKustomizationName() string {
	if x != nil {
		return x.KustomizationName
	}
	return ""
}

type RemoveKustomizationRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveKustomizationRes) Reset() {
	*x = RemoveKustomizationRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKustomizationRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKustomizationRes) ProtoMessage() {}

func (x *RemoveKustomizationRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKustomizationRes.ProtoReflect.Descriptor instead.
func (*RemoveKustomizationRes) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveKustomizationRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_app_kustomize_proto protoreflect.FileDescriptor

var file_api_app_kustomize_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x69, 0x74, 0x6f, 0x70,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x0d, 0x4b,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x36, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0xe9, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4b, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x39,
	0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x09,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x36, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x69,
	0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x22, 0x76, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x45, 0x0a, 0x0d, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6b, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67,
	0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6b,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x80, 0x01,
	0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0x9c, 0x06, 0x0a, 0x04, 0x46, 0x6c, 0x75, 0x78, 0x12, 0x94, 0x01,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x7d, 0x2f, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x2e, 0x67, 0x69,
	0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x22, 0x2f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d,
	0x2f, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xb2,
	0x01, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x41, 0x2a, 0x3c, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x6b, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6b, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x3a, 0x01, 0x2a, 0x12, 0x94, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x47, 0x69, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x47,
	0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x25, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x27,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x67, 0x69, 0x74, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x26, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x67, 0x69, 0x74,
	0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x67, 0x69, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x42, 0xca, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x2d, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x92, 0x41, 0x94, 0x01, 0x12, 0x7c, 0x0a, 0x22,
	0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70,
	0x20, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41,
	0x50, 0x49, 0x12, 0x51, 0x54, 0x68, 0x65, 0x20, 0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69,
	0x74, 0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x20, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x57, 0x65, 0x61, 0x76,
	0x65, 0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x32, 0x09, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_app_kustomize_proto_rawDescOnce sync.Once
	file_api_app_kustomize_proto_rawDescData = file_api_app_kustomize_proto_rawDesc
)

func file_api_app_kustomize_proto_rawDescGZIP() []byte {
	file_api_app_kustomize_proto_rawDescOnce.Do(func() {
		file_api_app_kustomize_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_kustomize_proto_rawDescData)
	})
	return file_api_app_kustomize_proto_rawDescData
}

var file_api_app_kustomize_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_app_kustomize_proto_goTypes = []interface{}{
	(*Kustomization)(nil),          // 0: gitops_server.v1.Kustomization
	(*AddKustomizationReq)(nil),    // 1: gitops_server.v1.AddKustomizationReq
	(*AddKustomizationRes)(nil),    // 2: gitops_server.v1.AddKustomizationRes
	(*ListKustomizationsReq)(nil),  // 3: gitops_server.v1.ListKustomizationsReq
	(*ListKustomizationsRes)(nil),  // 4: gitops_server.v1.ListKustomizationsRes
	(*RemoveKustomizationReq)(nil), // 5: gitops_server.v1.RemoveKustomizationReq
	(*RemoveKustomizationRes)(nil), // 6: gitops_server.v1.RemoveKustomizationRes
	(*SourceRef)(nil),              // 7: gitops_server.v1.SourceRef
	(*Interval)(nil),               // 8: gitops_server.v1.Interval
	(*AddGitRepositoryReq)(nil),    // 9: gitops_server.v1.AddGitRepositoryReq
	(*ListGitRepositoryReq)(nil),   // 10: gitops_server.v1.ListGitRepositoryReq
	(*AddGitRepositoryRes)(nil),    // 11: gitops_server.v1.AddGitRepositoryRes
	(*ListGitRepositoryRes)(nil),   // 12: gitops_server.v1.ListGitRepositoryRes
}
var file_api_app_kustomize_proto_depIdxs = []int32{
	7,  // 0: gitops_server.v1.Kustomization.sourceRef:type_name -> gitops_server.v1.SourceRef
	8,  // 1: gitops_server.v1.Kustomization.interval:type_name -> gitops_server.v1.Interval
	7,  // 2: gitops_server.v1.AddKustomizationReq.sourceRef:type_name -> gitops_server.v1.SourceRef
	8,  // 3: gitops_server.v1.AddKustomizationReq.interval:type_name -> gitops_server.v1.Interval
	0,  // 4: gitops_server.v1.AddKustomizationRes.kustomization:type_name -> gitops_server.v1.Kustomization
	0,  // 5: gitops_server.v1.ListKustomizationsRes.kustomizations:type_name -> gitops_server.v1.Kustomization
	1,  // 6: gitops_server.v1.Flux.AddKustomization:input_type -> gitops_server.v1.AddKustomizationReq
	3,  // 7: gitops_server.v1.Flux.ListKustomizations:input_type -> gitops_server.v1.ListKustomizationsReq
	5,  // 8: gitops_server.v1.Flux.RemoveKustomization:input_type -> gitops_server.v1.RemoveKustomizationReq
	9,  // 9: gitops_server.v1.Flux.AddGitRepository:input_type -> gitops_server.v1.AddGitRepositoryReq
	10, // 10: gitops_server.v1.Flux.ListGitRepositories:input_type -> gitops_server.v1.ListGitRepositoryReq
	2,  // 11: gitops_server.v1.Flux.AddKustomization:output_type -> gitops_server.v1.AddKustomizationRes
	4,  // 12: gitops_server.v1.Flux.ListKustomizations:output_type -> gitops_server.v1.ListKustomizationsRes
	6,  // 13: gitops_server.v1.Flux.RemoveKustomization:output_type -> gitops_server.v1.RemoveKustomizationRes
	11, // 14: gitops_server.v1.Flux.AddGitRepository:output_type -> gitops_server.v1.AddGitRepositoryRes
	12, // 15: gitops_server.v1.Flux.ListGitRepositories:output_type -> gitops_server.v1.ListGitRepositoryRes
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_app_kustomize_proto_init() }
func file_api_app_kustomize_proto_init() {
	if File_api_app_kustomize_proto != nil {
		return
	}
	file_api_app_source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_app_kustomize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kustomization); i {
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
		file_api_app_kustomize_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKustomizationReq); i {
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
		file_api_app_kustomize_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKustomizationRes); i {
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
		file_api_app_kustomize_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKustomizationsReq); i {
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
		file_api_app_kustomize_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKustomizationsRes); i {
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
		file_api_app_kustomize_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKustomizationReq); i {
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
		file_api_app_kustomize_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKustomizationRes); i {
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
			RawDescriptor: file_api_app_kustomize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_app_kustomize_proto_goTypes,
		DependencyIndexes: file_api_app_kustomize_proto_depIdxs,
		MessageInfos:      file_api_app_kustomize_proto_msgTypes,
	}.Build()
	File_api_app_kustomize_proto = out.File
	file_api_app_kustomize_proto_rawDesc = nil
	file_api_app_kustomize_proto_goTypes = nil
	file_api_app_kustomize_proto_depIdxs = nil
}
