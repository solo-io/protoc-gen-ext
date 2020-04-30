package templates

import (
	"fmt"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

func Template(params pgs.Parameters, register RegisterFn) *template.Template {
	tpl := template.New("go")
	register(tpl, params)
	return tpl
}

func FilePathFor(tpl *template.Template, moduleName string) FilePathFn {
	switch tpl.Name() {
	default:
		return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = out.SetExt(fmt.Sprintf(".%s.", moduleName) + tpl.Name())
			return &out
		}
	}
}
