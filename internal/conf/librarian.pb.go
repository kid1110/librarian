// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: conf/librarian.proto

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

type Librarian struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableServiceDiscovery *Librarian_EnableServiceDiscovery `protobuf:"bytes,10,opt,name=enable_service_discovery,json=enableServiceDiscovery,proto3" json:"enable_service_discovery,omitempty"`
	Server                 *SephirahServer                   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Database               *Database                         `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	S3                     *S3                               `protobuf:"bytes,3,opt,name=s3,proto3" json:"s3,omitempty"`
	Auth                   *Auth                             `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
	Mq                     *MQ                               `protobuf:"bytes,5,opt,name=mq,proto3" json:"mq,omitempty"`
	Cache                  *Cache                            `protobuf:"bytes,6,opt,name=cache,proto3" json:"cache,omitempty"`
	Porter                 *Porter                           `protobuf:"bytes,7,opt,name=porter,proto3" json:"porter,omitempty"`
	Consul                 *Consul                           `protobuf:"bytes,8,opt,name=consul,proto3" json:"consul,omitempty"`
	Sentry                 *Sentry                           `protobuf:"bytes,9,opt,name=sentry,proto3" json:"sentry,omitempty"`
	Otlp                   *OTLP                             `protobuf:"bytes,14,opt,name=otlp,proto3" json:"otlp,omitempty"`
	Mapper                 *Mapper                           `protobuf:"bytes,11,opt,name=mapper,proto3" json:"mapper,omitempty"`
	Searcher               *Searcher                         `protobuf:"bytes,12,opt,name=searcher,proto3" json:"searcher,omitempty"`
	Miner                  *Miner                            `protobuf:"bytes,13,opt,name=miner,proto3" json:"miner,omitempty"`
}

func (x *Librarian) Reset() {
	*x = Librarian{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_librarian_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Librarian) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Librarian) ProtoMessage() {}

func (x *Librarian) ProtoReflect() protoreflect.Message {
	mi := &file_conf_librarian_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Librarian.ProtoReflect.Descriptor instead.
func (*Librarian) Descriptor() ([]byte, []int) {
	return file_conf_librarian_proto_rawDescGZIP(), []int{0}
}

func (x *Librarian) GetEnableServiceDiscovery() *Librarian_EnableServiceDiscovery {
	if x != nil {
		return x.EnableServiceDiscovery
	}
	return nil
}

func (x *Librarian) GetServer() *SephirahServer {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Librarian) GetDatabase() *Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Librarian) GetS3() *S3 {
	if x != nil {
		return x.S3
	}
	return nil
}

func (x *Librarian) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Librarian) GetMq() *MQ {
	if x != nil {
		return x.Mq
	}
	return nil
}

func (x *Librarian) GetCache() *Cache {
	if x != nil {
		return x.Cache
	}
	return nil
}

func (x *Librarian) GetPorter() *Porter {
	if x != nil {
		return x.Porter
	}
	return nil
}

func (x *Librarian) GetConsul() *Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Librarian) GetSentry() *Sentry {
	if x != nil {
		return x.Sentry
	}
	return nil
}

func (x *Librarian) GetOtlp() *OTLP {
	if x != nil {
		return x.Otlp
	}
	return nil
}

func (x *Librarian) GetMapper() *Mapper {
	if x != nil {
		return x.Mapper
	}
	return nil
}

func (x *Librarian) GetSearcher() *Searcher {
	if x != nil {
		return x.Searcher
	}
	return nil
}

func (x *Librarian) GetMiner() *Miner {
	if x != nil {
		return x.Miner
	}
	return nil
}

type Librarian_EnableServiceDiscovery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mapper   bool `protobuf:"varint,1,opt,name=mapper,proto3" json:"mapper,omitempty"`
	Searcher bool `protobuf:"varint,2,opt,name=searcher,proto3" json:"searcher,omitempty"`
	Porter   bool `protobuf:"varint,3,opt,name=porter,proto3" json:"porter,omitempty"`
	Miner    bool `protobuf:"varint,4,opt,name=miner,proto3" json:"miner,omitempty"`
}

func (x *Librarian_EnableServiceDiscovery) Reset() {
	*x = Librarian_EnableServiceDiscovery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_librarian_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Librarian_EnableServiceDiscovery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Librarian_EnableServiceDiscovery) ProtoMessage() {}

