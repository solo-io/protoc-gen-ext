package main

import (
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/solo-io/protoc-gen-ext/module/clone"
	"github.com/solo-io/protoc-gen-ext/module/equal"
	"github.com/solo-io/protoc-gen-ext/module/hash"
	"github.com/solo-io/protoc-gen-ext/module/merge"

	pgs "github.com/lyft/protoc-gen-star"
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
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
