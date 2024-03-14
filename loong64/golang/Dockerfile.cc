ARG IMG=debian:12-slim
FROM ${IMG}

RUN if [ -f /etc/apt/sources.list ];then sed -ri 's/(ports|deb|security|archive).(debian.org|ubuntu.com)/mirrors.aliyun.com/g' /etc/apt/sources.list; fi; \
    set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        cmake \
        curl \
        gnupg \
        automake \
        file \
        git \
        dpkg-dev \
        make pkg-config \
        wget xz-utils; \
    apt clean; \
    rm -rf /var/cache/apk/* /tmp/* /var/cache/apt/* /var/lib/apt/lists/*; 

# http://www.loongnix.cn/zh/toolchain/GNU/
RUN set -eux; \
    cd /; \
    wget http://ftp.loongnix.cn/toolchain/gcc/release/loongarch/gcc8/loongson-gnu-toolchain-8.3-x86_64-loongarch64-linux-gnu-rc1.1.tar.xz; \
    tar xfJp loongson-gnu-toolchain-*.tar.xz; \
    rm -f loongson-gnu-toolchain-*.tar.xz; \
    mv loongson-gnu-toolchain-*-x86_64-* loongson-gnu-toolchain

ENV PATH /loongson-gnu-toolchain/bin:$PATH
