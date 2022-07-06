PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION = v$(shell cat version/version)
COMMIT_SHA ?= $(shell git describe --always)-devel

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 go build -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"
GOINSTALL=CGO_ENABLED=0 go install -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

MAIN_ROOT ?= ./cmd/tools
OUTPUT=$(abspath bin)
ASSERT=$(OUTPUT)/tools

.PHONY: install
install: download
	@cd $(MAIN_ROOT) && $(GOINSTALL)

.PHONY: build
build: download
	@mkdir -p $(OUTPUT)
	@cd $(MAIN_ROOT) && $(GOBUILD) -o $(ASSERT)

.PHONY: download
download:
	go mod download -x

.PHONY:dep
dep:
	go get -u ./...

.PHONY: clean
clean:
	@rm -rf $(OUTPUT)

test:
	@echo $(OUTPUT)
	@echo $(ASSERT)
