// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pool/pool_with_slice_reuse.proto

package pool

import (
	_ "github.com/planetscale/vtprotobuf/vtproto"
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

type Slice1E int32

const (
	Slice1_A Slice1E = 0
	Slice1_B Slice1E = 1
)

// Enum value maps for Slice1E.
var (
	Slice1E_name = map[int32]string{
		0: "A",
		1: "B",
	}
	Slice1E_value = map[string]int32{
		"A": 0,
		"B": 1,
	}
)

func (x Slice1E) Enum() *Slice1E {
	p := new(Slice1E)
	*p = x
	return p
}

func (x Slice1E) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Slice1E) Descriptor() protoreflect.EnumDescriptor {
	return file_pool_pool_with_slice_reuse_proto_enumTypes[0].Descriptor()
}

func (Slice1E) Type() protoreflect.EnumType {
	return &file_pool_pool_with_slice_reuse_proto_enumTypes[0]
}

func (x Slice1E) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Slice1E.Descriptor instead.
func (Slice1E) EnumDescriptor() ([]byte, []int) {
	return file_pool_pool_with_slice_reuse_proto_rawDescGZIP(), []int{1, 0}
}

type Test1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sl []*Slice1 `protobuf:"bytes,1,rep,name=Sl,proto3" json:"Sl,omitempty"`
}

func (x *Test1) Reset() {
	*x = Test1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test1) ProtoMessage() {}

func (x *Test1) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test1.ProtoReflect.Descriptor instead.
func (*Test1) Descriptor() ([]byte, []int) {
	return file_pool_pool_with_slice_reuse_proto_rawDescGZIP(), []int{0}
}

func (x *Test1) GetSl() []*Slice1 {
	if x != nil {
		return x.Sl
	}
	return nil
}

type Slice1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64  `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C uint64 `protobuf:"fixed64,3,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *Slice1) Reset() {
	*x = Slice1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slice1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slice1) ProtoMessage() {}

func (x *Slice1) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slice1.ProtoReflect.Descriptor instead.
func (*Slice1) Descriptor() ([]byte, []int) {
	return file_pool_pool_with_slice_reuse_proto_rawDescGZIP(), []int{1}
}

func (x *Slice1) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Slice1) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

func (x *Slice1) GetC() uint64 {
	if x != nil {
		return x.C
	}
	return 0
}

type Test2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sl []*Slice2 `protobuf:"bytes,1,rep,name=Sl,proto3" json:"Sl,omitempty"`
}

func (x *Test2) Reset() {
	*x = Test2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2) ProtoMessage() {}

func (x *Test2) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2.ProtoReflect.Descriptor instead.
func (*Test2) Descriptor() ([]byte, []int) {
	return file_pool_pool_with_slice_reuse_proto_rawDescGZIP(), []int{2}
}

func (x *Test2) GetSl() []*Slice2 {
	if x != nil {
		return x.Sl
	}
	return nil
}

type Slice2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A map[int64]int64 `protobuf:"bytes,1,rep,name=a,proto3" json:"a,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	B *int32          `protobuf:"varint,2,opt,name=b,proto3,oneof" json:"b,omitempty"`
	C []string        `protobuf:"bytes,3,rep,name=c,proto3" json:"c,omitempty"`
	D *Element2       `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *Slice2) Reset() {
	*x = Slice2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slice2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slice2) ProtoMessage() {}

func (x *Slice2) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slice2.ProtoReflect.Descriptor instead.
func (*Slice2) Descriptor() ([]byte, []int) {
	return file_pool_pool_with_slice_reuse_proto_rawDescGZIP(), []int{3}
}

func (x *Slice2) GetA() map[int64]int64 {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Slice2) GetB() int32 {
	if x != nil && x.B != nil {
		return *x.B
	}
	return 0
}

func (x *Slice2) GetC() []string {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *Slice2) GetD() *Element2 {
	if x != nil {
		return x.D
	}
	return nil
}

type Element2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Element2) Reset() {
	*x = Element2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element2) ProtoMessage() {}

func (x *Element2) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_with_slice_reuse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Element2.ProtoReflect.Descriptor instead.
func (*Element2) Descriptor() ([]byte, []int) {
	return file_pool_pool_with_slice_reuse_proto_rawDescGZIP(), []int{4}
}

