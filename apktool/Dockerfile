FROM zhangguanzhang/alpine as down

ARG APKTOOL_VERSION=2.5.0

WORKDIR /opt

RUN curl -LO https://raw.githubusercontent.com/iBotPeaches/Apktool/master/scripts/linux/apktool && \
    chmod +x apktool && \
    curl -L -o apktool.jar https://github.com/iBotPeaches/Apktool/releases/download/v${APKTOOL_VERSION}/apktool_${APKTOOL_VERSION}.jar && \
    chmod +x apktool.jar


FROM openjdk:8-alpine

LABEL MAINTAINER zhangguanzhang (zhangguanzhang@qq.com)

COPY --from=down /opt/apktool.jar /opt/apktool /usr/local/bin/

WORKDIR /opt

ENTRYPOINT ["apktool"]
