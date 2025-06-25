.PHONY: fmt format lint test

fmt:
	gofmt -w checkpoint.go checkpoint_test.go

format: fmt

lint:
	go vet ./...

test:
	go test ./...
