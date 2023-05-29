# https://github.com/jschaf/ci_speed_test/blob/master/docker/debian-ci/Dockerfile
# Builder image for SSH to copy onto the base Debian image.

FROM python:3.9-bullseye AS build-env

ARG ANSIBLE_CORE_VERSION=2.14.6
ARG SSHPASS_VERSION=1.10

# https://pypi.org/project/ansible-core/#history
RUN pip3 install --no-cache-dir --disable-pip-version-check \
       -i http://mirrors.aliyun.com/pypi/simple/ \
       --trusted-host mirrors.aliyun.com \
              ansible-core==${ANSIBLE_CORE_VERSION}; \
    pip3 list | grep ansible;

RUN set -eux; \
    mkdir -p /opt/usr/local/bin/ /opt/usr/local/lib/python3/; \
    cp -a /usr/local/lib/python${PYTHON_VERSION%.*}/site-packages /opt/usr/local/lib/python3/site-packages/; \
    cp /usr/local/bin/ansible* /opt/usr/local/bin/; \
    sed -ri '1s#/.+/python#/usr/bin/python3#' /opt/usr/local/bin/ansible*; 
# 理论上 sshpass 也可以包管理扣 so，但是包管理版本没源码编译新
# 也可以非静态编译，但是不想折腾了
RUN set -eux; \
    wget https://udomain.dl.sourceforge.net/project/sshpass/sshpass/${SSHPASS_VERSION}/sshpass-${SSHPASS_VERSION}.tar.gz; \
    tar zxf sshpass-${SSHPASS_VERSION}.tar.gz && cd sshpass-*; \
    CFLAGS="-static -s" LDFLAGS=-static ./configure; \
    make && make install; \
    cp `which sshpass` /opt/usr/local/bin/;

FROM debian:11 as builder-ssh
RUN set -eux; \
    if [  -e /etc/apt/sources.list ];then sed -ri 's/[a-zA-Z0-9.]+(debian.org|ubuntu.com)/mirrors.aliyun.com/g' /etc/apt/sources.list; fi; \
    apt update; \
    # debian 的清理关闭，下一步解压 deb 包，主要是为了 openssh-client
    echo > /etc/apt/apt.conf.d/docker-clean; \
    DEBIAN_FRONTEND=noninteractive apt -o APT::Keep-Downloaded-Packages="true" install -y openssh-client;
RUN cd /var/cache/apt/archives; \
    apt download \
        libselinux1 \
        libpcre2-8-0 \
        # libk5crypto3 \
        ; \
    ls *.deb | while read deb;do dpkg-deb -x $deb /opt/; done; \
    rm -rf /opt/usr/share/{man,doc} /opt/usr/lib/systemd/;

FROM gcr.io/distroless/python3-debian11:debug
LABEL maintainer=zhangguanzhang@qq.com:zhangguanzhang/Dockerfile
SHELL ["/busybox/sh", "-c"]
COPY --from=build-env /opt /
COPY --from=builder-ssh /opt /
ENV PYTHONPATH=/usr/local/lib/python3/site-packages

RUN set -eux; \
    # https://github.com/ansible/ansible/issues/77840
    # 避免 suse12 出现 py 版本错误
    cd `find /usr -type f -name base.yml -path '*/ansible/config/*' -exec dirname {} \;`; \
    if ! grep 'suse:' base.yml;then \
        #cd /usr/lib/python3/dist-packages/ansible/config/; \
        cp base.yml base.yml.bak; \
        sed_num=$(awk '/INTERPRETER_PYTHON_DISTRO_MAP:/,/default:/{a=NR}END{print a+1}'  base.yml); \
        sed -ri $sed_num"i \ \ \ \ suse:\n      '12': /usr/bin/python2.7"  base.yml; \
    fi; \
    mkdir -p /etc/ansible/; \
    { echo '[local]'; \
    echo 'localhost ansible_connection=local'; } >> /etc/ansible/hosts; \
    ansible all -m setup | tail -n5; \
    ansible all -m shell -a 'pwd'; \
    > /etc/ansible/hosts; \
    sshpass -V; \
    ssh -V; \
    ansible --version
