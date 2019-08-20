# Set Project Name
NAME=golang-library
# Set Go version
GO_VERSION=$(go version)
# Set OS
GO_OS=$(go env GOOS)
# Set architecture
GO_ARCH=$(go env GOARCH)

# Install development tools
devtools:
	# Dependency management tool
	go get github.com/golang/dep/cmd/dep
	# Manage remote repository clones
	go get github.com/motemen/ghq
	# gore
	go get github.com/motemen/gore/cmd/gore
	go get github.com/k0kubun/pp
	go get github.com/mdempsky/gocode
	go get github.com/rogpeppe/godef
	go get github.com/jstemmer/gotags

# Install go tools
gotools:
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/godoc
	go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/tools/cmd/gorename
	go get golang.org/x/tools/cmd/guru
	go get golang.org/x/tools/cmd/gopls

# Init project
init:
	dep init

# Install external dependencies
deps:
	dep ensure -vendor-only

# Build programs
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

# Format Files
fmt:
	goimports -w ./cmd

# Run Lint
lint: deps
	go vet ./cmd/...
	golint ./cmd/...

# Run tests
test: deps
	go test -cover -v ./cmd/...

doc:
	go doc -all ./cmd

# Show help
help:
	@make2help #(MAKEFILE_LIST)

# Clean files
clean:
	go clean
	rm -rf bin/*

.PHONY: devtools gotools init deps fmt lint test help clean
