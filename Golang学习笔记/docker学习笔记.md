# Docker概述

**虚拟化技术**

**发展史**

虚拟化技术出现之前，需要做的工作

1. 买硬件服务器
2. 在服务器上配置操作系统
3. 在操作系统上部署应用环境
4. 部署应用并运行
5. 如果要迁移应用，需要重复以上4步

是什么

> 把一套计算机硬件资源（cpu、内存、网络、存储等）虚拟化出多套硬件资源环境，可以把一台服务器当多台服务器使用

**分类**

- 硬件级虚拟化

  > 核心技术是Hypervisor，是运行在基础物理服务器硬件之上的软件层，可以虚拟化硬件资源，再在此之上安装os
  >
  > 虚拟化出的虚拟机是分隔独立的，可以当作多套硬件资源使用
  >
  > 比如VMWare、VirtualBox等虚拟机

  ![image-20210423115436360](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210423115436360.png)

- 操作系统级虚拟化

  > 即容器技术

**优缺点**

- 优点：一台服务器能虚拟出多台服务器，计算机资源得以充分利用
- 缺点
  - 每创建一个虚拟机，都会创建一个操作系统，太耗费资源
  - 不同虚拟机的环境不一样，部署时会有兼容性问题

**容器技术**

**发展史**

> 操作系统级的虚拟化技术（在同一套服务器的物理资源上，用容器把进程隔离开）
>
> 模拟的是运行在一个os上的不同进程，并将其封闭在一个密闭的容器里

**Docker**

**历史**

- 2013年问世，继续LXC技术（Linux Container），是一种内核虚拟化技术（轻量级虚拟化，隔离进程和资源，和宿主机使用同一个内核）
- docerk起初是做PaaS（platform as a service 平台及服务）的业务（云平台），后来开源，把PaaS业务出售，收购的公司后来倒闭

**介绍**

- `github.com/docker/docker-ce`
- 开发者把应用及依赖打包到一个可移植的容器中，打包好的容器可以在任何流行的linux服务器上运行，不存在开发环境和运维环境不一致的情况
- **Docker就是对软件及其依赖环境的标准化打包，应用之间互相隔离，共享一个os kernel，可以运行在很多主流os上；是内核级的虚拟化**
- **Docker本身不是容器，而是容器管理的引擎**

**容器和虚拟机的区别**

> **docker是容器间共享一个操作系统，启动快、占磁盘空间少、性能接近原生、单机支持量大、与宿主机共享os**
>
> 虚拟机是每套虚拟机都需要创建自己的操作系统，docker的有点，虚拟机都没有

# Docker环境搭建

![image-20210423161121905](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210423161121905.png)

## 名词介绍

- 镜像（image）

  docker镜像好比一个模板，可以通过这个模板来创建容器（通过镜像可以创建多个容器，最终服务/项目就是在容器中运行）

  > Tomcat镜像\===> run \==\=> tomcat01容器（真正提供服务的）

- 容器（container）
  - docker利用容器技术，独立运行一个或一组应用，它是通过镜像来创建的
  - 有启动、停止、删除等基本命令（可以先把容器理解为一个简易的linux系统）

- 仓库（repository）

  仓库就是存放镜像的地方，分为公有仓和私有仓

  > Docker Hub，国内使用阿里云配置镜像加速

## 阿里云服务器

- 控制台\=\=\==>云服务器ECS====>实例
  - 更多\=\=\==>密码/密钥====>修改远程连接密码
  - 安全组\=\=\==>创建安全组====>添加3306/6379/22/80/443端口
  - 密钥对\=\=\==>创建密钥对\=\=\==>输入名称、选择自动创建\=\=\==>下载`.pem`密钥文件====>绑定密钥对
  
- 重启服务器

- X shell新建会话，输入阿里云的公网Ip地址，连接后密钥文件选前面下载的`.pem`并输入修改后的密码即可

- 修改登录id（root@后面的一串）：`hostnamectl set-hostname xxx`，然后退出登录，重启服务器，重新连接，重新启动docker

