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
  - NoSQL没有固定的查询语言，存储方式多（键值对存储、列存储、文档存储、图形数据库)
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

## CAP

## BASE

# Redis介绍

> Remote Dictionary Server，远程字典服务，Key-Value存储
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

# 通用命令

- 默认16个数据库，默认使用第0个，可以用`select n`切换，切换后域名后面会有`[n]`显示，没有的话就是在第0个；`DBSIZE`查看数据库大小

- 官网查看命令[帮助文档](http://redis.cn/commands.html)

  ```bash
  ping				#测试连接
  EXPIRE <key> sec	#sec秒后key过期
  ttl <key>			#查看key的剩余过期时间，-1是不会过期，-2是已经过期，过期后就没了
  type <key>			#查看key的类型
  keys *				#查看所有的key
  flushdb / FLUSHALL	#清空当前数据库 / 清空全部数据库
  move <key> 1		#把此key移动到数据库1
  EXISTS <key>		#查看这个key是否存在
  config get/set xxx	#查看/修改某配置项
  del <key>			#删除key
  ```

# key的基本数据类型

- String

  > 可以是字符串，也可以是数字；场景：计数器、对象缓存存储等

  ```bash
  set <key> <val>						#插入键值对
  get <key>							#查看key对应的值
  APPEND <key> <val>					#在key对应的字符串值后追加val的值，如果当前key不存在，就自动添加一个键值对
  STRLEN <key>						#key对应的字符串值的长度
  incr <key> / decr <key>				#key对应的值自增1/自减1
  INCRBY <key> n / DECRBY <key> n		#key对应的值增加n/减少n
  GETRANGE <key> start end			#查看值的下标start到end的内容（闭区间），end是-1的话就是到尾
  SETRANGE <key> idx <val>			#修改指定下标位置的字符，把下标为idx开始的等长字符串替换成val
  setex <key> sec <val>				#设置过期时间（原子操作，设置值和设置过期实际一起完成）
  setnx <key> <val>					#如果不存在则设置（set if not exists）,成功返回1；如果key存在则返回0
  mset <key1> <val1> <key2> <val2>	#批量设置键值对
  mget <key1> <key2> <key3>			#批量查看key对应的值
  msetnx <key1> <val1> <key4> <val4>	#如果不存在则批量设置，是原子性操作，要么一起成功，要么一起失败
  getset <key> <val>					#先get再set，若存在则获取原来的值并设新值；若不存在则返回nil并设置值
  # 设置对象
  # 方式1：<obj>:<id>作为key，后面{}括起来的json是val，后面再去解析这个json
  set <obj>:<id> {<key1>:<val>,<key2>:<val2>}				#如set usr:1 {name:zs,age:10}
  # 方式2：<obj>:<id>:<key1>，<obj>:<id>:<key2>分别为两个key
  mset <obj>:<id>:<key1> <val1> <obj>:<id>:<key2> <val2>	#如mset usr:1:name:zs usr:1:age:10
  mget <obj>:<id>:<key1> <obj>:<id>:<key2>
  ```
  
- List

  > 所有的list命令都是L开头，LPUSH LPOP则是栈，LPUSH RPOP则是队列

  ```bash
  LPUSH <key1> <val1>	<val2>		#把val塞入key这个list；入栈
  LPOP <key>								#弹出栈顶元素，可以跟数字，弹出几个；出栈
  RPOP <key>								#移除栈底元素，可以跟数字，移除几个；出队
  RPUSH <key> <val>						#从另一个方向塞入list（从栈底塞进去）
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

> 修改配置后，重启redis生效

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
  save 900 1					#如果900s内，至少有一个key进行了修改，则900s后进行持久化操作
  save 300 10					#同上
  save 60 10000				#同上
  stop-writes-on-bgsave-error	#持久化出错后，redis是否继续工作
  rdbcompression				#是否压缩rdb文件（会消耗cpu资源）
  rdbchecksum					#保存rdb文件的时候，是否进行错误校验
  dir							#rdb保存的目录
  ```

- `REPLICATION`主从复制，只用配置从机的这部分内容

  - `replicaof <masterip> <masterport>`，主机的ip和端口
  - `masterauto <password>`主机的密码，可不配

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
  - `no-appendfsync-on-rewrite`、`auto-aof-rewrite-percentage 100`、`auto-aof-rewrite-min-size 64mb`重写aof文件的三个配置项

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

## AOF（Append Only File）

> 把所有命令都记录下来，当做history，恢复的时候就把这个文件中的命令全部执行一遍

- 只记录写操作，不记录读操作
- 只追加文件不改写文件，redis启动时会读取此文件来重新构建数据库
- 手动修改`appendonly.aof`的话，启动redis会失败，可以运行`redis-check-aof --fix`来修复

![image-20210503214656330](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210503214656330.png)

# 事务

> Redis单条命令能保证原子性，但事务不保证原子性；没有隔离级别的概念
>
> Redis事务的本质：一组命令的集合，是一个队列；事务中的命令并没有被直接执行，执行事务的时候才会按序执行命令

**redis事务的特性**

- 单独的隔离操作：事务中的所有命令都会被序列化，按序执行。事务执行过程中不会被其他客户端发来的命令打断
- 没有隔离级别：队列中的命令在提交之前都不会被执行
- 不保证原子性：事务执行过程中某条命令失败，则只是此条报错，其他命令仍执行

**redis事务的过程**

- 开启事务：`multi`
- 命令入队：`一系列命令`（组队阶段）
- 执行事务 / 放弃事务：`exec` / `discard`（执行阶段/放弃组队）

**特殊情况**

- 命令入队过程中有一条命令语法错误（编译异常），则执行事务后，所有命令都不会被执行（`exex`后会报错）
- 命令入队过程中有一条命令执行的不对（运行时异常，如对字符串`incr`），则执行事务后，此条命令不会执行，其他命令不受影响（即redis事务不保证原子性）

**模拟并发的工具**：`yum install httpd-tools`，使用方法`ab [options] [http[s]://]hostname[:port]/path`

> 悲观锁：认为什么时候都会出问题，每次取数据钱都要先加锁
>
> 乐观锁：认为什么时候都不会出问题，不会上锁。只是更新的时候去判断一下，在此期间是否修改过数据
>
> - `watch <key>`事务开启之前监视某个key，事务执行之前其他线程操作了这个key，则执行事务时会失败，返回nil
> - 执行失败的话说明有人更新过值，先`unwacth`再重新`watch <key>`

**库存问题**

1. 超时连接问题：用连接池
2. 超卖问题：即订单数为负值，用乐观锁
3. 库存遗留问题：即加了乐观锁后一些抢单操作不会被执行，导致库存剩了很多，用LUA脚本

# 发布订阅

- 三个角色：消息发送者，频道，消息接收者
- 命令：`subscribe <chann>`订阅频道，`publish <channel> <msg>`向频道发送消息，`unsubscribe <chann>`取消订阅频道

![image-20210503221505925](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210503221505925.png)

# 主从复制

## 概念

> 将一台redis服务器的数据，复制到其他redis的服务器。前者为主节点（master/leader)，后者为从节点（slave/follower）
>
> 数据的复制是单向的，只能由主节点到从节点，==master以写为主，slave以读为主==，读写分离可以减轻服务器压力
>
> ==默认情况下每台redis服务器都是主节点，最低配是一主二从，一般单台redis的最大内存不应超20G==

