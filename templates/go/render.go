package golang

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	pgs "github.com/lyft/protoc-gen-star"
)

type Value struct {
	Name           string
	Hasher         string
	Nullable       string
	InnerTemplates struct {
		Key   string
		Value string
	}
}

func (fns goSharedFuncs) render(field pgs.Field) (string, error) {
	var tpl *template.Template
	var nullable string

	if field.Type().ProtoType().IsNumeric() ||
		field.Type().ProtoType() == pgs.BoolT ||
		field.Type().IsEnum() {

		tpl = template.Must(fns.tpl.New("primitive").Parse(primitiveTmpl))
	} else if field.Type().IsMap() {
		return fns.renderMap(field)
	} else if field.Type().IsRepeated() {
		return fns.renderRepeated(field)
	} else {
		switch field.Type().ProtoType() {
		case pgs.BytesT:
			tpl = template.Must(fns.tpl.New("bytes").Parse(bytesTpl))
		case pgs.StringT:
			tpl = template.Must(fns.tpl.New("string").Parse(stringTpl))
		case pgs.MessageT:
			tpl = template.Must(fns.tpl.New("message").Parse(messageTpl))
		default:
			return "", errors.New("unknown type")
		}
	}

	var b bytes.Buffer
	err := tpl.Execute(&b, Value{
		Name:     fns.accessor(field),
		Hasher:   "hasher",
		Nullable: nullable,
	})
	return b.String(), err
}

func isNullable(options []*descriptor.UninterpretedOption) bool {
	for range options {
		return true
	}
	return false
}

func (fns goSharedFuncs) renderMap(field pgs.Field) (string, error) {

	var b bytes.Buffer
	valueTemplate, err := fns.simpleRender(field.Type().Element(), "v", "innerHash")
	if err != nil {
		return "", err
	}
	keyTemplate, err := fns.simpleRender(field.Type().Key(), "k", "innerHash")
	if err != nil {
		return "", err
	}
	values := Value{
		Name:   fns.accessor(field),
		Hasher: "innerHash",
		InnerTemplates: struct {
			Key   string
			Value string
		}{Value: valueTemplate, Key: keyTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("repeated").Parse(mapTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) renderRepeated(field pgs.Field) (string, error) {
	var b bytes.Buffer
	innerTemplate, err := fns.simpleRender(field.Type().Element(), "v", "hasher")
	if err != nil {
		return "", err
	}
	values := Value{
		Name:   fns.accessor(field),
		Hasher: "innerHash",
		InnerTemplates: struct {
			Key   string
			Value string
		}{Value: innerTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("map").Parse(repeatedTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) simpleRender(field pgs.FieldTypeElem, valueName, hasherName string) (string, error) {
	var tpl *template.Template
	if field.ProtoType().IsNumeric() ||
		field.ProtoType() == pgs.BoolT ||
		field.IsEnum() {
		tpl = template.Must(fns.tpl.New("primitive").Parse(primitiveTmpl))
	} else {
		switch field.ProtoType() {
		case pgs.BytesT:
			tpl = template.Must(fns.tpl.New("bytes").Parse(bytesTpl))
		case pgs.StringT:
			tpl = template.Must(fns.tpl.New("string").Parse(stringTpl))
		case pgs.MessageT:
			tpl = template.Must(fns.tpl.New("message").Parse(messageTpl))
		default:
			return "", errors.New("unknown type")
		}
	}

	var b bytes.Buffer
	err := tpl.Execute(&b, Value{Name: valueName, Hasher: hasherName})
	return b.String(), err
}
