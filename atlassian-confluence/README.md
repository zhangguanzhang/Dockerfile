配置建议高些,1c2g带不起来坑了我一天,4c8g测试正常

https://hub.docker.com/r/cptactionhank/atlassian-confluence/dockerfile
改的上面这个镜像,相关问就在我github上 https://github.com/zhangguanzhang/Dockerfile/blob/master/atlassian-confluence
大致和jira一样,破解在网上搜到了两种方式
1. 是市面上替换decoder那个jar的流程,路径为`/opt/atlassian/confluence/confluence/WEB-INF/lib/atlassian-extras-decoder-api-3.4.1.jar`。市面上是把安装完路径的jar拿出来用注册机打pathc后换回去
2. 我在市面上下载的破解包里有个类似jira的破解jar文件`atlassian-extras-3.2.jar`,网上搜到了另一种破解思路是这个文件扔进去前删掉`atlassian-extras*.jar`.

3.变量`$JVM_MEMORY`可以设置为`-Xms1024m -Xmx1024m`，但是这个基础镜像的jdk好像可以自动识别cgroup限制

4. 这里找 https://mritd.me/ 漠然大佬帮我把注册机生成许可和打patch做成了不需要图形界面的cli工具用法是自行看命令帮助`atlassianctl --help`,用cli打patch和获取注册码,官网申请的使用许可证不行,因为反编译的时候改了签名
confluence的compose.yml如下
```yaml
version: '3.7'
services:
  confluence:
    image: zhangguanzhang/atlassian-confluence:6.14.1
    container_name: confluence
    hostname: confluence
    init: true
    volumes:
      - CONF_HOME_data:/var/atlassian/confluence
      - CONF_log_data:/opt/atlassian/confluence/logs
      - /etc/localtime:/etc/localtime:ro
    ports:
      - '8090:8090'
    logging:
      driver: json-file
      options:
        max-file: '3'
        max-size: 100m
volumes:
  CONF_HOME_data: {}
  CONF_log_data: {}
```
起来后访问ip:8090进到初次的页面设置
![9](https://images2017.cnblogs.com/blog/907596/201709/907596-20170928163446169-500307045.jpg)

![9](https://images2017.cnblogs.com/blog/907596/201709/907596-20170928163453434-1803555365.jpg)

下面两个不要勾选
![9](https://images2017.cnblogs.com/blog/907596/201709/907596-20170928163715059-1486445585.jpg)


然后用命令生成注册码填进去即可
`docker exec confluence atlassianctl license -s <id>`
设置完后admin登陆进去后在在右上角的齿轮小图标里一般设置左侧栏里往下翻到授权细节查看如下图所示:

![6](https://github.com/zhangguanzhang/Image-Hosting/blob/master/atlassian/6.png?raw=true)

和jira一样修改数据库连接参数,文件路径为`/var/lib/docker/volumes/confluence_CONF_HOME_data/_data/confluence.cfg.xml`

`jdbc:mysql://10.20.4.17:3306/confluence`改为`jdbc:mysql://10.20.4.17:3306/confluence?useSSL=false`改完后重启下容器即可
