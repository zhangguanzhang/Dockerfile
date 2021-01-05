FROM alpine

COPY entrypoint.sh /usr/local/bin/

RUN apk add --no-cache tftp-hpa && \
    rm -rf /var/cache/apk/* /tmp/* && \
    chmod a+x /usr/local/bin/entrypoint.sh

EXPOSE 69/udp
VOLUME [/var/tftpboot]
ENTRYPOINT ["entrypoint.sh"]
CMD ["-L","--verbose","--foreground","--secure","/var/tftpboot"]
