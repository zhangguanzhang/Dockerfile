FROM debian:stable-slim

ENV TZ Asia/Shanghai

RUN if [  -e /etc/apt/sources.list ];then sed -ri 's/[a-zA-Z0-9.]+(debian.org|ubuntu.com)/mirrors.aliyun.com/g' /etc/apt/sources.list; fi && \
    if [  -e /etc/apk/repositories ];then sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories; fi && \
    if command -v apk 2> /dev/bull 1> /dev/null; then \
      apk update && \
      apk add --no-cache \
        curl \
        ca-certificates \
        bash \
        iproute2 \
        tzdata; \
      fi && \
    if command -v apt 2> /dev/bull 1> /dev/null; then \
      apt-get update && \
      apt-get install -y \
        curl \
        tzdata && \
      apt-get clean; \
      fi && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone && \
    if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns myhostname' > /etc/nsswitch.conf; fi && \
   rm -rf /var/cache/apk/* /tmp/* /var/cache/apt/* /var/lib/apt/lists/*
