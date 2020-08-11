FROM golang:1.14.7-alpine

RUN apk add git

WORKDIR /usr/app

RUN go get github.com/cosmtrek/air