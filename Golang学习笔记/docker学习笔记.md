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

  > Tomcat镜像===> run ===> tomcat01容器（真正提供服务的）

- 容器（container）
  - docker利用容器技术，独立运行一个或一组应用，它是通过镜像来创建的
  - 有启动、停止、删除等基本命令（可以先把容器理解为一个简易的linux系统）

- 仓库（repository）

  仓库就是存放镜像的地方，分为公有仓和私有仓

  > Docker Hub，国内使用阿里云配置镜像加速

## 阿里云服务器

- 控制台====>云服务器ECS====>实例
  - 更多====>密码/密钥====>修改远程连接密码
  - 安全组====>创建安全组====>添加3306/6379/22/80/443端口
  - 密钥对====>创建密钥对====>输入名称、选择自动创建====>下载`.pem`密钥文件====>绑定密钥对
  
- 重启服务器

- X shell新建会话，输入阿里云的公网Ip地址，连接后密钥文件选前面下载的`.pem`并输入修改后的密码即可

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

  - `-e`后跟配置参数：

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

- `docker cp xxx:容器内路径 主机的目的路径`：拷贝容器内的文件到主机上

- `docker stats xxx`：查看xxx容器的占用资源情况

## docker可视化（不推荐用）

- `portainer`提供一个后台面板供我们操作

- ```shell
  docker run -d -p 8088:9000 \
  --restart=always -v /var/run/docker.sock:/var/run/docker.sock --privileged=true portainer/portainer
  ```

- 然后访问`https://linux的ip:8088`

## 小结

![image-20210426164918133](D:\资料\Go\src\studygo\Golang学习笔记\docker学习笔记.assets\image-20210426164918133.png)

# Docker镜像

## 概述



## 加载原理

# 容器数据卷

# DockerFile

# Docker网络原理

# IDEA整合Docker

# Docker Compose

# Docker Swarm

# CI\CD Jenkins