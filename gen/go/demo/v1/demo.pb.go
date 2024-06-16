// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: demo/v1/demo.proto

package demov1

import (
	_ "github.com/xiexianbin/go-grpc-demo/gen/go/google/api"
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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nums []int64 `protobuf:"varint,1,rep,packed,name=nums,proto3" json:"nums,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetNums() []int64 {
	if x != nil {
		return x.Nums
	}
	return nil
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type DiffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nums []int64 `protobuf:"varint,1,rep,packed,name=nums,proto3" json:"nums,omitempty"`
}

func (x *DiffRequest) Reset() {
	*x = DiffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffRequest) ProtoMessage() {}

func (x *DiffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffRequest.ProtoReflect.Descriptor instead.
func (*DiffRequest) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{2}
}

func (x *DiffRequest) GetNums() []int64 {
	if x != nil {
		return x.Nums
	}
	return nil
}

type DiffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DiffResponse) Reset() {
	*x = DiffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffResponse) ProtoMessage() {}

func (x *DiffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffResponse.ProtoReflect.Descriptor instead.
func (*DiffResponse) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{3}
}

func (x *DiffResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type ReadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ReadFileRequest) Reset() {
	*x = ReadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileRequest) ProtoMessage() {}

func (x *ReadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileRequest.ProtoReflect.Descriptor instead.
func (*ReadFileRequest) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{4}
}

func (x *ReadFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ReadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ReadFileResponse) Reset() {
	*x = ReadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileResponse) ProtoMessage() {}

func (x *ReadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileResponse.ProtoReflect.Descriptor instead.
func (*ReadFileResponse) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{5}
}

func (x *ReadFileResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{6}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{7}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_demo_v1_demo_proto protoreflect.FileDescriptor

var file_demo_v1_demo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0a, 0x53,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x25, 0x0a,
	0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x44, 0x69, 0x66, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x25, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x32, 0x9a, 0x02, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x3a, 0x01, 0x2a,
	0x22, 0x04, 0x2f, 0x73, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x04, 0x44, 0x69, 0x66, 0x66, 0x12, 0x14,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x08, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x8f, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31,
	0x42, 0x09, 0x44, 0x65, 0x6d, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x65, 0x78, 0x69, 0x61,
	0x6e, 0x62, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31,
	0x3b, 0x64, 0x65, 0x6d, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x07,
	0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x44, 0x65, 0x6d, 0x6f, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x13, 0x44, 0x65, 0x6d, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x44, 0x65, 0x6d, 0x6f, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_v1_demo_proto_rawDescOnce sync.Once
	file_demo_v1_demo_proto_rawDescData = file_demo_v1_demo_proto_rawDesc
)

func file_demo_v1_demo_proto_rawDescGZIP() []byte {
	file_demo_v1_demo_proto_rawDescOnce.Do(func() {
		file_demo_v1_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_v1_demo_proto_rawDescData)
	})
	return file_demo_v1_demo_proto_rawDescData
}

var file_demo_v1_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_demo_v1_demo_proto_goTypes = []interface{}{
	(*SumRequest)(nil),       // 0: demo.v1.SumRequest
	(*SumResponse)(nil),      // 1: demo.v1.SumResponse
	(*DiffRequest)(nil),      // 2: demo.v1.DiffRequest
	(*DiffResponse)(nil),     // 3: demo.v1.DiffResponse
	(*ReadFileRequest)(nil),  // 4: demo.v1.ReadFileRequest
	(*ReadFileResponse)(nil), // 5: demo.v1.ReadFileResponse
	(*VersionRequest)(nil),   // 6: demo.v1.VersionRequest
	(*VersionResponse)(nil),  // 7: demo.v1.VersionResponse
}
var file_demo_v1_demo_proto_depIdxs = []int32{
	0, // 0: demo.v1.DemoService.Sum:input_type -> demo.v1.SumRequest
	2, // 1: demo.v1.DemoService.Diff:input_type -> demo.v1.DiffRequest
	4, // 2: demo.v1.DemoService.ReadFile:input_type -> demo.v1.ReadFileRequest
	6, // 3: demo.v1.DemoService.Version:input_type -> demo.v1.VersionRequest
	1, // 4: demo.v1.DemoService.Sum:output_type -> demo.v1.SumResponse
	3, // 5: demo.v1.DemoService.Diff:output_type -> demo.v1.DiffResponse
	5, // 6: demo.v1.DemoService.ReadFile:output_type -> demo.v1.ReadFileResponse
	7, // 7: demo.v1.DemoService.Version:output_type -> demo.v1.VersionResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_v1_demo_proto_init() }
func file_demo_v1_demo_proto_init() {
	if File_demo_v1_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_v1_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_demo_v1_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_demo_v1_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffRequest); i {
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
		file_demo_v1_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffResponse); i {
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
		file_demo_v1_demo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadFileRequest); i {
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
		file_demo_v1_demo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadFileResponse); i {
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
		file_demo_v1_demo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_demo_v1_demo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
			RawDescriptor: file_demo_v1_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_v1_demo_proto_goTypes,
		DependencyIndexes: file_demo_v1_demo_proto_depIdxs,
		MessageInfos:      file_demo_v1_demo_proto_msgTypes,
	}.Build()
	File_demo_v1_demo_proto = out.File
	file_demo_v1_demo_proto_rawDesc = nil
	file_demo_v1_demo_proto_goTypes = nil
	file_demo_v1_demo_proto_depIdxs = nil
}
