// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: conf/searcher.proto

package conf

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

type Searcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Searcher_Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Searcher_Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Searcher) Reset() {
	*x = Searcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_searcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Searcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Searcher) ProtoMessage() {}

func (x *Searcher) ProtoReflect() protoreflect.Message {
	mi := &file_conf_searcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Searcher.ProtoReflect.Descriptor instead.
func (*Searcher) Descriptor() ([]byte, []int) {
	return file_conf_searcher_proto_rawDescGZIP(), []int{0}
}

func (x *Searcher) GetServer() *Searcher_Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Searcher) GetData() *Searcher_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Searcher_Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grpc *GRPC `protobuf:"bytes,1,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Searcher_Server) Reset() {
	*x = Searcher_Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_searcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Searcher_Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Searcher_Server) ProtoMessage() {}

func (x *Searcher_Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_searcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Searcher_Server.ProtoReflect.Descriptor instead.
func (*Searcher_Server) Descriptor() ([]byte, []int) {
	return file_conf_searcher_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Searcher_Server) GetGrpc() *GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Searcher_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Searcher_Data) Reset() {
	*x = Searcher_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_searcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Searcher_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Searcher_Data) ProtoMessage() {}

func (x *Searcher_Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_searcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Searcher_Data.ProtoReflect.Descriptor instead.
func (*Searcher_Data) Descriptor() ([]byte, []int) {
	return file_conf_searcher_proto_rawDescGZIP(), []int{0, 1}
}

var File_conf_searcher_proto protoreflect.FileDescriptor

var file_conf_searcher_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x2e, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x1a, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x1e, 0x5a, 0x1c, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_conf_searcher_proto_rawDescOnce sync.Once
	file_conf_searcher_proto_rawDescData = file_conf_searcher_proto_rawDesc
)

func file_conf_searcher_proto_rawDescGZIP() []byte {
	file_conf_searcher_proto_rawDescOnce.Do(func() {
		file_conf_searcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_searcher_proto_rawDescData)
	})
	return file_conf_searcher_proto_rawDescData
}

var file_conf_searcher_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_conf_searcher_proto_goTypes = []interface{}{
	(*Searcher)(nil),        // 0: kratos.api.Searcher
	(*Searcher_Server)(nil), // 1: kratos.api.Searcher.Server
	(*Searcher_Data)(nil),   // 2: kratos.api.Searcher.Data
	(*GRPC)(nil),            // 3: kratos.api.GRPC
}
var file_conf_searcher_proto_depIdxs = []int32{
	1, // 0: kratos.api.Searcher.server:type_name -> kratos.api.Searcher.Server
	2, // 1: kratos.api.Searcher.data:type_name -> kratos.api.Searcher.Data
	3, // 2: kratos.api.Searcher.Server.grpc:type_name -> kratos.api.GRPC
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conf_searcher_proto_init() }
func file_conf_searcher_proto_init() {
	if File_conf_searcher_proto != nil {
		return
	}
	file_conf_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conf_searcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Searcher); i {
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
		file_conf_searcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Searcher_Server); i {
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
		file_conf_searcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Searcher_Data); i {
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
			RawDescriptor: file_conf_searcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_searcher_proto_goTypes,
		DependencyIndexes: file_conf_searcher_proto_depIdxs,
		MessageInfos:      file_conf_searcher_proto_msgTypes,
	}.Build()
	File_conf_searcher_proto = out.File
	file_conf_searcher_proto_rawDesc = nil
	file_conf_searcher_proto_goTypes = nil
	file_conf_searcher_proto_depIdxs = nil
}
