package templates

import (
	"fmt"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	golang "github.com/solo-io/protoc-gen-ext/templates/go"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

func Template(params pgs.Parameters) *template.Template {
	tpl := template.New("go")
	// shared.RegisterFunctions(tpl, params)
	golang.Register(tpl, params)
	return tpl
}

func FilePathFor(tpl *template.Template, moduleName string) FilePathFn {
	switch tpl.Name() {
	// case "h":
	// 	return cc.CcFilePath
	// case "cc":
	// 	return cc.CcFilePath
	// case "java":
	// 	return java.JavaFilePath
	default:
		return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = out.SetExt(fmt.Sprintf(".%s.", moduleName) + tpl.Name())
			return &out
		}
	}
}
