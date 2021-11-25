ARG ALPINE_VERSION 3.14.3
FROM alpine:${ALPINE_VERSION}

LABEL MAINTAINER="zhangguanzhang zhangguanzhang@qq.com" \
    URL="https://github.com/zhangguanzhang/Dockerfile/tree/master/keepalived"

ENV TZ Asia/Shanghai

COPY docker-entrypoint.sh /entrypoint.sh

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache keepalived curl bash iproute2 grep sed ipvsadm ipset tcpdump tzdata libnl && \
    chmod a+x /entrypoint.sh && \
    mkdir -p /etc/keepalived && \
    echo Asia/Shanghai > /etc/timezone && \
    if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns myhostname' > /etc/nsswitch.conf; fi && \
    rm -rf /var/cache/apk/* /tmp/*

ENTRYPOINT ["/entrypoint.sh"]
