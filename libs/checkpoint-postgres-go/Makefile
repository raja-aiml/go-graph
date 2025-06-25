.PHONY: fmt format lint test

fmt:
	gofmt -w postgres.go postgres_test.go

format: fmt

lint:
	go vet ./...

test:
	go test ./...
