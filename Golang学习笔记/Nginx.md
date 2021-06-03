# 概述

> Nginx是一个高性能的web服务器和web反向代理服务器，也可以用来做邮件代理服务器，用C语言开发的
>
> 特点是占内存少，并发处理能力强，官方测试是5万并发请求，同类型的还有Apache、Lighttpd等

- 正向代理：代理客户端，用户访问代理，代理再去访问服务器，服务器把数据返回给代理，再返回给用户（VPN）。需要配置代理服务器的地址，服务端不知道实际发起请求的客户端。

- 反向代理：代理服务端，用户访问固定的域名，代理再把请求转发到相应的服务器，即使服务器动态扩容了，用户也感知不到。不用配置代理服务器地址，用户不知道实际提供服务的服务端

  ![image-20210504151053963](D:\资料\Go\src\studygo\Golang学习笔记\Nginx.assets\image-20210504151053963.png)

# 功能

- 静态网站部署

- 静态代理

- 虚拟主机

- 负载均衡：内存大的服务器，权重就大一些，有两种策略

  - 内置策略：轮询，加权轮询，Ip hash
  - 扩展策略：各种各样

- 动静分离

  ![image-20210429214757132](D:\资料\Go\src\studygo\Golang学习笔记\Nginx.assets\image-20210429214757132.png)

# 安装

0. `yum install gcc openssl openssl-devel pcre pcre-devel zlib zlib-devel -y`一次性安装环境

1. `tar -zxvf nginx.tar.gz`解压缩`nginx`的压缩包

2. 进入解压后的目录，`./configure --prefix=/www/server/nginx`配置`nginx`的安装路径，路径根据自己需要改

3. `make`编译，然后`make install`安装

4. 进入安装的目录，有四个文件夹

   - `conf/nginx.conf`，配置文件

   - `html/50x.html`，nginx访问错误时自动转发到这个页面；`html/index.html`欢迎首页面

   - `logs/`日志文件目录

   - `sbin目录下`启动

     - 直接启动`./nginx`
     - 配置文件启动`./nginx -c ../conf/nginx.conf`，启动前先用此命令跟`-t`检查下配置文件的语法错误

     >  启动后两个进程，master和worker，后面worker会有多个
     >
     > 主进程主要负责维护所有的工作进程，工作进程具体完成请求转发等工作

   - 关闭

     - `kill -QUIT pid`，杀掉`master id`即可，杀掉后`nginx`不能接受新的请求，但是不会立即关闭，而是等所有用户请求都响应成功后才关闭
     - `kill -TERM pid`，或者不加`-TERM`暴力快速关闭（常用）。`-9`是强制杀掉，但不会杀掉子进程
     - `./nginx -s reload`修改配置文件后需要此命令重启来生效

   - 查看版本

     - `nginx -v`显示版本
     - `nginx -V`显示版本、编译器版本和配置参数

# 使用

## 基本配置

- 通用配置
  - `user nobody`：nobody是运行nginx worker进程的用户，没有密码，用来在后台运行一些进程
  - `worker_processes`：工作进程的数量，默认1，通常cpu数量或其2倍（新版本默认值是auto了），面试时回答中间值，（4核CPU的话就是6个进程）
  - `error_log`：错误日志的路径和级别，共8个级别，默认`error`级别（即不写级别）
  - `pid`：存放进程id的文件

- `events`：配置工作模式和连接数

  - `worker_connections`，工作进程的连接上限，最大65535；nginx的连接上限是工作进程数量*此值

- `http`：配置http服务器，利用它的反向代理功能提供负载均衡支持

  - `include`：nginx可支持的多媒体类型，在`/conf/mime.types`里有具体的类型
  - `default_type`：默认文件类型（如果include里没有，则以此配置项的值的方式来解析，一般配成流类型`octet-stream`）
  - `log_format`：配置日志格式
  - `access_log`：配置access.log日志及存放路径，并使用上面定义的日志格式，和上面搭配使用
  - `sendfile`：开启高效文件传输模式
  - `tcp_nopush`：防止网络阻塞，上线后要开启
  - `keepalive_timeout`：长连接超时时间，单位是秒
  - `gzip`：开启gzip压缩输出，上线后要开启

- `server`：配置虚拟主机，一个http中有若干server，各server的端口和域名不能都相同

  - 端口号、服务名（域名）、字符集、access.log的路径（不管访问哪个server，都会输出http的此log，只有访问本server时才输出本server的此log）

  - `location`拦截请求，根据浏览器输入的地址返回配置的html文件

    ```bash
    location /ace		#浏览器请求的路径，/代表跟路径，这行意思是ip后面跟/ace的请求会被这个location拦截
    {
        root /opt/www/ace;	#磁盘目录的根路径，nginx拦截到请求后会去这个路径下面找东西，此路径必须有/ace
        index login.html;	#找到后需要打开的静态页面文件
    }
    
    upstream www.myweb.com {			#必须和proxy_pass后的域名一致
    	server 192.168.115.128：8081;
    	server 192.168.115.128：8082;	#按轮询的方式提供负载均衡，其他方式见下面
    }
    location /myweb
    {
        proxy_pass http://www.myweb.com	#固定写法，域名随便
    }
    ```

