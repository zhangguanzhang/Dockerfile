FROM alpine:3.18

LABEL maintainer="zhangguanzhang <zhangguanzhang@qq.com>"

ARG FileName=udp2raw_binaries.tar.gz
ARG VERSION=latest
# 支持硬件 aes 加密
ARG AES_EN=true

RUN set -eux; \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories; \
    apk add curl \
          bash \
          iptables \
          tcpdump \
          jq \
          --no-cache; \
    curl -L https://raw.githubusercontent.com/kubernetes-sigs/iptables-wrappers/v2/iptables-wrapper-installer.sh > /iptables-wrapper-installer.sh; \
    bash /iptables-wrapper-installer.sh --no-sanity-check; \
    arch=$(arch); \
    if [ "${arch}" = "x86_64" ]; then \
            arch="amd64"; \
       elif [ "${arch}" = "aarch64" ]; then \
            arch="arm"; \
       fi; \
    if [ "$VERSION" != latest ];then \
          url=https://github.com/wangyu-/udp2raw/releases/download/${VERSION}/$FileName; \
     else \
          url=$( curl -sL https://api.github.com/repos/wangyu-/udp2raw/releases/${VERSION} | jq -r '.assets[]| select(.name=="'${FileName}'") | .browser_download_url' ); \
     fi; \
    wget -O $FileName $url ; \
    tar zxvf $FileName -C /tmp/; \
    rm -f $FileName; \
    cp /tmp/version.txt /; \
    cd /tmp; \
    if [ "$AES_EN" = true ];then \
          echo $(ls /tmp/udp2raw_${arch}_*_aes) >> /version.txt; \
          mv /tmp/udp2raw_${arch}_*_aes /udp2raw; \
     else \
          echo udp2raw_${arch} >> /version.txt; \
          mv /tmp/udp2raw_${arch} /udp2raw; \
     fi; \
    cd -; \
    chmod a+x /udp2raw; \
    rm -rf /var/cache/apk/* /tmp/*
ENTRYPOINT ["/udp2raw"]