- 基本命令

  - `uname -r`，查看内核版本，`3.10.0-1062.18.1.el7.x86_64`

  - `cat /etc/os-release`，查看系统版本

    ```shell
    NAME="CentOS Linux"
    VERSION="7 (Core)"
    ID="centos"
    ID_LIKE="rhel fedora"
    VERSION_ID="7"
    PRETTY_NAME="CentOS Linux 7 (Core)"
    ANSI_COLOR="0;31"
    CPE_NAME="cpe:/o:centos:centos:7"
    HOME_URL="https://www.centos.org/"
    BUG_REPORT_URL="https://bugs.centos.org/"
    
    CENTOS_MANTISBT_PROJECT="CentOS-7"
    CENTOS_MANTISBT_PROJECT_VERSION="7"
    REDHAT_SUPPORT_PRODUCT="centos"
    REDHAT_SUPPORT_PRODUCT_VERSION="7"
    ```

- 配置镜像加速器

  ```shell
  sudo mkdir -p /etc/docker
  sudo tee /etc/docker/daemon.json <<-'EOF'
  {
    "registry-mirrors": ["https://gtkcr9ph.mirror.aliyuncs.com"]
  }
  EOF
  sudo systemctl daemon-reload
  sudo systemctl restart docker
  ```

## 安装Docker

- 帮助文档`https://docs.docker.com/`

- `/var/lib/docker`docker的默认工作路径

```shell
# 1.卸载旧版本
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
# 2.需要安装的包
yum install -y yum-utils
# 3.设置镜像的仓库
yum-config-manager \
    --add-repo \
    https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo # 推荐使用阿里云
# 3.1更新yum软件包索引
yum makecache fast
# 4.安装 docker-ce：社区版（ee是企业版）
yum install docker-ce docker-ce-cli containerd.io
# 5.启动docker
systemctl start docker
# 6.查看docker版本
docker version
# 7.hello-world
docker run hello-world
```

![image-20210425134834378](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210425134834378.png)

```shell
# 8.查看下载的镜像
docker images
# 0. 卸载docker
yum remove docker-ce docker-ce-cli containerd.io # 卸载软件
rm -rf /var/lib/docker # 删除资源
```

# Docker命令

## 底层原理

- docker是一个client-server结构的系统，docker的守护进程运行在主机上（即服务在后台运行，主机是宿主机）

- 通过socket从客户端访问

- docker-server接收到docker-client的命令，就会执行

![image-20210425135559996](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210425135559996.png)

>  docker为什么比VM快？
>
> 新建一个容器，docker不需要像虚拟机一样重载操作系统内核。虚拟机要加载Guest OS，而docker是利用宿主机的os

- docker有着比虚拟机更少的抽象层
- docker利用的是宿主机的内核

## 常用命令

- 帮助命令
  - `docker version`：显示docker的版本信息
  - `docker info`：显示docker的系统信息，包括镜像、容器的数量、插件等
  - `docker events`
  - `docker 某命令 --help` ：万能命令
  - `https://docs.docker.com/reference/commandline`

## 镜像命令

- `docker images`：查看所有本机的镜像，后面可以跟如下可选项
  - `-a`：列出所有镜像
  - `-f`：过滤
  - `-q`：只显示镜像的ID（常用），和a联合起来用，`-aq`
- `docker search mysql`：搜索镜像，后面可跟如下可选项
  - `--filter=stars=3000`：stars数量在3000以上的
  - `https://hub.docker/com/_/mysql`：对应的docker hub上的地址
- `docker pull mysql`：下载镜像，后面可跟`:tag`（按指定tag下载，如果不写，默认是latest）
  - 分层下载，docker image的核心，联合文件系统（软件的不同版本共享一部分重复的文件）
  - `Digest`是签名防伪标志
  - 等价于`docker pull mysql docker.io/library/mysql:latest`
- `docker rmi -f xxx aaa bbb`：删除`image id`为xxx、aaa、bbb的镜像的所有内容
  - 批量删除`docker rmi -f $(docker images -aq)`：先通过括号中的命令查出所有镜像的`image id`，再替换`$`的占位

## 容器命令

> 有了镜像才能创建容器，容器id永远放在最后

