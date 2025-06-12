FROM golang:1.24.4-alpine3.22 as base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o lemonAuth ./cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY  --from=base  /app/lemonAuth .
COPY .env .env
CMD [ "./lemonAuth" ]
