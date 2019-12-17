// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tests/api/hello.proto

package api

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/solo-io/protoc-gen-ext/ext"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Test int32

const (
	Test_HELLO Test = 0
	Test_WORLD Test = 1
)

var Test_name = map[int32]string{
	0: "HELLO",
	1: "WORLD",
}

var Test_value = map[string]int32{
	"HELLO": 0,
	"WORLD": 1,
}

func (x Test) String() string {
	return proto.EnumName(Test_name, int32(x))
}

func (Test) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e23aa1f3bc52fd3, []int{0}
}

// proto object to test all base objects
type Simple struct {
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
	StrSkipped           string   `protobuf:"bytes,16,opt,name=str_skipped,json=strSkipped,proto3" json:"str_skipped,omitempty"`
	IntSkipped           uint32   `protobuf:"varint,17,opt,name=int_skipped,json=intSkipped,proto3" json:"int_skipped,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Simple) Reset()         { *m = Simple{} }
func (m *Simple) String() string { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()    {}
func (*Simple) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e23aa1f3bc52fd3, []int{0}
}

func (m *Simple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple.Unmarshal(m, b)
}
func (m *Simple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple.Marshal(b, m, deterministic)
}
func (m *Simple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple.Merge(m, src)
}
func (m *Simple) XXX_Size() int {
	return xxx_messageInfo_Simple.Size(m)
}
func (m *Simple) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple.DiscardUnknown(m)
}

var xxx_messageInfo_Simple proto.InternalMessageInfo

func (m *Simple) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Simple) GetByt() []byte {
	if m != nil {
		return m.Byt
	}
	return nil
}

func (m *Simple) GetTestUint32() uint32 {
	if m != nil {
		return m.TestUint32
	}
	return 0
}

func (m *Simple) GetTestUint64() uint64 {
	if m != nil {
		return m.TestUint64
	}
	return 0
}

func (m *Simple) GetTestBool() bool {
	if m != nil {
		return m.TestBool
	}
	return false
}

func (m *Simple) GetDoubleTest() float64 {
	if m != nil {
		return m.DoubleTest
	}
	return 0
}

func (m *Simple) GetFloatTest() float32 {
	if m != nil {
		return m.FloatTest
	}
	return 0
}

func (m *Simple) GetInt32Test() int32 {
	if m != nil {
		return m.Int32Test
	}
	return 0
}

func (m *Simple) GetInt64Test() int64 {
	if m != nil {
		return m.Int64Test
	}
	return 0
}

func (m *Simple) GetSint32Test() int32 {
	if m != nil {
		return m.Sint32Test
	}
	return 0
}

func (m *Simple) GetSint64Test() int64 {
	if m != nil {
		return m.Sint64Test
	}
	return 0
}

func (m *Simple) GetFixed32Test() uint32 {
	if m != nil {
		return m.Fixed32Test
	}
	return 0
}

func (m *Simple) GetFixed64Test() uint64 {
	if m != nil {
		return m.Fixed64Test
	}
	return 0
}

func (m *Simple) GetSfixed32Test() int32 {
	if m != nil {
		return m.Sfixed32Test
	}
	return 0
}

func (m *Simple) GetSfixed64Test() int64 {
	if m != nil {
		return m.Sfixed64Test
	}
	return 0
}

func (m *Simple) GetStrSkipped() string {
	if m != nil {
		return m.StrSkipped
	}
	return ""
}

func (m *Simple) GetIntSkipped() uint32 {
	if m != nil {
		return m.IntSkipped
	}
	return 0
}

type Nested struct {
	Simple               *Simple            `protobuf:"bytes,1,opt,name=simple,proto3" json:"simple,omitempty"`
	Details              *_struct.Struct    `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
	Test                 Test               `protobuf:"varint,2,opt,name=test,proto3,enum=envoy.type.Test" json:"test,omitempty"`
	Empty                *Empty             `protobuf:"bytes,3,opt,name=empty,proto3" json:"empty,omitempty"`
	Hello                []string           `protobuf:"bytes,4,rep,name=hello,proto3" json:"hello,omitempty"`
	Skipper              *Simple            `protobuf:"bytes,6,opt,name=skipper,proto3" json:"skipper,omitempty"`
	X                    []*Simple          `protobuf:"bytes,7,rep,name=x,proto3" json:"x,omitempty"`
	Initial              map[string]*Simple `protobuf:"bytes,9,rep,name=initial,proto3" json:"initial,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EmptyNonNullable     *Empty             `protobuf:"bytes,10,opt,name=emptyNonNullable,proto3" json:"emptyNonNullable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e23aa1f3bc52fd3, []int{1}
}

func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetSimple() *Simple {
	if m != nil {
		return m.Simple
	}
	return nil
}

func (m *Nested) GetDetails() *_struct.Struct {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Nested) GetTest() Test {
	if m != nil {
		return m.Test
	}
	return Test_HELLO
}

