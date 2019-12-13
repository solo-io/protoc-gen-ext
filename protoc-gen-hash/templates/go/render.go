package golang

import (
	"bytes"
	"errors"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

type Value struct {
	Name           string
	Hasher         string
	InnerTemplates struct {
		Key   string
		Value string
	}
}

func (fns goSharedFuncs) render(field pgs.Field) (string, error) {
	var tpl *template.Template

	if field.Type().ProtoType().IsNumeric() ||
		field.Type().ProtoType() == pgs.BoolT ||
		field.Type().IsEnum() {

		tpl = template.Must(fns.tpl.New("primitive").Parse(primitiveTmpl))
	} else if field.Type().IsRepeated() {
		if field.Type().IsMap() {
			return fns.renderMap(field)
		}
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
	err := tpl.Execute(&b, Value{Name: fns.accessor(field), Hasher: "hasher"})
	return b.String(), err
}


func (fns goSharedFuncs) renderMap(field pgs.Field) (string, error) {

	var b bytes.Buffer
	innerTemplate, err := fns.simpleRender(field, "v", "hasher")
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


func (fns goSharedFuncs) renderRepeated(field pgs.Field) (string, error) {
	var b bytes.Buffer
	innerTemplate, err := fns.simpleRender(field, "v", "hasher")
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


func (fns goSharedFuncs) simpleRender(field pgs.Field, valueName, hasherName string) (string, error) {

	var tpl *template.Template
	if field.Type().ProtoType().IsNumeric() ||
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
			tpl = template.Must(fns.tpl.New("message").Parse(messageTpl))
		default:
			return "", errors.New("unknown type")
		}
	}

	var b bytes.Buffer
	err := tpl.Execute(&b, Value{Name: valueName, Hasher: hasherName})
	return b.String(), err
}
