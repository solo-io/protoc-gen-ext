// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: tests/api/hello.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type Test int32

const (
	Test_HELLO Test = 0
	Test_WORLD Test = 1
)

// Enum value maps for Test.
var (
	Test_name = map[int32]string{
		0: "HELLO",
		1: "WORLD",
	}
	Test_value = map[string]int32{
		"HELLO": 0,
		"WORLD": 1,
	}
)

func (x Test) Enum() *Test {
	p := new(Test)
	*p = x
	return p
}

func (x Test) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Test) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_api_hello_proto_enumTypes[0].Descriptor()
}

func (Test) Type() protoreflect.EnumType {
	return &file_tests_api_hello_proto_enumTypes[0]
}

func (x Test) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Test.Descriptor instead.
func (Test) EnumDescriptor() ([]byte, []int) {
	return file_tests_api_hello_proto_rawDescGZIP(), []int{0}
}

// proto object to test all base objects
type Simple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str          string  `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	Byt          []byte  `protobuf:"bytes,2,opt,name=byt,proto3" json:"byt,omitempty"`
	TestUint32   uint32  `protobuf:"varint,3,opt,name=testUint32,proto3" json:"testUint32,omitempty"`
	TestUint64   uint64  `protobuf:"varint,4,opt,name=testUint64,proto3" json:"testUint64,omitempty"`
	TestBool     bool    `protobuf:"varint,5,opt,name=testBool,proto3" json:"testBool,omitempty"`
	DoubleTest   float64 `protobuf:"fixed64,6,opt,name=doubleTest,proto3" json:"doubleTest,omitempty"`
	FloatTest    float32 `protobuf:"fixed32,7,opt,name=floatTest,proto3" json:"floatTest,omitempty"`
	Int32Test    int32   `protobuf:"varint,8,opt,name=int32Test,proto3" json:"int32Test,omitempty"`
	Int64Test    int64   `protobuf:"varint,9,opt,name=int64Test,proto3" json:"int64Test,omitempty"`
	Sint32Test   int32   `protobuf:"zigzag32,10,opt,name=sint32Test,proto3" json:"sint32Test,omitempty"`
	Sint64Test   int64   `protobuf:"zigzag64,11,opt,name=sint64Test,proto3" json:"sint64Test,omitempty"`
	Fixed32Test  uint32  `protobuf:"fixed32,12,opt,name=fixed32Test,proto3" json:"fixed32Test,omitempty"`
	Fixed64Test  uint64  `protobuf:"fixed64,13,opt,name=fixed64Test,proto3" json:"fixed64Test,omitempty"`
	Sfixed32Test int32   `protobuf:"fixed32,14,opt,name=sfixed32Test,proto3" json:"sfixed32Test,omitempty"`
	Sfixed64Test int64   `protobuf:"fixed64,15,opt,name=sfixed64Test,proto3" json:"sfixed64Test,omitempty"`
	// skipped primitives
	StrSkipped string `protobuf:"bytes,16,opt,name=str_skipped,json=strSkipped,proto3" json:"str_skipped,omitempty"`
	IntSkipped uint32 `protobuf:"varint,17,opt,name=int_skipped,json=intSkipped,proto3" json:"int_skipped,omitempty"`
}

func (x *Simple) Reset() {
	*x = Simple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_api_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simple) ProtoMessage() {}

func (x *Simple) ProtoReflect() protoreflect.Message {
	mi := &file_tests_api_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simple.ProtoReflect.Descriptor instead.
func (*Simple) Descriptor() ([]byte, []int) {
	return file_tests_api_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Simple) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Simple) GetByt() []byte {
	if x != nil {
		return x.Byt
	}
	return nil
}

func (x *Simple) GetTestUint32() uint32 {
	if x != nil {
		return x.TestUint32
	}
	return 0
}

func (x *Simple) GetTestUint64() uint64 {
	if x != nil {
		return x.TestUint64
	}
	return 0
}

func (x *Simple) GetTestBool() bool {
	if x != nil {
		return x.TestBool
	}
	return false
}

func (x *Simple) GetDoubleTest() float64 {
	if x != nil {
		return x.DoubleTest
	}
	return 0
}

func (x *Simple) GetFloatTest() float32 {
	if x != nil {
		return x.FloatTest
	}
	return 0
}

func (x *Simple) GetInt32Test() int32 {
	if x != nil {
		return x.Int32Test
	}
	return 0
}

func (x *Simple) GetInt64Test() int64 {
	if x != nil {
		return x.Int64Test
	}
	return 0
}

func (x *Simple) GetSint32Test() int32 {
	if x != nil {
		return x.Sint32Test
	}
	return 0
}

func (x *Simple) GetSint64Test() int64 {
	if x != nil {
		return x.Sint64Test
	}
	return 0
}

func (x *Simple) GetFixed32Test() uint32 {
	if x != nil {
		return x.Fixed32Test
	}
	return 0
}

func (x *Simple) GetFixed64Test() uint64 {
	if x != nil {
		return x.Fixed64Test
	}
	return 0
}

func (x *Simple) GetSfixed32Test() int32 {
	if x != nil {
		return x.Sfixed32Test
	}
	return 0
}

func (x *Simple) GetSfixed64Test() int64 {
	if x != nil {
		return x.Sfixed64Test
	}
	return 0
}

func (x *Simple) GetStrSkipped() string {
	if x != nil {
		return x.StrSkipped
	}
	return ""
}

func (x *Simple) GetIntSkipped() uint32 {
	if x != nil {
		return x.IntSkipped
	}
	return 0
}

type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Simple  *Simple            `protobuf:"bytes,1,opt,name=simple,proto3" json:"simple,omitempty"`
	Details *_struct.Struct    `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
	Test    Test               `protobuf:"varint,2,opt,name=test,proto3,enum=envoy.type.Test" json:"test,omitempty"`
	Empty   *Empty             `protobuf:"bytes,3,opt,name=empty,proto3" json:"empty,omitempty"`
	Hello   []string           `protobuf:"bytes,4,rep,name=hello,proto3" json:"hello,omitempty"`
	Skipper *Simple            `protobuf:"bytes,6,opt,name=skipper,proto3" json:"skipper,omitempty"`
	X       []*Simple          `protobuf:"bytes,7,rep,name=x,proto3" json:"x,omitempty"`
	Initial map[string]*Simple `protobuf:"bytes,9,rep,name=initial,proto3" json:"initial,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to TestOneOf:
	//	*Nested_EmptyOneOf
	//	*Nested_NestedOneOf
	TestOneOf isNested_TestOneOf `protobuf_oneof:"test_one_of"`
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_api_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_tests_api_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_tests_api_hello_proto_rawDescGZIP(), []int{1}
}

func (x *Nested) GetSimple() *Simple {
	if x != nil {
		return x.Simple
	}
	return nil
}

func (x *Nested) GetDetails() *_struct.Struct {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Nested) GetTest() Test {
	if x != nil {
		return x.Test
	}
	return Test_HELLO
}

func (x *Nested) GetEmpty() *Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

func (x *Nested) GetHello() []string {
	if x != nil {
		return x.Hello
	}
	return nil
}

func (x *Nested) GetSkipper() *Simple {
	if x != nil {
		return x.Skipper
	}
	return nil
}

func (x *Nested) GetX() []*Simple {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *Nested) GetInitial() map[string]*Simple {
	if x != nil {
		return x.Initial
	}
	return nil
}

func (m *Nested) GetTestOneOf() isNested_TestOneOf {
	if m != nil {
		return m.TestOneOf
	}
	return nil
}

func (x *Nested) GetEmptyOneOf() *Empty {
	if x, ok := x.GetTestOneOf().(*Nested_EmptyOneOf); ok {
		return x.EmptyOneOf
	}
	return nil
}

func (x *Nested) GetNestedOneOf() *NestedEmpty {
	if x, ok := x.GetTestOneOf().(*Nested_NestedOneOf); ok {
		return x.NestedOneOf
	}
	return nil
}

type isNested_TestOneOf interface {
	isNested_TestOneOf()
}

type Nested_EmptyOneOf struct {
	EmptyOneOf *Empty `protobuf:"bytes,11,opt,name=empty_one_of,json=emptyOneOf,proto3,oneof"`
}

type Nested_NestedOneOf struct {
	NestedOneOf *NestedEmpty `protobuf:"bytes,12,opt,name=nested_one_of,json=nestedOneOf,proto3,oneof"`
}

func (*Nested_EmptyOneOf) isNested_TestOneOf() {}

func (*Nested_NestedOneOf) isNested_TestOneOf() {}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_api_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_tests_api_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_tests_api_hello_proto_rawDescGZIP(), []int{2}
}

type NestedEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nested *Nested `protobuf:"bytes,1,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *NestedEmpty) Reset() {
	*x = NestedEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_api_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedEmpty) ProtoMessage() {}

