package golang

import (
	"bytes"
	"errors"
	"text/template"


	pgs "github.com/lyft/protoc-gen-star"
)

func (fns goSharedFuncs) render(field pgs.Field) (string, error) {
	var tpl *template.Template
	switch field.Type().ProtoType() {
	case pgs.DoubleT:
		tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	case pgs.BoolT:
		tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	case pgs.BytesT:
	// 	tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	// case pgs.EnumT:
	// 	tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	case pgs.Fixed32T:
		tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	case pgs.Fixed64T:
		tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	case pgs.Int32T:
		tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	case pgs.Int64T:
		tpl = template.Must(fns.tpl.New("").Parse(primitiveTmpl))
	case pgs.MessageT:
		tpl = template.Must(fns.tpl.New("").Parse(messageTpl))
	// case pgs.StringT:
	// 	tpl = template.Must(tplmessageTpl))
	default:
		return "", errors.New("unknown type")
	}
	var b bytes.Buffer
	err := tpl.Execute(&b, field)
	return b.String(), err
}
