VERSION=$(shell cat VERSION)
GO_BUILD_LDFLAGS=-ldflags '-X main.Version=v${VERSION}' 

.PHONY: build
build:
	go build -buildmode=c-shared -o pangea.so ${GO_BUILD_LDFLAGS} .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: build
	deno test -A --unstable
