package clone

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/solo-io/protoc-gen-ext/templates"
)

func Register(common pgs.DebuggerCommon) templates.RegisterFn {
	return func(tpl *template.Template, params pgs.Parameters) {
		register(tpl, params, common)
		template.Must(tpl.Parse(fileTpl))
		template.Must(tpl.New("msg").Parse(msgTpl))
	}
}
