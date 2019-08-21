# Set meta information
NAME:=golang-library
PROJECT_PATH=github.com/nagatax/$(NAME)
VERSION:=$(gobump show -r)
REVISION:=$(shell git rev-parse --short HEAD)
LDFLAGS:="-X main.revision=$(REVISION)"

export GO111MODULE=on

# Initialize a project
init:
	go mod init $(PROJECT_PATH)

# Install development tools
devtools:
	GO111MODULE=off \
				go get github.com/motemen/ghq \
				github.com/motemen/gore/cmd/gore \
				github.com/k0kubun/pp \
				github.com/mdempsky/gocode \
				github.com/rogpeppe/godef \
				github.com/jstemmer/gotags \
				github.com/motemen/gobump/cmd/gobump \
				github.com/Songmu/make2help/cmd/make2help

# Install go tools
gotools:
	GO111MODULE=off \
				go get golang.org/x/lint/golint \
				golang.org/x/tools/cmd/godoc \
				golang.org/x/tools/cmd/goimports \
				golang.org/x/tools/cmd/gorename \
				golang.org/x/tools/cmd/guru \
				golang.org/x/tools/cmd/gopls

# Install dependencies
deps:
	go get -v -d

# Build programs
bin/%: cmd/%/main.go deps
	go build -ldflags $(LDFLAGS) -o $@ $<

build: bin/myprof

# Format Files
fmt:
	goimports -w ./cmd

# Run Lint
lint: deps
	go vet ./cmd/...
	golint -set_exit_status ./cmd/...

# Run tests
test: deps
	go test -cover -v ./cmd/...

doc:
	go doc -all ./cmd

# Show help
help:
	@make2help $(MAKEFILE_LIST)

# Clean files
clean:
	go clean
	rm -rf bin/*

.PHONY: init devtools gotools deps build fmt lint test doc help clean