func (x *Element2) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

var File_pool_pool_with_slice_reuse_proto protoreflect.FileDescriptor

var file_pool_pool_with_slice_reuse_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x31,
	0x12, 0x17, 0x0a, 0x02, 0x53, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0x31, 0x52, 0x02, 0x53, 0x6c, 0x3a, 0x04, 0xa8, 0xa6, 0x1f, 0x01, 0x22,
	0x45, 0x0a, 0x06, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x01, 0x63, 0x22, 0x11, 0x0a, 0x01, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12,
	0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x22, 0x26, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0x12,
	0x17, 0x0a, 0x02, 0x53, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x32, 0x52, 0x02, 0x53, 0x6c, 0x3a, 0x04, 0xa8, 0xa6, 0x1f, 0x01, 0x22, 0x9c,
	0x01, 0x0a, 0x06, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x32, 0x12, 0x1c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x32, 0x2e, 0x41, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x01, 0x61, 0x12, 0x11, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x01, 0x62, 0x88, 0x01, 0x01, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x01, 0x63, 0x12, 0x17, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x52, 0x01,
	0x64, 0x1a, 0x34, 0x0a, 0x06, 0x41, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x62, 0x22, 0x18, 0x0a,
	0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x42, 0x10, 0x5a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pool_pool_with_slice_reuse_proto_rawDescOnce sync.Once
	file_pool_pool_with_slice_reuse_proto_rawDescData = file_pool_pool_with_slice_reuse_proto_rawDesc
)

func file_pool_pool_with_slice_reuse_proto_rawDescGZIP() []byte {
	file_pool_pool_with_slice_reuse_proto_rawDescOnce.Do(func() {
		file_pool_pool_with_slice_reuse_proto_rawDescData = protoimpl.X.CompressGZIP(file_pool_pool_with_slice_reuse_proto_rawDescData)
	})
	return file_pool_pool_with_slice_reuse_proto_rawDescData
}

var file_pool_pool_with_slice_reuse_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pool_pool_with_slice_reuse_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pool_pool_with_slice_reuse_proto_goTypes = []interface{}{
	(Slice1E)(0),     // 0: Slice1.e
	(*Test1)(nil),    // 1: Test1
	(*Slice1)(nil),   // 2: Slice1
	(*Test2)(nil),    // 3: Test2
	(*Slice2)(nil),   // 4: Slice2
	(*Element2)(nil), // 5: Element2
	nil,              // 6: Slice2.AEntry
}
var file_pool_pool_with_slice_reuse_proto_depIdxs = []int32{
	2, // 0: Test1.Sl:type_name -> Slice1
	4, // 1: Test2.Sl:type_name -> Slice2
	6, // 2: Slice2.a:type_name -> Slice2.AEntry
	5, // 3: Slice2.d:type_name -> Element2
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pool_pool_with_slice_reuse_proto_init() }
func file_pool_pool_with_slice_reuse_proto_init() {
	if File_pool_pool_with_slice_reuse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pool_pool_with_slice_reuse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test1); i {
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
		file_pool_pool_with_slice_reuse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slice1); i {
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
		file_pool_pool_with_slice_reuse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2); i {
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
		file_pool_pool_with_slice_reuse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slice2); i {
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
		file_pool_pool_with_slice_reuse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Element2); i {
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
	file_pool_pool_with_slice_reuse_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pool_pool_with_slice_reuse_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pool_pool_with_slice_reuse_proto_goTypes,
		DependencyIndexes: file_pool_pool_with_slice_reuse_proto_depIdxs,
		EnumInfos:         file_pool_pool_with_slice_reuse_proto_enumTypes,
		MessageInfos:      file_pool_pool_with_slice_reuse_proto_msgTypes,
	}.Build()
	File_pool_pool_with_slice_reuse_proto = out.File
	file_pool_pool_with_slice_reuse_proto_rawDesc = nil
	file_pool_pool_with_slice_reuse_proto_goTypes = nil
	file_pool_pool_with_slice_reuse_proto_depIdxs = nil
}
