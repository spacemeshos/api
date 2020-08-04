# Source: https://github.com/bufbuild/buf-example/blob/master/Makefile

SHELL := /usr/bin/env bash -o pipefail

# This controls the location of the cache.
PROJECT := api
# This controls the remote HTTPS git location to compare against for breaking changes in CI.
#
# Most CI providers only clone the branch under test and to a certain depth, so when
# running buf check breaking in CI, it is generally preferable to compare against
# the remote repository directly.
#
# Basic authentication is available, see https://buf.build/docs/inputs#https for more details.
HTTPS_GIT := https://github.com/spacemesh/api.git
# This controls the remote SSH git location to compare against for breaking changes in CI.
#
# CI providers will typically have an SSH key installed as part of your setup for both
# public and private repositories. Buf knows how to look for an SSH key at ~/.ssh/id_rsa
# and a known hosts file at ~/.ssh/known_hosts or /etc/ssh/known_hosts without any further
# configuration. We demo this with CircleCI.
#
# See https://buf.build/docs/inputs#ssh for more details.
SSH_GIT := ssh://git@github.com/spacemesh/api.git
# This controls the version of buf to install and use.
BUF_VERSION := 0.16.0

# This controls the version of protoc to install and use.
PROTOC_VERSION = 3.12.3

# Version of go protoc tools
PROTOC_GEN_GO_VERSION = v1.4.2
PROTOC_GEN_GRPC_GATEWAY_VERSION = v1.14.6

# The include flags to pass to protoc.
PROTOC_INCLUDES := -I ./proto -I ./third_party

# The files to run protoc on
PROTOC_INPUTS := $(shell find ./proto -name *.proto)

# The directory to store go builds
PROTOC_GO_BUILD_DIR := ./release/go

# Name of go module corresponding to this path
GO_MODULE_PATH := github.com/spacemeshos/api/release/go

# Plugins string for go builds (must end in ':')
PROTOC_GO_PLUGINS := plugins=grpc:

# Options string appended to go build command (optional, obviously)
PROTOC_GO_OPT := --go_opt=paths=source_relative

# The directory to store python builds
PROTOC_PYTHON_BUILD_DIR := ./release/python

# Plugins string for grpc-gateway (must end in ':')
PROTOC_GATEWAY_PLUGINS := logtostderr=true:

# Service configuration file
PROTOC_GATEWAY_CONFIG := ./proto/spacemesh/v1/api_config.yaml

# Options string appended to grpc-gateway build command (optional)
PROTOC_GATEWAY_OPT := --grpc-gateway_opt=paths=source_relative,grpc_api_configuration=$(PROTOC_GATEWAY_CONFIG)

### Everything below this line is meant to be static, i.e. only adjust the above variables. ###

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
UNAME_OS_PROTOC := $(shell uname -s | tr '[:upper:]' '[:lower:]' | sed 's/darwin/osx/')
# Buf will be cached to ~/.cache/buf-example.
CACHE_BASE := $(HOME)/.cache/$(PROJECT)
# This allows switching between i.e a Docker container and your local setup without overwriting.
CACHE := $(CACHE_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
# The location where buf will be installed.
CACHE_BIN := $(CACHE)/bin
# Marker files are put into this directory to denote the current version of binaries that are installed.
CACHE_VERSIONS := $(CACHE)/versions
# We do some temporary work here
CACHE_TMP := $(CACHE_BASE)/tmp

# Go tools require that this be set
ifndef GOPATH
	export GOPATH=$(shell go env GOPATH)
endif

# Update the $PATH so we can use buf and protoc directly
export PATH := $(abspath $(CACHE_BIN)):$(abspath $(GOPATH)/bin):$(PATH)

# BUF points to the marker file for the installed version.
#
# If BUF_VERSION is changed, the binary will be re-downloaded.
BUF := $(CACHE_VERSIONS)/buf/$(BUF_VERSION)
$(BUF):
	@rm -f $(CACHE_BIN)/buf
	@mkdir -p $(CACHE_BIN)
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_OS)-$(UNAME_ARCH)" \
		-o "$(CACHE_BIN)/buf"
	chmod +x "$(CACHE_BIN)/buf"
	@rm -rf $(dir $(BUF))
	@mkdir -p $(dir $(BUF))
	@touch $(BUF)

PROTOC := $(CACHE_VERSIONS)/protoc/$(PROTOC_VERSION)
$(PROTOC):
	@rm -f $(CACHE_TMP)/protoc.zip
	@rm -f $(CACHE_BIN)/protoc
	@mkdir -p $(CACHE_TMP)
	@mkdir -p $(CACHE_BIN)
	curl -sSL \
		"https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(UNAME_OS_PROTOC)-$(UNAME_ARCH).zip" \
		-o "$(CACHE_TMP)/protoc.zip"
	unzip $(CACHE_TMP)/protoc.zip bin/protoc -d $(CACHE)
	@rm -rf $(dir $(PROTOC))
	@mkdir -p $(dir $(PROTOC))
	@touch $(PROTOC)

