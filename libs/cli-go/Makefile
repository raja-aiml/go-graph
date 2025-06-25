.PHONY: fmt format lint test

fmt:
	gofmt -w main.go main_test.go

format: fmt

lint:
	go vet ./...

test:
	go test ./...