func (m *Nested) GetEmpty() *Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

func (m *Nested) GetHello() []string {
	if m != nil {
		return m.Hello
	}
	return nil
}

func (m *Nested) GetSkipper() *Simple {
	if m != nil {
		return m.Skipper
	}
	return nil
}

func (m *Nested) GetX() []*Simple {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *Nested) GetInitial() map[string]*Simple {
	if m != nil {
		return m.Initial
	}
	return nil
}

func (m *Nested) GetEmptyNonNullable() *Empty {
	if m != nil {
		return m.EmptyNonNullable
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e23aa1f3bc52fd3, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.type.Test", Test_name, Test_value)
	proto.RegisterType((*Simple)(nil), "envoy.type.Simple")
	proto.RegisterType((*Nested)(nil), "envoy.type.Nested")
	proto.RegisterMapType((map[string]*Simple)(nil), "envoy.type.Nested.InitialEntry")
	proto.RegisterType((*Empty)(nil), "envoy.type.Empty")
}

func init() { proto.RegisterFile("tests/api/hello.proto", fileDescriptor_1e23aa1f3bc52fd3) }

var fileDescriptor_1e23aa1f3bc52fd3 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xab, 0xda, 0xce, 0x9f, 0x93, 0xb4, 0x75, 0x45, 0xc7, 0x4c, 0x28, 0xab, 0x08, 0x85,
	0x89, 0x5e, 0x38, 0x2c, 0x2d, 0x65, 0xdb, 0x65, 0xb6, 0xc2, 0x06, 0x25, 0x05, 0x75, 0x63, 0xb0,
	0x9b, 0xe1, 0x2c, 0x6a, 0x66, 0xaa, 0x5a, 0xc6, 0x52, 0x4a, 0xfc, 0x2a, 0x7b, 0xa2, 0x3d, 0xc5,
	0xde, 0x61, 0x77, 0xbd, 0x1c, 0x92, 0x9d, 0x5a, 0x65, 0xd9, 0xdd, 0xf1, 0xf7, 0xfd, 0x74, 0xac,
	0x23, 0x7d, 0x82, 0x67, 0x9a, 0x2b, 0xad, 0x46, 0x49, 0x9e, 0x8e, 0x7e, 0x70, 0x21, 0x64, 0x9c,
	0x17, 0x52, 0x4b, 0x0c, 0x3c, 0xbb, 0x97, 0x65, 0xac, 0xcb, 0x9c, 0x0f, 0x0e, 0x17, 0x52, 0x2e,
	0x04, 0x1f, 0x59, 0x67, 0xb6, 0xbc, 0x19, 0x29, 0x5d, 0x2c, 0xbf, 0xeb, 0x8a, 0x1c, 0x1c, 0x2c,
	0xe4, 0x42, 0xda, 0x72, 0x64, 0xaa, 0x5a, 0xdd, 0xe1, 0x2b, 0x3d, 0xe2, 0xab, 0x1a, 0x1a, 0xfe,
	0xf4, 0xa1, 0x75, 0x9d, 0xde, 0xe5, 0x82, 0xe3, 0x10, 0x3c, 0xa5, 0x8b, 0x08, 0x11, 0x44, 0xbb,
	0xcc, 0x94, 0x46, 0x99, 0x95, 0x3a, 0xda, 0x26, 0x88, 0xf6, 0x99, 0x29, 0xf1, 0x0b, 0x00, 0xb3,
	0xad, 0xcf, 0x69, 0xa6, 0x4f, 0xc7, 0x91, 0x47, 0x10, 0xdd, 0x61, 0x8e, 0xe2, 0xfa, 0xe7, 0x67,
	0x91, 0x4f, 0x10, 0xf5, 0x99, 0xa3, 0xe0, 0x01, 0x74, 0xcc, 0xd7, 0x44, 0x4a, 0x11, 0x05, 0x04,
	0xd1, 0x0e, 0x7b, 0xfc, 0x36, 0x6b, 0xe7, 0x72, 0x39, 0x13, 0xfc, 0x13, 0x57, 0x3a, 0x6a, 0x11,
	0x44, 0x11, 0x73, 0x14, 0x7c, 0x08, 0xdd, 0x1b, 0x21, 0x13, 0x6d, 0xed, 0x36, 0x41, 0x74, 0x9b,
	0x35, 0x82, 0x71, 0xed, 0x16, 0xac, 0xdb, 0x21, 0x88, 0x06, 0xac, 0x11, 0x6a, 0xf7, 0xfc, 0xcc,
	0xba, 0x5d, 0x82, 0xa8, 0xc7, 0x1a, 0xc1, 0xfc, 0x59, 0x35, 0x8b, 0x81, 0x20, 0xba, 0xcf, 0x1c,
	0x65, 0xed, 0xd7, 0xcb, 0x7b, 0x04, 0x51, 0xcc, 0x1c, 0x05, 0x13, 0xe8, 0xdd, 0xa4, 0x2b, 0x3e,
	0xaf, 0x1b, 0xf4, 0x09, 0xa2, 0x6d, 0xe6, 0x4a, 0x8f, 0x44, 0xdd, 0x62, 0x87, 0x20, 0xda, 0x62,
	0xae, 0x84, 0x87, 0xd0, 0x57, 0x6e, 0x93, 0x5d, 0x82, 0xe8, 0x1e, 0x7b, 0xa2, 0x35, 0x4c, 0xdd,
	0x66, 0x8f, 0x20, 0x1a, 0xb2, 0x27, 0x1a, 0x3e, 0x86, 0x9e, 0xd2, 0xc5, 0x37, 0x75, 0x9b, 0xe6,
	0x39, 0x9f, 0x47, 0xa1, 0xb9, 0xcd, 0x89, 0xf7, 0x30, 0x41, 0x0c, 0x94, 0x2e, 0xae, 0x2b, 0xd9,
	0x50, 0x69, 0xa6, 0x1f, 0xa9, 0x7d, 0x73, 0x91, 0x35, 0x95, 0x66, 0xba, 0xa6, 0x86, 0x7f, 0x3c,
	0x68, 0x4d, 0xb9, 0xd2, 0x7c, 0x8e, 0x4f, 0xa0, 0xa5, 0x6c, 0x4c, 0x6c, 0x3e, 0x7a, 0x63, 0x1c,
	0x37, 0x39, 0x8c, 0xab, 0x00, 0xb1, 0x9a, 0xc0, 0xaf, 0xa0, 0x3d, 0xe7, 0x3a, 0x49, 0x85, 0xb2,
	0x77, 0xdc, 0x1b, 0x3f, 0x8f, 0xab, 0xa0, 0xc6, 0xeb, 0xa0, 0xc6, 0xd7, 0x36, 0xa8, 0x6c, 0xcd,
	0xe1, 0x63, 0xf0, 0x4d, 0x0e, 0x6c, 0xd4, 0x76, 0xc7, 0xa1, 0xdb, 0xdc, 0x4c, 0xc5, 0xac, 0x8b,
	0x5f, 0x42, 0xc0, 0xef, 0x72, 0x5d, 0xda, 0xe0, 0xf5, 0xc6, 0xfb, 0x2e, 0x76, 0x61, 0x0c, 0x56,
	0xf9, 0xf8, 0x00, 0x02, 0xfb, 0x66, 0x22, 0x9f, 0x78, 0xb4, 0xcb, 0xaa, 0x0f, 0xb3, 0xaf, 0x6a,
	0xe0, 0xc2, 0xa6, 0x6b, 0xe3, 0x10, 0xd5, 0x21, 0xac, 0x39, 0x4c, 0x00, 0xad, 0xa2, 0x36, 0xf1,
	0xfe, 0x33, 0x31, 0x5a, 0xe1, 0x37, 0xd0, 0x4e, 0xb3, 0x54, 0xa7, 0x89, 0x88, 0xba, 0x96, 0x3b,
	0x72, 0xb9, 0xea, 0xf4, 0xe2, 0x8f, 0x15, 0x71, 0x91, 0xe9, 0xa2, 0x64, 0x6b, 0x1e, 0xbf, 0x83,
	0xd0, 0x6e, 0x77, 0x2a, 0xb3, 0xe9, 0x52, 0x88, 0x64, 0x26, 0xb8, 0x0d, 0xdf, 0xa6, 0xc9, 0x26,
	0xfe, 0xaf, 0xdf, 0x47, 0x5b, 0xec, 0x9f, 0x05, 0x83, 0x29, 0xf4, 0xdd, 0xee, 0xe6, 0xcd, 0xde,
	0xf2, 0x72, 0xfd, 0x8a, 0x6f, 0x79, 0x89, 0x29, 0x04, 0xf7, 0x89, 0x58, 0x72, 0x7b, 0xb8, 0x9b,
	0xe7, 0xa8, 0x80, 0xb7, 0xdb, 0xaf, 0xd1, 0xb0, 0x0d, 0x81, 0xfd, 0xe1, 0xc9, 0x21, 0xf8, 0x36,
	0x50, 0x5d, 0x08, 0x3e, 0x5c, 0x5c, 0x5e, 0x5e, 0x85, 0x5b, 0xa6, 0xfc, 0x72, 0xc5, 0x2e, 0xdf,
	0x87, 0x68, 0xd2, 0x79, 0x98, 0xa0, 0xaf, 0x5e, 0x92, 0xa7, 0xb3, 0x96, 0xbd, 0xd4, 0xd3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x6d, 0x6d, 0x8e, 0xb0, 0x04, 0x00, 0x00,
}
