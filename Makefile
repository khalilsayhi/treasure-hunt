FILEPATH?=

run:
	go run . $(FILEPATH)
.PHONY: run

install-golangci-lint:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
.PHONY: install-golangci-lint

lint: install-golangci-lint
	golangci-lint run -v --timeout 5m

test:
	go test ./tests/... -v -race
.PHONY: test