.PHONY: all build test test-integration lint tidy vet clean run

BINARY := offat
PKG    := github.com/owasp-offat/offat
LDFLAGS := -s -w -X $(PKG)/internal/cli.Version=$(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

all: build

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/offat

test:
	go test ./...

test-integration:
	go test -tags=integration ./test/...

lint:
	golangci-lint run

tidy:
	go mod tidy

vet:
	go vet ./...

run: build
	./$(BINARY) --help

clean:
	rm -f $(BINARY) coverage.out