**作用**

- 数据冗余：实现了数据的热备份，是持久化之外的一种数据冗余方式
- 故障恢复：主节点有问题时可由从节点提供服务，实现快速的故障恢复，实际上是一种服务的冗余
- 负载均衡：读写分离，提高redis服务器的并发量
- 高可用基石：主从复制是哨兵和集群能够实施的基础

## 配置

- 复制出三个配置文件，一主二从，分别为79、80、81

- 改端口、pid文件名、log文件名、dump.rdb文件名（按79、80、81改）

- 以配置文件启动三个`redis-server`

- 只配置从库，不配置主库

  - 一主（79）二从（80、81），主机写，从机读，从机写的时候会报错
  - 链路，79主，80是79的节点，81是80的从节点，但80仍然无法写操作
    - 此模式下，主机断开后，80可自己当主机，即使79重新连接，也没用，必须手动重新配

  ```bash
  #主机断开连接后，从机仍然连着，但是无法写操作；重启后，从机可继续获取主机的信息
  #从机断开连接重启后，如果配置文件配置了主机信息，可继续获取主机信息；如果是命令行配置的，就不会，因为默认是主机
  slaveof <ip> <port>	#认老大
  info replication	#查看当前库的主从信息
  slaveof no one		#从机自己当主机
  ```

- 实际开发不用命令配置，而是修改配置文件中`REPLICATION`部分的配置项

