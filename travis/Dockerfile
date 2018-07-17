FROM ruby:alpine
RUN apk add --no-cache build-base && \
    gem install travis && \
    gem install travis-lint && \
    apk del build-base && \
    apk add --no-cache git
WORKDIR project
#VOLUME ["/project"]
ENTRYPOINT ["travis"]
