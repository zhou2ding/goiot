# NoSQL

## 概述

> 单机MySQL\===>Memcached（缓存）+读写分离+垂直拆分=\==>分库分表+水平拆分+MySQL集群
>
> 为什么用NoSQL？数据爆发式增长，位置信息、热榜、图片等等，web2.0的诞生，超大规模的高并发，传统的关系型数据库无法对付

![image-20210430111738570](E:\study\studygo\Golang学习笔记\redis笔记.assets\image-20210430111738570.png)

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

![image-20210430160658134](E:\study\studygo\Golang学习笔记\redis笔记.assets\image-20210430160658134.png)

## CAP

**一个分布式系统最多只能同时满足一致性（Consistency）、可用性（Availability）和分区容错性（Partition tolerance）这三项中的两项**。由于网络的不可靠性质，大多数开源的分布式系统都会实现P，也就是分区容忍性，之后在C和A中做抉择。

虽然说我们设计系统时不能同时保证拥有三点。但是也并不是说，保证了其中2点后，就要完全抛弃另外一点。只是相对的要做一些牺牲。比如在保证CP的情况下，虽然没办法保证高可用性，但这不意味着可用性为0，我们可以通过合理的设计尽量的提高可用性，让可用性尽可能的接近100%。同理，在AP的情况下，也可以尽量的保证数据的一致性，或者实现弱一致性，即最终一致性。