![image-20210503232030687](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210503232030687.png)

# 哨兵模式

> 自动选举主机的模式，哨兵是一个独立的进程。原理是哨兵通过发送命令，等待redis服务器响应，从而监控运行的多个redis实例
>
> 主机宕机后，通过发布订阅模式通知其他服务器，修改配置文件，切换主机；原主机恢复后，只能给新主机当从机
>
> 一个哨兵进程可能会出问题，因此需要多个哨兵进行监控，各哨兵之间也会进行监控

1. 配置哨兵，`sentinel.conf`
   - `sentinel monitor 被监控的主机名称 host port quorum`，quorum代表多少个哨兵认为主机失联时就客观确认主机失联
2. 启动哨兵`redis-sentinel sentinel.conf`

![image-20210504105343034](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210504105343034.png)

![image-20210504104856796](D:\资料\Go\src\studygo\Golang学习笔记\redis笔记.assets\image-20210504104856796.png)

缺点

- 不好在线扩容，集群容量到达上限后，在线扩容和麻烦
- 配置项太多

# 缓存问题

## 缓存穿透

- 概念

  > 现象：大量请求应用服务器压力变大，redis命中率降低，缓存中查不到，一直去持续去持久层数据库查询
  >
  > 原因：被攻击而出现大量非正常等url会导致此问题

- 解决
  - 实时监控，发现redis命中率急速降低时设置黑名单
  - 布隆过滤器(底层也是bitmap)
  - 缓存空对象，但要设置过期时间，最长不超过5分钟
  - 设置可访问的白名单，id作为偏移量存档bitmap中

## 缓存击穿

- 概念

  > 现象：对mysql的访问压力瞬时增加，但redis中没有出现大量key过期，而redis正常运行
  >
  > 原因：redis中某个key过期，大量访问使用这个key

- 解决
    - 预先设置热门数据到redis中，并加长这些key的时长
    - 实时监控热门数据，现场调整key的过期时间
    - 加互斥锁

## 缓存雪崩

- 概念

  > 现象：数据库压力变大，服务器崩溃
  >
  > 原因：一个时间段内缓存集中过期失效

- 解决

  - 构建多级缓存架构：nginx缓存+redis缓存+其他缓存（ehcache等）
  - 使用锁或队列，但不适合高并发
  - 记录缓存是否过期（设置提前量），如果过期就起另外线程更新缓存
  - 将缓存的失效时间分散开

## 分布式锁

- 概念：加一把锁，不仅对单机有效，而且对整个集群都有效
- 方案
  - 基于数据库实现分布式锁
  - 基于缓存
    - `setnx <key> <val>`来设置锁，`del <key>`或key过期后即释放了锁（`expire <key> sec`）
    - 或`set <key> <val> nx ex sec`来完成设置值和设置锁的原子操作
    - `set lock uuid nx ex sec `设置uuid来标识锁，避免误删锁
  - 基于zookeeper

# go操作redis

## 安装

