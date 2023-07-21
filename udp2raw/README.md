## 关于

https://github.com/wangyu-/udp2raw-tunnel/blob/master/doc/README.zh-cn.md#%E8%BF%90%E8%A1%8C

## 构建

```
docker buildx build --platform linux/amd64,linux/arm64 \
    --push --progress plain \
    -t zhangguanzhang/udp2raw:20230206.0 --build-arg VERSION=20230206.0 .
```

## 运行

在server端运行:

```shell
docker run \
    -d --name udp2raw  \
    --net host \
    --cap-add NET_RAW \
    --cap-add NET_ADMIN  \
    zhangguanzhang/udp2raw \
    -s -l 0.0.0.0:86 \
    -r 127.0.0.1:19919 \
    -k passwd \
    --raw-mode faketcp  \
    --cipher-mode xor  -a
```

在client端运行:

```
docker run --net host \
    -d --name udp2raw  \
    --cap-add NET_RAW \
    --cap-add NET_ADMIN    \
    zhangguanzhang/udp2raw \
    -c -l 0.0.0.0:19919 \
    -r <public_ip>:86 \
    -k passwd \
    --raw-mode faketcp   \
    --cipher-mode xor  -a
```

windows 客户端

```shell
 ./udp2raw_mp.exe -c -l 0.0.0.0:19919 -r <public_ip>:86 -k passwd --raw-mode faketcp --cipher-mode xor
```

用在 `wg` 伪装的话，`windows` 客户端 `endpoint` 的连接 `ip` 不能写 127.0.0.1 , 必须写网卡`ip`，否则会下面这样

```log
2021-01-06 17:44:29.897: [TUN] [4-gw] peer(vpTm…Rxy8) - Failed to send handshake initiation write udp4 0.0.0.0:58007->127.0.0.1:19919: wsasendto: The requested address is not valid in its context.
```

另外 `wg` 的 MTU 最小是 1280，`windows` 客户端少于这个值会无法启动