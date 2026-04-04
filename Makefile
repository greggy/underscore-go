.PHONY: build test lint vet fmt check

build:
	go build ./...

test:
	go test ./... -count=1 -v

lint: vet fmt
	@echo "Lint passed"

vet:
	go vet ./...

fmt:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && echo "Run 'make fmt-fix' to fix" && exit 1)

fmt-fix:
	gofmt -w .

check: build lint test
