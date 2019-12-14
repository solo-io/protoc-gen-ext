empty :=
space := $(empty) $(empty)
PACKAGE := github.com/solo-io/protoc-gen/protoc-gen-hash

# protoc-gen-go parameters for properly generating the import path for PGV
VALIDATE_IMPORT := Mhash/hash.proto=${PACKAGE}/hash
GO_IMPORT_SPACES := ${VALIDATE_IMPORT},\
	Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,\
	Mgogoproto/gogo.proto=${PACKAGE}/gogoproto
GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))


PHONY: protoc-gen-hash
protoc-gen-hash:
	go build -o _output/protoc-gen-hash protoc-gen-hash/main.go
