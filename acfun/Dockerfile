FROM golang:1.15.6 as builder

WORKDIR $GOPATH/src/acfun
COPY . $GOPATH/src/acfun
ENV GO111MODULE=on
ARG GOPROXY=https://goproxy.cn,https://goproxy.io,https://mirrors.aliyun.com/goproxy/,direct

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /root/acfun main.go

FROM zhangguanzhang/youtube-dl
WORKDIR /root/data
COPY --from=builder /root/acfun /usr/local/bin
COPY entrypoint.sh /
COPY down.sh /etc/profile.d/down.sh
RUN apk add bash libwebp-tools && \
    rm -rf /var/cache/apk/* /tmp/* && \
    chmod a+x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