![img](https://pic4.zhimg.com/80/v2-2f26a48f5549c2bc4932fdf88ba4f72f_720w.webp)

*Consistency 一致性*

一致性指“`all nodes see the same data at the same time`”，即所有节点在同一时间的数据完全一致。

一致性是因为多个数据拷贝下并发读写才有的问题，因此理解时一定要注意结合考虑多个数据拷贝下并发读写的场景。

对于一致性，可以分为强/弱/最终一致性三类

- 强一致性（Strict Consistency）：系统中的某个数据被成功更新后，后续任何对该数据的读取操作都将得到更新后的值；也称为原子一致性（Atomic Consistency）或线性一致性（Linearizable Consistency）两个要求：
  - 任何一次读都能读到某个数据的最近一次写的数据。
  - 系统中的所有进程，看到的操作顺序，都和全局时钟下的顺序一致。**简言之，在任意时刻，所有节点中的数据是一样的。**

- 弱一致性：系统中的某个数据被更新后，后续对该数据的读取操作可能得到更新后的值，也可能是更改前的值。但即使过了“不一致时间窗口”这段时间后，后续对该数据的读取也不一定是最新之；所以说，可以理解为数据更新后，如果能容忍后续的访问只能访问到部分或者全部访问不到，则是弱一致性。

- 最终一致性：是弱一致性的特殊形式，存储系统保证在没有新的更新的条件下，最终所有的访问都是最后更新的值。不保证在任意时刻任意节点上的同一份数据都是相同的，但是随着时间的迁移，不同节点上的同一份数据总是在向趋同的方向变化。简单说，**就是在一段时间后，节点间的数据会最终达到一致状态**。

*Availability 可用性*

可用性指“`Reads and writes always succeed`”，即服务在正常响应时间内一直可用。

好的可用性主要是指系统能够很好的为用户服务，不出现用户操作失败或者访问超时等用户体验不好的情况。可用性通常情况下可用性和分布式数据冗余，负载均衡等有着很大的关联。

*Partition Tolerance分区容错性*

分区容错性指“`the system continues to operate despite arbitrary message loss or failure of part of the system`”，由于分布式系统通过网络进行通信，网络是不可靠的。当任意数量的消息丢失或延迟到达时，系统仍会继续提供服务，不会挂掉。即分布式系统在遇到某节点或网络分区故障的时候（如节点之间无法通信），仍然能够对外提供满足一致性或可用性的服务。

## BASE

> 核心思想：既是无法做到强一致性（Strong consistency），但每个应用都可以根据自身的业务特点，采用适当的方式来使系统达到最终一致性（Eventual consistency）。

*Basically Available 基本可用*

假设系统，出现了不可预知的故障，但还是能用，相比较正常的系统而言会有响应时间和功能上的损失：

1. **响应时间上的损失**：正常情况下的搜索引擎0.5秒即返回给用户结果，而基本可用的搜索引擎可以在2秒作用返回结果。
2. **功能上的损失**：在一个电商网站上，正常情况下，用户可以顺利完成每一笔订单。但是到了大促期间，为了保护购物系统的稳定性，部分消费者可能会被引导到一个降级页面。

*Soft state(软状态)*

相对于原子性而言，要求多个节点的数据副本都是一致的，这是一种“硬状态”。

软状态指的是：允许系统中的数据存在中间状态，并认为该状态不影响系统的整体可用性，即允许系统在多个不同节点的数据副本存在数据延时。

如先把订单状态改成已支付成功，然后告诉用户已经成功了；剩下在异步发送mq消息通知库存服务减库存，即使消费失败,MQ消息也会重新发送（重试）。

注：不可能一直是软状态，必须有个时间期限。在期限过后，应当保证所有副本保持数据一致性，从而达到数据的最终一致性。这个时间期限取决于网络延时、系统负载、数据复制方案设计等等因素。

*Eventually consistent(最终一致性)*

- 强一致性读操作要么处于阻塞状态，要么读到的是最新的数据
- 最终一致性通常是异步完成的，读到的数据刚开始可能是旧数据，但是过一段时间后读到的就是最新的数据

## 分布式事务

### X/A

XA 是由 X/Open 组织提出的分布式事务的规范，XA 规范主要定义了(全局)事务管理器(TM)和(局部)资源管理器(RM)之间的接口。本地的数据库如 mysql 在 XA 中扮演的是 RM 角色

XA 一共分为两阶段：

> 第一阶段（prepare）：即所有的参与者 RM 准备执行事务并锁住需要的资源。参与者 ready 时，向 TM 报告已准备就绪。
>
> 第二阶段 (commit/rollback)：当事务管理者(TM)确认所有参与者(RM)都 ready 后，向所有参与者发送 commit 命令。

目前主流的数据库基本都支持 XA 事务，包括 mysql、oracle、sqlserver、postgre

XA 事务由一个或多个资源管理器（RM）、一个事务管理器（TM）和一个应用程序（ApplicationProgram）组成。

把上面的转账作为例子，一个成功完成的 XA 事务时序图如下：

![img](https://static001.geekbang.org/infoq/29/2955cd73080bed2e0858ed5074e77445.jpeg?x-oss-process=image%2Fresize%2Cp_80%2Fauto-orient%2C1)

如果有任何一个参与者 prepare 失败，那么 TM 会通知所有完成 prepare 的参与者进行回滚。

XA 事务的特点是：

> 简单易理解，开发较容易
>
> 对资源进行了长时间的锁定，并发度低

go 语言可参考[DTM](https://github.com/dtm-labs/dtm)

### SAGA

其核心思想是将长事务拆分为多个本地短事务，由 Saga 事务协调器协调，如果正常结束那就正常完成，如果某个步骤失败，则根据相反顺序一次调用补偿操作。

把上面的转账作为例子，一个成功完成的 SAGA 事务时序图如下：

![img](https://static001.geekbang.org/infoq/11/1164e5613acd9230805b7ee9d8e9c948.jpeg?x-oss-process=image%2Fresize%2Cp_80%2Fauto-orient%2C1)

SAGA 事务的特点：

> 并发度高，不用像 XA 事务那样长期锁定资源
>
> 需要定义正常操作以及补偿操作，开发量比 XA 大
>
> 一致性较弱，对于转账，可能发生 A 用户已扣款，最后转账又失败的情况

SAGA 内容较多，包括两种恢复策略，包括分支事务并发执行，这里仅包括最简单的 SAGA

SAGA 适用的场景较多，长事务适用，对中间结果不敏感的业务场景适用

go 语言可参考[DTM](https://github.com/dtm-labs/dtm)

### TCC

TCC 分为 3 个阶段

- Try 阶段：尝试执行，完成所有业务检查（一致性）, 预留必须业务资源（准隔离性）
- Confirm 阶段：确认执行真正执行业务，不作任何业务检查，只使用 Try 阶段预留的业务资源，Confirm 操作要求具备幂等设计，Confirm 失败后需要进行重试。
- Cancel 阶段：取消执行，释放 Try 阶段预留的业务资源。Cancel 阶段的异常和 Confirm 阶段异常处理方案基本上一致，要求满足幂等设计。

把上面的转账作为例子，通常会在 Try 里面冻结金额，但不扣款，Confirm 里面扣款，Cancel 里面解冻金额，一个成功完成的 TCC 事务时序图如下：

![img](https://static001.geekbang.org/infoq/d0/d038f3c4bb392307a395d12513a29418.jpeg?x-oss-process=image%2Fresize%2Cp_80%2Fauto-orient%2C1)

TCC 特点如下：

> 并发度较高，无长期资源锁定。
>
> 开发量较大，需要提供 Try/Confirm/Cancel 接口。
>
> 一致性较好，不会发生 SAGA 已扣款最后又转账失败的情况

TCC 适用于订单类业务，对中间状态有约束的业务

### 事务消息

阿里开源的 RocketMQ 4.3 之后的版本正式支持事务消息，该事务消息本质上是把本地消息表放到 RocketMQ 上，解决生产端的消息发送与本地事务执行的原子性问题。

事务消息发送及提交：

- 发送消息（half 消息）
- 服务端存储消息，并响应消息的写入结果
- 根据发送结果执行本地事务（如果写入失败，此时 half 消息对业务不可见，本地逻辑不执行）
- 根据本地事务状态执行 Commit 或者 Rollback（Commit 操作发布消息，消息对消费者可见）

正常发送的流程图如下：

![img](https://static001.geekbang.org/infoq/44/44fdbd27567440fd21334c78c91d56a0.jpeg?x-oss-process=image%2Fresize%2Cp_80%2Fauto-orient%2C1)

补偿流程：

- 对没有 Commit/Rollback 的事务消息（pending 状态的消息），从服务端发起一次“回查”
- Producer 收到回查消息，返回消息对应的本地事务的状态，为 Commit 或者 Rollback

事务消息方案与本地消息表机制非常类似，区别主要在于原先相关的本地表操作替换成了一个反查接口

事务消息特点如下：

> 长事务仅需要分拆成多个任务，并提供一个反查接口，使用简单
>
> 消费者的逻辑如果无法通过重试成功，那么还需要更多的机制，来回滚操作

适用于可异步执行的业务，且后续操作无需回滚的业务

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

  ![image-20210430172506575](E:\study\studygo\Golang学习笔记\redis笔记.assets\image-20210430172506575.png)

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

# 线程模型

**Redis 单线程指的是「接收客户端请求->解析请求 ->进行数据读写等操作->发送数据给客户端」这个过程是由一个线程（主线程）来完成的**，这也是我们常说 Redis 是单线程的原因。

但是，**Redis 程序并不是单线程的**，Redis 在启动的时候，是会**启动后台线程**（BIO）来处理关闭文件、AOF 刷盘、异步释放 Redis 内存（lazyfree线程）

之所以 Redis 采用单线程（网络 I/O 和执行命令）那么快，有如下几个原因：

- Redis 的大部分操作**都在内存中完成**，并且**采用了高效的数据结构**，因此 Redis 瓶颈可能是机器的内存或者网络带宽，而并非 CPU，既然 CPU 不是瓶颈，那么自然就采用单线程的解决方案了；
- Redis 采用单线程模型可以**避免了多线程之间的竞争**，省去了多线程切换带来的时间和性能上的开销，而且也不会导致死锁问题。
- Redis 采用了 **I/O 多路复用机制**处理大量的客户端 Socket 请求，IO 多路复用机制是指一个线程处理多个 IO 流，就是我们经常听到的 select/epoll 机制。简单来说，在 Redis 只运行单线程的情况下，该机制允许内核中，同时存在多个监听 Socket 和已连接 Socket。内核会一直监听这些 Socket 上的连接请求或数据请求。一旦有请求到达，就会交给 Redis 线程处理，这就实现了一个 Redis 线程处理多个 IO 流的效果。

# key的基本数据类型

- String

  > String 类型的底层的数据结构实现主要是 SDS（简单动态字符串）。
>
  > - **SDS 不仅可以保存文本数据，还可以保存图片、音频、视频、压缩文件这样的二进制数据**。
  > - **SDS 获取字符串长度的时间复杂度是 O(1)**。
  > - **Redis 的 SDS API 是安全的，拼接字符串不会造成缓冲区溢出**。
  
  ```bash
  set <key> <val>						#插入键值对
  get <key>							#查看key对应的值
  APPEND <key> <val>					#在key对应的字符串值后追加val的值，如果当前key不存在，就自动添加一个键值对
  STRLEN <key>						#key对应的字符串值的长度
  incr <key> / decr <key>				#key对应的值自增1/自减1
  INCRBY <key> n / DECRBY <key> n		#key对应的值增加n/减少n
  GETRANGE <key> start end			#查看值的下标start到end的内容（闭区间），end是-1的话就是到尾，第一个字符下标是0
  SETRANGE <key> idx <val>			#修改指定下标位置的字符，把下标为idx开始的等长字符串（和val）替换成val
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
  >
  > **List 数据类型底层数据结构就只由 quicklist （快速列表）实现**。

  ```bash
  LPUSH <key1> <val1>	<val2>				#把val塞入key这个list；入栈
  LPOP <key>								#弹出栈顶元素，可以跟数字（6.2以上版本才支持），弹出几个；出栈
  RPOP <key>								#移除栈底元素，可以跟数字（6.2以上版本才支持），移除几个；出队
  RPUSH <key> <val>						#从另一个方向塞入list（从栈底塞进去）
  LRANGE <key> start end					#获取下标区间内的值，按出栈顺序获，第一个元素下标是0，end为-1表示到栈底
  LINDEX <key> idx						#查看idx下标对应的值，下标顺序是从栈顶（下标0）到栈底（-1或栈底的实际下标）
  LLEN <key>								#查看长度
  LREM <key> <count> <val>				#移除指定个数的某个值，是从栈顶往下移除的
  LTRIM <key> start end					#截取指定长度区间的一串值，其余的值删掉。如果是截取1~0，则清空list
  RPOPLPUSH <src> <dst>					#把src的栈底元素移动到dst中，dst不存在的话新建一个dst
  LSET <key> idx <val>					#把指定下标处的值替换为val，key或val不存在的话都会报错
  LINSERT <key> direction <val> <newval>	#把newval插入val的前面或后面，direction是before/after
  ```

- Set

  > 无序不重复集合， 添加已有元素时会添加失败
  >
  > Set 类型的底层数据结构是由**哈希表或整数集合**实现的：如果集合中的元素都是整数且元素个数小于 512 个则用整数集合，否则使用哈希表

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

  > 存的值是一个map集合，也有自己的key，和string一样，只不过元素是键值对。以下命令中mkey为map中的key
  >
  > Hash 类型的底层数据结构是由**压缩列表或哈希表**实现的；在 Redis 7.0 中，压缩列表数据结构已经废弃了，交由 **listpack** 数据结构来实现了。

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
>
  > Zset 类型的底层数据结构是由**压缩列表或跳表**实现的；在 Redis 7.0 中，压缩列表数据结构已经废弃了，交由 **listpack** 数据结构来实现了。
  
  ```bash
  zadd <key> score <val>				#存值，score是此有序集合中的排序权重
  #3、4行的命令可以跟"withscores"把socre也输出
  zrange/zrevrange <key> start end	#和range类似，只不过是按score升序/降序输出，0~-1是查看所有
  zrangebyscore <key> min max			#类似上条，上面的是下标，这个是score，无穷是+inf/-inf
  zrem <key> <val>					#删除指定集合中val和它的score
  zcard <key>							#获取数量
  zcount <key> min max				#获取指定score区间范围的数量
  zscore <key> <val>					#获取指定元素的分数
  ```

# key的特殊数据类型

- geospecial

  > 存储地理位置信息的场景
  >
  > 底层是一个：Zset，因此能用Zset的命令来操作geospecial

  ```bash
  geoadd <key> longitude latitude region			#经度、纬度、地区，可以批量添加，南北极无法添加
  geopos <key> region1 region2					#获取指定城市的经纬度
  geodist <key> region1 region2 unit				#获取指定两个位置间的距离，单位：m km mi ft，省略时默认km
  georadius <key> longitude latitude radius unit	#以给定的经纬度为中心，得到指定半径范围内的元素，可以跟													 withcoord、withdist、count（指定查出的数量）
  georadiusbymember <key> region radius unit		#以region的名字为中心，而不需要以经纬度为中心
  #geohash，返回经纬度转换成的11个字符的字符串，一般用不到
  ```

- hyperloglog

  > 海量数据基数统计的场景，比如百万级网页 UV 计数等；
  
  ```bash
pfadd <key> <val1> <val2>		#存值
  pfcount <key>					#统计元素数量
  pfmerge <newkey> <key1> <key2>	#合并两个key到newkey中，合并后的元素会去重
  ```
  
- bitmap

  > 位存储，操作二进制位来记录。二值状态统计的场景

  ```bash
  setbit <key> offset <val>	#val只能是0或1，offset从0开始，不能批量设
  getbit <key> offset			#查看某个offset上的值
  bitcount <key> start end	#查看值为1的数量，start和end可省略
  ```

- stream

  > 消息队列，相比于基于 List 类型实现的消息队列，有这两个特有的特性：自动生成全局唯一消息ID，支持以消费组形式消费数据。

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

**触发生成`dump.rdb`**

- 满足save规则时
- 执行flushall时
- 退出redis时

### 快照

Redis 提供了两个命令来生成 RDB 文件，分别是 `save` 和 `bgsave`，他们的区别就在于是否在「主线程」里执行：

- 执行了 save 命令，就会在主线程生成 RDB 文件，由于和执行操作命令在同一个线程，所以如果写入 RDB 文件的时间太长，**会阻塞主线程**；
- 执行了 bgsave 命令，会创建一个子进程来生成 RDB 文件，这样可以**避免主线程的阻塞**；

RDB 文件的加载工作是在服务器启动时自动执行的，Redis 并没有提供专门用于加载 RDB 文件的命令。

Redis 还可以通过配置文件的选项来实现每隔一段时间自动执行一次 bgsave 命令，默认会提供以下配置，只要满足任意一个，就会执行 bgsave：

```text
save 900 1		# 900 秒之内，对数据库进行了至少 1 次修改；下同
save 300 10
save 60 10000
```

Redis 的快照是**全量快照**，也就是说每次执行快照，都是把内存中的「所有数据」都记录到磁盘中。执行快照是一个比较重的操作，如果频率太频繁，可能会对 Redis 性能产生影响。如果频率太低，服务器故障时，丢失的数据会更多。

### 快照实现

执行 bgsave 过程中，Redis 依然**可以继续处理操作命令**的，也就是数据是能被修改的。关键的技术就在于**写时复制技术（Copy-On-Write, COW）。**

执行 bgsave 命令的时候，会通过 `fork()` 创建子进程，此时子进程和父进程是共享同一片内存数据的，因为创建子进程的时候，会复制父进程的页表，但是页表指向的物理内存还是一个。

![图片](https://img-blog.csdnimg.cn/img_convert/c34a9d1f58d602ff1fe8601f7270baa7.png)

所以，创建 bgsave 子进程后，由于共享父进程的所有内存数据，于是就可以直接读取主线程（父进程）里的内存数据，并将数据写入到 RDB 文件，执行时不会阻塞主线程，这就使得主线程同时可以修改数据。

当主线程（父进程）要**修改共享数据里的某一块数据**（比如键值对 `A`）时，就会发生写时复制，于是这块数据的**物理内存就会被复制一份（键值对 `A'`）**，然后**主线程在这个数据副本（键值对 `A'`）进行修改操作**。与此同时，**bgsave 子进程可以继续把原来的数据（键值对 `A`）写入到 RDB 文件**，而主线程刚修改的数据，是没办法在这一时间写入 RDB 文件的，只能交由下一次的 bgsave 快照。如果系统恰好在 RDB 快照文件创建完毕后崩溃了，那么 Redis 将会丢失主线程在快照期间修改的数据。

## AOF（Append Only File）

> 把所有命令都记录下来，当做history，恢复的时候就把这个文件中的命令全部执行一遍

- 只记录写操作，不记录读操作
- 只追加文件不改写文件，redis启动时会读取此文件来重新构建数据库
- 手动修改`appendonly.aof`的话，启动redis会失败，可以运行`redis-check-aof --fix`来修复

Redis 写入 AOF 日志的过程，如下图：

![img](https://img-blog.csdnimg.cn/img_convert/4eeef4dd1bedd2ffe0b84d4eaa0dbdea.png)

1. Redis 执行完写操作命令后，会将命令追加到 AOF 缓冲区；
2. 然后通过 write() 系统调用，将 AOF 缓冲区的数据写入到 AOF 文件，此时数据并没有写入到硬盘，而是拷贝到了内核缓冲区 page cache，等待内核将数据写入硬盘；
3. 具体内核缓冲区的数据什么时候写入到硬盘，由内核决定。

### 写回策略

在 `redis.conf` 配置文件中的 `appendfsync` 配置项可以有以下 3 种参数可填：

- **Always**，这个单词的意思是「总是」，所以它的意思是每次写操作命令执行完后，同步将 AOF 日志数据写回硬盘；Always 策略就是每次写入 AOF 文件数据后，就执行 fsync() 函数。
- **Everysec**，这个单词的意思是「每秒」，所以它的意思是每次写操作命令执行完后，先将命令写入到 AOF 文件的内核缓冲区，然后每隔一秒将缓冲区里的内容写回到硬盘；Everysec 策略就会创建一个异步任务来执行 fsync() 函数。
- **No**，意味着不由 Redis 控制写回硬盘的时机，转交给操作系统控制写回的时机，也就是每次写操作命令执行完后，先将命令写入到 AOF 文件的内核缓冲区，再由操作系统决定何时将缓冲区内容写回硬盘。No 策略就是永不执行 fsync() 函数。

![img](https://img-blog.csdnimg.cn/img_convert/98987d9417b2bab43087f45fc959d32a.png)

### AOF重写

AOF 日志是一个文件，随着执行的写操作命令越来越多，文件的大小会越来越大。

如果当 AOF 日志文件过大就会带来性能问题，比如重启 Redis 后，需要读 AOF 文件的内容以恢复数据，如果文件过大，整个恢复的过程就会很慢。

所以，Redis 为了避免 AOF 文件越写越大，提供了 **AOF 重写机制**，当 AOF 文件的大小超过所设定的阈值后，Redis 就会启用 AOF 重写机制，来压缩 AOF 文件。

AOF 重写机制是在重写时，读取当前数据库中的所有键值对，然后将每一个键值对用一条命令记录到「新的 AOF 文件」，等到全部记录完后，就将新的 AOF 文件替换掉现有的 AOF 文件。

重写机制的妙处在于，尽管某个键值对被多条写命令反复修改，**最终也只需要根据这个「键值对」当前的最新状态，然后用一条命令去记录键值对**，代替之前记录这个键值对的多条命令，这样就减少了 AOF 文件中的命令数量。最后在重写工作完成后，将新的 AOF 文件覆盖现有的 AOF 文件。

Redis 的**重写 AOF 过程是由后台子进程 *bgrewriteaof* 来完成的（由主进程通过fork系统调用生成）**，在 bgrewriteaof 子进程执行 AOF 重写期间，主进程需要执行以下三个工作:

- 执行客户端发来的命令；
- 将执行后的写命令追加到 「AOF 缓冲区」；
- 将执行后的写命令追加到 「AOF 重写缓冲区」；

当子进程完成 AOF 重写工作（*扫描数据库中所有数据，逐一把内存数据的键值对转换成一条命令，再将命令记录到重写日志*）后，会向主进程发送一条信号，信号是进程间通讯的一种方式，且是异步的。

主进程收到该信号后，会调用一个信号处理函数，该函数主要做以下工作：

- 将 AOF 重写缓冲区中的所有内容追加到新的 AOF 的文件中，使得新旧两个 AOF 文件所保存的数据库状态一致；
- 新的 AOF 的文件进行改名，覆盖现有的 AOF 文件。

信号函数执行完后，主进程就可以继续像往常一样处理命令了。

在整个 AOF 后台重写过程中，除了发生写时复制会对主进程造成阻塞，还有信号处理函数执行时也会对主进程造成阻塞，在其他时候，AOF 后台重写都不会阻塞主进程。

## 对比

- RDB 快照就是记录某一个瞬间的内存数据，记录的是实际数据，而 AOF 文件记录的是命令操作的日志，而不是实际的数据。因此在 Redis 恢复数据时， RDB 恢复数据的效率会比 AOF 高些，因为直接将 RDB 文件读入内存就可以，不需要像 AOF 那样还需要额外执行操作命令的步骤才能恢复数据。
- RDB通常可能设置至少 5 分钟才保存一次快照，这时如果 Redis 出现宕机等情况，则意味着最多可能丢失 5 分钟数据。这就是 RDB 快照的缺点，在服务器发生故障时，丢失的数据会比 AOF 持久化的方式更多，因为 RDB 快照是全量快照的方式，因此执行的频率不能太频繁，否则会影响 Redis 性能，而 AOF 日志可以以秒级的方式记录操作命令，所以丢失的数据就相对更少。

## 混合持久化

在 Redis 4.0 提出的，该方法叫**混合使用 AOF 日志和内存快照**，也叫混合持久化。

如果想要开启混合持久化功能，可以在 Redis 配置文件将下面这个配置项设置成 yes：

```text
aof-use-rdb-preamble yes
```

当开启了混合持久化时，**在 AOF 重写日志时**，`fork` 出来的重写子进程会先将与主线程共享的内存数据以 RDB 方式写入到 AOF 文件，然后主线程处理的操作命令会被记录在重写缓冲区里，重写缓冲区里的增量命令会以 AOF 方式写入到 AOF 文件，写入完成后通知主进程将新的含有 RDB 格式和 AOF 格式的 AOF 文件替换旧的的 AOF 文件。

​	，使用了混合持久化，AOF 文件的**前半部分是 RDB 格式的全量数据，后半部分是 AOF 格式的增量数据**。

![图片](https://img-blog.csdnimg.cn/img_convert/f67379b60d151262753fec3b817b8617.png)

重启 Redis 加载数据的时候，由于前半部分是 RDB 内容，这样**加载的时候速度会很快**。

加载完 RDB 的内容后，才会加载后半部分的 AOF 内容，这里的内容是 Redis 后台子进程重写 AOF 期间，主线程处理的操作命令，可以使得**数据更少的丢失**。

# 过期删除

每当我们对一个 key 设置了过期时间时，Redis 会把该 key 带上过期时间存储到一个**过期字典**（expires dict）中，它是一个哈希表。

常见的三种过期删除策略：

- 定时删除：在设置 key 的过期时间时，同时创建一个定时事件，当时间到达时，由事件处理器自动执行 key 的删除操作。
- 惰性删除：不主动删除过期键，每次从数据库访问 key 时，都检测 key 是否过期，如果过期则删除该 key。
- 定期删除：每隔一段时间「随机」从数据库中取出一定数量的 key 进行检查，并删除其中的过期key。

**Redis 选择「惰性删除+定期删除」这两种策略配和使用**，以求在合理使用 CPU 时间和避免内存浪费之间取得平衡。

> Redis 是怎么实现惰性删除的？

惰性删除的流程图如下：

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/redis/%E8%BF%87%E6%9C%9F%E7%AD%96%E7%95%A5/%E6%83%B0%E6%80%A7%E5%88%A0%E9%99%A4.jpg)

> Redis 是怎么实现定期删除的？

每隔一段时间「随机」从数据库中取出一定数量的 key 进行检查，并删除其中的过期key。

*1、这个间隔检查的时间是多长呢？*

在 Redis 中，默认每秒进行 10 次过期检查一次数据库，此配置可通过 redis.conf 进行配置，配置键为 hz 它的默认值是 hz 10。

*2、随机抽查的数量是多少呢？*

它是写死在代码中的，数值是 20。也就是说，数据库每轮抽查时，会随机选择 20 个 key 判断是否过期。

Redis 的定期删除的流程：

1. 从过期字典中随机抽取 20 个 key；
2. 检查这 20 个 key 是否过期，并删除已过期的 key；
3. 如果本轮检查的已过期 key 的数量，超过 5 个（20/4），也就是「已过期 key 的数量」占比「随机抽取 key 的数量」大于 25%，则继续重复步骤 1；如果已过期的 key 比例小于 25%，则停止继续删除过期 key，然后等待下一轮再检查。

那 Redis 为了保证定期删除不会出现循环过度，导致线程卡死现象，为此增加了定期删除循环流程的时间上限，默认不会超过 25ms。

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

![image-20210503221505925](E:\study\studygo\Golang学习笔记\redis笔记.assets\image-20210503221505925.png)

## 过期后通知

订阅通知的一个典型场景：key过期后通知客户端

```bash
config set notify-keyspace-events Ex	#过期时间监听生效

#以下两种方法二选一
PSUBSCRIBE __keyevent@*__:expired	#订阅一个或者多个符合pattern格式的频道
SUBSCRIBE __keyevent@0__:expired	#0是第几个数据库，如果监听的是第二个数据库就改为1
```

# 主从复制

作用

- 数据冗余：实现了数据的热备份，是持久化之外的一种数据冗余方式
- 故障恢复：主节点有问题时可由从节点提供服务，实现快速的故障恢复，实际上是一种服务的冗余
- 负载均衡：读写分离，提高redis服务器的并发量
- 高可用基石：主从复制是哨兵和集群能够实施的基础

配置

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

## 原理

使用 `replicaof`（Redis 5.0 之前使用 slaveof）命令形成主服务器和从服务器的关系。

比如，现在有服务器 A 和 服务器 B，我们在服务器 B 上执行下面这条命令：

```text
# 服务器 B 执行这条命令
replicaof <服务器 A 的 IP 地址> <服务器 A 的 Redis 端口号>
```

接着，服务器 B 就会变成服务器 A 的「从服务器」，然后与主服务器进行第一次同步。

主从服务器间的第一次同步的过程可分为三个阶段：

- 第一阶段是建立链接、协商同步；
- 第二阶段是主服务器同步数据给从服务器（全量复制）；
- 第三阶段是主服务器发送新写操作命令给从服务器（命令传播）。

![图片](https://img-blog.csdnimg.cn/img_convert/ea4f7e86baf2435af3999e5cd38b6a26.png)

主从复制共有三种模式：**全量复制、基于长连接的命令传播、增量复制**。

主从服务器第一次同步的时候，就是采用全量复制，此时主服务器会两个耗时的地方，分别是生成 RDB 文件和传输 RDB 文件。为了避免过多的从服务器和主服务器进行全量复制，可以把一部分从服务器升级为「经理角色」，让它也有自己的从服务器，通过这样可以分摊主服务器的压力。

> 在「从服务器」上执行下面这条命令，使其作为目标服务器的从服务器：
>
> ```text
> replicaof <目标服务器的IP> 6379
> ```
>
> 此时如果目标服务器本身也是「从服务器」，那么该目标服务器就会成为「经理」的角色，不仅可以接受主服务器同步的数据，也会把数据同步给自己旗下的从服务器，从而减轻主服务器的负担。

第一次同步完成后，主从服务器都会维护着一个长连接，主服务器在接收到写操作命令后，就会通过这个连接将写命令传播给从服务器，来保证主从服务器的数据一致性。

如果遇到网络断开，增量复制就可以上场了，主服务将主从服务器断线期间，所执行的写命令发送给从服务器，然后从服务器执行这些命令：

- **repl_backlog_buffer**，是一个「**环形**」缓冲区，用于主从服务器断连后，从中找到差异的数据；repl_backlog_buffer 大小不能低于 5 MB
- **replication offset**，标记上面那个缓冲区的同步进度，主从服务器都有各自的偏移量，主服务器使用 master_repl_offset 来记录自己「*写*」到的位置，从服务器使用 slave_repl_offset 来记录自己「*读*」到的位置。

# 哨兵模式

> 自动选举主机的模式，哨兵是一个独立的进程。原理是哨兵通过发送命令，等待redis服务器响应，从而监控运行的多个redis实例
>
> 主机宕机后，通过发布订阅模式通知其他服务器，修改配置文件，切换主机；原主机恢复后，只能给新主机当从机
>
> 一个哨兵进程可能会出问题，因此需要多个哨兵进行监控，各哨兵之间也会进行监控

1. 配置哨兵，`sentinel.conf`
   - `sentinel monitor 被监控的主机名称 host port quorum`，quorum代表多少个哨兵认为主机失联时就客观确认主机失联
2. 启动哨兵`redis-sentinel sentinel.conf`

Redis 在 2.8 版本以后提供的**哨兵（*Sentinel*）机制**，它的作用是实现**主从节点故障转移**。它会监测主节点是否存活，如果发现主节点挂了，它就会选举一个从节点切换为主节点，并且把新主节点的相关信息通知给从节点和客户端。

哨兵节点主要负责三件事情：**监控、选主、通知**。

![image-20210504105343034](E:\study\studygo\Golang学习笔记\redis笔记.assets\image-20210504105343034.png)

![image-20210504104856796](E:\study\studygo\Golang学习笔记\redis笔记.assets\image-20210504104856796.png)

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