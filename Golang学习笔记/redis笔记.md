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
> 可以用作==数据库==、==缓存==、==消息中间件==
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

- 官网查看命令[帮助文档](http://redis.cn/commands.html)

  ```bash
  ping				#测试连接
  EXPIRE <key> sec	#sec秒后key过期
  ttl <key>			#查看key的剩余过期时间，-1是不会过期，-2是已经过期，过期后就没了
  type <key>			#查看key的乐西
  keys *				#查看所有的key
  flushdb / FLUSHALL	#清空当前数据库 / 清空全部数据库
  move <key> 1		#把此key移动到数据库1
  EXISTS <key>		#查看这个key是否存在
  config get/set xxx	#查看/修改某配置项
  ```

# key的基本数据类型

- String

  > 可以是字符串，也可以是数字，场景：计数器、对象缓存存储等

  ```bash
  set <key> <val>						#插入键值对
  get <key>							#查看key对应的值
  del <key>							#删除key
  APPEND <key> <val>					#在key对应的字符串值后追加ccc，如果当前key不存在，就自动添加一个键值对
  STRLEN <key>						#key对应的字符串值的长度
  incr <key> / decr <key>				#key对应的值自增1/自减1
  INCRBY <key> n / DECRBY <key> n		#key对应的值增加n/减少n
  GETRANGE <key> start end			#查看值的下标start到end的内容（闭区间），end是-1的话就是到尾
  SETRANGE <key> idx <val>			#修改指定下标位置的字符，把下标为idx开始的等长字符串替换成val
  setex <key> sec <val>				#设置过期时间（set expire）
  setnx <key> <val>					#如果不存在则设置（set if not exists）,成功返回1；如果key存在则返回0
  mset <key1> <val1> <key2> <val2>	#批量设置键值对
  mget <key1> <key2> <key3>			#批量查看key对应的值
  msetnx <key1> <val1> <key4> <val4>	#如果不存在则批量设置，是原子性操作，要么一起成功，要么一起失败
  getset <key> <val>					#先get再set，若存在则获取原来的值并设新值；若不存在则返回nil并设置值
  # 设置对象
  # 方式1：<obj>:<id>作为key，后面{}括起来的json是val，后面再去解析这个json
  set <obj>:<id> {<key1>:<val>,<key2>:<val2>}				#如set usr:1 {name:zs,age:10}
  # 方式2：<obj>:<id>:<key1>，<obj>:<id>:<key2>分别为两个key
  mset <obj>:<id>:<key1> <val1> <obj>:<id>:<key2> <val2>	#如mset usr:1:name:zs usr1:1:age:10
  mget <obj>:<id>:<key1> <obj>:<id>:<key2>
  ```

- List

  > 所有的list命令都是L开头，LPUSH LPOP则是栈，LPUSH RPOP则是队列

  ```bash
  LPUSH <key1> <val1>	<key1> <val1>		#把val塞入key这个list，入栈
  RPUSH <key> <val>						#从另一个方向塞入list（从栈底塞进去）
  LPOP <key>								#弹出栈顶元素，可以跟数字，弹出几个
  RPOP <key>								#移除栈底元素，可以跟数字
  LRANGE <key> start end					#获取下标区间内的值，按出栈顺序获取
  LINDEX <key> idx						#查看idx下标对应的值，下标顺序是从栈顶到栈底
  LLEN <key>								#查看长度
  LREM <key> <count> <val>				#移除指定个数的某个值，是从栈顶往下移除的
  LTRIM <key> start end					#截取指定长度区间的一串值。如果是截取1~0，则清空list
  RPOPLPUSH <src> <dst>					#把src的栈底元素移动到dst中，dst不存在的话新建一个dst
  LSET <key> idx <val>					#把指定下标处的值替换为val，key或val不存在的话都会报错
  LINSERT <key> direction <val> <newval>	#把newval插入val的前面或后面，direction是before/after
  ```

- Set

  > 无序不重复集合， 添加已有元素时会添加失败

  ```bash
  sadd <key> <val>			#把val添加到set中
  SMEMBERS <key>				#查看此set的所有值
  SISMEMBER <key> <val>		#判断此set中是否有此值
  scard <key>					#获取set中元素的数量
  srem <key> <val1> <val2>	#移除set中的某值
  SRANDMEMBER <key> count		#随机抽出指定个数的元素，count省略的话就是1
  spop <key>					#随机移除set中的元素
  smove <key1> <key2> <val>	#把val从key1的set移动到key2的set中
  #差集 交集 并集
  sdiff <key1> <key2>
  sinter <key1> <key2>
  sunion <key1> <key2>
  ```

- Hash

  > 存的值是一个map集合，也有自己的key，和string一样，只不过元素是键值对。以下mkey为map中的key
  >
  > 相比string更适合对象的存储

  ```bash
  hset <key> <mkey> <mval>					#存值，redis 4.0.0之后支持批量存值
  hget <key> <meky>							#取值
  hmset <key1> <mkey1> <mval1> <mkey2> <val2>	#对应string的mset，和hset的批量存值效果一样
  hmget <key1> <mkey1> <mkey2>				#批量取值（只查看出来值）
  hgetall <key>								#查看所有map键值对
  hdel <key> <mkey>							#删除指定的map的键值对
  hlen <key>									#类似string同类型的方法
  hexists <key> <mkey>						#类似string同类型的方法
  hkeys <key>									#获取所有的map的key
  hvals <key>									#获取所有的map的value
  hincrby <key> <mkey> n						#类似string同类型的方法，没有自增1的方法
  hsetnx <key> <mkey> <mval>					#类似string同类型的方法
  ```

- Zset

  > 有序集合，在set的基础上多了排序的权重

  ```bash
  zadd <key> score <val>				#存值，score是此有序集合中的排序权重
  #3、4行的命令可以跟"withscores"把socre也输出
  zrange/zrevrange <key> start end	#和range类似，只不过是按score升序/降序输出，0~-1是查看所有
  zrangebyscore <key> min max			#类似上条，上面的是下标，这个是score，无穷是+inf/-inf
  zrem <key> <val>					#删除指定集合中val和它的score
  zcard <key>							#获取数量
  zcount <key> min max				#获取指定score区间范围的数量
  ```

# key的特殊数据类型

- geospecial

  > 底层是一个：Zset，因此能用Zset的命令来操作geospecial
  >
  > 有效的经度只能从-180\~180，纬度只能从-85.05112878\~85.05112878

  ```bash
  geoadd <key> longitude latitude region			#经度、纬度、地区，可以批量添加，南北极无法添加
  geopos <key> region1 region2					#获取指定城市的经纬度
  geodist <key> region1 region2 unit				#获取指定两个位置间的距离，单位：m km mi ft，省略时默认km
  georadius <key> longitude latitude radius unit	#以给定的经纬度为中心，得到指定半径范围内的元素，可以跟													 withcoord、withdist、count（指定查出的数量）
  georadiusbymember <key> region radius unit		#以region的名字为中心，而不需要以经纬度为中心
  #geohash，返回经纬度转换成的11个字符的字符串，一般用不到
  ```

- hyperloglog

  > 一种数据结构，是做基数统计的算法；只需要固定12KB内存，就能统计2^64个不同的元素，存在0.81%的误差
  >
  > 基数：集合中部重复元素的数量

  ```bash
  pfadd <key> <val1> <val2>		#存值
  pfcount <key>					#统计元素数量
  pfmerge <newkey> <key1> <key2>	#合并两个key到newkey中，合并后的元素会去重
  ```

- bitmap

  > 位存储，操作二进制位来记录。使用场景：统计巨大数量，存在两个状态的场景，都能使用

  ```bash
  setbit <key> offset <val>	#val只能是0或1，offset从0开始，不能批量设
  getbit <key> offset			#查看某个offset上的值
  bitcount <key> start end	#查看值为1的数量，start和end可省略
  ```

# 配置（redis.config）

- `unit`单位对大小写不敏感

- `INCLUDES`可以包含其他配置文件，可以把多个`redis.config`组合成一个

- `NETWORK`设置绑定的ip、保护模式和端口号

- `GENERAL`通用配置

  - `deamonize`，设为yes（以守护进程的方式运行，即可后台运行）
  - `pidfile`，如果上条设为yes，就需要指定一个pid文件
  - `loglevel`，日志级别，默认为`notice`，一般不改
  - `logfile`，日志的文件名和路径
  - `database`，数据库数量，默认为16
  - `always-show-logo`，是否显示redis的logo

- `SNAPSHOTTING`快照（rdb的配置）

  ```bash
  save 900 1					#如果900s内，至少有一个key进行了修改，则进行持久化操作
  save 300 10					#同上
  save 60 10000				#同上
  stop-writes-on-bgsave-error	#持久化出错后，redis是否继续工作
  rdbcompression				#是否压缩rdb文件（会消耗cpu资源）
  rdbchecksum					#保存rdb文件的时候，是否进行错误校验
  dir							#rdb保存的目录
  ```

- `REPLICATION`：主从复制

- `SECURITY`

  - 修改`requirepass`字段可以设置是否需要密码，后面直接跟密码
  - 也可通过`config set requirepass`来修改，登录时`auth <password>`来登录

- `CLIENTS`：客户端的限制，`maxclients`（最大连接的客户端数量）等配置

- `MEMORY MANAGEMENT`

  - 最大的内存容量（`maxmemory`配置项）

  - 内存达到上限后的处理策略（`maxmemory-policy`），六种策略

    > 1、volatile-lru：只对设置了过期时间的key进行LRU（默认值） 
    >
    > 2、allkeys-lru ： 删除lru算法的key  
    >
    > 3、volatile-random：随机删除即将过期key  
    >
    > 4、allkeys-random：随机删除  
    >
    > 5、volatile-ttl ： 删除即将过期的  
    >
    > 6、noeviction ： 永不过期，返回错误

- `APPEND ONLY MODE`：AOF的配置

  - `appendonly`是否开启`AOF`模式，默认不开启而使用`rdb`。一般场景rdb完全够用
  - `appendfilename`aof的文件名
  - `appendfsync`同步的时间，每次修改都同步 / 每秒都同步 / 不同步

# 持久化

> redis是内存数据库，如果没持久化，则断电即失

## RDB（Redis Database）

> 原理：在指定时间间隔内把内存中的数据集当做快照写入磁盘，恢复时则将快照文件读到内存中
>
> 工作过程：redis会创建（fork）一个子进程来进行持久化，会将数据写入临时文件，待持久化过程结束，用此临时文件替换上次持久化好的文件。整个过程中主进程不进行任何IO操作。
>
> 优点：大规模时效率较高
>
> 缺点：最后一次持久化时redis意外宕机，则最后一次修改的数据会丢失，恢复时也不保证完整性；fork进程会占内存空间

**触发生成`dump.rdb`**（生产环境中会备份此文件）

- 满足save规则时
- 执行flushall时
- 退出redis时

**恢复rdb文件中的数据**

- 把rdb放到redis的启动目录即可，会自动检查dump.rdb并恢复其中的数据
- 默认目录根据配置文件中的`dir`来，默认值就是redis的启动目录

![image-20210503162343047](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210503162343047.png)

## AOF

# 事务

> Redis单条命令能保证原子性，但事务不保证原子性；没有隔离级别的概念
>
> Redis事务的本质：一组命令的集合，是一个队列；事务中的命令并没有被直接执行，执行事务的时候才会按序执行命令

**redis事务的特性**

- 一次性
- 顺序性：一个事务中的所有命令都会被序列化，会按照顺序执行一系列的命令
- 排他性：事务执行之前，命令不会执行

**redis事务的过程**

- 开启事务：`multi`
- 命令入队：`一系列命令`
- 执行事务 / 放弃事务：`exec` / `discard`

**特殊情况**

- 命令入队过程中有一条命令语法错误（编译异常），则执行事务后，所有命令都不会被执行（`exex`后会报错）
- 命令入队过程中有一条命令执行的不对（运行时异常，如对字符串`incr`），则执行事务后，此条命令不会执行，其他命令不受影响（即redis事务不保证原子性）

> 悲观锁：认为什么时候都会出问题，无论做什么都枷锁
>
> 乐观锁：认为什么时候都不会出问题，不会上锁。只是更新的时候去判断一下，在此期间是否修改过数据
>
> - `watch <key>`事务开启之前监视某个key，事务执行之前其他线程操作了这个key，则执行事务时会失败，返回nil
> - 执行失败的话说明有人更新过值，先`unwacth`再重新`watch <key>`

# 订阅发布实现

# 主从复制

# 哨兵模式

# 缓存问题

# 缓存穿透及解决方案

# 缓存击穿及解决方案

# 缓存雪崩及解决方案