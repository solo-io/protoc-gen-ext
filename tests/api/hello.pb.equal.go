// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: tests/api/hello.proto

package api

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Simple) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Simple)
	if !ok {
		that2, ok := that.(Simple)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetStr(), target.GetStr()) != 0 {
		return false
	}

	if bytes.Compare(m.GetByt(), target.GetByt()) != 0 {
		return false
	}

	if m.GetTestUint32() != target.GetTestUint32() {
		return false
	}

	if m.GetTestUint64() != target.GetTestUint64() {
		return false
	}

	if m.GetTestBool() != target.GetTestBool() {
		return false
	}

	if m.GetDoubleTest() != target.GetDoubleTest() {
		return false
	}

	if m.GetFloatTest() != target.GetFloatTest() {
		return false
	}

	if m.GetInt32Test() != target.GetInt32Test() {
		return false
	}

	if m.GetInt64Test() != target.GetInt64Test() {
		return false
	}

	if m.GetSint32Test() != target.GetSint32Test() {
		return false
	}

	if m.GetSint64Test() != target.GetSint64Test() {
		return false
	}

	if m.GetFixed32Test() != target.GetFixed32Test() {
		return false
	}

	if m.GetFixed64Test() != target.GetFixed64Test() {
		return false
	}

	if m.GetSfixed32Test() != target.GetSfixed32Test() {
		return false
	}

	if m.GetSfixed64Test() != target.GetSfixed64Test() {
		return false
	}

	if strings.Compare(m.GetStrSkipped(), target.GetStrSkipped()) != 0 {
		return false
	}

	if m.GetIntSkipped() != target.GetIntSkipped() {
		return false
	}

	return true
}

// Equal function
func (m *Nested) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Nested)
	if !ok {
		that2, ok := that.(Nested)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSimple()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSimple()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSimple(), target.GetSimple()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOtherSimple()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOtherSimple()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOtherSimple(), target.GetOtherSimple()) {
			return false
		}
	}

	if m.GetTest() != target.GetTest() {
		return false
	}

	if h, ok := interface{}(m.GetEmpty()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEmpty()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEmpty(), target.GetEmpty()) {
			return false
		}
	}

	if len(m.GetHello()) != len(target.GetHello()) {
		return false
	}
	for idx, v := range m.GetHello() {

		if strings.Compare(v, target.GetHello()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDetails(), target.GetDetails()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSkipper()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSkipper()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSkipper(), target.GetSkipper()) {
			return false
		}
	}

	if len(m.GetX()) != len(target.GetX()) {
		return false
	}
	for idx, v := range m.GetX() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetX()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetX()[idx]) {
				return false
			}
		}

	}

	if len(m.GetInitial()) != len(target.GetInitial()) {
		return false
	}
	for k, v := range m.GetInitial() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetInitial()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetInitial()[k]) {
				return false
			}
		}

	}

	if len(m.GetSimpleMap()) != len(target.GetSimpleMap()) {
		return false
	}
	for k, v := range m.GetSimpleMap() {

		if strings.Compare(v, target.GetSimpleMap()[k]) != 0 {
			return false
		}

	}

	if len(m.GetRepeatedPrimitive()) != len(target.GetRepeatedPrimitive()) {
		return false
	}
	for idx, v := range m.GetRepeatedPrimitive() {

		if v != target.GetRepeatedPrimitive()[idx] {
			return false
		}

	}

	switch m.TestOneOf.(type) {

	case *Nested_EmptyOneOf:
		if _, ok := target.TestOneOf.(*Nested_EmptyOneOf); !ok {
			return false
		}

		if h, ok := interface{}(m.GetEmptyOneOf()).(equality.Equalizer); ok {
			if !h.Equal(target.GetEmptyOneOf()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetEmptyOneOf(), target.GetEmptyOneOf()) {
				return false
			}
		}

	case *Nested_NestedOneOf:
		if _, ok := target.TestOneOf.(*Nested_NestedOneOf); !ok {
			return false
		}

		if h, ok := interface{}(m.GetNestedOneOf()).(equality.Equalizer); ok {
			if !h.Equal(target.GetNestedOneOf()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetNestedOneOf(), target.GetNestedOneOf()) {
				return false
			}
		}

	case *Nested_PrimitiveOneOf:
		if _, ok := target.TestOneOf.(*Nested_PrimitiveOneOf); !ok {
			return false
		}

		if strings.Compare(m.GetPrimitiveOneOf(), target.GetPrimitiveOneOf()) != 0 {
			return false
		}

	case *Nested_BytesOneOf:
		if _, ok := target.TestOneOf.(*Nested_BytesOneOf); !ok {
			return false
		}

		if bytes.Compare(m.GetBytesOneOf(), target.GetBytesOneOf()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.TestOneOf != target.TestOneOf {
			return false
		}
	}

	return true
}

// Equal function
func (m *Empty) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Empty)
	if !ok {
		that2, ok := that.(Empty)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *NestedEmpty) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*NestedEmpty)
	if !ok {
		that2, ok := that.(NestedEmpty)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetNested()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNested()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNested(), target.GetNested()) {
			return false
		}
	}

	return true
}
