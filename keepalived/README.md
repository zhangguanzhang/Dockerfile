## 版本

版本:
- `v2.0.20`
- `v2.2.0`

拉取:

```
 docker pull zhangguanzhang/keepalived:v2.2.0
 docker pull registry.aliyuncs.com/zhangguanzhang/keepalived:v2.2.0
```

## keepalived

### ipvs svc 的参考

```yaml
version: '3.5'
services:
  keepalived: 
    image: 'registry.aliyuncs.com/zhangguanzhang/keepalived:v2.2.0'
    hostname: 'keepalived-ipvs'
    restart: unless-stopped
    container_name: "keepalived-ipvs"
    labels: 
      - app=keepalived
    network_mode: host
    privileged: true
    cap_drop:
      - ALL
    cap_add:
      - NET_BIND_SERVICE
    volumes:
      - /usr/share/zoneinfo/Asia/Shanghai:/etc/localtime:ro
      - /lib/modules:/lib/modules
      - /run/xtables.lock:/run/xtables.lock
      - ./conf.d/:/etc/keepalived/conf.d/
      - ./keepalived.conf:/etc/keepalived/keepalived.conf
      - ./always-initsh.d:/always-initsh.d
    command: 
      - --dont-fork
      - --log-console
      - --log-detail
      - --use-file=/etc/keepalived/keepalived.conf
    logging:
      driver: json-file
      options:
        max-file: '3'
        max-size: 20m
```


### 我博客的 dns 方案的话

前台主要是运行的几个选项，exec 让 keepalived 主进程能够感知到信号，--net=host 运行，lvs 的话可以下面类似


```yaml
...
        command:
        - keepalived
        - --dont-fork
        - --log-console 
        - --log-detail
        - --all
        - -f
        - /etc/keepalived/keepalived.conf
...
        volumeMounts:
        - mountPath: /etc/localtime
          name: host-localtime
        - name: config-volume
          mountPath: /etc/keepalived
      volumes:
      - name: config-volume
        configMap:
          name: node-local-dns
...
apiVersion: v1
kind: ConfigMap
metadata:
  name: node-local-dns
  namespace: kube-system
  labels:
      addonmanager.kubernetes.io/mode: EnsureExists
data:
  keepalived.conf: |
    global_defs {
      router_id LVS_DEVEL
    }
    virtual_server {{ CLUSTER_DNS_SVC_IP }} 53 {
        delay_loop 2
        lb_algo rr
        lb_kind NAT
        protocol TCP
{% for host in groups['kube_master'] %}
        real_server {{ host }} {{ coredns_port }} {
            weight 1
            HTTP_GET {
                url {
                  path /health
                  status_code 200
                }
                connect_port    {{ coredns_healthz_port }}
                connect_timeout 1
                retry 1
                delay_before_retry 2
            }
        }
{% endfor %}
    }

    virtual_server {{ CLUSTER_DNS_SVC_IP }} 53 {
        delay_loop 1
        lb_algo rr
        lb_kind NAT
        protocol UDP
{% for host in groups['kube_master'] %}
        real_server {{ host }} {{ coredns_port }} {
            weight 1
            HTTP_GET {
                url {
                  path /health
                  status_code 200
                }
                connect_port    {{ coredns_healthz_port }}
                connect_timeout 1
                retry 1
                delay_before_retry 2
            }
        }
{% endfor %}

    }
```