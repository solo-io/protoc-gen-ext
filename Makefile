empty :=
space := $(empty) $(empty)
PACKAGE := github.com/solo-io/protoc-gen-ext

# protoc-gen-go parameters for properly generating the import path for PGV
EXT_IMPORT := Mhash/hash.proto=${PACKAGE}/ext
GO_IMPORT_SPACES := ${EXT_IMPORT},\
	Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,
GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))



PHONE: generated-code
generated-code:
	protoc -I=. --go_out="${EXT_IMPORT}:." ext/ext.proto
	mv ${PACKAGE}/ext/ext.pb.go ext; rm -rf github.com
	protoc -I=. --go_out="." --ext_out="."  tests/api/hello.proto

PHONY: protoc-gen-ext
protoc-gen-ext:
	go build -o _output/protoc-gen-ext main.go
