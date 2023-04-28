.PHONY: build
build:
	go build -buildmode=c-shared -o pangea.so .

.PHONY: lint
lint:
	golangci-lint run
