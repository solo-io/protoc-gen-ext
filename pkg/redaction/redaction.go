package redaction

import (
	"github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/extproto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Redact clears every sensitive field in pb.
func Redact(m protoreflect.Message) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		opts := fd.Options().(*descriptorpb.FieldOptions)

		switch typed := v.Interface().(type) {
		case protoreflect.Message:
			Redact(typed)
		case protoreflect.Map:
			typed.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
				if _, ok := value.Interface().(protoreflect.Message); ok {
					Redact(value.Message())
				}
				if _, ok := key.Interface().(protoreflect.Message); ok {
					Redact(key.Value().Message())
				}
				return true
			})
		case protoreflect.List:
			for i := 0; i < typed.Len(); i++ {
				if _, ok := typed.Get(i).Interface().(protoreflect.Message); ok {
					Redact(typed.Get(i).Message())
				}
			}
		}

		// Get extension from field
		ext, err := proto.GetExtension(opts, extproto.E_Sensitive)
		if err != nil {
			return true
		}
		// Check if equal to bool as expected
		extVal, ok := ext.(*bool)
		if !ok {
			return true
		}

		// If true clear field and move on
		if extVal != nil && *extVal == true {
			m.Clear(fd)
		}

		return true
	})
}
