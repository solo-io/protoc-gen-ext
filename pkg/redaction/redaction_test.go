package redaction_test

import (
	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/protoc-gen-ext/pkg/redaction"
	"github.com/solo-io/protoc-gen-ext/tests/api"
)

var _ = Describe("Redaction", func() {

	It("Will clear a redacted map", func() {
		obj := &api.Nested{
			Test: api.Test_WORLD,
			SimpleMap: map[string]string{
				"hello": "world",
			},
		}

		redaction.Redact(proto.MessageReflect(obj))

		Expect(obj.GetSimpleMap()).To(BeNil())
		Expect(obj.GetTest()).To(Equal(api.Test_WORLD))
	})

	It("Will clear a redacted list", func() {
		obj := &api.Nested{
			Test: api.Test_WORLD,
			X: []*api.Simple{
				{
					Str: "hello",
				},
			},
		}

		redaction.Redact(proto.MessageReflect(obj))

		Expect(obj.GetX()).To(BeNil())
		Expect(obj.GetTest()).To(Equal(api.Test_WORLD))
	})

	It("Will clear a redacted message", func() {
		obj := &api.Nested{
			Test: api.Test_WORLD,
			Empty: &api.Empty{},
		}

		redaction.Redact(proto.MessageReflect(obj))

		Expect(obj.GetEmpty()).To(BeNil())
		Expect(obj.GetTest()).To(Equal(api.Test_WORLD))
	})

	It("Will clear a redacted field", func() {
		obj := &api.Nested{
			Test: api.Test_WORLD,
			Simple: &api.Simple{
				Str: "will remain",
				Byt: []byte("to be deleted"),
			},
		}

		redaction.Redact(proto.MessageReflect(obj))

		Expect(obj.GetSimple().GetByt()).To(BeNil())
		Expect(obj.GetSimple().GetStr()).To(Equal("will remain"))
		Expect(obj.GetTest()).To(Equal(api.Test_WORLD))
	})

})
