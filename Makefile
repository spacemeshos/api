PROJECT := api

# This controls the remote HTTPS git location to compare against for breaking changes in CI.
#
# Most CI providers only clone the branch under test and to a certain depth, so when
# running buf check breaking in CI, it is generally preferable to compare against
# the remote repository directly.
#
# Basic authentication is available, see https://buf.build/docs/inputs#https for more details.
HTTPS_GIT := https://github.com/spacemeshos/api.git

# This controls the remote SSH git location to compare against for breaking changes in CI.
#
# CI providers will typically have an SSH key installed as part of your setup for both
# public and private repositories. Buf knows how to look for an SSH key at ~/.ssh/id_rsa
# and a known hosts file at ~/.ssh/known_hosts or /etc/ssh/known_hosts without any further
# configuration. We demo this with CircleCI.
#
# See https://buf.build/docs/inputs#ssh for more details.
SSH_GIT := ssh://git@github.com/spacemeshos/api.git

# This controls the version of buf to install and use.
BUF_VERSION := 1.50.1

GRPC_JSON_PROXY_VERSION := v2.26.3
PROTOC_GO_VERSION := v1.36.6
PROTOC_GEN_GO_VERSION := v1.5.1
GO_SWAGGER_VERSION := v0.31.0

# Everything below this line is meant to be static, i.e. only adjust the above variables. ###

ifeq ($(OS),Windows_NT)
	UNAME_OS := windows
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		UNAME_ARCH := x86_64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),ARM64)
		UNAME_ARCH := aarch64
	endif
	PROTOC_BUILD := win64

	BIN_DIR := $(abspath .)/bin
	export PATH := $(BIN_DIR);$(PATH)
else
	UNAME_OS := $(shell uname -s)
	UNAME_ARCH := $(shell uname -m)
	PROTOC_BUILD := $(shell echo ${UNAME_OS}-${UNAME_ARCH} | tr '[:upper:]' '[:lower:]' | sed 's/darwin/osx/' | sed 's/arm64/aarch_64/' | sed 's/aarch64/aarch_64/')

 	BIN_DIR := $(abspath .)/bin
 	export PATH := $(BIN_DIR):$(PATH)
endif

# `go install` will put binaries in $(GOBIN), avoiding
# messing up with global environment.
export GOBIN := $(BIN_DIR)

BUF := $(BIN_DIR)/buf_$(BUF_VERSION)
$(BUF): protoc-plugins deps
	@mkdir -p $(BIN_DIR)
	@rm -f $(BIN_DIR)/buf_*
	@curl -sSL "https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_OS)-$(UNAME_ARCH)" -o $@
	@chmod +x $@
	@ln -sf $@ $(BIN_DIR)/buf

# Download protoc plugins
protoc-plugins:
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(GRPC_JSON_PROXY_VERSION)
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@$(GRPC_JSON_PROXY_VERSION)
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GO_VERSION)
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_VERSION)
.PHONY: protoc-plugins

# Download go deps
.PHONY: deps
deps:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@$(GO_SWAGGER_VERSION)

# local is what we run when testing locally.
# This does breaking change detection against our local git repository.

.PHONY: local
local: $(BUF)
	buf breaking --against '.git#branch=node-split-poc'

# Linter only. This does not do breaking change detection.

.PHONY: lint
lint: $(BUF)
	buf lint

# https is what we run when testing in most CI providers.
# This does breaking change detection against our remote HTTPS git repository.

.PHONY: breaking
breaking: $(BUF)
	buf breaking --against "$(HTTPS_GIT)#branch=node-split-poc"

# Run all builds
.PHONY: build
build: $(BUF)
	buf generate

	sed -i'.bak' 's/"version": "version not set"/"version": "v2alpha1"/' release/openapi/swagger/src/api.swagger.json && rm -f release/openapi/swagger/src/api.swagger.json.bak
	swagger flatten --with-flatten=remove-unused release/openapi/swagger/src/api.swagger.json -o release/openapi/swagger/src/api.swagger.json

	sed -i'.bak' 's/"version": "version not set"/"version": "v2beta1"/' release/openapi/swagger/src/api_v2beta1.swagger.json && rm -f release/openapi/swagger/src/api_v2beta1.swagger.json.bak
	swagger flatten --with-flatten=remove-unused release/openapi/swagger/src/api_v2beta1.swagger.json -o release/openapi/swagger/src/api_v2beta1.swagger.json

# Make sure build is up to date
.PHONY: check
check: $(BUF)
	@git diff --quiet || (echo "\033[0;31mWorking directory not clean!\033[0m" && git --no-pager diff && exit 1)
	@make build
	@git diff --name-only --diff-filter=AM --exit-code . || { echo "\nPlease rerun 'make build' and commit changes.\n"; exit 1; }
