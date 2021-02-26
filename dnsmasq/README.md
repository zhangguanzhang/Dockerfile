## docker-dnsmasq

参考 [andyshinn/docker-dnsmasq](https://github.com/andyshinn/docker-dnsmasq) , 配置文件可以:

```conf
no-resolv
all-servers
{% for host in groups['kube_master'] %}
server={{ host }}#{{ coredns_port }}
{% endfor %}
#log-queries
```

```yaml
...
        command:
        - dnsmasq
        - -d
        - --conf-file=/etc/dnsmasq/dnsmasq.conf
```

### 镜像

```
registry.aliyuncs.com/zhangguanzhang/dnsmasq:2.83
```

相关链接：
- http://www.thekelleys.org.uk/dnsmasq/