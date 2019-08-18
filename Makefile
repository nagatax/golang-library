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
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/jstemmer/gotags
	go get -u github.com/nsf/gocode
	go get -u github.com/rogpeppe/godef

# Install go tools
gotools:
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/godoc
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/tools/cmd/gorename
	go get -u golang.org/x/tools/cmd/guru

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
	go fmt ./...

# Run Lint
lint: deps
	go vet ./...
	golint ./...

# Run tests
test: deps
	go test -cover -v ./...

# Show help
help:
	@make2help #(MAKEFILE_LIST)

# Clean files
clean:
	go clean
	rm -rf bin/*

.PHONY: devtools gotools init deps fmt lint test help clean
