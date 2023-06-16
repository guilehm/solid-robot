FROM golang:1.20-alpine

ADD . /go/src

WORKDIR /go/src

RUN go get -d -v ./...
RUN go install -v ./...
