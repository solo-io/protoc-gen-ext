package hash

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("hash").Parse(hashTpl))
	template.Must(tpl.New("uniquehash").Parse(uniqueHashTpl))
}
