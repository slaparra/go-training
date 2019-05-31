FROM golang:1.12-alpine AS build

MAINTAINER <slaparram@gmail.com>

WORKDIR /go/src/app

# install git to get after the app dependencies
RUN apk add --no-cache git

COPY web-app.go .
COPY layout.gohtml .

# get dependencies
RUN go get -d -v ./...
RUN go build -o /go/bin/myapp web-app.go

ENTRYPOINT ["/go/bin/myapp"]


