package tests

import (
	"fmt"

	structpb "github.com/golang/protobuf/ptypes/struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/solo-io/protoc-gen-ext/tests/api"
	"google.golang.org/protobuf/proto"
)

var _ = Describe("clone", func() {

	Measure("it should do something hard efficiently", func(b Benchmarker) {
		externalStruct := &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"I'm a field": {
					Kind: &structpb.Value_StructValue{
						StructValue: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"another field": {
									Kind: &structpb.Value_StringValue{
										StringValue: "ashjedsah",
									},
								},
							},
						},
					},
				},
			},
		}
		test := &api.Nested{
			Simple: &api.Simple{
				Byt: []byte("hello"),
			},
			OtherSimple: &api.Simple{},
			Test:        50,
			Empty:       &api.Empty{},
			Hello:       []string{"one", "two", "three"},
			Details:     externalStruct,
			Skipper:     nil,
			X: []*api.Simple{
				{
					Str: "ajskldja",
				},
				{
					Byt: []byte("jaklsdjkals"),
				},
				{
					DoubleTest: 1231,
				},
			},
			Initial: map[string]*api.Simple{
				"hello": {
					Byt: []byte("stuff"),
				},
				"world": {
					Byt: []byte("things"),
				},
			},
			SimpleMap: map[string]string{
				"hello":     "world",
				"something": "test",
			},
			TestOneOf: &api.Nested_NestedOneOf{
				NestedOneOf: &api.NestedEmpty{
					Nested: &api.Nested{
						Simple: &api.Simple{},
					},
				},
			},
			RepeatedPrimitive: []uint64{1, 2, 3, 4},
			RepeatedExternal:  []*structpb.Struct{externalStruct, externalStruct},
			MapExternal: map[string]*structpb.Struct{
				"hello": externalStruct,
				"world": externalStruct,
			},
		}
		const times = 10000
		reflectionBased := b.Time(fmt.Sprintf("runtime of %d reflect-based hash calls", times), func() {
			for i := 0; i < times; i++ {
				proto.Clone(test)
			}
		})
		generated := b.Time(fmt.Sprintf("runtime of %d generated hash calls", times), func() {
			for i := 0; i < times; i++ {
				test.Clone()
			}
		})
		// divide by 1e3 to get time in micro seconds instead of nano seconds
		b.RecordValue("Runtime per reflection call in µ seconds", float64(int64(reflectionBased)/times)/1e3)
		b.RecordValue("Runtime per generated call in µ seconds", float64(int64(generated)/times)/1e3)

	}, 10)

	DescribeTable(
		"simple",
		func(x *api.Simple) {
			cloned := x.Clone()
			Expect(x.Equal(cloned)).To(BeTrue())
		},
		Entry(
			"nil",
			nil,
		),
		Entry(
			"empty",
			&api.Simple{},
		),
		Entry(
			"filled struct",
			&api.Simple{
				Str:          "world",
				Byt:          []byte("hello"),
				TestUint32:   1,
				TestUint64:   2,
				TestBool:     true,
				DoubleTest:   3,
				FloatTest:    4,
				Int32Test:    5,
				Int64Test:    6,
				Sint32Test:   7,
				Sint64Test:   8,
				Fixed32Test:  9,
				Fixed64Test:  10,
				Sfixed32Test: 11,
				Sfixed64Test: 12,
				StrSkipped:   "asda",
				IntSkipped:   13,
			},
		),
	)

	DescribeTable(
		"nested",
		func(x *api.Nested) {
			cloned := x.Clone()
			Expect(x.Equal(cloned)).To(BeTrue())
		},
		Entry(
			"non-nil oneOf",
			&api.Nested{
				TestOneOf: &api.Nested_NestedOneOf{
					NestedOneOf: &api.NestedEmpty{
						Nested: &api.Nested{
							Simple: &api.Simple{Str: "hello"},
						},
					},
				},
			},
		),
		Entry(
			"foreign struct",
			&api.Nested{
				Details: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"hello": {},
					},
				},
			},
		),
		Entry(
			"map",
			&api.Nested{
				Initial: map[string]*api.Simple{
					"hello": {
						Str: "hello",
					},
				},
			},
		),
		Entry(
			"slice",
			&api.Nested{
				X: []*api.Simple{
					{
						Str: "hello",
					},
				},
			},
		),
		Entry(
			"oneof message",
			&api.Nested{
				TestOneOf: &api.Nested_NestedOneOf{
					NestedOneOf: &api.NestedEmpty{
						Nested: &api.Nested{
							Simple: &api.Simple{},
							TestOneOf: &api.Nested_NestedOneOf{
								NestedOneOf: &api.NestedEmpty{
									Nested: &api.Nested{
										Empty: &api.Empty{},
									},
								},
							},
						},
					},
				},
			},
		),
		Entry(
			"oneof bytes",
			&api.Nested{
				TestOneOf: &api.Nested_BytesOneOf{
					BytesOneOf: []byte("hjsakda"),
				},
			},
		),
		Entry(
			"oneof primitive",
			&api.Nested{
				TestOneOf: &api.Nested_PrimitiveOneOf{
					PrimitiveOneOf: "hello",
				},
			},
		),
	)
})