- `docker run [可选参数] image名字`：新建容器并启动，可选参数说明

  - `--name xxx`，容器名字，用来区分容器

  - `-d`，后台方式运行

  - `-it`，使用交互方式运行，进入容器查看内容（`-i`、`-t`联合起来使用）

  - `-p`，指定容器端口

    - `-p ip:宿主机端口:容器端口`
    - `-p 宿主机端口:容器端口`（常用）
    - `-p 容器端口`
    - `容器端口`

  - `-P`，随机指定端口

  - `-e`后跟配置参数：登录密码、端口、es的运行最大内存等配置

  - `-rm`，退出容器后自动删除，用完即删（不推荐使用）

    ```shell
    docker run -it centos /bin/bash # 通过bash进行交互运行
    ```

- `docker ps`：列出当前正在运行的容器，后接

  - `-a`：列出当前+历史运行的容器
  - `-n=5`：显示最近创建的5个容器
  - `-q`：只显示容器的编号

- `exit`：从容器退出到主机并停止容器

  - `ctrl+p+q`：只退出不停止

- `docker rm xxx`：删除容器id为xxx的容器

  - 不能删除正在运行的容器，如果要强删，用`-f`
  - `docker rm -f $(docker ps -aq)`：删除所有容器，也可用（`docker ps -a -q|xargs docker rm`）

- 启停容器，xxx为容器id，可以是历史id

  - `docker start xxx`：启动容器
  - `docker restart xxx`：重启容器
  - `docker stop xxx`：停止当前运行的容器
  - `docker kill xxx`：强制停止当前运行的容器
  - `docker pause/unpause`：暂停/恢复容器

## 常用其他命令

- `docker run -d centos`：后台启动容器

  > 容器要在后台运行，就必须要有一个前台进程，如果没有的话docker就会自动停止容器

- `docker logs xxx`：查看日志，后跟`-tf --tail n`来显示具体的信息

  - `n`是显示的行数，不带`--tail`就显示所有行
  - `-t`是显示时间戳

- `docker top xxx`：查看容器中的进程信息

- `docker inspect xxx` ：查看容器的元数据

- 进入当前正在运行的容器:

  - `docker exec -it xxx /bin/bash`：进入容器后开启一个新的终端，可以在里面操作（常用），且用此方法进入容器后再`exit`，不会停止容器
  - `docker attach xxx`：进入容器正在执行的终端，不会启动新的进程，`exit`后会停止容器

- `docker cp xxx:容器内路径 主机的目的路径`：拷贝容器内的文件到主机上，也可以颠过来

- `docker stats xxx`：查看xxx容器的占用资源情况

- `docker volume`：查看卷

  - `create`创建一个卷
  - `inspect xxx`查看卷的详细信息
  - `ls`列出所有卷
  - `rm`移除一个或多个卷
  - `prune`移除所有在使用的卷

- `docker history 镜像id`：查看镜像的构建过程

- 保存导入

  - 把容器导出为tar：`docker export container id/name  > latest.tar`
  - 把镜像保存为tar：` docker save nginx > nginx.tar`
  - import和load则为导入和加载

## 小结

![image-20210426164918133](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210426164918133.png)

# Docker镜像

## 概述

> 镜像是一种轻量级、可执行的独立软件包，用来打包软件运行环境和基于运行环境开发的软件，包含运行某个软件所需的所有内容，包括代码、运行时、库、环境变量和配置文件

![image-20210427145001669](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210427145001669.png)

- 所有的应用，直接打包成docker镜像，就可以直接跑起来
- 获取镜像方法
  - 从远程仓库下载
  - 从别处拷贝
  - 自己制作`DockerFile`

## 加载原理

### UFS(联合文件系统)

![image-20210427093301969](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210427093301969.png)

### docker镜像加载原理

![image-20210427093452844](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210427093452844.png)

### 分层理解

![image-20210427094658700](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210427094658700.png)

![image-20210427094751568](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210427094751568.png)

## 特点

Docker镜像层都是只读的，当容器启动时（docker run xxx），一个新的可写层被加载到镜像的顶部！

这一层就是容器层，容器之下的都叫镜像层

![image-20210427095624983](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210427095624983.png)

## commit容器

```shell
docker commit -m="提交的描述信息" -a"作者" 容器id 目标镜像名:TAG（TAG可选）
```

# 容器数据卷

