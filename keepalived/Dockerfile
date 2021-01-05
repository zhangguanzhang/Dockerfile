FROM alpine:3.12.1

COPY docker-entrypoint.sh /entrypoint.sh

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache keepalived curl bash iproute2 grep sed tcpdump && \
    chmod a+x /entrypoint.sh && \
    mkdir -p /etc/keepalived && \
    rm -rf /var/cache/apk/* /tmp/*

ENTRYPOINT ["/entrypoint.sh"]
