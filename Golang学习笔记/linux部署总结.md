# Linux部署总结

## Linux系统分类

一般来说linux系统基本上分两大类：

- RedHat系列：Redhat、Centos、Fedora等
- Debian系列：Debian、Ubuntu等

### RedHat 系列

- 常见的安装包格式：rpm包，安装rpm包的命令是“rpm -参数”
- 包管理工具：yum
- 支持tar包

### Debian系列

- 常见的安装包格式：deb包，安装deb包的命令是“dpkg -参数”
- 包管理工具：apt-get
- 支持tar包

## Linux安装包

### rpm包

rpm是redhat公司的一种软件包管理机制，直接通过rpm命令进行安装删除等操作，最大的优点是自己内部自动处理了各种软件包可能的依赖关系。

```shell
rpm -ivh *.rpm			#安装
rpm -e packge.rpm		#卸载
rpm -Uvh packge.rpm		#升级
rpm -q [Query options]	#查询
```

### dgb包

dpkg是为Debian专门开发的套件管理系统，方便软件的安装、更新及移除。

```shell
dpkg -i packge.deb		#安装
dpkg -r packge.deb		#卸载
dpkg -i -R packges		#递归安装
```

## Linux包管理工具

### yum

> Yellow dog Updater, Modified的缩写

yum是RedHat系列的高级软件包管理工具

- 主要功能是更方便的添加/删除/更新RPM包。
- 它能自动解决包的依赖性问题。
- 它能便于管理大量系统的更新问题。
- yum源的存放路径：`/etc/yum.repos.d/`。

yum的特点

- 可以同时配置多个资源库(Repository)
- 简洁的配置文件(/etc/yum.conf)
- 自动解决增加或删除rpm包时遇到的倚赖性问题
- 保持与RPM数据库的一致性

yum可以用于运作rpm包，例如在CentOS/RedHat系统上对某个软件的管理：

```bash
yum install <package_name>	#安装
yum remove <package_name>	#卸载
yum update <package_name>	#更新
```

### apt-get

apt-get是Debian系列的高级软件包管理工具

- 配置文件/etc/apt/sources.list

apt-get可以用于运作deb包，例如在Ubuntu上对某个软件的管理：

```shell
apt-get install <package_name>	#安装
apt-get remove <package_name>	#卸载
apt-get update <package_name>	#更新
```

## CentOS的离线安装

> 软件如果部署在无法连接公网的服务器上，则无法使用yum或apt-get等方式联网下载，需要使用离线安装的方式

### 制作本地yum源

准备比较全的源仓库，如[阿里源](https://developer.aliyun.com/mirror/)或[中科大源](https://mirrors.ustc.edu.cn/help/)，以下已阿里源为例，制作本地EPEL yum源

> EPEL 的全称叫 Extra Packages for Enterprise Linux。EPEL 是由 Fedora 社区打造，为 RHEL 及衍生发行版如 CentOS、Scientific Linux 等提供高质量软件包的项目。装上了 EPEL 之后，就相当于添加了一个第三方源。且CentOS 源包含的大多数的库都是比较旧的，很多流行的库也不存在。而EPEL 在其基础上不仅全，而且还够新。

- 替换本地repo

  ```shell
  curl -o /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo	#替换CentOS-Base.repo
  curl -o /etc/yum.repos.d/epel-7.repo https://mirrors.aliyun.com/repo/epel-7.repo 		#替换epel-7.repo
  yum clean all	#清除缓存
  yum makecache	# 重建yum源缓存
  createrepo -v /data/soft/epel
  mkdir -p /mnt/cdrom			#创建目录
  mount /dev/sr0 /mnt			#挂载镜像
  cd /mnt/cdrom/Packages/		#Packages里面包含CentOS7所有的安装包，安装软件可以从这个目录中获取
  ```

- 修改本地yum源

  ```shell
  mkdir -p /data/soft/epel    /data/soft/centos7	#创建制作yum源的目录
  yum install yum-utils createrepo -y	#安装yum仓库相关软件
  reposync -r epel -p /data/soft/epel	#同步yum源
  cd /etc/yum.repos.d/		#切换到yum.repos.d目录
  rename .repo .repo.bak *	#备份所有配置文件
  vim CentOS-Local.repo		#新增本地yum源
  #在本地yum源中新增以下内容（去掉注释）
  [base]
  name=CentOS-Local   		#yum名称
  baseurl=file:///mnt/cdrom   #协议类型和软件包目录地址
  gpgcheck=0   				#要不要检查
  enabled=1 					#要不要使用
  ```

- 清空缓存，安装所需软件

  ```shell
  yum clean all
  yum -y install
  ```

### 离线安装神器

### 常用软件的rpm包

