FROM ubuntu:20.04 as build
ARG VERSION=V0.14.01
RUN  if [ -f /etc/apt/sources.list ];then sed -ri 's/(ports|deb|security|archive).(debian.org|ubuntu.com)/mirrors.aliyun.com/g' /etc/apt/sources.list; fi && \
    apt-get update && apt-get install -y make gcc git

WORKDIR /stress
# https://kernel.ubuntu.com/~cking/tarballs/
RUN git clone  https://github.com/ColinIanKing/stress-ng.git -b ${VERSION}

RUN cd stress-ng/ && \
    STATIC=1 make

FROM alpine:3.14

COPY --from=build /stress/stress-ng/stress-ng /

RUN if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns myhostname' > /etc/nsswitch.conf; fi 

ENTRYPOINT ["/stress-ng"]
