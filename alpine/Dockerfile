FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache curl busybox-extras unzip && \
    rm -rf /var/cache/apk/* /tmp/* && \
    if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns' > /etc/nsswitch.conf; fi
