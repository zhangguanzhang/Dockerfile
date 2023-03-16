ARG ALPINE_VERSION=3.16
FROM alpine:${ALPINE_VERSION}

LABEL MAINTAINER="zhangguanzhang zhangguanzhang@qq.com" \
    URL="https://github.com/zhangguanzhang/Dockerfile/tree/master/keepalived"

ENV TZ Asia/Shanghai

COPY docker-entrypoint.sh /entrypoint.sh

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache keepalived curl bash iproute2 grep sed ipvsadm ipset tcpdump tzdata libnl && \
    # 避免容器内和宿主机使用 iptables-nft iptables-legacy 模式不一致造成网络问题
    curl -L https://raw.githubusercontent.com/kubernetes-sigs/iptables-wrappers/master/iptables-wrapper-installer.sh > /iptables-wrapper-installer.sh && \
    bash /iptables-wrapper-installer.sh && \
    chmod a+x /entrypoint.sh && \
    mkdir -p /etc/keepalived && \
    echo Asia/Shanghai > /etc/timezone && \
    if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns myhostname' > /etc/nsswitch.conf; fi && \
    rm -rf /var/cache/apk/* /tmp/*

ENTRYPOINT ["/entrypoint.sh"]
