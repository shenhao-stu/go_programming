# Docker

## MySQL

### 创建MySQL镜像

```bash
docker pull mysql
```

### 查看MySQL镜像

```bash
docker images
```

### 创建MySQL容器

```bash
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pwd --name mysql CONTANINER
```

- **-d :**代表后台运行容器
- **-p 宿主机端口:容器端口 :**将宿主机和容器端口进行映射
- **–-name 容器名称 :**指定容器名称

```bash
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pwd -v /root/docker/mysql/data:/var/lib/mysql --name mysql mysql:latest
```

### 查看正在运行的容器

```bash
docker ps [-qa]
```

- **-a :**查看全部的容器，包括没有运行
- **-q :**只查看容器得到标识

### 查看容器的日志

```bash
docker logs -f CONTANINER
```

- **-f :**可以滚动查看日志的最后几行

## Redis

### 简单的启动一个Redis
使用docker可以简单的启动一个Redis，只需要从docker仓库中拉取到redis的镜像，然后直接运行即可，依次执行如下脚本。
```bash
# 从docker仓库中拉取redis最新的镜像
docker pull redis
# 运行一个redis容器
docker run --name redis -p 6380:6379 -d redis
# 运行redis容器，并指定密码
docker run --name redis -p 6380:6379 -d redis --requirepass "123456"
```

启动好了之后，就可以通过redis的客户端进行操作了，使用exec命令可以打开docker容器。
```bash
# 打开redis客户端；redis表示的是容器名称；redis-cli是容器里面的小型linux系统指令，用来打开redis客户端
docker exec -it redis redis-cli
```

- **-i:**表示交互式的运行

- **-t:**表示分配可以输入指令的终端。

  两者一般是一同使用的。

### redis的持久化

#### redis.conf下载

