.PHONY: fmt format lint test

fmt:
	gofmt -w graph.go graph_test.go

format: fmt

lint:
	go vet ./...

test:
	go test ./...
