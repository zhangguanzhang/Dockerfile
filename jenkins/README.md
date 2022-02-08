
## RUN

有些版本的 docker 组 id 不是 995，自行更改 Dockerfile

### 构建

```bash
docker build \
    --build-arg DOCKER_GID=994 \
    --build-arg DOCKER_CLIENT=docker-20.10.12.tgz \
    -t registry.aliyuncs.com/zhangguanzhang/jenkins:lts-docker .
```

### 使用

```bash
mkdir data
chown -R 1000:1000 data
docker-compose up -d
```