> 容器的持久化和同步操作：将容器中产生的数据同步到本地（即目录的挂载，将容器内的目录挂载到linux上）
>
> 容器间也是可以数据共享的
>
> - docker run --name aaa **--volumes-from** xxx imagename
>
> - xxx和aaa称为数据卷容器，把父容器（xxx）删了，子容器（aaa）的数据不会消失
> - 可以用来容器之间配置信息的传递，数据卷容器的生命周期持续到没有容器使用为止

## 使用

- `docker run -v 主机目录:容器目录 xxx`（-v是volume）（是双向同步的）
- 所有的卷，不指定本机目录时，都存放在`/var/lib/docker/volumes/卷名/_data/`里面
- 增删改文件，会同步到挂载的目录中，删除容器不会影响本机已经同步的数据
- 只要容器存在，即使停止了，也能同步数据
- 在最后跟上`:ro`或`:rw`可以设置容器内目录的权限
  - `:ro`是只读，只能从宿主机的目录来改数据，容器内的目录改不了
  - `:rw`是读写

## 挂载方式1

### 指定目录挂载

- 挂载的时候指定了宿主机的目录

  `docker run -v 主机目录:容器目录 xxx`：容器目录带`/`

### 具名挂载（指定卷名）

- 挂载的时候指定了卷名，不指定本机的目录：卷名不带`/`

  `docker run -v 卷名:容器目录 xxx`

### 匿名挂载

- 挂载的时候不指定本机的目录，也不指定卷名，只写了容器内的目录（最终的卷名是一堆哈希值）
  - `docker run -v 主机目录:容器目录 xxx`
  - `docker run -v 容器目录 xxx`

## 挂载方式2（常用）

- 通过dockerfile的命令来挂载

  ```shell
  FROM centos
  # 在容器内生成了两个目录，是和宿主机有挂载关系的目录
  # 且是匿名挂载，宿主机的目录是在/var/lib/docker/volumes/哈希值/_data下
  VOLUME ["volume1","volume2"]
  CMD echo -----end------
  ```

# DockerFile

## 概述

- dockerfile是用来构建docker镜像的构建文件，是一个命令脚本！
- 镜像是一层一层的，所以脚本也是一条一条命令（命令都是大写的），每条命令是一层

## 使用

> 编写好DockerFile后，通过它构建生成镜像，发布后run这个镜像得到容器，容器提供最终的服务

### 常用命令

- `FROM`：基础镜像，如centos、ubuntu等

- `MAINTAINER`：镜像是谁写的，留`name<email>`

- `RUN`：镜像构建的时候需要运行的命令，一般用来安装其他软件，可以用`&&`连接起来，这样安装时就只有一层镜像

- `ADD`：需要的其他环境，一般是压缩包，构建时会自动解压（可添加解压后的目录）

- `WORKDIR`：镜像的工作目录

- `VOLUME`：挂载的目录（容器内的目录）

- `EXPOSE`：指定暴露的端口

- `CMD`：指定这个容器启动时要运行的命令（即跟在`docker run`后面的东西）

  - 命令可以用`[]`括起来：`CMD ["ls","-a"]`（推荐）
  - 也可以直接跟命令名字：`CMD ls`
  - 只有最后一个CMD会生效，用CMD打印的东西不受影响
  - 手动`docker run`时，后面跟的命令，也会替换掉dockerfile中的CMD命令

  ```shell
  CMD shell命令
  # 推荐如下格式
  CMD ["可执行文件或命令","param1","param2",...] 
  CMD ["param1","<param2",...]  # 该写法是为 ENTRYPOINT 指令指定的程序提供默认参数
  ```

- `ENTRYPOINT`：指定这个容器启动时要运行的命令，可以**追加命令**

  ```shell
  ENTRYPOINT ["可执行文件或命令","param1","param2",...]
  ```

- `ONBUILD`：触发指令，当构建一个继承此镜像的镜像时，就会触发这个命令（延迟构建）

- `COPY`：类似ADD，将文件拷贝到镜像中

- `ENV`：构建的时候设置环境变量，可以是`key value`或`key=value`的格式

### 构建

`docker build -f dockerfile文件路径 -t 目标镜像名:版本号 .`