redis的每个版本都有一个默认的`redis.conf`配置文件，这个配置文件可以在[官网](http://download.redis.io/releases/)去下载，也可以下载对应版本的tar包，从tar包中获取。

在下载之前，先查看docker容器中的redis版本，打开redis-cli，通过info指令可以获取到版本号。

```bash
(base) root@ubuntu:~# docker exec -it gitmodel_redis redis-cli
127.0.0.1:6379> info
NOAUTH Authentication required.
127.0.0.1:6379> auth password
OK
127.0.0.1:6379> info
# Server
redis_version:6.2.6
```

> 注：redis.conf的版本需要和redis容器的版本一致，否则可能会启动失败。

现在有了redis.conf文件，要完成将rdb和aof文件保存到宿主机中，还需要两个步骤：

1. 准备两个文件路径，一个存放redis.conf文件，一个用于存放持久化文件，并修改redis.conf 配置持久化策略。
2. 将redis容器中的配置文件和持久化文件路劲映射到上面准备的路径中。

#### 配置持久化策略

先准备两个文件路径以备用，后面的映射需要用到：


- 配置文件路径：/root/docker/redis/conf
- 持久化文件路径：/root/docker/redis/data

然后将下载好的`redis.conf`文件扔到配置文件路径中，做持久化策略，只需要修改两处位置，其他的可以使用默认配置。

**RDB配置：**
通过`vim`编辑配置文件，输入`/SNAP`回车，找到快照配置。

```bash
# You can set these explicitly by uncommenting the three following lines.
# 表示在一个小时内至少有1个key发生变化
# save 3600 1
# save 300 100
# save 60 10000
save 60 1
```

同时需要配置持久化文件的保存路径，在SNAPSHOTING的最后一行。

```bash
# The working directory.
#
# The DB will be written inside this directory, with the filename specified
# above using the 'dbfilename' configuration directive.
#
# The Append Only File will also be created inside this directory.
#
# Note that you must specify a directory here, not a file name.
dir /data
```

> **注：这里修改为/data，是因为redis容器中存放快照文件的路径就是/data，方便后面做映射，当然也可以修改为其他的可用路径。**

**AOF配置：**
使用`/appendonly`，回车找到AOF配置的位置，打开AOF策略。

```bash
# AOF and RDB persistence can be enabled at the same time without problems.
# If the AOF is enabled on startup Redis will load the AOF, that is the file
# with the better durability guarantees.
#
# Please check https://redis.io/topics/persistence for more information.

appendonly yes
```

> **注：如果需要远程访问redis，还需要将配置文件中绑定127.0.0.1注释掉**。
> 输入`/bind 127.0.0.1`，并注释

#### 配置映射并启动

```bash
docker run --name redis -p 6379:6379 -v /root/docker/redis/data:/data -v /root/docker/redis/conf/redis.conf:/etc/redis/redis.conf --privileged=true -d redis redis-server /etc/redis/redis.conf --requirepass "gitmodelAbcde" --loglevel warning
```

- **--name:**指定容器的名字
- **-p:**做端口映射，6379:6379
- **-d:**表示在后台启动容器
- **--requirepass:redis**访问密码
- **--privileged=true:**表示给容器root权限，否则无法使用appenonly。
- **-v:**docker的Volume指令
- **redis-server:**表示使用右侧路径中的配置文件启动。

容器启动时会将宿主机路径`/root/docker/redis/conf/redis.conf`同步到容器的`/etc/redis/redis.conf/`中，然后容器通过这个etc路径下的配置文件进行启动。同时容器`/data`目录也会在redis运行过程中，将持久化文件同步到宿主机中去。

## Tomcat

### 创建Tomcat镜像

```bash
docker pull tomcat
```

### 查看Tomcat镜像

```bash
docker images
```

### 创建Tomcat容器

```bash
docker run -d -p 8080:8080 --name tomcat CONTANINER
```

### 上传项目

通过FTP上传项目到服务器

### 进入容器目录

```bash
docker exec -it CONTANINER bash
```

- **-i :**即使没有附加也保持STDIN 打开
- **-t :**分配一个伪终端

### 应用数据卷

```bash
docker run -d -p 8080:8080 -v /home/tomcat_webapps:/usr/local/tomcat/webapps/ --name tomcat CONTANINER
```

## Dockerfile

```dockerfile
FROM tomcat
ADD demo.war /usr/local/tomcat/webapps
```

```bash
docker build -t tomcat:1.0 -f /path/to/a/Dockerfile
```

- **-t NAME :**定义镜像的tag
- **-f PATH :**Dockerfile的路径

## Docker-Compose管理MySQL和Tomcat容器

```yml
version:'3.8'
services:
	mysql:														# 服务的名称
		restart: always											# 代表只要Docker启动，那么这个容器就跟着一起启动
		image: mysql											# 指定镜像路径
		container_name: mysql									# 指定容器名称
		ports:
			- 3306:3306											# 指定端口号的映射
		environment:
			MYSQL_ROOT_PASSWORD: 123456							# 指定MySQL的ROOT用户登陆密码
			TZ: Asia/Shanghai									# 指定时区
		volumes:
			- /home/docker_mysql:/var/lib/mysql					# 映射数据卷
	tomcat:
		restart: always
		image: tomcat
		conatiner_name: tomcat
		ports:
			- 8080:8080
		environment:
			TZ: Asia/Shanghai
		volumes:
			- /home/tomcat_webapps:/usr/local/tomcat/webapps
      - /home/tomcat_logs:/usr/local/tomcat/logs
```

### 使用docker-compose命令管理容器 

1. 基于docker-compose.yml启动管理的容器

在测试目录中，执行以下命令来启动应用程序：

```
docker-compose up
```

如果你想在后台执行该服务可以加上 **-d** 参数：

```
docker-compose up -d
```

2. 关闭并删除容器


```
docker-compose down
```

3. 开启关闭重启已经存在的由docker-compose维护的容器

```
docker-compose start|stop|restart
```

4. 查看由docker-compose管理的容器

```
docker-compose ps
```

5. 查看日志

```
docker-compose logs -f
```

## docker-compose结合dockerfile使用

```yml
version: '3.1'
services:
	web-demo:
		restart: always
		build:										# 构建自定义镜像
			context: ../							# 指定Dockerfile文件所在路径
			dockerfile: Dockerfile					# 指定Dockerfile文件名称
		image: demo:1.0
		container_name: demo
		ports:
			8081:8080
		environment:
			TZ: Asia/Shanghai
```

```dockerfile
FROM tomcat
ADD demo.war /usr/local/tomcat/webapps
```

> https://hub.daocloud.io/