func (x *Librarian_EnableServiceDiscovery) ProtoReflect() protoreflect.Message {
	mi := &file_conf_librarian_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Librarian_EnableServiceDiscovery.ProtoReflect.Descriptor instead.
func (*Librarian_EnableServiceDiscovery) Descriptor() ([]byte, []int) {
	return file_conf_librarian_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Librarian_EnableServiceDiscovery) GetMapper() bool {
	if x != nil {
		return x.Mapper
	}
	return false
}

func (x *Librarian_EnableServiceDiscovery) GetSearcher() bool {
	if x != nil {
		return x.Searcher
	}
	return false
}

func (x *Librarian_EnableServiceDiscovery) GetPorter() bool {
	if x != nil {
		return x.Porter
	}
	return false
}

func (x *Librarian_EnableServiceDiscovery) GetMiner() bool {
	if x != nil {
		return x.Miner
	}
	return false
}

var File_conf_librarian_proto protoreflect.FileDescriptor

var file_conf_librarian_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x73, 0x65, 0x70, 0x68, 0x69, 0x72,
	0x61, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6d,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x06, 0x0a, 0x09, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x12, 0x66, 0x0a, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x70, 0x68, 0x69, 0x72, 0x61, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x02, 0x73, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x33, 0x52, 0x02, 0x73, 0x33, 0x12, 0x24,
	0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x02, 0x6d, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x51,
	0x52, 0x02, 0x6d, 0x71, 0x12, 0x27, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x2a, 0x0a,
	0x06, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x24, 0x0a, 0x04, 0x6f, 0x74, 0x6c, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x54, 0x4c,
	0x50, 0x52, 0x04, 0x6f, 0x74, 0x6c, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x1a, 0x7a,
	0x0a, 0x16, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x1e, 0x5a, 0x1c, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_conf_librarian_proto_rawDescOnce sync.Once
	file_conf_librarian_proto_rawDescData = file_conf_librarian_proto_rawDesc
)

func file_conf_librarian_proto_rawDescGZIP() []byte {
	file_conf_librarian_proto_rawDescOnce.Do(func() {
		file_conf_librarian_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_librarian_proto_rawDescData)
	})
	return file_conf_librarian_proto_rawDescData
}

var file_conf_librarian_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_conf_librarian_proto_goTypes = []interface{}{
	(*Librarian)(nil),                        // 0: kratos.api.Librarian
	(*Librarian_EnableServiceDiscovery)(nil), // 1: kratos.api.Librarian.EnableServiceDiscovery
	(*SephirahServer)(nil),                   // 2: kratos.api.SephirahServer
	(*Database)(nil),                         // 3: kratos.api.Database
	(*S3)(nil),                               // 4: kratos.api.S3
	(*Auth)(nil),                             // 5: kratos.api.Auth
	(*MQ)(nil),                               // 6: kratos.api.MQ
	(*Cache)(nil),                            // 7: kratos.api.Cache
	(*Porter)(nil),                           // 8: kratos.api.Porter
	(*Consul)(nil),                           // 9: kratos.api.Consul
	(*Sentry)(nil),                           // 10: kratos.api.Sentry
	(*OTLP)(nil),                             // 11: kratos.api.OTLP
	(*Mapper)(nil),                           // 12: kratos.api.Mapper
	(*Searcher)(nil),                         // 13: kratos.api.Searcher
	(*Miner)(nil),                            // 14: kratos.api.Miner
}
var file_conf_librarian_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Librarian.enable_service_discovery:type_name -> kratos.api.Librarian.EnableServiceDiscovery
	2,  // 1: kratos.api.Librarian.server:type_name -> kratos.api.SephirahServer
	3,  // 2: kratos.api.Librarian.database:type_name -> kratos.api.Database
	4,  // 3: kratos.api.Librarian.s3:type_name -> kratos.api.S3
	5,  // 4: kratos.api.Librarian.auth:type_name -> kratos.api.Auth
	6,  // 5: kratos.api.Librarian.mq:type_name -> kratos.api.MQ
	7,  // 6: kratos.api.Librarian.cache:type_name -> kratos.api.Cache
	8,  // 7: kratos.api.Librarian.porter:type_name -> kratos.api.Porter
	9,  // 8: kratos.api.Librarian.consul:type_name -> kratos.api.Consul
	10, // 9: kratos.api.Librarian.sentry:type_name -> kratos.api.Sentry
	11, // 10: kratos.api.Librarian.otlp:type_name -> kratos.api.OTLP
	12, // 11: kratos.api.Librarian.mapper:type_name -> kratos.api.Mapper
	13, // 12: kratos.api.Librarian.searcher:type_name -> kratos.api.Searcher
	14, // 13: kratos.api.Librarian.miner:type_name -> kratos.api.Miner
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_conf_librarian_proto_init() }
func file_conf_librarian_proto_init() {
	if File_conf_librarian_proto != nil {
		return
	}
	file_conf_base_proto_init()
	file_conf_sephirah_proto_init()
	file_conf_mapper_proto_init()
	file_conf_searcher_proto_init()
	file_conf_miner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conf_librarian_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Librarian); i {
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
		file_conf_librarian_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Librarian_EnableServiceDiscovery); i {
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
			RawDescriptor: file_conf_librarian_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_librarian_proto_goTypes,
		DependencyIndexes: file_conf_librarian_proto_depIdxs,
		MessageInfos:      file_conf_librarian_proto_msgTypes,
	}.Build()
	File_conf_librarian_proto = out.File
	file_conf_librarian_proto_rawDesc = nil
	file_conf_librarian_proto_goTypes = nil
	file_conf_librarian_proto_depIdxs = nil
}
