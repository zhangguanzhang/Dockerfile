FROM alpine

LABEL maintainer="zhangguanzhang <zhangguanzhang@qq.com>"

ARG FileName=udp2raw_binaries.tar.gz
ARG VERSION=latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add curl iptables tcpdump jq \
    && arch=$(arch) \
    && if [ "${arch}" = "x86_64" ]; then \
            arch="amd64"; \
       elif [ "${arch}" = "aarch64" ]; then \
            arch="arm"; \
       fi \
    && url=$( curl -s https://api.github.com/repos/wangyu-/udp2raw-tunnel/releases/${VERSION} | jq -r '.assets[]| select(.name=="'${FileName}'") | .browser_download_url' ) \
    && wget $url -O $FileName \
    && tar zxvf $FileName -C /tmp/ \
    && rm -f $FileName \
    && mv /tmp/udp2raw_${arch} /udp2raw \
    && echo udp2raw_${arch} >> /tmp/version.txt \
    && chmod a+x /udp2raw \
    && mv /tmp/version.txt / \
    && rm -rf /var/cache/apk/* /tmp/*
ENTRYPOINT ["/udp2raw"]