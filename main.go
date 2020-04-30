package main

import (
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen-ext/module/equal"
	"github.com/solo-io/protoc-gen-ext/module/hash"
	"github.com/solo-io/protoc-gen-ext/module/merge"

	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("PROTO_DEBUG"),
	).RegisterModule(
		hash.Hash(),
		equal.Equal(),
		merge.Merge(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
