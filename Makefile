.PHONY: build test lint fmt fmt-fix check

build:
	go build ./...

test:
	go test ./... -count=1 -v

lint:
	golangci-lint run ./...

fmt:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && echo "Run 'make fmt-fix' to fix" && exit 1)

fmt-fix:
	gofmt -w .

check: build fmt lint test
