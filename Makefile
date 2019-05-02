# Set Project Name
NAME=golang-library
# Set Go version
GO_VERSION=$(go version)
# Set OS
GO_OS=$(go env GOOS)
# Set architecture
GO_ARCH=$(go env GOARCH)

# Install external dependencies
.PHONY: deps
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

# Run tests
test: deps
	go test -cover -v ./...

# Run Lint
lint: deps
	go vet ./...
	golint ./...

# Clean files
clean:
	go clean
	rm -rf bin/*

# Build programs
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

# Show help
help:
	@make2help #(MAKEFILE_LIST)


.PHONY: build test link clean help
