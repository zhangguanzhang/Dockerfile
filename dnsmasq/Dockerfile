FROM alpine:3.13

ARG VERSION=2.83

LABEL MAINTAINER="zhangguanzhang zhangguanzhang@qq.com" \
    URL="https://github.com/zhangguanzhang/Dockerfile/tree/master/dnsmasq"

RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.12/main' >> /etc/apk/repositories && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add "dnsmasq-dnssec==${VERSION}-r0" && \
    rm -rf /var/cache/apk/* /tmp/*

EXPOSE 53 53/udp

ENTRYPOINT ["dnsmasq", "-k"]