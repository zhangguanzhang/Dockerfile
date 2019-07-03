#!/bin/sh


set -e
[[ "${DEBUG}" == "true" ]] && set -x

echo 'option docs|https://manpages.debian.org/stretch/tftpd-hpa/tftpd.8.en.html'

syslogd -n -O /dev/stdout &

if [ "${1:0:1}" = '-' ];then
    exec in.tftpd "$@"
else
    exec "$@"
fi
