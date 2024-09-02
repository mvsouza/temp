FROM golang:1.22 AS base

RUN mkdir -p /go/src/adjoe.io

WORKDIR /go/src/adjoe.io

FROM adjoe-test/golang-base AS dev

ENV GO111MODULE=on

RUN go install github.com/codegangsta/gin@latest
RUN go install github.com/onsi/ginkgo/ginkgo@latest
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1
RUN go install github.com/onsi/gomega/...@latest

# Set the local PORT environment variable inside the container
ENV PORT 3000