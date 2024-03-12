FROM golang:1.22.0-alpine3.19  as builder

WORKDIR /go/src/github.com/Lichmaker/ladder-admin-v2/v2ray-manager
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o v2ray-manager .

FROM alpine:latest

LABEL MAINTAINER="lich.wu2014@gmail.com"

WORKDIR /go/src/github.com/Lichmaker/ladder-admin-v2/v2ray-manager

COPY --from=0 /go/src/github.com/Lichmaker/ladder-admin-v2/v2ray-manager/v2ray-manager ./
COPY --from=0 /go/src/github.com/Lichmaker/ladder-admin-v2/v2ray-manager/.env.docker ./.env

EXPOSE 10088
ENTRYPOINT ./v2ray-manager serve
