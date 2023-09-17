.PHONY: run
run: fmt
	go run cmd/main.go
fmt: 
	go fmt ./...