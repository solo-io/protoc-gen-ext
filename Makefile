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
GOGO_OUTPUT := github.com/gogo/protobuf/gogoproto/*

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
	PATH=$(DEPSGOBIN):$$PATH protoc -I=. --go_out="${EXT_IMPORT}:." extproto/ext.proto
	PATH=$(DEPSGOBIN):$$PATH cp -r ${PACKAGE}/extproto/ extproto
	PATH=$(DEPSGOBIN):$$PATH protoc -I=. -I=./extproto --go_out="." --ext_out="."  tests/api/hello.proto
	PATH=$(DEPSGOBIN):$$PATH cp -r ${PACKAGE}/tests/api/ tests/api
	PATH=$(DEPSGOBIN):$$PATH rm -rf github.com
	PATH=$(DEPSGOBIN):$$PATH goimports -w .


DEPSGOBIN=$(shell pwd)/_output/.bin

# https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
.PHONY: install-go-tools
install-go-tools: install
	mkdir -p $(DEPSGOBIN)groups.sh
	GOBIN=$(DEPSGOBIN) go install github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(DEPSGOBIN) go install golang.org/x/tools/cmd/goimports
	GOBIN=$(DEPSGOBIN) go install github.com/onsi/ginkgo/ginkgo

.PHONY: run-tests
run-tests: install-go-tools
	$(DEPSGOBIN)/ginkgo -r -failFast -trace -progress -race -compilers=4 -failOnPending -noColor $(TEST_PKG)

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
