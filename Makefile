# Set meta information
NAME:=golang-library
PROJECT_PATH=github.com/nagatax/$(NAME)
VERSION:=$(gobump show -r)
REVISION:=$(shell git rev-parse --short HEAD)
LDFLAGS:="-X main.revision=$(REVISION)"

export GO111MODULE=on

## Initialize a project
init:
	go mod init $(PROJECT_PATH)

## Install development tools
devtools:
	GO111MODULE=off \
				go get github.com/motemen/gore/cmd/gore \
				github.com/k0kubun/pp \
				github.com/mdempsky/gocode \
				github.com/rogpeppe/godef \
				github.com/jstemmer/gotags \
				github.com/motemen/gobump/cmd/gobump \
				github.com/Songmu/make2help/cmd/make2help

## Install go tools
gotools:
	GO111MODULE=off \
				go get golang.org/x/lint/golint \
				golang.org/x/tools/cmd/godoc \
				golang.org/x/tools/cmd/goimports \
				golang.org/x/tools/cmd/gorename \
				golang.org/x/tools/cmd/guru \
				golang.org/x/tools/gopls

## Install dependencies
deps: fmt
	go mod tidy

## Build programs
bin/%: cmd/%/main.go deps
	go build -ldflags $(LDFLAGS) -o $@ $<

## Format Files
fmt:
	goimports -w ./cmd ./utils

## Run Lint
lint: deps
	go vet ./cmd/... ./utils/...
	golint -set_exit_status -min_confidence=0.1 ./cmd/... ./utils/...

## Run tests
test: deps lint
	go test -tags=uint -cover -v ./cmd/... ./utils/...
	go test -tags=integration -cover -v ./cmd/... ./utils/...

## Show docs
doc:
	go doc -all ./cmd ./utils

## Show help
help:
	@make2help $(MAKEFILE_LIST)

## Clean files
clean:
	go clean
	rm -rf bin/*

.PHONY: init devtools gotools deps build fmt lint test doc help clean
