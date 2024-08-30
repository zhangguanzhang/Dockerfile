FROM python:2-slim
RUN set -eux; \
    apt-get -y update; \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        curl; \
    savedAptMark="$(apt-mark showmanual)"; \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        git \
        ca-certificates \
        unzip \
        make \
        gcc \
        libreadline-dev \
        pkg-config \
        ; \
    mkdir /opt/test; \
    cd /opt/test; \
    curl -LS https://github.com/protocolbuffers/protobuf/releases/download/v3.0.2/protoc-3.0.2-linux-x86_64.zip \
        -o protoc-3.0.2-linux-x86_64.zip; \
    unzip protoc-*.zip; \
    cp bin/protoc /usr/bin/; \
    curl -LS https://www.lua.org/ftp/lua-5.1.5.tar.gz -o lua-5.1.5.tar.gz; \
    tar zxf lua-5.1.5.tar.gz; \
    cd lua-*; \
    make linux; \
    make install; \
    cd /usr/local/; \
    git clone https://github.com/sean-lin/protoc-gen-lua; \
    cd protoc-gen-lua/protobuf; \
    make; \
    cd ..; \
    rm -rf .git; \
    cd plugin; \
    ln -sf `readlink -f protoc-gen-lua` /usr/bin/protoc-gen-lua; \
    python2 -m pip install --no-cache protobuf; \
    rm -rf /opt/test; \
    apt-mark auto '.*' > /dev/null; \
	[ -z "$savedAptMark" ] || apt-mark manual $savedAptMark > /dev/null; \
	apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false; \
    apt-get clean; \
    rm -rf /var/cache/apt/* /var/lib/apt/lists/*; \
    apt-get autoclean; \
    apt-get autoremove;

COPY proto-reverse-lua-v2 /usr/bin/
