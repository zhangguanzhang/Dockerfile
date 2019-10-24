FROM golang:1.14.1 as builder

WORKDIR $GOPATH/src/acfun
COPY . $GOPATH/src/acfun
ENV GO111MODULE=on
ARG GOPROXY=https://goproxy.cn,https://goproxy.io,https://mirrors.aliyun.com/goproxy/,direct

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /root/acfun main.go

FROM zhangguanzhang/youtube-dl
WORKDIR /root/data
COPY --from=builder /root/acfun /usr/local/bin
COPY entrypoint.sh /usr/local/bin
ENTRYPOINT ["entrypoint.sh"]
