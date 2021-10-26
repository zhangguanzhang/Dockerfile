## 版本

```
export VERSION=0.13.05
docker build --build-arg VERSION=V"${VERSION}" -t zhangguanzhang/stress-ng:${VERSION} .
docker buildx build --build-arg VERSION=V"${VERSION}" \
    -t registry.aliyuncs.com/zhangguanzhang/stress-ng:${VERSION} .  \
     --push   --platform linux/amd64,linux/arm64
```

版本:
- `0.13.05`

拉取:

```
 docker pull zhangguanzhang/stress-ng:0.13.05
 docker pull registry.aliyuncs.com/zhangguanzhang/stress-ng:0.13.05
```

测试参考:
- https://cloud.tencent.com/developer/article/1513544
