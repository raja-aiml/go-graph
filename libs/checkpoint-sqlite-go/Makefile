.PHONY: fmt format lint test

fmt:
	gofmt -w sqlite.go sqlite_test.go

format: fmt

lint:
	go vet ./...

test:
	go test ./...
