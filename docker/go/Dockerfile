FROM golang:1.15.12-alpine

WORKDIR $GOPATH/src

ENV GO111MODULE=off

RUN apk update && \
    apk add --no-cache \
    alpine-sdk \
    git \
    tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

RUN go get -u github.com/cosmtrek/air

ENV GO111MODULE=on
