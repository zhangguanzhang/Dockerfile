FROM alpine

RUN set -xe \
    && apk add --no-cache ca-certificates \
                          curl \
                          ffmpeg \
                          openssl \
                          python3 \
                          jq \
    && python3 -m ensurepip \
    && python3 -m pip install youtube-dl

WORKDIR /data

ENTRYPOINT ["youtube-dl"]
CMD ["--help"]