func (x *NestedEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_tests_api_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedEmpty.ProtoReflect.Descriptor instead.
func (*NestedEmpty) Descriptor() ([]byte, []int) {
	return file_tests_api_hello_proto_rawDescGZIP(), []int{3}
}

func (x *NestedEmpty) GetNested() *Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

var File_tests_api_hello_proto protoreflect.FileDescriptor

var file_tests_api_hello_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x79, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x62, 0x79, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x55, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x55, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x54, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x65, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x54, 0x65, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0a,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x54, 0x65, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0a,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x54, 0x65, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x0b, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x54, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x54, 0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x06, 0x52, 0x0b, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x54, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x54, 0x65, 0x73, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0f, 0x52, 0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x54, 0x65,
	0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x54, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x5f, 0x73, 0x6b,
	0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xb8, 0xf5, 0x04,
	0x01, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x53, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x04, 0xb8, 0xf5, 0x04, 0x01, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x53, 0x6b, 0x69,
	0x70, 0x70, 0x65, 0x64, 0x22, 0xb2, 0x04, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12,
	0x2a, 0x0a, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x74, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x32, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x04, 0xb8, 0xf5, 0x04, 0x01, 0x52, 0x07,
	0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x01, 0x78, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x01, 0x78, 0x12, 0x39, 0x0a, 0x07, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x0c, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x6f, 0x6e,
	0x65, 0x5f, 0x6f, 0x66, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x12, 0x3d, 0x0a, 0x0d, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x1a, 0x4e, 0x0a, 0x0c, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x39, 0x0a, 0x0b, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x2a, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2a, 0x1c, 0x0a,
	0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x01, 0x42, 0x39, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x65, 0x78, 0x74,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0xb8, 0xf5, 0x04, 0x01, 0xc8, 0xf5,
	0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_api_hello_proto_rawDescOnce sync.Once
	file_tests_api_hello_proto_rawDescData = file_tests_api_hello_proto_rawDesc
)

