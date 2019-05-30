FROM golang:1.12-alpine

MAINTAINER <slaparram@gmail.com>

WORKDIR /go/src/app

RUN apk add --no-cache git

COPY web-app.go .
COPY layout.gohtml .

ENV GOPATH "/go"
ENV GOBIN  "$GOPATH/bin"
ENV PATH "$GOBIN:$PATH"

RUN go get -d -v ./...