- 使用dockerfile来构建目标镜像，构建的结果是在当前目录下，版本号是可选的
- dockerfile的国际通用命名`Dockerfile`，按这个命名的话就不需要`-f`参数了

```dockerfile
# 分阶段构建，不需要依赖资源等等。先把构建过程作为builder
FORM golang:alpine AS builder

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    
# 移动工作目录到/build
WORKDIR /build

# 把dockerfile所在目录的相关内容拷贝到镜像的当前目录（当前目录在上一步已移动到/build)
COPY go.mod .
COPY go.sum .
RUN go mod download

# docker run时跟的命令，指定可执行文件名为web_app，可执行文件在当前目录下
RUN go build -o bluebell_app .

#创建一个小镜像，debian是为了执行wait-for.sh脚本
FROM debian:stretch-slim

#拷贝静态资源，第四个是docker-compose.yaml中要执行的脚本
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf
COPY ./wait-for.sh /

#拷贝builder中编译好的可执行文件到镜像的根目录下
COPY --from=builder /build/bluebell_app /

#暴露端口
EXPOSE 8888

#执行wait-for.sh所需要的环境和配置
RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
		chmod 755 wait-for.sh
		
#ENTRYPOINT ["bluebell_app"] 不需要了，后续使用docker-compose up运行
```

### 发布

1. `docker login -u username -p password`
2. `docker tag 镜像ID tag`
3. `docker push 镜像名:tag`（不手动加tag的话，tag就是`latest`）

