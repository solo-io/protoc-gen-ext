#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------

ROOTDIR := $(shell pwd)
OUTPUT_DIR ?= $(ROOTDIR)/_output
EXEC_NAME := $(OUTPUT_DIR)/protoc-gen-ext
SOURCES := $(shell find . -name "*.go" | grep -v test.go)
VERSION ?= $(shell git describe --tags)

GO_BUILD_FLAGS := GO111MODULE=on CGO_ENABLED=0 GOARCH=$(shell uname -m)
GCFLAGS := 'all=-N -l'


#----------------------------------------------------------------------------------
# Repo init
#----------------------------------------------------------------------------------

# https://www.viget.com/articles/two-ways-to-share-git-hooks-with-your-team/
.PHONY: init
init:
	git config core.hooksPath .githooks

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)

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


.PHONY: generated-code
generated-code:
	PATH=$(DEPSGOBIN):$$PATH $(DEPSGOBIN)/protoc -I=. -I=./external --go_out="${EXT_IMPORT}:." extproto/ext.proto
	PATH=$(DEPSGOBIN):$$PATH cp -r ${PACKAGE}/extproto/* extproto
	PATH=$(DEPSGOBIN):$$PATH $(DEPSGOBIN)/protoc -I=. -I=./extproto -I=./external --go_out="." --ext_out="." tests/api/hello.proto
	PATH=$(DEPSGOBIN):$$PATH cp -r ${PACKAGE}/tests/api/* tests/api/
	PATH=$(DEPSGOBIN):$$PATH rm -rf github.com
	PATH=$(DEPSGOBIN):$$PATH $(DEPSGOBIN)/goimports -w .


DEPSGOBIN=$(shell pwd)/_output/.bin

# https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
.PHONY: install-go-tools
install-go-tools: install
	mkdir -p $(DEPSGOBIN)
	GOBIN=$(DEPSGOBIN) go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
	GOBIN=$(DEPSGOBIN) go install golang.org/x/tools/cmd/goimports@v0.1.2
	GOBIN=$(DEPSGOBIN) go install github.com/onsi/ginkgo/ginkgo@v1.16.5

# proto compiler installation
PROTOC_URL:=https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8
.PHONY: install-protoc
install-protoc:
ifeq ($(shell uname),Darwin)
	@echo downloading protoc for osx
	wget $(PROTOC_URL)-osx-x86_64.zip -O $(DEPSGOBIN)/protoc-3.15.8.zip
else
ifeq ($(shell uname -m),aarch64)
	@echo downloading protoc for linux aarch64
	wget $(PROTOC_URL)-linux-aarch_64.zip -O $(DEPSGOBIN)/protoc-3.15.8.zip
else
	@echo downloading protoc for linux x86-64
	wget $(PROTOC_URL)-linux-x86_64.zip -O $(DEPSGOBIN)/protoc-3.15.8.zip
endif
endif

	unzip $(DEPSGOBIN)/protoc-3.15.8.zip -d $(DEPSGOBIN)/protoc-3.15.8
	mv $(DEPSGOBIN)/protoc-3.15.8/bin/protoc $(DEPSGOBIN)/protoc
	chmod +x $(DEPSGOBIN)/protoc
	rm -rf $(DEPSGOBIN)/protoc-3.15.8 $(DEPSGOBIN)/protoc-3.15.8.zip

	# manage google protos too, since we have a folder of them based on the protoc version

.PHONY: install-tools
install-tools: install-go-tools install-protoc

.PHONY: run-tests
run-tests: install-tools
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
	GOBIN=$(DEPSGOBIN) $(GO_BUILD_FLAGS) go install .
