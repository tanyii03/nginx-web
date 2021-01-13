FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /go/src/nginx-web
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY . .
RUN go env && go build -o server .

FROM alpine:latest

WORKDIR /go/src/nginx-web
COPY --from=0 /go/src/nginx-web/server ./
COPY --from=0 /go/src/nginx-web/config.yaml ./
COPY --from=0 /go/src/nginx-web/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server
