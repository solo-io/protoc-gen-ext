package golang

import (
	"text/template"

	"github.com/lyft/protoc-gen-star"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))
	// template.Must(tpl.New("message").Parse(messageTpl))
}
