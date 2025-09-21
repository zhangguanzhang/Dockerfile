ARG IMG=debian:11
FROM ${IMG}

RUN if [ -f /etc/apt/sources.list ];then sed -ri 's/(ports|deb|security|archive).(debian.org|ubuntu.com)/mirrors.aliyun.com/g' /etc/apt/sources.list; fi; \
    set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        gnupg \
        automake \
        file \
        git \
        dpkg-dev \
        pigz \
        make pkg-config \
        wget xz-utils; \
    apt clean; \
    rm -rf /var/cache/apk/* /tmp/* /var/cache/apt/* /var/lib/apt/lists/*; 

ENV PATH /usr/local/go/bin:$PATH
ENV GOLANG_VERSION 1.23.12
#ENV GOLANG_VERSION 1.20.12

RUN set -eux; \
    # http://www.loongnix.cn/zh/toolchain/Golang/
    wget http://ftp.loongnix.cn/toolchain/golang/go-${GOLANG_VERSION%.*}/abi1.0/go${GOLANG_VERSION}.linux-amd64.tar.gz; \
    tar zxf go*.linux-amd64.tar.gz -C /usr/local/; \
    rm -f go*.linux-amd64.tar.gz; \
    rm -rf /var/cache/apk/* /tmp/* /var/cache/apt/* /var/lib/apt/lists/*; \
    go version

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

# # http://www.loongnix.cn/zh/toolchain/GNU/
# RUN set -eux; \
#     cd /; \
#     wget http://ftp.loongnix.cn/toolchain/gcc/release/loongarch/gcc8/loongson-gnu-toolchain-8.3-x86_64-loongarch64-linux-gnu-rc1.1.tar.xz; \
#     tar xfJp loongson-gnu-toolchain-*.tar.xz; \
#     rm -f loongson-gnu-toolchain-*.tar.xz; \
#     mv loongson-gnu-toolchain-*-x86_64-* loongson-gnu-toolchain

# ENV PATH /loongson-gnu-toolchain/bin:$PATH

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct

RUN set -eux; \
    mkdir -p ~/.ssh/; \
    chmod 0700 ~/.ssh/; \
    touch ~/.ssh/known_hosts ~/.ssh/authorized_keys; \
    chmod 0600 ~/.ssh/known_hosts ~/.ssh/authorized_keys; 
