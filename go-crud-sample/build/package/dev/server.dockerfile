FROM golang:1.15.5-alpine3.12 AS builder

WORKDIR /go-crud-sample

COPY ./ ./

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache bash curl \
  && rm -rf /var/cache/apk/*

# NOTE: マルチステージビルド見本
# FROM alpine:3.12.1 AS binary
# WORKDIR /app
# COPY --from=builder ./bin/server ./server
# EXPOSE 8080
# CMD ["/app/server"]