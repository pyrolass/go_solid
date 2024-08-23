.DEFAULT_GOAL := run

.PHONY: run vet fmt


fmt:
	go fmt ./...

vet: fmt
	go vet ./...

run: vet
	@go build -o bin/main main.go
	@go run main.go

