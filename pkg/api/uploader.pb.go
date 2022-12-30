// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/uploader.proto

package api

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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk    []byte `protobuf:"bytes,1,opt,name=Chunk,proto3" json:"Chunk,omitempty"`
	ChunkId  int32  `protobuf:"varint,2,opt,name=ChunkId,proto3" json:"ChunkId,omitempty"`
	ClientId string `protobuf:"bytes,3,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	Filename string `protobuf:"bytes,4,opt,name=Filename,proto3" json:"Filename,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uploader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uploader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_uploader_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *UploadRequest) GetChunkId() int32 {
	if x != nil {
		return x.ChunkId
	}
	return 0
}

func (x *UploadRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UploadRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uploader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uploader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_uploader_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UploadInitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    string `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	ModifiedAt  string `protobuf:"bytes,2,opt,name=ModifiedAt,proto3" json:"ModifiedAt,omitempty"`
	Size        int32  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	ChunksCount int32  `protobuf:"varint,4,opt,name=ChunksCount,proto3" json:"ChunksCount,omitempty"`
	ClientId    string `protobuf:"bytes,5,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *UploadInitRequest) Reset() {
	*x = UploadInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uploader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInitRequest) ProtoMessage() {}

func (x *UploadInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uploader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInitRequest.ProtoReflect.Descriptor instead.
func (*UploadInitRequest) Descriptor() ([]byte, []int) {
	return file_proto_uploader_proto_rawDescGZIP(), []int{2}
}

func (x *UploadInitRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadInitRequest) GetModifiedAt() string {
	if x != nil {
		return x.ModifiedAt
	}
	return ""
}

func (x *UploadInitRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UploadInitRequest) GetChunksCount() int32 {
	if x != nil {
		return x.ChunksCount
	}
	return 0
}

func (x *UploadInitRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type UploadInitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkStartFrom int32 `protobuf:"varint,1,opt,name=ChunkStartFrom,proto3" json:"ChunkStartFrom,omitempty"`
}

func (x *UploadInitResponse) Reset() {
	*x = UploadInitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uploader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInitResponse) ProtoMessage() {}

func (x *UploadInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uploader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInitResponse.ProtoReflect.Descriptor instead.
func (*UploadInitResponse) Descriptor() ([]byte, []int) {
	return file_proto_uploader_proto_rawDescGZIP(), []int{3}
}

func (x *UploadInitResponse) GetChunkStartFrom() int32 {
	if x != nil {
		return x.ChunkStartFrom
	}
	return 0
}

var File_proto_uploader_proto protoreflect.FileDescriptor

var file_proto_uploader_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x77, 0x0a, 0x0d, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa1,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x32, 0x82, 0x01, 0x0a, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x0a, 0x69, 0x6e, 0x69, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_uploader_proto_rawDescOnce sync.Once
	file_proto_uploader_proto_rawDescData = file_proto_uploader_proto_rawDesc
)

func file_proto_uploader_proto_rawDescGZIP() []byte {
	file_proto_uploader_proto_rawDescOnce.Do(func() {
		file_proto_uploader_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_uploader_proto_rawDescData)
	})
	return file_proto_uploader_proto_rawDescData
}

var file_proto_uploader_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_uploader_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),      // 0: api.UploadRequest
	(*UploadResponse)(nil),     // 1: api.UploadResponse
	(*UploadInitRequest)(nil),  // 2: api.UploadInitRequest
	(*UploadInitResponse)(nil), // 3: api.UploadInitResponse
}
var file_proto_uploader_proto_depIdxs = []int32{
	2, // 0: api.Uploader.initUpload:input_type -> api.UploadInitRequest
	0, // 1: api.Uploader.upload:input_type -> api.UploadRequest
	3, // 2: api.Uploader.initUpload:output_type -> api.UploadInitResponse
	1, // 3: api.Uploader.upload:output_type -> api.UploadResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_uploader_proto_init() }
func file_proto_uploader_proto_init() {
	if File_proto_uploader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_uploader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_proto_uploader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_proto_uploader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInitRequest); i {
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
		file_proto_uploader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInitResponse); i {
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
			RawDescriptor: file_proto_uploader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_uploader_proto_goTypes,
		DependencyIndexes: file_proto_uploader_proto_depIdxs,
		MessageInfos:      file_proto_uploader_proto_msgTypes,
	}.Build()
	File_proto_uploader_proto = out.File
	file_proto_uploader_proto_rawDesc = nil
	file_proto_uploader_proto_goTypes = nil
	file_proto_uploader_proto_depIdxs = nil
}
