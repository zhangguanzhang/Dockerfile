## about

mysql8 官方镜像开始去掉了 debian，只基于 oraclelinux 制作，并且 8.4 开始没有内置 `mysqlbinlog` 命令了，尝试验证后可行，主要步骤：

### repo

根据官方 Dockerfile 内添加的 repo `https://repo.mysql.com/yum/mysql-8.4-community/docker/el/9/$basearch/` 上级 web 路由，找到了 mysqlbinlog 命令属于 `mysql-community-client-$MYSQL_VERSION` ，所以添加 repo：

```shell
$ cat /etc/yum.repos.d/mysql-community.repo 
[mysql8.4-server]
name=MySQL 8.4 Server
enabled=1
baseurl=https://repo.mysql.com/yum/mysql-8.4-community/el/9/$basearch/
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql
module_hotfixes=true
```

### RUN 修改

然后发现安装 `mysql-community-client-$MYSQL_VERSION` 会报错文件冲突，理论上覆盖的话也没问题，但是优先保证 mysql-server，所以 Dockerfile 内改为先安装 client，最后再 rpm --force 安装 mysql-server 覆盖 mysql-client，测试 `mysqlbinlog` 没问题。

### 结果

```
$ docker images | grep 8.4
mysql                                                              8.4.6                                            1e55ff196100   2 days ago       831MB
$ docker run --rm -ti --entrypoint bash mysql:8.4.6
bash-5.1# mysqld --help
mysqld  Ver 8.4.6 for Linux on x86_64 (MySQL Community Server - GPL)
BuildID[sha1]=0d785fe97d9ab7838fcd9648862a50fbd9b30bdc
Copyright (c) 2000, 2025, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Starts the MySQL database server.

Usage: mysqld [OPTIONS]

For more help options (several pages), use mysqld --verbose --help.
bash-5.1# mysqlbinlog --version
mysqlbinlog  Ver 8.4.6 for Linux on x86_64 (MySQL Community Server - GPL)
```
