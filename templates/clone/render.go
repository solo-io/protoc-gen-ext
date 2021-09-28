package clone

import (
	"bytes"
	"errors"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/solo-io/protoc-gen-ext/extproto"
)

type Value struct {
	Name           string
	TargetName     string
	TypeName       string
	Field          pgs.Field
	OneOfInterface string
	InnerTemplates struct {
		Value string
	}
	ProtoLibrary string
}

func (fns goSharedFuncs) render(field pgs.Field) (string, error) {
	var tpl *template.Template
	var typeName string

	protoLibrary := "proto" // By default use the Google proto library

	if field.Type().IsRepeated() {
		return fns.renderRepeated(field)
	} else if field.Type().IsMap() {
		return fns.renderMap(field)
	} else if field.Type().ProtoType().IsNumeric() ||
		field.Type().ProtoType() == pgs.BoolT ||
		field.Type().IsEnum() {

		tpl = template.Must(fns.tpl.New("primitive").Parse(primitiveTmpl))
	} else {
		switch field.Type().ProtoType() {
		case pgs.BytesT:
			tpl = template.Must(fns.tpl.New("bytes").Parse(bytesTpl))
		case pgs.StringT:
			tpl = template.Must(fns.tpl.New("string").Parse(stringTpl))
		case pgs.MessageT:
			var gogoClone bool
			if _, err := field.Extension(extproto.E_GogoClone, &gogoClone); err != nil {
				return "", err
			}

			if gogoClone {
				protoLibrary = "gogo_proto"
			}

			tpl = template.Must(fns.tpl.New("message").Parse(messageTpl))
			typeName = fns.typeName(field)
		default:
			return "", errors.New("unknown type")
		}
	}

	var b bytes.Buffer
	err := tpl.Execute(&b, Value{
		Name:         fns.accessor(field),
		TargetName:   fns.targetAccessor(field),
		TypeName:     typeName,
		Field:        field,
		ProtoLibrary: protoLibrary,
	})
	return b.String(), err
}

func (fns goSharedFuncs) oneofRender(field pgs.Field, oneofInterface pgs.Name) (string, error) {
	var tpl *template.Template
	var typeName string

	protoLibrary := "proto" // By default use the Google proto library

	if field.Type().IsRepeated() {
		return fns.renderRepeated(field)
	} else if field.Type().IsMap() {
		return fns.renderMap(field)
	} else if field.Type().ProtoType().IsNumeric() ||
		field.Type().ProtoType() == pgs.BoolT ||
		field.Type().IsEnum() {
		tpl = template.Must(fns.tpl.New("primitive").Parse(oneofPrimitiveTmpl))
	} else {
		switch field.Type().ProtoType() {
		case pgs.BytesT:
			tpl = template.Must(fns.tpl.New("bytes").Parse(oneofBytesTpl))
		case pgs.StringT:
			tpl = template.Must(fns.tpl.New("string").Parse(oneofStringTpl))
		case pgs.MessageT:
			var gogoClone bool
			if _, err := field.Extension(extproto.E_GogoClone, &gogoClone); err != nil {
				return "", err
			}

			if gogoClone {
				protoLibrary = "gogo_proto"
			}

			tpl = template.Must(fns.tpl.New("message").Parse(oneofMessageTpl))
			typeName = fns.typeName(field)
		default:
			return "", errors.New("unknown type")
		}
	}

	var b bytes.Buffer
	err := tpl.Execute(&b, Value{
		Name:           fns.accessor(field),
		TargetName:     fns.targetAccessor(field),
		TypeName:       typeName,
		Field:          field,
		OneOfInterface: oneofInterface.String(),
		ProtoLibrary:   protoLibrary,
	})
	return b.String(), err
}

func (fns goSharedFuncs) renderMap(field pgs.Field) (string, error) {
	var b bytes.Buffer
	valueTemplate, err := fns.simpleRender(
		field,
		field.Type().Element(),
		"v", fns.targetAccessor(field)+"[k]",
	)
	if err != nil {
		return "", err
	}
	values := Value{
		Name:       fns.accessor(field),
		TargetName: fns.targetAccessor(field),
		TypeName:   fns.typeName(field),
		Field:      field,
		InnerTemplates: struct {
			Value string
		}{Value: valueTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("repeated").Parse(mapTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) renderRepeated(field pgs.Field) (string, error) {
	var b bytes.Buffer
	innerTemplate, err := fns.simpleRender(
		field,
		field.Type().Element(),
		"v",
		fns.targetAccessor(field)+"[idx]",
	)
	if err != nil {
		return "", err
	}
	values := Value{
		Name:       fns.accessor(field),
		TargetName: fns.targetAccessor(field),
		TypeName:   fns.typeName(field),
		Field:      field,
		InnerTemplates: struct {
			Value string
		}{Value: innerTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("map").Parse(repeatedTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) simpleRender(
	field pgs.Field,
	typeElem pgs.FieldTypeElem,
	valueName, targetName string,
) (string, error) {
	var tpl *template.Template
	var typeName string

	protoLibrary := "proto" // By default use the Google proto library

	if typeElem.ProtoType().IsNumeric() ||
		typeElem.ProtoType() == pgs.BoolT ||
		typeElem.IsEnum() {
		tpl = template.Must(fns.tpl.New("primitive").Parse(primitiveTmpl))
	} else {
		switch typeElem.ProtoType() {
		case pgs.BytesT:
			tpl = template.Must(fns.tpl.New("bytes").Parse(bytesTpl))
		case pgs.StringT:
			tpl = template.Must(fns.tpl.New("string").Parse(stringTpl))
		case pgs.MessageT:
			var gogoClone bool
			if _, err := field.Extension(extproto.E_GogoClone, &gogoClone); err != nil {
				return "", err
			}

			if gogoClone {
				protoLibrary = "gogo_proto"
			}

			tpl = template.Must(fns.tpl.New("message").Parse(messageTpl))
			typeName = fns.entityTypeName(field, typeElem)
		default:
			return "", errors.New("unknown type")
		}
	}
	var b bytes.Buffer
	err := tpl.Execute(&b, Value{
		Name:         valueName,
		TargetName:   targetName,
		TypeName:     typeName,
		Field:        field,
		ProtoLibrary: protoLibrary,
	})
	return b.String(), err
}
