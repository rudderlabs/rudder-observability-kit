# Generate labels for all language runtimes
.PHONY: generate
generate:
	go run cmd/generate/main.go

.PHONY: test
test:
	go test -v -count 1 ./...