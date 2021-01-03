// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: openconfig/golang/kingpinv2/v1/extensions.proto

// Extensions based on https://pkg.go.dev/gopkg.in/alecthomas/kingpin.v2

package kingpinv2

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Content is byte steam created from PathOrContent flag, a custom extension built on top of kingpin.v2 flags type.
type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *any.Any `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"` // any proto.Message
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetContent() *any.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

// ExistingFile represents https://pkg.go.dev/gopkg.in/alecthomas/kingpin.v2#ArgClause.ExistingFile.
// repeated ExistingFile represents https://pkg.go.dev/gopkg.in/alecthomas/kingpin.v2#ArgClause.ExistingFiles.
type ExistingFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *any.Any `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"` // *os.File
}

func (x *ExistingFile) Reset() {
	*x = ExistingFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistingFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistingFile) ProtoMessage() {}

func (x *ExistingFile) ProtoReflect() protoreflect.Message {
	mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistingFile.ProtoReflect.Descriptor instead.
func (*ExistingFile) Descriptor() ([]byte, []int) {
	return file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescGZIP(), []int{1}
}

func (x *ExistingFile) GetFile() *any.Any {
	if x != nil {
		return x.File
	}
	return nil
}

// IP represents https://pkg.go.dev/gopkg.in/alecthomas/kingpin.v2#ArgClause.IP.
// repeated IP represents https://pkg.go.dev/gopkg.in/alecthomas/kingpin.v2#ArgClause.IPList.
type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip *any.Any `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"` // *net.IP
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescGZIP(), []int{2}
}

func (x *IP) GetIp() *any.Any {
	if x != nil {
		return x.Ip
	}
	return nil
}

// Regexp represents https://pkg.go.dev/gopkg.in/alecthomas/kingpin.v2#ArgClause.Regexp.
// repeated Regexp represents https://pkg.go.dev/gopkg.in/alecthomas/kingpin.v2#ArgClause.RegexpList.
type Regexp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regexp *any.Any `protobuf:"bytes,1,opt,name=regexp,proto3" json:"regexp,omitempty"` // **regexp.Regexp
}

func (x *Regexp) Reset() {
	*x = Regexp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regexp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regexp) ProtoMessage() {}

func (x *Regexp) ProtoReflect() protoreflect.Message {
	mi := &file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regexp.ProtoReflect.Descriptor instead.
func (*Regexp) Descriptor() ([]byte, []int) {
	return file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescGZIP(), []int{3}
}

func (x *Regexp) GetRegexp() *any.Any {
	if x != nil {
		return x.Regexp
	}
	return nil
}

var file_openconfig_golang_kingpinv2_v1_extensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         6000,
		Name:          "golang.kingpinv2.v1.placeholder",
		Tag:           "bytes,6000,opt,name=placeholder",
		Filename:      "openconfig/golang/kingpinv2/v1/extensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         6001,
		Name:          "golang.kingpinv2.v1.envvar",
		Tag:           "bytes,6001,opt,name=envvar",
		Filename:      "openconfig/golang/kingpinv2/v1/extensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         6002,
		Name:          "golang.kingpinv2.v1.argument",
		Tag:           "varint,6002,opt,name=argument",
		Filename:      "openconfig/golang/kingpinv2/v1/extensions.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         6002,
		Name:          "golang.kingpinv2.v1.command_name",
		Tag:           "bytes,6002,opt,name=command_name",
		Filename:      "openconfig/golang/kingpinv2/v1/extensions.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional string placeholder = 6000;
	E_Placeholder = &file_openconfig_golang_kingpinv2_v1_extensions_proto_extTypes[0]
	// optional string envvar = 6001;
	E_Envvar = &file_openconfig_golang_kingpinv2_v1_extensions_proto_extTypes[1]
	// By default field represents a flag.
	//
	// optional bool argument = 6002;
	E_Argument = &file_openconfig_golang_kingpinv2_v1_extensions_proto_extTypes[2]
)

// Extension fields to descriptor.MessageOptions.
var (
	// By default message represents just complex configuration type. If command_name is specified
	// such message becomes a kingpin.v2 command.
	// Name has to be unique within single openconfig.Configuration type.
	//
	// optional string command_name = 6002;
	E_CommandName = &file_openconfig_golang_kingpinv2_v1_extensions_proto_extTypes[3]
)

