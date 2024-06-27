package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/solo-io/protoc-gen-ext/module/clone"
	"github.com/solo-io/protoc-gen-ext/module/equal"
	"github.com/solo-io/protoc-gen-ext/module/hash"
	"github.com/solo-io/protoc-gen-ext/module/merge"
	"github.com/solo-io/protoc-gen-ext/module/uniquehash"
)

func main() {
	feat := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(
		pgs.DebugEnv("PROTO_DEBUG"),
		pgs.SupportedFeatures(&feat),
	).RegisterModule(
		hash.Hash(),
		equal.Equal(),
		merge.Merge(),
		clone.Clone(),
		uniquehash.Hash(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
