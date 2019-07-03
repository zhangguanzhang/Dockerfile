FROM ruby:alpine
RUN apk add --no-cache build-base ca-certificates && \
    gem install travis && \
    gem install travis-lint && \
    apk del build-base && \
    apk add --no-cache git && \
    rm -f /var/cache/apk/* /tmp/*
WORKDIR project
#VOLUME ["/project"]
ENTRYPOINT ["travis"]
