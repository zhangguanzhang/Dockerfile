# 改善镜像fauria/vsftpd

This Docker container implements a vsftpd server, with the following features:

 * Centos 7 base image.
 * vsftpd 3.0
 * Virtual users
 * Passive mode
 * Logging to a file or STDOUT.

### Installation from [Docker registry hub](https://registry.hub.docker.com/u/zhangguanzhang/vsftpd/).

You can download the image with the following command:

```bash
docker pull zhangguanzhang/vsftpd
```

Environment variables
----

This image uses environment variables to allow the configuration of some parameteres at run time:

----

* Variable name: `PASV_ADDRESS`
* Default value: Docker host IP.
* Accepted values: Any IPv4 address.
* Description: If you don't specify an IP address to be used in passive mode, the routed IP address of the Docker host will be used. Bear in mind that this could be a local address.

----

* Variable name: `PASV_MIN_PORT`
* Default value: 21100.
* Accepted values: Any valid port number.
* Description: This will be used as the lower bound of the passive mode port range. Remember to publish your ports with `docker -p` parameter.

----

* Variable name: `PASV_MAX_PORT`
* Default value: 21110.
* Accepted values: Any valid port number.
* Description: This will be used as the upper bound of the passive mode port range. It will take longer to start a container with a high number of published ports.

----

* Variable name: `LOG_STDOUT`
* Default value: null.
* Accepted values: Any string to enable, empty string or not defined to disable.
* Description: Output vsftpd log through STDOUT, so that it can be accessed through the [container logs](https://docs.docker.com/reference/commandline/logs/).

----

Exposed ports and volumes
----

The image exposes ports `20` and `21`. Also, exports two volumes: `/home/vsftpd`, which contains users home directories, and `/var/log/vsftpd`, used to store logs.

When sharing a homes directory between the host and the container (`/home/vsftpd`) the owner user id and group id should be 14 and 80 respectively. This correspond ftp user and ftp group on the container, but may match something else on the host.

Use cases
----

1) Create a temporary container for testing purposes:

```bash
  docker run --rm -v $PWD/virtual_users.txt:/etc/vsftpd/virtual_users.txt zhangguanzhang/vsftpd
```

2) Create a container in active mode using the default user account, with a binded data directory:

```bash
$ cat virtual_users.txt 
myuser
mypass
$ docker run -d -p 21:21 -v /my/data/directory:/home/vsftpd -v $PWD/virtual_users.txt:/etc/vsftpd/virtual_users.txt --name vsftpd zhangguanzhang/vsftpd
# see logs for credentials:
docker logs vsftpd
```

3) Create a **production container** with a custom user account, binding a data directory and enabling both active and passive mode:

```bash
$ cat virtual_users.txt 
myuser
mypass
$ docker run -d -v /my/data/directory:/home/vsftpd \
-p 20-21:20-21 -p 21100-21110:21100-21110 \
-e PASV_ADDRESS=127.0.0.1 -v $PWD/virtual_users.txt:/etc/vsftpd/virtual_users.txt \
--name vsftpd --restart=always zhangguanzhang/vsftpd
```

4) Manually add a new FTP user to an existing container:
```bash

docker exec -i -t vsftpd bash
mkdir /home/vsftpd/myuser
echo -e "myuser\nmypass" >> /etc/vsftpd/virtual_users.txt
/usr/bin/db_load -T -t hash -f /etc/vsftpd/virtual_users.txt /etc/vsftpd/virtual_users.db
exit
docker restart vsftpd
```
