package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/protoc-gen-ext/tests/api"
)

var _ = Describe("merge", func() {
	Context("object with oneofs", func() {
		It("will merge non-nil fields", func() {

			baseObj := &api.Nested{
				Simple: &api.Simple{
					Str: "hello",
				},
			}
			overrides := &api.Nested{
				Hello: []string{"jello"},
				TestOneOf: &api.Nested_EmptyOneOf{
					EmptyOneOf: &api.Empty{},
				},
			}

			baseObj.Merge(overrides)

			Expect(baseObj).To(Equal(&api.Nested{
				Simple: &api.Simple{
					Str: "hello",
				},
				Hello: []string{"jello"},
				TestOneOf: &api.Nested_EmptyOneOf{
					EmptyOneOf: &api.Empty{},
				},
			}))
		})
	})
})
