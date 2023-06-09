## about

不要 qemu 构建，会报错。龙芯没提供 openjdk:11-slim-buster，自己制作个，参考 [Loongson-Cloud-Community/dockerfiles](https://github.com/Loongson-Cloud-Community/dockerfiles/blob/main/library/openjdk/11-buster/Makefile)

```
$ docker build -t zhangguanzhang/openjdk:11-slim-buster-loong64  \
    --build-arg DOWNLOAD_URL=http://ftp.loongnix.cn/Java/openjdk11/loongson11.5.0-fx-jdk11.0.19_7-linux-loongarch64.tar.gz  \
    --build-arg JAVA_VERSION=11.5.0 \
    .
```
