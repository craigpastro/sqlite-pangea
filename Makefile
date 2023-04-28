.PHONY: build
build:
	go build -buildmode=c-shared -o pangea.so .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	PANGEA_EXTENSION=pangea.so go test ./...
