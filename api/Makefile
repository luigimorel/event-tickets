.PHONY: run dev lint fmt test

run:
	go run cmd/app/main.go

dev:
	nodemon --exec go run cmd/app/main.go --signal SIGTERM

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

test:
	go test -v ./...
