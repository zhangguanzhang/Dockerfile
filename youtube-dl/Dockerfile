FROM alpine

RUN set -xe \
    && apk add --no-cache ca-certificates \
                          curl \
                          ffmpeg \
                          openssl \
                          python3 \
                          jq \
    && python3 -m ensurepip \
    && ln -s /usr/bin/python3 /usr/bin/python \
    && curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/bin/youtube-dl \
    && chmod a+x /usr/bin/youtube-dl \
    && rm -rf /var/cache/apk/* /tmp/*

WORKDIR /data

ENTRYPOINT ["youtube-dl"]
CMD ["--help"]
