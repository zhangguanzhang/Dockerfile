
# hexo的archer容器化

## 过程

假设目录为`/root/blog`，目录树大体为下面
```
.
├── .deploy_git
├── data
│   ├── _config.yml
│   ├── source
│   │   ├── bash
│   │   ├── curl-to-go
│   │   ├── json-to-go
│   │   └── _posts
│   └── themes
│       └── archer
|           ├── _config.yml
│           └── source
│               ├── assets
│               └── avatar
└── public
```

- public是缓存运行中生成的静态文件，防止多次生成，不用备份
- data是基于hexo的init目录下需要备份的所有文件，markdown和配置文件还有图片啥的，需要覆盖默认的，所以先挂载到容器里非workdir后通过`entrypoint.sh`拷贝覆盖

## 运行

把相关认证文件挂载进去会自动认证
```
docker run --rm -ti \
    -v ~/blog/data:/tmp/blog \
    -v ~/.ssh:/root/.ssh/ \
    -v ~/.gitconfig:/root/.gitconfig \
    -v ~/blog/public:/root/blog/public  \
    -v ~/blog/.deploy_git:/root/blog/.deploy_git \
    zhangguanzhang/hexo-archer
```
