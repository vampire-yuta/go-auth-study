FROM golang:1.18-alpine3.16

RUN apk update && \
    apk upgrade && \
    apk add git

# RUN go get github.com/cespare/reflex@latest
RUN go install github.com/cespare/reflex@latest
ENV CGO_ENABLED=0

WORKDIR /go/src/api
COPY ./api ./

RUN go mod download
