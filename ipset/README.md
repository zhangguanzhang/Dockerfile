
## 使用

```
docker run -ti --name test \
    --net host \
    --cap-add=NET_ADMIN \
    --cap-add=NET_RAW \
    -v  /run/xtables.lock:/run/xtables.lock:rw registry.aliyuncs.com/zhangguanzhang/ipset
```

用 inotifywait 来同步 ipset 规则

```
inotifywait -mrq \
    --timefmt '%d/%m/%y/%H:%M' \
    --format '%T %w %f' \
    -e modify,delete,create,attrib /data/kube/ | while read line;do
echo $line
bash /iptables.sh
done
```