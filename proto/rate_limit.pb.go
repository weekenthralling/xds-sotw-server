// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/rate_limit.proto

package proto

import (
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

type ListRateLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRateLimitRequest) Reset() {
	*x = ListRateLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rate_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRateLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRateLimitRequest) ProtoMessage() {}

func (x *ListRateLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rate_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRateLimitRequest.ProtoReflect.Descriptor instead.
func (*ListRateLimitRequest) Descriptor() ([]byte, []int) {
	return file_proto_rate_limit_proto_rawDescGZIP(), []int{0}
}

type ListRateLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateLimit []*RateLimit `protobuf:"bytes,1,rep,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *ListRateLimitResponse) Reset() {
	*x = ListRateLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rate_limit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRateLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRateLimitResponse) ProtoMessage() {}

func (x *ListRateLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rate_limit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRateLimitResponse.ProtoReflect.Descriptor instead.
func (*ListRateLimitResponse) Descriptor() ([]byte, []int) {
	return file_proto_rate_limit_proto_rawDescGZIP(), []int{1}
}

func (x *ListRateLimitResponse) GetRateLimit() []*RateLimit {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

type RateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain      string                 `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Descriptor_ []*RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptor,proto3" json:"descriptor,omitempty"`
}

func (x *RateLimit) Reset() {
	*x = RateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rate_limit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimit) ProtoMessage() {}

func (x *RateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rate_limit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimit.ProtoReflect.Descriptor instead.
func (*RateLimit) Descriptor() ([]byte, []int) {
	return file_proto_rate_limit_proto_rawDescGZIP(), []int{2}
}

func (x *RateLimit) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RateLimit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RateLimit) GetDescriptor_() []*RateLimitDescriptor {
	if x != nil {
		return x.Descriptor_
	}
	return nil
}

type RateLimitDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  *string          `protobuf:"bytes,2,opt,name=value,proto3,oneof" json:"value,omitempty"`
	Policy *RateLimitPolicy `protobuf:"bytes,3,opt,name=policy,proto3,oneof" json:"policy,omitempty"`
}

func (x *RateLimitDescriptor) Reset() {
	*x = RateLimitDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rate_limit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitDescriptor) ProtoMessage() {}

func (x *RateLimitDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rate_limit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitDescriptor.ProtoReflect.Descriptor instead.
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return file_proto_rate_limit_proto_rawDescGZIP(), []int{3}
}

func (x *RateLimitDescriptor) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RateLimitDescriptor) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *RateLimitDescriptor) GetPolicy() *RateLimitPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type RateLimitPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit           *int32 `protobuf:"varint,1,opt,name=unit,proto3,oneof" json:"unit,omitempty"`
	RequestPerUnit *int32 `protobuf:"varint,2,opt,name=request_per_unit,json=requestPerUnit,proto3,oneof" json:"request_per_unit,omitempty"`
}

func (x *RateLimitPolicy) Reset() {
	*x = RateLimitPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rate_limit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitPolicy) ProtoMessage() {}

