FROM debian:stable-slim

ENV TZ Asia/Shanghai

RUN if [  -e /etc/apt/sources.list ];then sed -ri 's/(deb|security).debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list; fi && \
    apt-get update && \
    apt-get install -y  --no-install-recommends \
        curl \
        iproute2 \
        psmisc \
        procps \
        net-tools \
        tzdata && \
    apt-get clean && \
    rm -rf /var/cache/apt/* /var/lib/apt/lists/* && \
    if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns myhostname' > /etc/nsswitch.conf; fi  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo Asia/Shanghai > /etc/timezone
