package clone

import (
	"bytes"
	"errors"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

type Value struct {
	Name           string
	TargetName     string
	InnerTemplates struct {
		Value string
	}
}

func (fns goSharedFuncs) render(field pgs.Field) (string, error) {
	var tpl *template.Template

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
			tpl = template.Must(fns.tpl.New("message").Parse(messageTpl))
		default:
			return "", errors.New("unknown type")
		}
	}

	var b bytes.Buffer
	err := tpl.Execute(&b, Value{
		Name:       fns.accessor(field),
		TargetName: fns.targetAccessor(field),
	})
	return b.String(), err
}

func (fns goSharedFuncs) renderMap(field pgs.Field) (string, error) {
	var b bytes.Buffer
	valueTemplate, err := fns.simpleRender(field.Type().Element(), "v", fns.targetAccessor(field)+"[k]")
	if err != nil {
		return "", err
	}
	values := Value{
		Name:       fns.accessor(field),
		TargetName: fns.targetAccessor(field),
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
	innerTemplate, err := fns.simpleRender(field.Type().Element(), "v", fns.targetAccessor(field)+"[idx]")
	if err != nil {
		return "", err
	}
	values := Value{
		Name:       fns.accessor(field),
		TargetName: fns.targetAccessor(field),
		InnerTemplates: struct {
			Value string
		}{Value: innerTemplate},
	}
	outerTpl := template.Must(fns.tpl.New("map").Parse(repeatedTpl))
	err = outerTpl.Execute(&b, values)
	return b.String(), err
}

func (fns goSharedFuncs) simpleRender(field pgs.FieldTypeElem, valueName, targetName string) (string, error) {
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
	err := tpl.Execute(&b, Value{Name: valueName, TargetName: targetName})
	return b.String(), err
}
