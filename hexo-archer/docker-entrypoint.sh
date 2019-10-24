#!/bin/sh
set -e
[ -n "$DEBUG" ] && set -x

if [ -n "$EMAIL" ];then
    git config --global user.email "$EMAIL"
fi
if [ -n "$NAME" ];then
    git config --global user.email "$NAME"
fi


count=0
if [ -d '/tmp/blog' ];then 
    count=`ls /tmp/blog | wc -l`
    if [ "$count" -gt 0 ];then
        cp -a /tmp/blog/* /root/blog/
    fi
fi

hexo d -g

if [ "$count" -gt 0 ];then
    \cp /root/blog/db.json /tmp/blog/
fi
