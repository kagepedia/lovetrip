FROM golang:1.13-alpine

WORKDIR /go/src/github.com/kagepedia/go-next-pj/backend

COPY . /go/src/github.com/kagepedia/go-next-pj/backend

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/cespare/reflex

COPY reflex.conf /
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]