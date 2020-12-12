FROM alpine:3.12.2

WORKDIR /go-crud-sample

COPY ./bin/server ./bin/server

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache bash curl mysql-client \
  && rm -rf /var/cache/apk/*

EXPOSE 8080

CMD ["/go-crud-sample/bin/server"]