区别于另一个比较常用的Go语言redis client库：[redigo](https://github.com/gomodule/redigo)，我们这里采用https://github.com/go-redis/redis连接Redis数据库并进行操作，因为`go-redis`支持连接哨兵及集群模式的Redis。

使用以下命令下载并安装:

```bash
go get -u github.com/go-redis/redis
```

## 连接

### 普通连接

```go
// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
```

## V8新版本相关

最新版本的`go-redis`库的相关命令都需要传递`context.Context`参数，例如：

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var (
	rdb *redis.Client
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
```

### 连接Redis哨兵模式

```go
func initClient()(err error){
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
```

### 连接Redis集群

```go
func initClient()(err error){
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
```

## 基本使用

### set/get示例

```go
func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}
```

### zset示例

```go
func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
```

输出结果如下：

```bash
$ ./06redis_demo 
zadd 0 succ.
Golang's score is 100.000000 now.
Golang 100
C/C++ 99
Java 98
JavaScript 97
Java 98
C/C++ 99
Golang 100
```

### 根据前缀获取Key

```go
vals, err := rdb.Keys(ctx, "prefix*").Result()
```

### 执行自定义命令

```go
res, err := rdb.Do(ctx, "set", "key", "value").Result()
```

### 按通配符删除key

当通配符匹配的key的数量不多时，可以使用`Keys()`得到所有的key在使用`Del`命令删除。 如果key的数量非常多的时候，我们可以搭配使用`Scan`命令和`Del`命令完成删除。

```go
ctx := context.Background()
iter := rdb.Scan(ctx, 0, "prefix*", 0).Iterator()
for iter.Next(ctx) {
	err := rdb.Del(ctx, iter.Val()).Err()
	if err != nil {
		panic(err)
	}
}
if err := iter.Err(); err != nil {
	panic(err)
}
```

### Pipeline

> 不适用的场景：pipeline的命令中，后面的命令依赖前面命令的结果，就不能用了

`Pipeline` 主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间（RTT）。

`Pipeline` 基本示例如下：

```go
pipe := rdb.Pipeline()

incr := pipe.Incr("pipeline_counter")
pipe.Expire("pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
```

上面的代码相当于将以下两个命令一次发给redis server端执行，与不使用`Pipeline`相比能减少一次RTT。

```bash
INCR pipeline_counter
EXPIRE pipeline_counts 3600
```

也可以使用`Pipelined`：

```go
var incr *redis.IntCmd
_, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
	incr = pipe.Incr("pipelined_counter")
	pipe.Expire("pipelined_counter", time.Hour)
	return nil
})
fmt.Println(incr.Val(), err)
```

在某些场景下，当我们有多条命令要执行时，就可以考虑使用pipeline来优化。

### 事务

Redis是单线程的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，例如在它们之间交替执行。但是，`Multi/exec`能够确保在`multi/exec`两个语句之间的命令之间没有其他客户端正在执行命令。

在这种场景我们需要使用`TxPipeline`。`TxPipeline`总体上类似于上面的`Pipeline`，但是它内部会使用`MULTI/EXEC`包裹排队的命令。例如：

```go
pipe := rdb.TxPipeline()

incr := pipe.Incr("tx_pipeline_counter")
pipe.Expire("tx_pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
```

上面代码相当于在一个RTT下执行了下面的redis命令：

```bash
MULTI
INCR pipeline_counter
EXPIRE pipeline_counts 3600
EXEC
```

还有一个与上文类似的`TxPipelined`方法，使用方法如下：

```go
var incr *redis.IntCmd
_, err := rdb.TxPipelined(func(pipe redis.Pipeliner) error {
	incr = pipe.Incr("tx_pipelined_counter")
	pipe.Expire("tx_pipelined_counter", time.Hour)
	return nil
})
fmt.Println(incr.Val(), err)
```

### Watch

在某些场景下，我们除了要使用`MULTI/EXEC`命令外，还需要配合使用`WATCH`命令。在用户使用`WATCH`命令监视某个键之后，直到该用户执行`EXEC`命令的这段时间里，如果有其他用户抢先对被监视的键进行了替换、更新、删除等操作，那么当用户尝试执行`EXEC`的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。

```go
Watch(fn func(*Tx) error, keys ...string) error
```

Watch方法接收一个函数和一个或多个key作为参数。基本使用示例如下：

```go
// 监视watch_count的值，并在值不变的前提下将其值+1
key := "watch_count"
err = client.Watch(func(tx *redis.Tx) error {
	n, err := tx.Get(key).Int()
	if err != nil && err != redis.Nil {
		return err
	}
	_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Set(key, n+1, 0)
		return nil
	})
	return err
}, key)
```

最后看一个V8版本官方文档中使用GET和SET命令以事务方式递增Key的值的示例，仅当Key的值不发生变化时提交一个事务。

```go
func transactionDemo() {
	var (
		maxRetries   = 1000
		routineCount = 10
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Increment 使用GET和SET命令以事务方式递增Key的值
	increment := func(key string) error {
		// 事务函数
		txf := func(tx *redis.Tx) error {
			// 获得key的当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际的操作代码（乐观锁定中的本地操作）
			n++

			// 操作仅在 Watch 的 Key 没发生变化的情况下提交
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// 最多重试 maxRetries 次
		for i := 0; i < maxRetries; i++ {
			err := rdb.Watch(ctx, txf, key)
			if err == nil {
				// 成功
				return nil
			}
			if err == redis.TxFailedErr {
				// 乐观锁丢失 重试
				continue
			}
			// 返回其他的错误
			return err
		}

		return errors.New("increment reached maximum number of retries")
	}

	// 模拟 routineCount 个并发同时去修改 counter3 的值
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()
			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(context.TODO(), "counter3").Int()
	fmt.Println("ended with", n, err)
}
```