# Set Go version
GO_VERSION=$(go version)
# Set OS
GO_OS=$(go env GOOS)
# Set architecture
GO_ARCH=$(go env GOARCH)

# Install external dependencies
.PHONY: deps
deps:
	go get github.com/stretchr/testify

# Build programs
.PHONY: build
build:
	go build -x -o bin/app

# Run tests
.PHONY: test
test:
	go test -cover -v ./...

# Clean files
.PHONY: clean
clean:
	go clean
	rm -rf bin/*
