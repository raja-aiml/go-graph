.PHONY: test lint format fmt

fmt:
	gofmt -w sdk.go sdk_test.go

format: fmt

lint:
	go vet ./...

test:
	go test ./...

