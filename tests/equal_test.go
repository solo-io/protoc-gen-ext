package tests

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/solo-io/protoc-gen-ext/tests/api"
)

var _ = Describe("equal", func() {

	DescribeTable(
		"simple",
		func(x *api.Simple, y interface{}, expected bool) {
			Expect(x.Equal(y)).To(Equal(expected))
		},
		Entry(
			"both nil (equal)",
			nil, nil, true,
		),
		Entry(
			"source non-nil, target nil (!equal)",
			&api.Simple{}, nil, false,
		),
		Entry(
			"source nil, target non-nil (!equal)",
			nil, &api.Simple{}, false,
		),
		Entry(
			"target non-pointer (equal)",
			&api.Simple{}, api.Simple{}, true,
		),
		Entry(
			"filled struct (equal)",
			&api.Simple{
				Byt: []byte("hello"),
			}, api.Simple{
				Byt: []byte("hello"),
			}, true,
		),
		Entry(
			"different bytes (unequal)",
			&api.Simple{
				Byt: []byte("hello"),
			}, api.Simple{
				Byt: []byte("world"),
			}, false,
		),
		Entry(
			"different string (unequal)",
			&api.Simple{
				Byt: []byte("hello"),
			}, api.Simple{
				Byt: []byte("world"),
			}, false,
		),
		Entry(
			"different string (unequal)",
			&api.Simple{
				Str: "hello",
			}, api.Simple{
				Str: "world",
			}, false,
		),
	)

	DescribeTable(
		"nested",
		func(x *api.Nested, y interface{}, expected bool) {
			Expect(x.Equal(y)).To(Equal(expected))
		},
		Entry(
			"both nil (equal)",
			nil, nil, true,
		),
		Entry(
			"source non-nil, target nil (!equal)",
			&api.Nested{}, nil, false,
		),
		Entry(
			"source nil, target non-nil (!equal)",
			nil, &api.Nested{}, false,
		),
		Entry(
			"target non-pointer (equal)",
			&api.Nested{}, api.Nested{}, true,
		),
		Entry(
			"foreign struct (equal)",
			&api.Nested{
				Details: &structpb.Struct{},
			}, api.Nested{
				Details: &structpb.Struct{},
			}, true,
		),
		Entry(
			"foreign struct (!equal)",
			&api.Nested{
				Details: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"hello": {},
					},
				},
			}, api.Nested{
				Details: &structpb.Struct{},
			}, false,
		),
		Entry(
			"map (equal)",
			&api.Nested{
				Initial: map[string]*api.Simple{
					"hello": {
						Str: "hello",
					},
				},
			}, api.Nested{
				Initial: map[string]*api.Simple{
					"hello": {
						Str: "hello",
					},
				},
			}, true,
		),
		Entry(
			"map (!equal)",
			&api.Nested{
				Initial: map[string]*api.Simple{
					"hello": {
						Str: "hello",
					},
				},
			}, api.Nested{
				Initial: map[string]*api.Simple{
					"world": {
						Str: "hello",
					},
				},
			}, false,
		),
		Entry(
			"slice (equal)",
			&api.Nested{
				X: []*api.Simple{
					{
						Str: "hello",
					},
				},
			}, api.Nested{
				X: []*api.Simple{
					{
						Str: "hello",
					},
				},
			}, true,
		),
		Entry(
			"slice (!equal)",
			&api.Nested{
				X: []*api.Simple{
					{
						Str: "hello",
					},
				},
			}, api.Nested{
				X: []*api.Simple{
					{
						Str: "world",
					},
				},
			}, false,
		),
	)
})
