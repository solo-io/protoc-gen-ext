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
	UniqueHash bool
}

func (fns goSharedFuncs) render(field pgs.Field, hashUnique bool) (string, error) {
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

	if field.Type().ProtoType().IsNumeric() ||
		field.Type().ProtoType() == pgs.BoolT ||
		field.Type().IsEnum() {

		tpl = template.Must(fns.tpl.New("primitive").Parse(primitiveTmpl))
	} else if field.Type().IsMap() {
		return fns.renderMap(field, hashUnique)
	} else if field.Type().IsRepeated() {
		return fns.renderRepeated(field, hashUnique)
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
		UniqueHash:    hashUnique,
	})
	return b.String(), err
}

func (fns goSharedFuncs) renderMap(field pgs.Field, uniqueHash bool) (string, error) {
	var b bytes.Buffer
	valueTemplate, err := fns.simpleRender(field.Type().Element(), "v", "innerHash", uniqueHash)
	if err != nil {
		return "", err
	}
	keyTemplate, err := fns.simpleRender(field.Type().Key(), "k", "innerHash", uniqueHash)
	if err != nil {
		return "", err
	}
	values := Value{
		FieldAccessor: fns.accessor(field),
		FieldName:     fns.fieldName(field),
		Hasher:        "innerHash",
		UniqueHash:    uniqueHash,
		InnerTemplates: struct {
			Key   string
			Value string
		}{Value: valueTemplate, Key: keyTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("repeated").Parse(mapTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) renderRepeated(field pgs.Field, hashUnique bool) (string, error) {
	var b bytes.Buffer
	innerTemplate, err := fns.simpleRender(field.Type().Element(), "v", "hasher", hashUnique)
	if err != nil {
		return "", err
	}
	values := Value{
		FieldName:     fns.fieldName(field),
		FieldAccessor: fns.accessor(field),
		Hasher:        "innerHash",
		UniqueHash:    hashUnique,
		InnerTemplates: struct {
			Key   string
			Value string
		}{Value: innerTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("map").Parse(repeatedTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) simpleRender(field pgs.FieldTypeElem, valueName, hasherName string, uniqueHash bool) (string, error) {
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
	err := tpl.Execute(&b, Value{
		FieldName:     valueName,
		FieldAccessor: valueName,
		Hasher:        hasherName,
		UniqueHash:    uniqueHash,
	})
	return b.String(), err
}
