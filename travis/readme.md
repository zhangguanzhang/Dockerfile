不想登陆的话可以先主进程sh进去travis login登陆完把家目录的/root/.travis/config.yml用docker cp出来或者挂载出来
后面直接
```bash
[root@k8s-n3 travis]# docker run --rm -v $PWD/config.yml:/root/.travis/config.yml zhangguanzhang/travis show -r zhangguanzhang/gcr.io
Job #17.1:  Update .travis.yml
State:         errored
Type:          push
Branch:        master
Compare URL:   https://github.com/zhangguanzhang/gcr.io/compare/8f0d32c322b394aa9******33d07289519****038e...533a0d97a4a0cf8a786be8ec7d2b58d944f55d03
Duration:      49 min 21 sec
Started:       2018-07-17 09:48:42
Finished:      2018-07-17 10:38:03
Allow Failure: false
Config:        os: linux, python: 2.7
[root@k8s-n3 travis]# docker run --rm -v $PWD/config.yml:/root/.travis/config.yml zhangguanzhang/travis restart -r zhangguanzhang/gcr.io
restarted
[root@k8s-n3 travis]# docker run --rm -v $PWD/config.yml:/root/.travis/config.yml zhangguanzhang/travis status -r zhangguanzhang/gcr.io
created
[root@k8s-n3 travis]# docker run --rm -v $PWD/config.yml:/root/.travis/config.yml zhangguanzhang/travis status -r zhangguanzhang/gcr.io
started
```
