// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: tests/api/hello.proto

package api

// Merge non-nil fields from overrides into m
func (m *Simple) Merge(overrides *Simple) {
	if m == nil || overrides == nil {
		return
	}

	m.Str = overrides.Str

	m.Byt = overrides.Byt

	m.TestUint32 = overrides.TestUint32

	m.TestUint64 = overrides.TestUint64

	m.TestBool = overrides.TestBool

	m.DoubleTest = overrides.DoubleTest

	m.FloatTest = overrides.FloatTest

	m.Int32Test = overrides.Int32Test

	m.Int64Test = overrides.Int64Test

	m.Sint32Test = overrides.Sint32Test

	m.Sint64Test = overrides.Sint64Test

	m.Fixed32Test = overrides.Fixed32Test

	m.Fixed64Test = overrides.Fixed64Test

	m.Sfixed32Test = overrides.Sfixed32Test

	m.Sfixed64Test = overrides.Sfixed64Test

	m.StrSkipped = overrides.StrSkipped

	m.IntSkipped = overrides.IntSkipped

}

// Merge non-nil fields from overrides into m
func (m *Nested) Merge(overrides *Nested) {
	if m == nil || overrides == nil {
		return
	}

	if overrides.Simple != nil {
		m.Simple = overrides.Simple
	}

	if overrides.OtherSimple != nil {
		m.OtherSimple = overrides.OtherSimple
	}

	m.Test = overrides.Test

	if overrides.Empty != nil {
		m.Empty = overrides.Empty
	}

	m.Hello = overrides.Hello

	if overrides.Details != nil {
		m.Details = overrides.Details
	}

	if overrides.Skipper != nil {
		m.Skipper = overrides.Skipper
	}

	if overrides.X != nil {
		m.X = overrides.X
	}

	if overrides.Initial != nil {
		m.Initial = overrides.Initial
	}

	if overrides.SimpleMap != nil {
		m.SimpleMap = overrides.SimpleMap
	}

	m.RepeatedPrimitive = overrides.RepeatedPrimitive

	if overrides.TestOneOf != nil {
		m.TestOneOf = overrides.TestOneOf
	}

}

// Merge non-nil fields from overrides into m
func (m *Empty) Merge(overrides *Empty) {
	if m == nil || overrides == nil {
		return
	}

}

// Merge non-nil fields from overrides into m
func (m *NestedEmpty) Merge(overrides *NestedEmpty) {
	if m == nil || overrides == nil {
		return
	}

	if overrides.Nested != nil {
		m.Nested = overrides.Nested
	}

}
