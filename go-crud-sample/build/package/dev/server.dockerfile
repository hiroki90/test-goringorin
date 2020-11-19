FROM golang:1.15.5-alpine3.12 AS builder

WORKDIR /go-crud-sample

ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY ./ ./

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache \
  build-base bash vim curl git mysql-client \
  && rm -rf /var/cache/apk/* \
  && go build -o /bin/server ./cmd/go-crud-sample/main.go

# NOTE: マルチステージビルド見本
# docker-compose を使うローカル環境において
# 以下コードは使用されないが見本として記述

FROM alpine:3.12.0

WORKDIR /app

COPY --from=builder ./bin/server ./server

CMD ["/app/server"]