> 发布到阿里云
>
> 1. 找到“容器镜像服务”
> 2. 创建命名空间
> 3. 创建镜像仓库，类型选私有，下一步，选择“本地仓库”
> 4. 推送到阿里云的镜像仓库，[方法](https://cr.console.aliyun.com/repository/cn-shenzhen/zhou2ding/docker-study/details)

# Docker流程总结

![image-20210427170953776](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210427170953776.png)

# Docker网络

## Docker0

### 概述

- `ip addr`

![image-20210429170025017](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210429170025017.png)

- `docker run`时追加`ip addr`命令，得到的就是本机地址和docker内部地址

- `docker network inspect 容器id`可以查看容器网络的详细信息

- 每启动一个docker容器，docker就会给容器分配一个ip，只要安装了docker，就会有一个网卡docker0（桥接模式，使用veth-pair技术）

### veth-pair

  ==直接查看ip==：linux上的Ip对

![image-20210429172945934](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210429172945934.png)

==查看容器的ip==：容器内部的ip对，和linux上的ip对是一致的

![image-20210429173027541](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210429173027541.png)

> veth-pair技术：一对的虚拟设备接口，都是成对出现的，一段连着协议，一段彼此相连
>
> veth-pair来充当桥梁，容器之间可以互相ping同，linux服务器可以ping同它创建的容器

![image-20210429174300557](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210429174300557.png)

## 容器互联

- run docker的时候添加`--link`连通另一个容器，就可以通过容器名字来ping通，A link B的话，A能ping通B，但是反向不行

- 真实开发中已经不建议使用`--link`了，一般自定义网络

  > 原理：`/etc/host中配置了连接容器的IP地址


==自定义网络==（网卡）

  ![image-20210429203047263](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210429203047263.png)

- 网络模式
  - bridge：桥接，在docker上搭桥（多个容器把docker0当做桥）（自定义网络用此模式）
  - none：不配置网络
  - host：和宿主机共享网络
  - container：容器内网络连通（局限性大，用得少）
- 前置知识点
  - `docker run`的时候默认添加了`--net bridge`命令，即默认连通了`docker0`（bridge就是docker0的NAME）
  - `docker0`特点：默认，不能通过名称或id访问，需要`--link`
- 自定义网络
  - `docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet`
    - `--driver bridge`可省略，默认就是bridge
    - `--subnet`是子网掩码，可以给容器分配的ip地址从`192.168.0.2~192.168.255.255`（docker run --net的时候分配）
    - `--gateway`是网关，`mynet`是网络名字
    - 自定义网络的名字要避开`bridge`
  - 自定义网络完成后，可以在`docker run`的时候手动`--net mynet`，`mynet`就会根据子网掩码分配`ip`
  - 启动两个容器都通过`mynet`启动的话，两个容器之间就能通过容器名字或容器id来连通了

## 连接容器到另外的自定义网络

> 容器A通过自定义网络mynet1启动，容器B通过自定义网络mynet2启动，容器A和B是无法连通的

- `docker network mynet1 containerB`
  - 连通之后，直接把`containerB`放到`mynet1`下，`docker network inspect mynet1`可以看到``containerB`的信息

# Docker Compose

除了像上面一样使用`--link`的方式来关联两个容器之外，我们还可以使用`Docker Compose`来定义和运行多个容器。

`Compose`是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，你可以使用 YML 文件来配置应用程序需要的所有服务。然后，使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。

使用Compose基本上是一个三步过程：

1. 使用`Dockerfile`定义你的应用环境以便可以在任何地方复制。
2. 定义组成应用程序的服务，`docker-compose.yml` 以便它们可以在隔离的环境中一起运行。
3. 执行 `docker-compose up`命令来启动并运行整个应用程序。

我们的项目需要两个容器分别运行`mysql`和`bubble_app`，我们编写的`docker-compose.yml`文件内容如下：

```yaml
# yaml 配置
version: "3.7"
services:
  mysql8019:
    image: "mysql:8.0.19"
  #端口，本机端口是33061，容器端口是3306
    ports:
      - "33061:3306"
  #使用默认密码，使用init.sql进行数据库的初始化
    command: "--default-authentication-plugin=mysql_native_password --init-file /data/application/init.sql"
    environment:
      MYSQL_ROOT_PASSWORD: "564710"
      MYSQL_DATABASE: "bluebell"
      MYSQL_PASSWORD: "564710"
  #本地的init.sql挂载到容器中
    volumes:
      - ./init.sql:/data/application/init.sql
      
  redis507:
  	image:"redis:5.0.7"
  	ports:
  		- "26379:6379"
  bluebell_app:
  #使用本地的dockerfile来构建
    build: .
  #执行脚本：等待mysql和redis都启动后，再启动bluebell
    command: sh -c "./wait-for.sh mysql8019:3306 redis507:6379 -- ./bluebell_app ./conf/config.yaml"
    depends_on:
      - mysql8019
      - redis507
    ports:
      - "8888:8888"
```

这个 Compose 文件定义了两个服务：`bubble_app` 和 `mysql8019`。其中：

##### bubble_app

使用当前目录下的`Dockerfile`文件构建镜像，并通过`depends_on`指定依赖`mysql8019`服务，声明服务端口8888并绑定对外8888端口。

##### mysql8019

mysql8019 服务使用 Docker Hub 的公共 mysql:8.0.19 镜像，内部端口3306，外部端口33061。

这里需要注意一个问题就是，我们的`bubble_app`容器需要等待`mysql8019`容器正常启动之后再尝试启动，因为我们的web程序在启动的时候会初始化MySQL连接。这里共有两个地方要更改，第一个就是我们`Dockerfile`中要把最后一句注释掉：

```dockerfile
# Dockerfile
...
# 需要运行的命令（注释掉这一句，因为需要等MySQL启动之后再启动我们的Web程序）
# ENTRYPOINT ["/bubble", "conf/config.ini"]
```

第二个地方是在`bubble_app`下面添加如下命令，使用提前编写的`wait-for.sh`脚本检测`mysql8019:3306`正常后再执行后续启动Web应用程序的命令：

```bash
command: sh -c "./wait-for.sh mysql8019:3306 -- ./bubble ./conf/config.ini"
```

当然，因为我们现在要在`bubble_app`镜像中执行sh命令，所以不能在使用`scratch`镜像构建了，这里改为使用`debian:stretch-slim`，同时还要安装`wait-for.sh`脚本用到的`netcat`，最后不要忘了把`wait-for.sh`脚本文件COPY到最终的镜像中，并赋予可执行权限哦。更新后的`Dockerfile`内容如下：

```dockerfile
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o bubble .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

COPY ./wait-for.sh /
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf


# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bubble /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 需要运行的命令
# ENTRYPOINT ["/bubble", "conf/config.ini"]
```

所有的条件都准备就绪后，就可以执行下面的命令跑起来了：

```bash
docker-compose up
```

# Docker Swarm

# CI\CD Jenkins