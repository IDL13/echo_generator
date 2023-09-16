package _const

const DockerfileBase = `FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o echo ./cmd/app/app.go

CMD ["./echo"]`

const DockerComposeBase = `version: '3.8'

services:
  echo:
    build: ./
    command: ./echo
    ports:
      - "8080:8080"
    environment:
      DB_PASSWORD: "root"
      DB_USER: "admin"
      DB_DB: "postgres"
      DB_HOST: "5435"`
