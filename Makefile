#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------

ROOTDIR := $(shell pwd)
OUTPUT_DIR ?= $(ROOTDIR)/_output
EXEC_NAME := $(OUTPUT_DIR)/protoc-gen-ext
SOURCES := $(shell find . -name "*.go" | grep -v test.go)
VERSION ?= $(shell git describe --tags)

GO_BUILD_FLAGS := GO111MODULE=on CGO_ENABLED=0 GOARCH=amd64
GCFLAGS := 'all=-N -l'


#----------------------------------------------------------------------------------
# Repo init
#----------------------------------------------------------------------------------

# https://www.viget.com/articles/two-ways-to-share-git-hooks-with-your-team/
.PHONY: init
init:
	git config core.hooksPath .githooks

#----------------------------------------------------------------------------------
# Protos
#----------------------------------------------------------------------------------
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
	protoc -I=. --go_out="${EXT_IMPORT}:." ext/gogoproto/gogo.proto
	cp -r ${PACKAGE}/ext/ ext; rm -rf github.com
	protoc -I=. -I=./vendor/github.com/gogo/protobuf --go_out="." --ext_out="."  tests/api/hello.proto


$(EXEC_NAME):
	$(GO_BUILD_FLAGS) go build -o $@ main.go

$(EXEC_NAME)-darwin-amd64:
	$(GO_BUILD_FLAGS) GOOS=darwin go build -gcflags=$(GCFLAGS) -o $@ main.go

$(EXEC_NAME)-linux-amd64:
	$(GO_BUILD_FLAGS) GOOS=darwin go build -gcflags=$(GCFLAGS) -o $@ main.go


PHONY: build
build: $(EXEC_NAME)

PHONY: install
install:
	$(GO_BUILD_FLAGS) go install .
