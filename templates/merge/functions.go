package merge

import (
	"text/template"

	"github.com/solo-io/protoc-gen-ext/extproto"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func register(tpl *template.Template, params pgs.Parameters) {
	fns := goSharedFuncs{
		Context: pgsgo.InitContext(params),
		tpl:     tpl,
	}

	tpl.Funcs(map[string]interface{}{
		"msgTyp":      fns.msgTyp,
		"pkg":         fns.PackageName,
		"renderField": fns.renderField,
		"renderOneOf": fns.renderOneOf,
		"name":        fns.Name,
	})
}

type goSharedFuncs struct {
	pgsgo.Context
	tpl *template.Template
}

func (fns goSharedFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}

func (fns goSharedFuncs) renderField(field pgs.Field) (string, error) {
	// check if merge hash is set on a given field
	var skipMerge bool
	_, err := field.Extension(extproto.E_SkipMerging, &skipMerge)
	if err != nil {
		return "", err
	}
	if skipMerge {
		return "", nil
	}

	fieldName := field.Name().UpperCamelCase().String()

	mergeField := `m.` + fieldName + ` = overrides.` + fieldName + ` `

	mergeNullableField := `if overrides.` + fieldName + ` != nil {` + mergeField + ` }`

	if field.Type().ProtoType() == pgs.MessageT {
		return mergeNullableField, nil
	}

	return mergeField, nil

}

func (fns goSharedFuncs) renderOneOf(oneof pgs.OneOf) (string, error) {
	// check if skip merge is set on a given oneof
	var skipMerge bool
	_, err := oneof.Extension(extproto.E_SkipMergingOneof, &skipMerge)
	if err != nil {
		return "", err
	}
	if skipMerge {
		return "", nil
	}

	fieldName := oneof.Name().UpperCamelCase().String()

	mergeField := `m.` + fieldName + ` = overrides.` + fieldName + ` `

	return `if overrides.` + fieldName + ` != nil {` + mergeField + ` }`, nil
}
