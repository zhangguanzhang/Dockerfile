## 版本

```
export VERSION=0.14.01
docker build --build-arg VERSION=V"${VERSION}" -t zhangguanzhang/stress-ng:${VERSION} .
docker buildx build --build-arg VERSION=V"${VERSION}" \
    -t registry.aliyuncs.com/zhangguanzhang/stress-ng:${VERSION} .  \
     --push   --platform linux/amd64,linux/arm64
```

版本:

- `0.14.01`
- `0.13.05`
- `0.13.03`

拉取:

```
 docker pull zhangguanzhang/stress-ng:0.14.01
 docker pull registry.aliyuncs.com/zhangguanzhang/stress-ng:0.14.01
```

测试参考:

- https://cloud.tencent.com/developer/article/1513544
