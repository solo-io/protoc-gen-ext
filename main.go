package main

import (
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen/protoc-gen-hash/module"

	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		module.Hasher(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}