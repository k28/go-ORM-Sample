FROM golang:1.14.11-alpine3.12

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache \
            alpine-sdk \
            git

# ホットリロード用
RUN go get github.com/pilu/fresh

# Goのライブラリを追加する
RUN go get github.com/lib/pq
RUN go get github.com/go-gorp/gorp
RUN go get github.com/labstack/echo

EXPOSE 8080

CMD ["fresh"]

