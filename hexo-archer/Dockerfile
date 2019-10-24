FROM node:13.14-alpine3.10
WORKDIR /root/blog
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add git && \
    npm install -g hexo-cli && \
    hexo init && \
    npm install hexo-deployer-git --save && \
    npm install && \
    npm i hexo-generator-json-content --save && \
    npm i --save hexo-wordcount && \
    git clone https://github.com/fi3ework/hexo-theme-archer.git themes/archer --depth=1 && \
    sed -ri '/^theme/s#landscape#archer#' _config.yml && \
    rm -rf themes/landscape source/_posts/hello-world.md && \
    mkdir -m 0700 ~/.ssh && \
    apk add --no-cache openssh
COPY docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
