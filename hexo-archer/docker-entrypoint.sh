#!/bin/sh
set -e
[ -n "$DEBUG" ] && set -x

if [ -n "$EMAIL" ];then
    git config --global user.email "$EMAIL"
fi
if [ -n "$NAME" ];then
    git config --global user.name "$NAME"
fi


count=0
if [ -d '/tmp/blog' ];then
    if [ -f /tmp/blog/env ];then
        source /tmp/blog/env
    fi
    count=`ls /tmp/blog | wc -l`
    if [ "$count" -gt 0 ];then
        cp -a /tmp/blog/* /root/blog/
    fi
fi

# 修复 fatal: not a git repository (or any parent up to mount point /root/blog)
if [ -d /root/blog/.deploy_git/ ];then
    if [ ! -d /root/blog/.deploy_git/.git/ ];then
        cd /root/blog/.deploy_git/
        git init
        cd -
    fi
fi

hexo d -g

if [ "$count" -gt 0 ];then
    \cp /root/blog/db.json /tmp/blog/
fi
