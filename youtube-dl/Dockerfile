FROM alpine

RUN set -xe \
    && apk add --no-cache ca-certificates \
                          curl \
                          ffmpeg \
                          openssl \
                          python3 \
    && pip3 install youtube-dl

WORKDIR /data

ENTRYPOINT ["youtube-dl"]
CMD ["--help"]
