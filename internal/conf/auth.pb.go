// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: conf/auth.proto

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

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Salt      string `protobuf:"bytes,1,opt,name=salt,proto3" json:"salt,omitempty"`
	Issuer    string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	JwtSecret string `protobuf:"bytes,3,opt,name=jwt_secret,json=jwtSecret,proto3" json:"jwt_secret,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_conf_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_conf_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *Auth) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Auth) GetJwtSecret() string {
	if x != nil {
		return x.JwtSecret
	}
	return ""
}

var File_conf_auth_proto protoreflect.FileDescriptor

var file_conf_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x51, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x77, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x77, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x42, 0x1e, 0x5a, 0x1c, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_auth_proto_rawDescOnce sync.Once
	file_conf_auth_proto_rawDescData = file_conf_auth_proto_rawDesc
)

func file_conf_auth_proto_rawDescGZIP() []byte {
	file_conf_auth_proto_rawDescOnce.Do(func() {
		file_conf_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_auth_proto_rawDescData)
	})
	return file_conf_auth_proto_rawDescData
}

var file_conf_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_conf_auth_proto_goTypes = []interface{}{
	(*Auth)(nil), // 0: kratos.api.Auth
}
var file_conf_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_conf_auth_proto_init() }
func file_conf_auth_proto_init() {
	if File_conf_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
			RawDescriptor: file_conf_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_auth_proto_goTypes,
		DependencyIndexes: file_conf_auth_proto_depIdxs,
		MessageInfos:      file_conf_auth_proto_msgTypes,
	}.Build()
	File_conf_auth_proto = out.File
	file_conf_auth_proto_rawDesc = nil
	file_conf_auth_proto_goTypes = nil
	file_conf_auth_proto_depIdxs = nil
}
