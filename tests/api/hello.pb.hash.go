// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: tests/api/hello.proto

package api

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
// Deprecated due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Use the HashUnique function instead.
func (m *Simple) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Simple")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetStr())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetByt()); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint32())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint64())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestBool())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDoubleTest())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFloatTest())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt32Test())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt64Test())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint32Test())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint64Test())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed32Test())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed64Test())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed32Test())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed64Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetStrOptional())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint32Optional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint64Optional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestBoolOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDoubleTestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFloatTestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt32TestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt64TestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint32TestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint64TestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed32TestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed64TestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed32TestOptional())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed64TestOptional())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Simple) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Simple")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Str")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetStr())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Byt")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetByt()); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TestUint32")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint32())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TestUint64")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint64())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TestBool")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestBool())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DoubleTest")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDoubleTest())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("FloatTest")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFloatTest())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Int32Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt32Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Int64Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt64Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sint32Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint32Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sint64Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint64Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Fixed32Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed32Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Fixed64Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed64Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sfixed32Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed32Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sfixed64Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed64Test())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("StrOptional")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetStrOptional())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TestUint32Optional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint32Optional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TestUint64Optional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestUint64Optional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TestBoolOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTestBoolOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DoubleTestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDoubleTestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("FloatTestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFloatTestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Int32TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt32TestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Int64TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInt64TestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sint32TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint32TestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sint64TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSint64TestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Fixed32TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed32TestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Fixed64TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetFixed64TestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sfixed32TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed32TestOptional())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Sfixed64TestOptional")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSfixed64TestOptional())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
// Deprecated due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Use the HashUnique function instead.
func (m *Nested) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Nested")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSimple()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Simple")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSimple(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Simple")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetOtherSimple()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OtherSimple")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOtherSimple(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OtherSimple")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTest())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEmpty()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Empty")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmpty(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Empty")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetHello() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetDetails()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Details")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDetails(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Details")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetX() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetInitial() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetSimpleMap() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRepeatedPrimitive())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetRepeatedExternal() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetMapExternal() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	switch m.TestOneOf.(type) {

	case *Nested_EmptyOneOf:

		if h, ok := interface{}(m.GetEmptyOneOf()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("EmptyOneOf")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetEmptyOneOf(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("EmptyOneOf")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Nested_NestedOneOf:

		if h, ok := interface{}(m.GetNestedOneOf()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("NestedOneOf")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetNestedOneOf(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("NestedOneOf")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Nested_PrimitiveOneOf:

		if _, err = hasher.Write([]byte(m.GetPrimitiveOneOf())); err != nil {
			return 0, err
		}

	case *Nested_BytesOneOf:

		if _, err = hasher.Write(m.GetBytesOneOf()); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Nested) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Nested")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSimple()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Simple")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSimple(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Simple")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetOtherSimple()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OtherSimple")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOtherSimple(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OtherSimple")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Test")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTest())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEmpty()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Empty")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmpty(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Empty")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Hello")); err != nil {
		return 0, err
	}
	for i, v := range m.GetHello() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetDetails()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Details")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDetails(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Details")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("X")); err != nil {
		return 0, err
	}
	for i, v := range m.GetX() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetInitial() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetSimpleMap() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte("v")); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(v)); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("RepeatedPrimitive")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRepeatedPrimitive())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RepeatedExternal")); err != nil {
		return 0, err
	}
	for i, v := range m.GetRepeatedExternal() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetMapExternal() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	switch m.TestOneOf.(type) {

	case *Nested_EmptyOneOf:

		if h, ok := interface{}(m.GetEmptyOneOf()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("EmptyOneOf")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetEmptyOneOf(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("EmptyOneOf")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Nested_NestedOneOf:

		if h, ok := interface{}(m.GetNestedOneOf()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("NestedOneOf")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetNestedOneOf(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("NestedOneOf")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Nested_PrimitiveOneOf:

		if _, err = hasher.Write([]byte("PrimitiveOneOf")); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte(m.GetPrimitiveOneOf())); err != nil {
			return 0, err
		}

	case *Nested_BytesOneOf:

		if _, err = hasher.Write([]byte("BytesOneOf")); err != nil {
			return 0, err
		}

		if _, err = hasher.Write(m.GetBytesOneOf()); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
// Deprecated due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Use the HashUnique function instead.
func (m *Empty) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Empty")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Empty) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Empty")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
// Deprecated due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Use the HashUnique function instead.
func (m *NestedEmpty) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.NestedEmpty")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetNested()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Nested")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetNested(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Nested")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *NestedEmpty) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.NestedEmpty")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetNested()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Nested")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetNested(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Nested")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
// Deprecated due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Use the HashUnique function instead.
func (m *MultipleStrings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.MultipleStrings")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetS1())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetS2())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *MultipleStrings) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.MultipleStrings")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("S1")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetS1())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("S2")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetS2())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
// Deprecated due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Use the HashUnique function instead.
func (m *Repeated) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Repeated")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFirst() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetSecond() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Repeated) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.type.github.com/solo-io/protoc-gen-ext/tests/api.Repeated")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("First")); err != nil {
		return 0, err
	}
	for i, v := range m.GetFirst() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("Second")); err != nil {
		return 0, err
	}
	for i, v := range m.GetSecond() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
