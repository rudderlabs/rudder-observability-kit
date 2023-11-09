# Generare labels for all language runtimes
.PHONY: generate
generate:
	go run cmd/generate/main.go
