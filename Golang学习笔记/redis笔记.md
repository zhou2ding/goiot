# NoSQL

## 概述

> 单机MySQL\===>Memcached（缓存）+读写分离+垂直拆分=\==>分库分表+水平拆分+MySQL集群
>
> 为什么用NoSQL？数据爆发式增长，位置信息、热榜、图片等等，web2.0的诞生，超大规模的高并发，传统的关系型数据库无法对付

![image-20210430111738570](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210430111738570.png)

- 概念

  - NoSQL=Not Only SQL，泛指非关系型数据库（关系型数据库是表、行、列）
  - NoSQL数据的存储不需要固定的格式，不需要多余的操作就能横向扩展

- 特点

  - 方便扩展，数据之间没有关系
  - 大数据量，高性能读写（Redis 1秒写8万次，读11万次）（缓存是一种细粒度的缓存）
  - 数据类型多样，不需要设计数据库

- RDBMS和NoSQL区别

  - 关系型数据库是结构化组织、结构化查询语言、数据和关系都存在单独的表中、严格的一致性、事务
  - NoSQL没有固定的查询语言，存储方式多（键值对存储、列存储、文档存储、图形数据库
  - NoSQL最终一致性（只要最终结果一致就行）
  - NoSQL的CAP、BASE理论
  - NoSQL有高性能、高可用、高可扩展性（机器不够时，随时水平拆分）

- 3V+3高（了解）

  海量、多样、实时；高并发、高可拓、高性能

## 数据模型

## 分类

- 键值对存储：KV键值对，Redis
- 列存储数据库：HBase，分布式文件系统
- 文档型数据库（bson格式）：MongoDB，ConthDB
- 图关系数据库：存储的是关系，如社交网络、广告推荐等，Neo4j、InfoGrid

![image-20210430160658134](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210430160658134.png)

## CQP

## BASE

# Redis介绍

> Remote Dictionary Server，远程服务字典，Key-Value存储
>
> 可以内存存储，也可持久化（rdb、aof），高速缓存，发布订阅系统，地图信息分析，计时器、计数器
>
> 特点：多样的数据类型、持久化、集群、事务
>
> redis的瓶颈不是CPU，而是内存和网络带宽，使用单线程（==redis将所有数据放在内存中，所以使用单线程效率是最高的，因为多线程会有CPU上下文切换比较耗时，对于内存系统来说没有上下文切换效率就是最高的==）

# 安装启动

- 官方推荐在Linux上使用，Windows版本只能在github上下载，很久不维护了
- 用`Xftp`把压缩包拷贝到服务器，`tar -xsvf`解压，`yum install gcc-c++`安装gcc环境，之后`make`两次，再`make install`
- 安装完成后`redis-cli`、`redis-server`、`redis-sentinel`等都在`/usr/local/bin`里面，把redis原始目录里的配置文件拷贝到此目录下，后续启动就用这个配置文件，`daemonize`要改为yes
- `redis-server rconfig/redis.conf`启动服务，`redis-cli -h localhost -p 6379`连接服务器，`-h`可省略
- 先`shutdown`进入`not connected`状态，再`exit`，就能结束`redis-server`和`redis-client`的进程

- `redis-benchmark`性能测试：`redis-benchmark -h localhost -p 6379 -c 100 -n 100000`，测试结果如下

  ![image-20210430172506575](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210430172506575.png)

| 序号 | 选项      | 描述                                       | 默认值    |
| :--- | :-------- | :----------------------------------------- | :-------- |
| 1    | **-h**    | 指定服务器主机名                           | 127.0.0.1 |
| 2    | **-p**    | 指定服务器端口                             | 6379      |
| 3    | **-s**    | 指定服务器 socket                          |           |
| 4    | **-c**    | 指定并发连接数                             | 50        |
| 5    | **-n**    | 指定请求数                                 | 10000     |
| 6    | **-d**    | 以字节的形式指定 SET/GET 值的数据大小      | 2         |
| 7    | **-k**    | 1=keep alive 0=reconnect                   | 1         |
| 8    | **-r**    | SET/GET/INCR 使用随机 key, SADD 使用随机值 |           |
| 9    | **-P**    | 通过管道传输 <numreq> 请求                 | 1         |
| 10   | **-q**    | 强制退出 redis。仅显示 query/sec 值        |           |
| 11   | **--csv** | 以 CSV 格式输出                            |           |
| 12   | **-l**    | 生成循环，永久执行测试                     |           |
| 13   | **-t**    | 仅运行以逗号分隔的测试命令列表。           |           |
| 14   | **-I**    | Idle 模式。仅打开 N 个 idle 连接并等待。   |           |

# 基本命令

- 默认16个数据库，默认使用第0个，可以用`select n`切换，`DBSIZE`查看数据库大小

  ```bash
  ping #测试连接
  set name xxx #插入键值对
  get name #查看键对应的值
  keys * #查看所有的key
  flushdb / FLUSHALL #清空当前数据库 / 清空全部数据库
  ```

# 基本数据类型

- Redis-Key
- String
- List
- Set
- Hash
- Zset

# 特殊数据类型

# 配置

# 持久化

# 事务

# 订阅发布实现

# 主从复制

# 哨兵模式

# 缓存问题

# 缓存穿透及解决方案

# 缓存击穿及解决方案

# 缓存雪崩及解决方案