GO_MOD := $(PROTOC_GO_BUILD_DIR)/go.mod
$(GO_MOD):
	cd $(PROTOC_GO_BUILD_DIR) && go mod init $(GO_MODULE_PATH)

PROTOC_GEN_GO := $(CACHE_VERSIONS)/protoc-gen-go/$(PROTOC_GEN_GO_VERSION)
$(PROTOC_GEN_GO): $(GO_MOD)
	cd $(PROTOC_GO_BUILD_DIR) && \
	  go get github.com/golang/protobuf/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@rm -rf $(dir $(PROTOC_GEN_GO))
	@mkdir -p $(dir $(PROTOC_GEN_GO))
	@touch $(PROTOC_GEN_GO)

PROTOC_GEN_GRPC_GATEWAY := $(CACHE_VERSIONS)/protoc-gen-grpc-gateway/$(PROTOC_GEN_GRPC_GATEWAY_VERSION)
$(PROTOC_GEN_GRPC_GATEWAY): $(GO_MOD)
	cd $(PROTOC_GO_BUILD_DIR) && \
	  go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@$(PROTOC_GEN_GRPC_GATEWAY_VERSION)
	@rm -rf $(dir $(PROTOC_GEN_GRPC_GATEWAY))
	@mkdir -p $(dir $(PROTOC_GEN_GRPC_GATEWAY))
	@touch $(PROTOC_GEN_GRPC_GATEWAY)

.DEFAULT_GOAL := local

# deps allows us to install deps without running any checks.

.PHONY: deps
deps: $(BUF) $(PROTOC) $(PROTOC_GEN_GO) $(PROTOC_GEN_GRPC_GATEWAY)

# local is what we run when testing locally.
# This does breaking change detection against our local git repository.

.PHONY: local
local: $(BUF)
	buf check lint
	buf check breaking --against-input '.git#branch=master'

# Linter only. This does not do breaking change detection.

.PHONY: lint
lint: $(BUF)
	buf check lint

# https is what we run when testing in most CI providers.
# This does breaking change detection against our remote HTTPS git repository.

.PHONY: https
https: $(BUF)
	buf check lint
	buf check breaking --against-input "$(HTTPS_GIT)#branch=master"

# ssh is what we run when testing in CI providers that provide ssh public key authentication.
# This does breaking change detection against our remote HTTPS ssh repository.
# This is especially useful for private repositories.

.PHONY: ssh
ssh: $(BUF)
	buf check lint
	buf check breaking --against-input "$(SSH_GIT)#branch=master"

# Try to build using protoc. This performs different checks and surfaces
# different errors than linting alone. We want this to fail on warnings as well
# as errors, for which purpose we use grep.
.PHONY: protoc
protoc: $(PROTOC)
	protoc $(PROTOC_INCLUDES) $(PROTOC_INPUTS) -o /dev/null
	(protoc $(PROTOC_INCLUDES) $(PROTOC_INPUTS) -o /dev/null 2>&1) | grep warning \
	  && { echo "one or more warnings detected"; exit 1; } || exit 0

## LANGUAGE-SPECIFIC BUILDS

# Golang
.PHONY: golang
golang: $(PROTOC) | $(PROTOC_GEN_GO)
	mkdir -p $(PROTOC_GO_BUILD_DIR)
	protoc $(PROTOC_INCLUDES) $(PROTOC_INPUTS) \
	  --go_out=$(PROTOC_GO_PLUGINS)$(PROTOC_GO_BUILD_DIR) $(PROTOC_GO_OPT)

# Python
.PHONY: python
python: 
	pip install grpcio grpcio-tools
	mkdir -p $(PROTOC_PYTHON_BUILD_DIR)
	python -m grpc_tools.protoc $(PROTOC_INCLUDES) $(PROTOC_INPUTS) \
	  --python_out=$(PROTOC_PYTHON_BUILD_DIR) --grpc_python_out=$(PROTOC_PYTHON_BUILD_DIR) 

.PHONY: test-python
test-python:
	pip install -r $(PROTOC_PYTHON_BUILD_DIR)/requirements.txt
	cd $(PROTOC_PYTHON_BUILD_DIR) && python -m unittest test.spacemesh_test

# grpc-gateway
.PHONY: grpc-gateway
grpc-gateway: $(PROTOC) | $(PROTOC_GEN_GO) $(PROTOC_GEN_GRPC_GATEWAY)
	mkdir -p $(PROTOC_GO_BUILD_DIR)
	protoc $(PROTOC_INCLUDES) $(PROTOC_INPUTS) \
	  --grpc-gateway_out=$(PROTOC_GATEWAY_PLUGINS)$(PROTOC_GO_BUILD_DIR) $(PROTOC_GATEWAY_OPT)

# Run all builds
.PHONY: build
build: golang grpc-gateway

# Make sure build is up to date
.PHONY: check
check: build
	git add -N $(PROTOC_GO_BUILD_DIR)
	git diff --name-only --diff-filter=AM --exit-code $(PROTOC_GO_BUILD_DIR) \
	  || { echo "please update build"; exit 1; }

# clean deletes any files not checked in and the cache for all platforms.

.PHONY: clean
clean:
	git clean -xdf
	rm -rf $(CACHE_BASE)
