FROM alpine:3.7

COPY keepalived.sh  /usr/bin/keepalived.sh

RUN apk --update -t --no-cache add keepalived iproute2 grep bash tcpdump sed  && \
    chmod +x /usr/bin/keepalived.sh && \
    rm -f /var/cache/apk/* /tmp/*

COPY keepalived.conf /etc/keepalived/keepalived.conf

ENTRYPOINT ["/usr/bin/keepalived.sh"]
