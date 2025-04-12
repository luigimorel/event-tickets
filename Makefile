.PHONY: start-api start-web lint-api lint-web fmt-api fmt-web test-api test-web install-web

# API Commands
start-api:
	cd api && go run cmd/app/main.go

dev-api:
	cd api && nodemon --exec go run cmd/app/main.go --signal SIGTERM

lint-api:
	cd api && golangci-lint run ./...

fmt-api:
	cd api && go fmt ./...

test-api:
	cd api && go test -v ./...
install-api:
	cd api && go mod download

# Web Commands
install-web:
	cd web && npm install

start-web:
	cd web && npm run dev

build-web:
	cd web && npm run build

lint-web:
	cd web && npm run lint

format-web:
	cd web && npm run format

test-web:
	cd web && npm run test

# Combined Commands
install: install-web
	cd ..
	install-api

start:
	make -j2 start-api start-web

lint: lint-api lint-web

fmt: fmt-api fmt-web

test: test-api test-web

.DEFAULT_GOAL := start