- 配置完成后在浏览器输入Ip地址+端口+/ace即可打开静态页面

## 负载均衡

- 配置

  ```bash
  #在浏览中访问192.168.115.128/myweb，每次刷新都会切换一个服务器
  upstream www.myweb.com {
  	server 192.168.115.128:8081;		#服务器的Tomcat地址
  	#backup作用：更新服务器代码先在备份上更新，然后关闭一批服务器，更新它们，然后关闭另一批同时启动更新好的，用户连接到backup和更新好的，待其他服务器也更新后，再启动其他服务器。即热部署
  	server 182.168.115.128:8082 backup;
  	server 182.168.115.128:8082 down;	#当前的服务器是down状态，不参与负载均衡，不处理用户请求
  }
  location /myweb {
  	proxy_pass http://www.myweb.com;
  }
  ```

- 策略

  - 轮询

    > 按照上面的配置两个服务器轮流访问，要保证多个机器的性能、处理事务要一致
    >
    > ![image-20210504213828308](D:\资料\Go\src\studygo\Golang学习笔记\Nginx.assets\image-20210504213828308.png)

  - 权重

    > 也是交替请求的
    >
    > ![image-20210504215034093](D:\资料\Go\src\studygo\Golang学习笔记\Nginx.assets\image-20210504215034093.png)

  - ip hash

    > 不存在session丢失问题，在nginx计算好ip的哈希值后，就只把请求转发到固定的服务器上
    >
    > （用户ip的hash、hash取模，以此为依据来分配服务器）
    >
    > 只对于用户ip不变的情况有用，用户换ip后可能会导致服务器过载
    >
    > upstream配置中增加ip_hash字段即可

  - 最少连接数

    > 谁的请求数量少，就把请求给谁
    >
    > ![image-20210504215128435](D:\资料\Go\src\studygo\Golang学习笔记\Nginx.assets\image-20210504215128435.png)

## 静态代理

- 方式一：以文件格式拦截

  ```bash
  #~表示正则开始，$表示正则结束
  #.表示任意非换行字符
  #*表示匹配一个或多个字符
  #\.表示转义的.
  #括号中是匹配的项
  location ~ .*\.(gif|jpg|jpeg|png|bmp|swf|css|html|js)$
  {
  	root /opt/static;
  }
  ```

- 方式二：匹配路径（相比文件格式的拦截，写的更少）

  ```bash
  #匹配一个路径，路径名包含了括号中的内容，就是静态资源
  location ~ .*/(js|css|img|images)
  {
  	root /opt/static;
  }
  ```

## 动静分离

> 静态资源由Nginx处理，动态资源由Tomcat处理
>
> 部署两个Nginx，一个负载均衡，一个静态代理（Tomcat有几个，静态代理就有几个）
>
> 负载均衡的配置中，也要加上对静态资源的负载均衡

## 虚拟主机

1. 配置Nginx

   ```bash
   #Nginx的端口一样，域名不一样
   upstream beijing.myweb.com {
   	server 192.168.115.128:8081;
   }
   server {
   	listen		80;
   	server_name beijing.myweb.com;
   	location / {
   		proxy_pass http://beijing.myweb.com;
   	}
   }
   
   upstream tianjin.myweb.com {
   	server 192.168.115.128:8082;
   }
   server {
   	listen		80;	#虽然也是80，只要server_name和端口不都一样就行
   	server_name tianjin.myweb.com;
   	location / {
   		proxy_pass http://tianjin.myweb.com;
   	}
   }
   ```
   
2. 修改`C:\Windows\System32\drivers\etc\hosts`，添加如下DNS映射

   ```bash
   192.168.115.128 beijing.myweb.com
   192.168.115.128 tianjin.myweb.com
   ```


## 完整配置

假设两个Tomcat，则启动三个Nginx，一个负载均衡，两个静态代理

```bash
#负载均衡的Nginx配置，upstream和location一一对应
upstream www.myweb.com {
	ip_hash;
	server 192.167.115.128:8081;	#Tomcat的端口
	server 192.167.115.128:8082;
}
upstream static.myweb.com {
	server 192.167.115.128:81;		#静态代理的Nginx的端口
	server 192.167.115.128:82;
}
server {
    listen		80;
    server_name	localhost;
    #整个项目的负载均衡
    location /myweb {
    	proxy_pass http://www.myweb.com;
    }
    #静态代理的负载均衡
    location ~ .*/(css|js|img|image|images) {
    	proxy_pass http://static.myweb.com
    }
}
```

```bash
#静态代理的Nginx的配置
server {
    listen		81;
    server_name	localhost;
    location ~ .*/(css|js|img|image|images) {
    	root /opt/static;
    }
}
```

