FROM golang:1.16.0-alpine

RUN \
  apk update && \
  apk --no-cache --update add \
    curl \
    git
