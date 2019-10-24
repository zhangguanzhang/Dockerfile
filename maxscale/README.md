
https://www.imooc.com/article/23664
maxscale的sql更改导致无法连接，最后一个未更改的版本为 https://github.com/mariadb-corporation/MaxScale/blob/2.1.13/server/modules/authenticator/MySQLAuth/dbusers.c#L45
而该版本dockerhub上官方没有提供并且apt的源里2.1.13版本为非release，且源码编译失败，Dockerfile.2.1.13此镜像是下载deb包安装。

主文件是修改自官方的，增加了curl和jq来后续给rest api做健康检查做准备
