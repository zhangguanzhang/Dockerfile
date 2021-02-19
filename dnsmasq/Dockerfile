FROM alpine:3.13
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add dnsmasq-dnssec && \
    rm -rf /var/cache/apk/* /tmp/*

EXPOSE 53 53/udp

ENTRYPOINT ["dnsmasq", "-k"]