var File_openconfig_golang_kingpinv2_v1_extensions_proto protoreflect.FileDescriptor

var file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x69, 0x6e, 0x67, 0x70, 0x69, 0x6e, 0x76, 0x32, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6b, 0x69, 0x6e, 0x67, 0x70, 0x69,
	0x6e, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x38,
	0x0a, 0x0c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x24,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x02, 0x69, 0x70, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x2c,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x3a, 0x40, 0x0a, 0x0b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf0, 0x2e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x3a, 0x36,
	0x0a, 0x06, 0x65, 0x6e, 0x76, 0x76, 0x61, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf1, 0x2e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x6e, 0x76, 0x76, 0x61, 0x72, 0x3a, 0x3a, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xf2, 0x2e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x3a, 0x43, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf2, 0x2e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2d, 0x69, 0x6f, 0x2f,
	0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x6b, 0x69, 0x6e, 0x67, 0x70, 0x69, 0x6e, 0x76, 0x32, 0x3b, 0x6b, 0x69, 0x6e, 0x67,
	0x70, 0x69, 0x6e, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescOnce sync.Once
	file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescData = file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDesc
)

func file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescGZIP() []byte {
	file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescOnce.Do(func() {
		file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescData)
	})
	return file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDescData
}

var file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_openconfig_golang_kingpinv2_v1_extensions_proto_goTypes = []interface{}{
	(*Content)(nil),                   // 0: golang.kingpinv2.v1.Content
	(*ExistingFile)(nil),              // 1: golang.kingpinv2.v1.ExistingFile
	(*IP)(nil),                        // 2: golang.kingpinv2.v1.IP
	(*Regexp)(nil),                    // 3: golang.kingpinv2.v1.Regexp
	(*any.Any)(nil),                   // 4: google.protobuf.Any
	(*descriptor.FieldOptions)(nil),   // 5: google.protobuf.FieldOptions
	(*descriptor.MessageOptions)(nil), // 6: google.protobuf.MessageOptions
}
var file_openconfig_golang_kingpinv2_v1_extensions_proto_depIdxs = []int32{
	4, // 0: golang.kingpinv2.v1.Content.content:type_name -> google.protobuf.Any
	4, // 1: golang.kingpinv2.v1.ExistingFile.file:type_name -> google.protobuf.Any
	4, // 2: golang.kingpinv2.v1.IP.ip:type_name -> google.protobuf.Any
	4, // 3: golang.kingpinv2.v1.Regexp.regexp:type_name -> google.protobuf.Any
	5, // 4: golang.kingpinv2.v1.placeholder:extendee -> google.protobuf.FieldOptions
	5, // 5: golang.kingpinv2.v1.envvar:extendee -> google.protobuf.FieldOptions
	5, // 6: golang.kingpinv2.v1.argument:extendee -> google.protobuf.FieldOptions
	6, // 7: golang.kingpinv2.v1.command_name:extendee -> google.protobuf.MessageOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	4, // [4:8] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_openconfig_golang_kingpinv2_v1_extensions_proto_init() }
func file_openconfig_golang_kingpinv2_v1_extensions_proto_init() {
	if File_openconfig_golang_kingpinv2_v1_extensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistingFile); i {
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
		file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
		file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regexp); i {
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
			RawDescriptor: file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_openconfig_golang_kingpinv2_v1_extensions_proto_goTypes,
		DependencyIndexes: file_openconfig_golang_kingpinv2_v1_extensions_proto_depIdxs,
		MessageInfos:      file_openconfig_golang_kingpinv2_v1_extensions_proto_msgTypes,
		ExtensionInfos:    file_openconfig_golang_kingpinv2_v1_extensions_proto_extTypes,
	}.Build()
	File_openconfig_golang_kingpinv2_v1_extensions_proto = out.File
	file_openconfig_golang_kingpinv2_v1_extensions_proto_rawDesc = nil
	file_openconfig_golang_kingpinv2_v1_extensions_proto_goTypes = nil
	file_openconfig_golang_kingpinv2_v1_extensions_proto_depIdxs = nil
}
