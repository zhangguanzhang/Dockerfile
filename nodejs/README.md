# about this

I'm trying to run a nodejs project with docker, but the project needs to first npm install the dependencies in the package.json file while the node is running, and the package.json file needs to be provided by the user ,so I'm going to make a generic image and I reference The official dockerfile for node image adds the entrypoint section to automate npm install and npm run start for common

我尝试用docker运行一个nodejs项目，但是在项目中node运行起来之前需要先npm install文件package.json里的依赖，而package.json文件需要用户提供，我打算制作一个通用型镜像，所以我参照了node镜像的官方dockerfile增加了entrypoint部分来自动执行npm install和npm run start从而做到通用

## what did I do

> * I added the following to node's official dockerfile(我在node官方的dockerfile里增加了下面内容)
```
WORKDIR /home/node
COPY entrypoint.sh /usr/local/bin/
ENTRYPOINT ["entrypoint.sh"]
```
> * and I wrote bash script as entrypoint and the docker-compose.yml

## how to start and use

-------------------------------------------------------------------------
**whatever,the app folder just for reference only,you should use yours**.
-------------------------------------------------------------------------
You can use either part of it or all the part (你可以使用其中的一部分，也可以整个部分)  
if you use the part of this(just only dockerfile):  
>1. only download the node folder  
>2. use the  docker to build it,After building you will get a image, such as nodejs  
>3. Put your node project into a folder, such as app  
>4. cd the app directory,you must mount the all file into the /home/node/ when you run the docker command to run a container.for example：  
>    
```
docker run -d -v $PWD:/home/node/ nodejs npm run start  
```

if you use the docker-compose to run a node project(You must use docker-compose like me)：  
>1. change the docker-compose.yml file and the entrypoint.sh if you want to change it  
>2. run the command : 
```
docker-compose up -d
```
>3. just wait a few secends,enjoy it  
