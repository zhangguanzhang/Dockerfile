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