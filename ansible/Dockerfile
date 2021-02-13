FROM alpine:3.13

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --update --no-cache add \
        ca-certificates \
        git \
        openssh-client \
        openssl \
        python3\
        py3-pip \
        py3-cryptography \
        rsync \
        sshpass

RUN apk --update add --virtual \
        .build-deps \
        python3-dev \
        libffi-dev \
        openssl-dev \
        build-base \
        curl \
 && pip3 install --upgrade \
        pip \
        cffi -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com \
 && pip3 install \
        ansible \
        ansible-lint -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com \
 && apk del \
        .build-deps \
 && rm -rf /var/cache/apk/* \
&& mkdir -p /etc/ansible \
 && echo 'localhost' > /etc/ansible/hosts \
 && echo -e """\
\n\
Host *\n\
    StrictHostKeyChecking no\n\
    UserKnownHostsFile=/dev/null\n\
""" >> /etc/ssh/ssh_config
