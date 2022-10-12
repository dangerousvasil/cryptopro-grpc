// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: grpc_service_sign.proto

package grpc_service

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

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage string `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Key     string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Flag    uint64 `protobuf:"varint,4,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_sign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_sign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_sign_proto_rawDescGZIP(), []int{0}
}

func (x *SignRequest) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *SignRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SignRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SignRequest) GetFlag() uint64 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content     []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Code        uint64 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_sign_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_sign_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_sign_proto_rawDescGZIP(), []int{1}
}

func (x *SignResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SignResponse) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_grpc_service_sign_proto protoreflect.FileDescriptor

var file_grpc_service_sign_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x70, 0x72, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x67, 0x0a, 0x0b, 0x53, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x22, 0x5e, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x4e, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1b, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x70, 0x72,
	0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0x5a, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1b, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_service_sign_proto_rawDescOnce sync.Once
	file_grpc_service_sign_proto_rawDescData = file_grpc_service_sign_proto_rawDesc
)

func file_grpc_service_sign_proto_rawDescGZIP() []byte {
	file_grpc_service_sign_proto_rawDescOnce.Do(func() {
		file_grpc_service_sign_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_service_sign_proto_rawDescData)
	})
	return file_grpc_service_sign_proto_rawDescData
}

var file_grpc_service_sign_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_service_sign_proto_goTypes = []interface{}{
	(*SignRequest)(nil),  // 0: cryptopro_grpc.SignRequest
	(*SignResponse)(nil), // 1: cryptopro_grpc.SignResponse
}
var file_grpc_service_sign_proto_depIdxs = []int32{
	0, // 0: cryptopro_grpc.Service.Sign:input_type -> cryptopro_grpc.SignRequest
	0, // 1: cryptopro_grpc.ServiceInternal.Sign:input_type -> cryptopro_grpc.SignRequest
	1, // 2: cryptopro_grpc.Service.Sign:output_type -> cryptopro_grpc.SignResponse
	1, // 3: cryptopro_grpc.ServiceInternal.Sign:output_type -> cryptopro_grpc.SignResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_service_sign_proto_init() }
func file_grpc_service_sign_proto_init() {
	if File_grpc_service_sign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_service_sign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_grpc_service_sign_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
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
			RawDescriptor: file_grpc_service_sign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpc_service_sign_proto_goTypes,
		DependencyIndexes: file_grpc_service_sign_proto_depIdxs,
		MessageInfos:      file_grpc_service_sign_proto_msgTypes,
	}.Build()
	File_grpc_service_sign_proto = out.File
	file_grpc_service_sign_proto_rawDesc = nil
	file_grpc_service_sign_proto_goTypes = nil
	file_grpc_service_sign_proto_depIdxs = nil
}