FROM alpine:3.13

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns' > /etc/nsswitch.conf; fi

RUN apk add -u \
    iptables \
    ipset \
    inotify-tools \
    bash \
    --no-cache 
    
CMD [ "bash" ]
