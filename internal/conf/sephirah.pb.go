// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: conf/sephirah.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sephirah struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Sephirah_Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Sephirah_Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Auth   *Auth            `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Sephirah) Reset() {
	*x = Sephirah{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_sephirah_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sephirah) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sephirah) ProtoMessage() {}

func (x *Sephirah) ProtoReflect() protoreflect.Message {
	mi := &file_conf_sephirah_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sephirah.ProtoReflect.Descriptor instead.
func (*Sephirah) Descriptor() ([]byte, []int) {
	return file_conf_sephirah_proto_rawDescGZIP(), []int{0}
}

func (x *Sephirah) GetServer() *Sephirah_Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Sephirah) GetData() *Sephirah_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Sephirah) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type Sephirah_Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grpc    *Sephirah_Server_GRPC `protobuf:"bytes,1,opt,name=grpc,proto3" json:"grpc,omitempty"`
	GrpcWeb *Sephirah_Server_GRPC `protobuf:"bytes,2,opt,name=grpc_web,json=grpcWeb,proto3" json:"grpc_web,omitempty"`
}

func (x *Sephirah_Server) Reset() {
	*x = Sephirah_Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_sephirah_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sephirah_Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sephirah_Server) ProtoMessage() {}

func (x *Sephirah_Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_sephirah_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sephirah_Server.ProtoReflect.Descriptor instead.
func (*Sephirah_Server) Descriptor() ([]byte, []int) {
	return file_conf_sephirah_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Sephirah_Server) GetGrpc() *Sephirah_Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Sephirah_Server) GetGrpcWeb() *Sephirah_Server_GRPC {
	if x != nil {
		return x.GrpcWeb
	}
	return nil
}

type Sephirah_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Sephirah_Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *Sephirah_Data) Reset() {
	*x = Sephirah_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_sephirah_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sephirah_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sephirah_Data) ProtoMessage() {}

func (x *Sephirah_Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_sephirah_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sephirah_Data.ProtoReflect.Descriptor instead.
func (*Sephirah_Data) Descriptor() ([]byte, []int) {
	return file_conf_sephirah_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Sephirah_Data) GetDatabase() *Sephirah_Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

type Sephirah_Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Sephirah_Server_GRPC) Reset() {
	*x = Sephirah_Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_sephirah_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sephirah_Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sephirah_Server_GRPC) ProtoMessage() {}

func (x *Sephirah_Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_sephirah_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sephirah_Server_GRPC.ProtoReflect.Descriptor instead.
func (*Sephirah_Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_sephirah_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Sephirah_Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Sephirah_Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Sephirah_Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Sephirah_Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
}

func (x *Sephirah_Data_Database) Reset() {
	*x = Sephirah_Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_sephirah_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sephirah_Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sephirah_Data_Database) ProtoMessage() {}

func (x *Sephirah_Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_sephirah_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sephirah_Data_Database.ProtoReflect.Descriptor instead.
func (*Sephirah_Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_sephirah_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Sephirah_Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

var File_conf_sephirah_proto protoreflect.FileDescriptor

var file_conf_sephirah_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x73, 0x65, 0x70, 0x68, 0x69, 0x72, 0x61, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe9, 0x03, 0x0a, 0x08, 0x53, 0x65, 0x70, 0x68, 0x69, 0x72, 0x61, 0x68, 0x12,
	0x33, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x70,
	0x68, 0x69, 0x72, 0x61, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x65, 0x70, 0x68, 0x69, 0x72, 0x61, 0x68, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0xe6, 0x01, 0x0a, 0x06, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x65, 0x70, 0x68, 0x69, 0x72, 0x61, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x3b, 0x0a, 0x08, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x77, 0x65, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x70, 0x68, 0x69, 0x72,
	0x61, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x07,
	0x67, 0x72, 0x70, 0x63, 0x57, 0x65, 0x62, 0x1a, 0x69, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x1a, 0x6a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x70, 0x68, 0x69, 0x72,
	0x61, 0x68, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x22, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x42, 0x1e,
	0x5a, 0x1c, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_sephirah_proto_rawDescOnce sync.Once
	file_conf_sephirah_proto_rawDescData = file_conf_sephirah_proto_rawDesc
)

func file_conf_sephirah_proto_rawDescGZIP() []byte {
	file_conf_sephirah_proto_rawDescOnce.Do(func() {
		file_conf_sephirah_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_sephirah_proto_rawDescData)
	})
	return file_conf_sephirah_proto_rawDescData
}

var file_conf_sephirah_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_conf_sephirah_proto_goTypes = []interface{}{
	(*Sephirah)(nil),               // 0: kratos.api.Sephirah
	(*Sephirah_Server)(nil),        // 1: kratos.api.Sephirah.Server
	(*Sephirah_Data)(nil),          // 2: kratos.api.Sephirah.Data
	(*Sephirah_Server_GRPC)(nil),   // 3: kratos.api.Sephirah.Server.GRPC
	(*Sephirah_Data_Database)(nil), // 4: kratos.api.Sephirah.Data.Database
	(*Auth)(nil),                   // 5: kratos.api.Auth
	(*durationpb.Duration)(nil),    // 6: google.protobuf.Duration
}
var file_conf_sephirah_proto_depIdxs = []int32{
	1, // 0: kratos.api.Sephirah.server:type_name -> kratos.api.Sephirah.Server
	2, // 1: kratos.api.Sephirah.data:type_name -> kratos.api.Sephirah.Data
	5, // 2: kratos.api.Sephirah.auth:type_name -> kratos.api.Auth
	3, // 3: kratos.api.Sephirah.Server.grpc:type_name -> kratos.api.Sephirah.Server.GRPC
	3, // 4: kratos.api.Sephirah.Server.grpc_web:type_name -> kratos.api.Sephirah.Server.GRPC
	4, // 5: kratos.api.Sephirah.Data.database:type_name -> kratos.api.Sephirah.Data.Database
	6, // 6: kratos.api.Sephirah.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_conf_sephirah_proto_init() }
func file_conf_sephirah_proto_init() {
	if File_conf_sephirah_proto != nil {
		return
	}
	file_conf_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conf_sephirah_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sephirah); i {
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
		file_conf_sephirah_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sephirah_Server); i {
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
		file_conf_sephirah_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sephirah_Data); i {
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
		file_conf_sephirah_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sephirah_Server_GRPC); i {
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
		file_conf_sephirah_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sephirah_Data_Database); i {
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
			RawDescriptor: file_conf_sephirah_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_sephirah_proto_goTypes,
		DependencyIndexes: file_conf_sephirah_proto_depIdxs,
		MessageInfos:      file_conf_sephirah_proto_msgTypes,
	}.Build()
	File_conf_sephirah_proto = out.File
	file_conf_sephirah_proto_rawDesc = nil
	file_conf_sephirah_proto_goTypes = nil
	file_conf_sephirah_proto_depIdxs = nil
}
