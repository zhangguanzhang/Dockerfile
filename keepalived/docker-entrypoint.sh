#!/bin/bash

chmod 644 /etc/keepalived/keepalived.conf

if [ -f /run/keepalived/keepalived.pid ];then
    > /run/keepalived/keepalived.pid
fi

if [ -f /run/keepalived/vrrp.pid ];then
    > /run/keepalived/vrrp.pid
fi

exec keepalived --dont-fork --log-console --log-detail --vrrp  -f /etc/keepalived/keepalived.conf
