package main

import (
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen-ext/module/equal"
	"github.com/solo-io/protoc-gen-ext/module/hash"

	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("PROTO_DEBUG"),
	).RegisterModule(
		hash.Hash(),
		equal.Equal(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}