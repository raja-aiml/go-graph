.PHONY: fmt format lint test

fmt:
	gofmt -w prebuilt.go prebuilt_test.go

format: fmt

lint:
	go vet ./...

test:
	go test ./...
