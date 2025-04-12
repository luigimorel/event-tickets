# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS build

LABEL maintainer="Luigi Morel <hi@luigimorel.co>"

WORKDIR /app

COPY go.mod go.sum  ./

RUN go mod download

COPY . .

RUN go build -o ./server ./main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /app/server ./

COPY .env .

EXPOSE 8080

CMD ["./server"]
