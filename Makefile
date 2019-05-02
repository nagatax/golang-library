# Set Project Name
NAME=golang-library
# Set Go version
GO_VERSION=$(go version)
# Set OS
GO_OS=$(go env GOOS)
# Set architecture
GO_ARCH=$(go env GOARCH)

# Build programs
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

# Clean files
clean:
	go clean
	rm -rf bin/*

# Install external dependencies
deps:
	go get github.com/jstemmer/gotags
	go get github.com/nsf/gocode
	go get github.com/rogpeppe/godef
	go get github.com/stretchr/testify
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/godoc
	go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/tools/cmd/gorename
	go get golang.org/x/tools/cmd/guru

# Format Files
fmt:
	go fmt ./...

# Show help
help:
	@make2help #(MAKEFILE_LIST)

# Run Lint
lint: deps
	go vet ./...
	golint ./...

# Run tests
test: deps
	go test -cover -v ./...


.PHONY: clean deps fmt help lint test
