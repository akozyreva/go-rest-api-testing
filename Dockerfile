FROM golang:1.19.1

WORKDIR /rest-api-basics

COPY go.mod .
COPY go.sum .

RUN go mod download