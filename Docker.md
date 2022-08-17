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
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pwd --name tomcat CONTANINER
```

- **-d :**代表后台运行容器
- **-p 宿主机端口:容器端口 :**将宿主机和容器端口进行映射
- **–-name 容器名称 :**指定容器名称

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
	mysql:																					# 服务的名称
		restart: always																# 代表只要Docker启动，那么这个容器就跟着一起启动
		image: mysql																	# 指定镜像路径
		container_name: mysql													# 指定容器名称
		ports:
			- 3306:3306																	# 指定端口号的映射
		environment:
			MYSQL_ROOT_PASSWORD: 123456									# 指定MySQL的ROOT用户登陆密码
			TZ: Asia/Shanghai														# 指定时区
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
		build:														# 构建自定义镜像
			context: ../										# 指定Dockerfile文件所在路径
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