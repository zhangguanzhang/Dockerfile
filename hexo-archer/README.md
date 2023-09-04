
# hexo的archer容器化

## 过程

假设目录为 `/root/blog` ，目录树大体为下面
```
.
├── .deploy_git
├── data
│   ├── _config.yml
│   ├── env  #环境变量文件，里面写 shell 的 export xxx=xxxxx
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

- public 是缓存运行中生成的静态文件，防止多次生成，不用备份
- data 是基于 hexo 的 init 目录下需要备份的所有文件，markdown 和配置文件还有图片啥的，需要覆盖默认的，所以先挂载到容器里非 workdir 后通过 `entrypoint.sh` 拷贝覆盖

## 运行

把相关认证文件挂载进去会自动认证，不要在 hexo init 的目录后执行下面的运行，否则 node_modules 会冲突报错 hexo 奇怪的问题，需要严格按照上面目录存放后执行

```
docker run --rm -ti \
    -v ~/blog/data:/tmp/blog \
    -v ~/.ssh:/root/.ssh/ \
    -v ~/.gitconfig:/root/.gitconfig \
    -v ~/blog/public:/root/blog/public  \
    -v ~/blog/.deploy_git:/root/blog/.deploy_git \
    zhangguanzhang/hexo-archer
```

最后建议在每次发布博客之前用 shell 脚本检查你的关键字，防止某些关键字带上去。

```bash
function hexo () 
{ 
    if grep --color=auto -Eq 'xxx|xxxx|xxxxx|xxxxx' /root/blog/data/source/_posts/*; then
        echo 'keyword matched, will not run';
        return;
    fi;
    docker run --rm -ti -v ~/blog/data:/tmp/blog -v ~/.ssh:/root/.ssh/ -v ~/.gitconfig:/root/.gitconfig -v ~/blog/public:/root/blog/public -v ~/blog/.deploy_git:/root/blog/.deploy_git zhangguanzhang/hexo-archer
}
```

## cdn.jsdelivr.net/npm 问题

在 `~/blog/data/env` 里追加下面内容

```
find /root/blog/themes/archer/layout/ -type f -name '*.ejs' -exec sed -ri 's#cdn.jsdelivr.net/npm#unpkg.zhimg.com#g' {} \;
```