func (x *RateLimitPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rate_limit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitPolicy.ProtoReflect.Descriptor instead.
func (*RateLimitPolicy) Descriptor() ([]byte, []int) {
	return file_proto_rate_limit_proto_rawDescGZIP(), []int{4}
}

func (x *RateLimitPolicy) GetUnit() int32 {
	if x != nil && x.Unit != nil {
		return *x.Unit
	}
	return 0
}

func (x *RateLimitPolicy) GetRequestPerUnit() int32 {
	if x != nil && x.RequestPerUnit != nil {
		return *x.RequestPerUnit
	}
	return 0
}

type SaveDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string           `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Key    string           `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value  *string          `protobuf:"bytes,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
	Policy *RateLimitPolicy `protobuf:"bytes,4,opt,name=policy,proto3,oneof" json:"policy,omitempty"`
}

func (x *SaveDescriptorRequest) Reset() {
	*x = SaveDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rate_limit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDescriptorRequest) ProtoMessage() {}

func (x *SaveDescriptorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rate_limit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDescriptorRequest.ProtoReflect.Descriptor instead.
func (*SaveDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_proto_rate_limit_proto_rawDescGZIP(), []int{5}
}

func (x *SaveDescriptorRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SaveDescriptorRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SaveDescriptorRequest) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *SaveDescriptorRequest) GetPolicy() *RateLimitPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type SaveDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SaveDescriptorResponse) Reset() {
	*x = SaveDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rate_limit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDescriptorResponse) ProtoMessage() {}

func (x *SaveDescriptorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rate_limit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDescriptorResponse.ProtoReflect.Descriptor instead.
func (*SaveDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_proto_rate_limit_proto_rawDescGZIP(), []int{6}
}

func (x *SaveDescriptorResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_rate_limit_proto protoreflect.FileDescriptor

var file_proto_rate_limit_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x74, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34,
	0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x77, 0x0a, 0x0f, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x17, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x01, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x42, 0x13, 0x0a,
	0x11, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x6e,
	0x69, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x01, 0x52, 0x06, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x32, 0x0a, 0x16,
	0x53, 0x61, 0x76, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xba, 0x01, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x17, 0x53, 0x61, 0x76, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rate_limit_proto_rawDescOnce sync.Once
	file_proto_rate_limit_proto_rawDescData = file_proto_rate_limit_proto_rawDesc
)

func file_proto_rate_limit_proto_rawDescGZIP() []byte {
	file_proto_rate_limit_proto_rawDescOnce.Do(func() {
		file_proto_rate_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rate_limit_proto_rawDescData)
	})
	return file_proto_rate_limit_proto_rawDescData
}

var file_proto_rate_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_rate_limit_proto_goTypes = []any{
	(*ListRateLimitRequest)(nil),   // 0: config.ListRateLimitRequest
	(*ListRateLimitResponse)(nil),  // 1: config.ListRateLimitResponse
	(*RateLimit)(nil),              // 2: config.RateLimit
	(*RateLimitDescriptor)(nil),    // 3: config.RateLimitDescriptor
	(*RateLimitPolicy)(nil),        // 4: config.RateLimitPolicy
	(*SaveDescriptorRequest)(nil),  // 5: config.SaveDescriptorRequest
	(*SaveDescriptorResponse)(nil), // 6: config.SaveDescriptorResponse
}
var file_proto_rate_limit_proto_depIdxs = []int32{
	2, // 0: config.ListRateLimitResponse.rate_limit:type_name -> config.RateLimit
	3, // 1: config.RateLimit.descriptor:type_name -> config.RateLimitDescriptor
	4, // 2: config.RateLimitDescriptor.policy:type_name -> config.RateLimitPolicy
	4, // 3: config.SaveDescriptorRequest.policy:type_name -> config.RateLimitPolicy
	0, // 4: config.RateLimitService.ListRateLimit:input_type -> config.ListRateLimitRequest
	5, // 5: config.RateLimitService.SaveRateLimitDescriptor:input_type -> config.SaveDescriptorRequest
	1, // 6: config.RateLimitService.ListRateLimit:output_type -> config.ListRateLimitResponse
	6, // 7: config.RateLimitService.SaveRateLimitDescriptor:output_type -> config.SaveDescriptorResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_rate_limit_proto_init() }
func file_proto_rate_limit_proto_init() {
	if File_proto_rate_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rate_limit_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListRateLimitRequest); i {
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
		file_proto_rate_limit_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListRateLimitResponse); i {
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
		file_proto_rate_limit_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RateLimit); i {
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
		file_proto_rate_limit_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RateLimitDescriptor); i {
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
		file_proto_rate_limit_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RateLimitPolicy); i {
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
		file_proto_rate_limit_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SaveDescriptorRequest); i {
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
		file_proto_rate_limit_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SaveDescriptorResponse); i {
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
	file_proto_rate_limit_proto_msgTypes[3].OneofWrappers = []any{}
	file_proto_rate_limit_proto_msgTypes[4].OneofWrappers = []any{}
	file_proto_rate_limit_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rate_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rate_limit_proto_goTypes,
		DependencyIndexes: file_proto_rate_limit_proto_depIdxs,
		MessageInfos:      file_proto_rate_limit_proto_msgTypes,
	}.Build()
	File_proto_rate_limit_proto = out.File
	file_proto_rate_limit_proto_rawDesc = nil
	file_proto_rate_limit_proto_goTypes = nil
	file_proto_rate_limit_proto_depIdxs = nil
}