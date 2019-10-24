
change from https://github.com/hashicorp/docker-consul/blob/master/0.X/docker-entrypoint.sh
- `CONSUL_BIND_INTERFACE`: which ip will use
- `CONSUL_CLIENT_INTERFACE`: 
- `CONSUL_LOCAL_CONFIG`: string config
- `CONFIG_OWNER`: chown to consul:consul if this not null
- `RETRY_JOIN`: if cmd is not empty to generate 
```
$1=agent RETRY_JOIN=192.168.1.1,192.168.1.2,192.168.1.3
$@=consul agent ... -retry-join=192.168.1.1 -retry-join=192.168.1.2 -retry-join=192.168.1.3
```
