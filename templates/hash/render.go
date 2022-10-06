package hash

import (
	"bytes"
	"errors"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/solo-io/protoc-gen-ext/extproto"
)

type Value struct {
	FieldAccessor  string
	FieldName      string
	Hasher         string
	InnerTemplates struct {
		Key   string
		Value string
	}
}

func (fns goSharedFuncs) render(field pgs.Field) (string, error) {
	var tpl *template.Template

	// check if skip hash is set on a given field
	var skipHash bool
	_, err := field.Extension(extproto.E_SkipHashing, &skipHash)
	if err != nil {
		return "", err
	}
	if skipHash {
		return "", nil
	}
	if field.Type().IsRepeated() {
		return fns.renderRepeated(field)
	} else if field.Type().ProtoType() == pgs.BoolT {
		tpl = template.Must(fns.tpl.New("primitiveBoolTmpl").Parse(primitiveBoolTmpl))
	} else if field.Type().ProtoType().IsInt() ||
		field.Type().IsEnum() {
		tpl = template.Must(fns.tpl.New("primitiveIntTmpl").Parse(primitiveIntTmpl))
	} else if field.Type().ProtoType().IsNumeric() {
		tpl = template.Must(fns.tpl.New("primitiveFloatTmpl").Parse(primitiveFloatTmpl))
	} else if field.Type().IsMap() {
		return fns.renderMap(field)
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
	err = tpl.Execute(&b, Value{
		FieldAccessor: fns.accessor(field),
		FieldName:     fns.fieldName(field),
		Hasher:        "hasher",
	})
	return b.String(), err
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
		FieldAccessor: fns.accessor(field),
		Hasher:        "innerHash",
		InnerTemplates: struct {
			Key   string
			Value string
		}{Value: valueTemplate, Key: keyTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("map").Parse(mapTpl))
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
		FieldAccessor: fns.accessor(field),
		Hasher:        "innerHash",
		InnerTemplates: struct {
			Key   string
			Value string
		}{Value: innerTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("repeated").Parse(repeatedTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) simpleRender(field pgs.FieldTypeElem, valueName, hasherName string) (string, error) {
	var tpl *template.Template
	if field.ProtoType().IsNumeric() ||
		field.ProtoType() == pgs.BoolT ||
		field.IsEnum() {

		if field.ProtoType() == pgs.BoolT {
			tpl = template.Must(fns.tpl.New("primitiveBoolTmpl").Parse(primitiveBoolTmpl))
		} else if field.ProtoType().IsInt() || field.IsEnum() {
			tpl = template.Must(fns.tpl.New("primitiveIntTmpl").Parse(primitiveIntTmpl))
		} else if field.ProtoType().IsNumeric() {
			tpl = template.Must(fns.tpl.New("primitiveFloatTmpl").Parse(primitiveFloatTmpl))
		} else {
			// should never get here
			tpl = template.Must(fns.tpl.New("primitive").Parse(primitiveTmpl))
		}
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
	err := tpl.Execute(&b, Value{
		FieldAccessor: valueName,
		Hasher:        hasherName,
	})
	return b.String(), err
}
