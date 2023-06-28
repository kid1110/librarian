// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: conf/miner.proto

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

type Miner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Miner_Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Miner_Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Miner) Reset() {
	*x = Miner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_miner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Miner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Miner) ProtoMessage() {}

func (x *Miner) ProtoReflect() protoreflect.Message {
	mi := &file_conf_miner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Miner.ProtoReflect.Descriptor instead.
func (*Miner) Descriptor() ([]byte, []int) {
	return file_conf_miner_proto_rawDescGZIP(), []int{0}
}

func (x *Miner) GetServer() *Miner_Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Miner) GetData() *Miner_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Miner_Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grpc *GRPC `protobuf:"bytes,1,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Miner_Server) Reset() {
	*x = Miner_Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_miner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Miner_Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Miner_Server) ProtoMessage() {}

func (x *Miner_Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_miner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Miner_Server.ProtoReflect.Descriptor instead.
func (*Miner_Server) Descriptor() ([]byte, []int) {
	return file_conf_miner_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Miner_Server) GetGrpc() *GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Miner_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ocr *Miner_Data_OCR `protobuf:"bytes,1,opt,name=ocr,proto3" json:"ocr,omitempty"`
}

func (x *Miner_Data) Reset() {
	*x = Miner_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_miner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Miner_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Miner_Data) ProtoMessage() {}

func (x *Miner_Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_miner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Miner_Data.ProtoReflect.Descriptor instead.
func (*Miner_Data) Descriptor() ([]byte, []int) {
	return file_conf_miner_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Miner_Data) GetOcr() *Miner_Data_OCR {
	if x != nil {
		return x.Ocr
	}
	return nil
}

type Miner_Data_OCR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Miner_Data_OCR) Reset() {
	*x = Miner_Data_OCR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_miner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Miner_Data_OCR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Miner_Data_OCR) ProtoMessage() {}

func (x *Miner_Data_OCR) ProtoReflect() protoreflect.Message {
	mi := &file_conf_miner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Miner_Data_OCR.ProtoReflect.Descriptor instead.
func (*Miner_Data_OCR) Descriptor() ([]byte, []int) {
	return file_conf_miner_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Miner_Data_OCR) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_conf_miner_proto protoreflect.FileDescriptor

var file_conf_miner_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x0f,
	0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xec, 0x01, 0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2e, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x52, 0x50,
	0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x55, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2c, 0x0a, 0x03, 0x6f, 0x63, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x43, 0x52, 0x52, 0x03, 0x6f, 0x63, 0x72, 0x1a, 0x1f, 0x0a,
	0x03, 0x4f, 0x43, 0x52, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x1e,
	0x5a, 0x1c, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_miner_proto_rawDescOnce sync.Once
	file_conf_miner_proto_rawDescData = file_conf_miner_proto_rawDesc
)

func file_conf_miner_proto_rawDescGZIP() []byte {
	file_conf_miner_proto_rawDescOnce.Do(func() {
		file_conf_miner_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_miner_proto_rawDescData)
	})
	return file_conf_miner_proto_rawDescData
}

var file_conf_miner_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conf_miner_proto_goTypes = []interface{}{
	(*Miner)(nil),          // 0: kratos.api.Miner
	(*Miner_Server)(nil),   // 1: kratos.api.Miner.Server
	(*Miner_Data)(nil),     // 2: kratos.api.Miner.Data
	(*Miner_Data_OCR)(nil), // 3: kratos.api.Miner.Data.OCR
	(*GRPC)(nil),           // 4: kratos.api.GRPC
}
var file_conf_miner_proto_depIdxs = []int32{
	1, // 0: kratos.api.Miner.server:type_name -> kratos.api.Miner.Server
	2, // 1: kratos.api.Miner.data:type_name -> kratos.api.Miner.Data
	4, // 2: kratos.api.Miner.Server.grpc:type_name -> kratos.api.GRPC
	3, // 3: kratos.api.Miner.Data.ocr:type_name -> kratos.api.Miner.Data.OCR
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_conf_miner_proto_init() }
func file_conf_miner_proto_init() {
	if File_conf_miner_proto != nil {
		return
	}
	file_conf_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conf_miner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Miner); i {
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
		file_conf_miner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Miner_Server); i {
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
		file_conf_miner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Miner_Data); i {
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
		file_conf_miner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Miner_Data_OCR); i {
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
			RawDescriptor: file_conf_miner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_miner_proto_goTypes,
		DependencyIndexes: file_conf_miner_proto_depIdxs,
		MessageInfos:      file_conf_miner_proto_msgTypes,
	}.Build()
	File_conf_miner_proto = out.File
	file_conf_miner_proto_rawDesc = nil
	file_conf_miner_proto_goTypes = nil
	file_conf_miner_proto_depIdxs = nil
}
