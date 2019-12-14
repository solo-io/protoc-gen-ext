package main

import (
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen-ext/module"

	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		module.Ext(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}