func file_tests_api_hello_proto_rawDescGZIP() []byte {
	file_tests_api_hello_proto_rawDescOnce.Do(func() {
		file_tests_api_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_api_hello_proto_rawDescData)
	})
	return file_tests_api_hello_proto_rawDescData
}

var file_tests_api_hello_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_api_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tests_api_hello_proto_goTypes = []interface{}{
	(Test)(0),              // 0: envoy.type.Test
	(*Simple)(nil),         // 1: envoy.type.Simple
	(*Nested)(nil),         // 2: envoy.type.Nested
	(*Empty)(nil),          // 3: envoy.type.Empty
	(*NestedEmpty)(nil),    // 4: envoy.type.NestedEmpty
	nil,                    // 5: envoy.type.Nested.InitialEntry
	(*_struct.Struct)(nil), // 6: google.protobuf.Struct
}
var file_tests_api_hello_proto_depIdxs = []int32{
	1,  // 0: envoy.type.Nested.simple:type_name -> envoy.type.Simple
	6,  // 1: envoy.type.Nested.details:type_name -> google.protobuf.Struct
	0,  // 2: envoy.type.Nested.test:type_name -> envoy.type.Test
	3,  // 3: envoy.type.Nested.empty:type_name -> envoy.type.Empty
	1,  // 4: envoy.type.Nested.skipper:type_name -> envoy.type.Simple
	1,  // 5: envoy.type.Nested.x:type_name -> envoy.type.Simple
	5,  // 6: envoy.type.Nested.initial:type_name -> envoy.type.Nested.InitialEntry
	3,  // 7: envoy.type.Nested.empty_one_of:type_name -> envoy.type.Empty
	4,  // 8: envoy.type.Nested.nested_one_of:type_name -> envoy.type.NestedEmpty
	2,  // 9: envoy.type.NestedEmpty.nested:type_name -> envoy.type.Nested
	1,  // 10: envoy.type.Nested.InitialEntry.value:type_name -> envoy.type.Simple
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tests_api_hello_proto_init() }
func file_tests_api_hello_proto_init() {
	if File_tests_api_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_api_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simple); i {
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
		file_tests_api_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested); i {
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
		file_tests_api_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_tests_api_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedEmpty); i {
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
	file_tests_api_hello_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Nested_EmptyOneOf)(nil),
		(*Nested_NestedOneOf)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_api_hello_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_api_hello_proto_goTypes,
		DependencyIndexes: file_tests_api_hello_proto_depIdxs,
		EnumInfos:         file_tests_api_hello_proto_enumTypes,
		MessageInfos:      file_tests_api_hello_proto_msgTypes,
	}.Build()
	File_tests_api_hello_proto = out.File
	file_tests_api_hello_proto_rawDesc = nil
	file_tests_api_hello_proto_goTypes = nil
	file_tests_api_hello_proto_depIdxs = nil
}
