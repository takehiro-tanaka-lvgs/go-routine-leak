FROM golang:1.22.2-alpine

RUN apk add --update --no-cache \
    vim \
    git \
    tar
