FROM alpine
RUN apk add --no-cache tftp-hpa 
COPY entrypoint.sh /usr/local/bin/

EXPOSE 69/udp
VOLUME [/var/tftpboot]
ENTRYPOINT ["entrypoint.sh"]
CMD ["-L","--verbose","--foreground","--secure","/var/tftpboot"]
