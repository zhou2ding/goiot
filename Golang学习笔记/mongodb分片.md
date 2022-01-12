[TOC]
# MongoDB - 分片

## 分片

### 分片的目的
* 分片是一种将数据分配到多个机器上的方法。MongoDB通过分片技术来支持具有海量数据集和高吞吐量操作的部署方案。
数据库系统的数据集或应用的吞吐量比较大的情况下，会给单台服务器的处理能力带来极大的挑战。例如，高查询率会耗尽服务器的CPU资源。工作的数据集大于系统的内存会给磁盘驱动器的I/O容量带来压力。
* 解决系统增长的方法有两种：垂直扩展和水平扩展。
  *  垂直扩展 通过增加单个服务器的能力来实现，例如使用更强大的CPU，增加更多的内存或存储空间量。由于现有技术的局限性，不能无限制地增加单个机器的配置。此外，云计算供应商提供可用的硬件配置具有严格的上限。其结果是，垂直扩展有一个实际的最大值。
  *  水平扩展是通过将系统数据集划分至多台机器，并根据需要添加服务器来提升容量。虽然单个机器的总体速度或容量可能不高，但每台机器只需处理整个数据集的某个子集，所以可能会提供比单个高速大容量服务器更高的效率，而且机器的数量只需要根据数据集大小来进行扩展，与单个机器的高端硬件相比，这个方案可以降低总体成本。不过，这种方式会提高基础设施部署维护的复杂性。
* MongoDB通过分片来实现水平扩展。

### 分片涉及到的主要概念
* 片键（shard key）： 文档中的一个或多个字段
* 文档（document）： 包含shard key的一行数据
* 块（chunk）： 包含n个文档 
* 分片（shard）： 包含n个chunk
* 区（zone）： 包含n个分片
* 集群（cluster）： 包含n个分片

### 分片的设计思想和优势
* 分片为应对高吞吐量与大数据量提供了方法。使用分片减少了每个分片需要处理的请求数，因此，通过水平扩展，集群可以提高自己的存储容量和吞吐量。
举例来说，当插入一条数据时，应用只需要访问存储这条数据的分片.使用分片减少了每个分片存储的数据。例如，如果数据库1tb的数据集，并有4个分片，然后每个分片可能仅持有256 GB的数据。如果有40个分片，那么每个切分可能只有25GB的数据。
![image-20220112141936551](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112141936551.png)
* 分片的优势
  * 读写负载
  MongoDB将读写工作负载分布在分片集群中的各个分片上，从而允许每个分片处理集群操作的子集。通过添加更多分片，可以在集群中水平扩展读写工作负载。
  对于包含分片键或复合分片键的前缀的查询，mongos可以将查询定位到特定的分片或一组分片。这些目标操作通常比广播到集群中的每个分片更有效。
  * 存储容量
  通过分片技术将数据分布到分片集群中的各个分片中，每个分片只需存储数据集中的部分子集。随着数据集的增长， 通过增加分片的数量即可增加整个集群的存储容量。
  * 高可用性
  将配置服务器和分片作为副本集进行部署可提高可用性。
  即使一个或多个分片副本集变得完全不可用，分片群集也可以继续提供部分的读取或写入服务。也就是说，虽然停机期间无法访问不可用分片中的数据子集，但是针对可用分片执行读取或写入操作仍然可以成功。

## 分片集群
### 分片集群架构
* 架构示意图
![image-20220112141954611](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112141954611.png)
如上图所示，分片集群由三种组件构成
分片（shard）: 存储应用数据记录。每个shard（分片）包含被分片的数据集中的一个子集。每个分片可以被部署为副本集架构。
配置服务器(config server): config servers存储了分片集群的元数据和配置信息。默认需要部署成副本集(包含3个Config Server节点）。
路由服务器(mongos): 充当查询路由器，在客户端应用程序和分片集群之间提供接口。


### 分片集群组件
#### 分片
##### 基本介绍
* 一个分片包含分片集群中被分片数据的一个子集。集群的分片数据一起保存了集群的整个数据集。
* 从MongoDB3.6开始，分片必须作为副本集部署，以提供冗余和高可用性。
* 用户、客户端或应用程序只能直接连接到shard以执行本地管理和维护操作。
* 对单个shard执行查询只返回数据的子集。连接到mongos以执行集群级操作，包括读或写操作。
* 主分片
  *  分片集群中的每个数据库都有一个主分片，其中包含该数据库的所有未分片集合。每个数据库都有自己的主分片。主分片与副本集中的主节点没有关系。
  *  mongos在创建新数据库时通过在集群中选取数据量最少的shard来选择主shard。mongos使用listDatabases命令返回的totalSize字段作为选择条件的一部分。
  *  要更改数据库的主分片，请使用movePrimary命令。迁移主分片的过程可能需要很长时间才能完成，在迁移完成之前，不应访问与数据库关联的集合。根据迁移的数据量，迁移可能会影响整个群集操作。在尝试更改主分片之前，请考虑对群集操作和网络负载的影响。
  *  当使用以前用作副本集的分片部署新的分片群集时，所有现有数据库将继续驻留在其原始副本集上。随后创建的数据库可以驻留在集群中的任何分片上。
##### 分片策略
MongoDB支持如下两种分片策略来实现分片集群中的分布数据
* 哈希分片
哈希分片涉及计算分片键字段值的哈希值。然后，根据散列的分片键值为每个块分配一个范围。
![image-20220112142011039](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142011039.png)
尽管一系列分片键可能是“接近”的，但它们的哈希值不太可能在同一块上。基于哈希值的数据分发有助于更均匀的数据分发，尤其是在分片键单调更改的数据集中。
但是，哈希分布意味着对分片键的基于范围的查询不太可能针对单个分片，从而导致更多集群范围的广播操作。
* 范围分片
范围分片根据分片键的值将数据划分为多个范围，然后基于分片键的值分配每个块的范围。
![image-20220112142024818](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142024818.png)
值“接近”的一系列分片键更有可能分布在同一块上。好处是便于mongos执行针对性的操作，可以仅将操作路由到包含所需数据的分片上。
范围分片的效率取决于选择的分片键。分片键考虑不周全会导致数据分布不均，这可能会削弱分片的某些优势或导致性能瓶颈。

#### 配置服务器
##### 基本介绍
* 配置服务器存储分片集群的元数据。元数据反映了分片集群中所有数据和组件的状态和组织。元数据包括了每个分片上的chunk列表以及定义chunk的范围。
* 路由器服务器实例缓存这些数据，并且根据这些信息将读和写操作路由到正确的分片。当集群的元数据发生变化时，如chunk分裂或增加分片，mongos会更新缓存。分片也从配置服务器中读取chunk元数据。
* 配置服务器还存储身份验证配置信息，如基于角色的访问控制（RBAC）或集群内部身份验证设置。
* MongoDB还使用配置服务器来管理分布式锁。
* 每个分片集群必须有自己的配置服务器。不同的分片集群不要使用相同的配置服务器。
##### 配置服务器副本集
配置服务器应该部署为副本集，并且该副本集有如下限制
* 不能包含仲裁节点
* 不能包含延时成员
* 必须构建索引
##### 配置服务器的读写
admin数据库和配置数据库存储在配置服务器上。
* 写操作
  *  Admin数据库包含了与认证、授权相关的集合，另外还有一些用于内部使用的system.*集合。
  * 配置数据库中的集合包含了分片集群的元数据。当元数据变化时变化时，mongodb向配置数据库写入数据。
  * 在正常数据库操作和维护的过程中，用户不要直接向配置数据库直接写入数据。
  * 当对config进行写入时，mongodb使用majority级别的写入确认([write concern](https://docs.mongodb.com/manual/reference/write-concern/))
* 读操作
  * MongoDB从admin数据库读认证、授权数据和其它内部使用的数据。
  * 当启动mongos或者元数据改变之后，MongoDB都会从配置数据库读数据。分片也会从配置服务器读取chunk元数据。
  * 读数据的时候，会使用majority级别的[read concern](https://docs.mongodb.com/manual/reference/read-concern/)。
##### 配置服务器的可用性
* 如果配置服务器副本集丢失了主节点并且不能选举出一个新的主节点，则集群的元数据变为只读。你仍然可以从分片中读写数据，但在副本集选出新的主节点之前，将不会发生chunk迁移或chuank分裂。
* 在分片集群中，mongod和mongos实例监视分片集群中的副本集（如分片副本集、配置服务副本集）。
* 如果所有的配置服务器都不可用，则集群也无法运行。为保证配置服务器的可用和完整，配置服务器的备份至关重要。与集群中存储的数据相比，配置服务器上的数据较小，并且配置服务器的活动负载相对较低。

##### 分片集群元数据
* 配置服务器在配置数据库（[Config Database](https://docs.mongodb.com/manual/reference/config-database/)）中存储元数据。
* 为了访问配置数据库，在mongo shell中使用命令：use config
* 通常情况下，不应该对配置数据库中的数据进行直接的更改。配置数据库包含的集合有：changelog、chuaks、collections、databases、lockpings、locks、mongos、settings、shards、version等。

#### 路由服务器
##### 基本介绍
* MongoDB mongos实例将查询操作和写操作路由到分片集群中的分片上。从应用程序的角度，mongos提供分片集群的唯一接口。应用程序从不直接与分片进行连接或通信。
* mongos通过缓存来自配置服务器的元数据来跟踪哪些数据在哪个shard上，并使用元数据将操作从应用程序和客户端路由到mongod实例。mogos没有持久化状态，消费的系统资源最少。
* 最常见的做法是在与应用服务器相同的系统上运行mongos实例，但是您可以在shard或其他专用资源上维护mongos实例。

##### 路由和结果处理
mongos实例通过以下方式将查询路由到集群：
* 确定必须接收查询的分片的列表
* 在所有目标分片上建立游标

mongos然后合并来自每个目标分片的数据并返回结果文档。在mongos检索结果之前，对每个shard执行某些查询修饰符，例如排序。
在版本3.6中更改：对于在多个分片上运行的聚合操作，如果这些操作不需要在数据库的主分片上运行，则这些操作可以将结果路由回mongos，然后在mongos中合并结果。
有两种情况下，管道无法在mongos上运行：
* 第一种情况发生在拆分管道的合并部分包含必须在主分片上运行的阶段时。例如，如\$lookup 需要访问一个数据库上的未分片集合，而在这个数据库上一个分片集合正在运行聚合操作，则合并必须在主分片上运行。
* 第二种情况发生在拆分管道的合并部分包含可能将临时数据写入磁盘的stage时，例如$group，并且客户机指定了字段owDiskUse为true。在这种情况下，假设合并管道中没有其它阶段需要主分片，则合并将在聚合目标分片集中随机选择的分片上运行。

如果要获取有关如何在分片集群查询的组件之间拆分聚合工作的更多信息，可以使用"explain:true"作为aggregate()调用的一个参数。返回将包括三个json对象:
* mergeType显示合并阶段发生的位置（"primaryShard","anyShard","mongos"）。
* splitPipeline显示管道中的哪些操作在单个的分片上运行。
* shards显示每个分片所做的工作。

在某些情况下，当片键(shard key)或片键的的前缀是查询的一部分时，mongos执行一个定向操作，将查询路由到集群中的一个子集。
mongos对不包含片键的查询执行广播操作，将查询路由到集群中的所有分片。一些包含片键的查询仍然可能导致广播操作，这取决于集群中数据的分布和查询的选择性。

##### Mongos如何处理查询修饰符
* Sorting
如果查询的结果没有排序，mongos实例将打开一个结果游标，循环获取分片上所有游标的结果。
* Limits
如果查询使用limit()游标方法限制结果集的大小，mongos实例会将该限制传递给shard，然后在将结果返回给客户端之前将该限制重新应用于结果。
* Skips
如果查询使用skip()游标方法指定要跳过的记录数，mongos将无法将跳过传递给shard，而是从shard中检索未跳过的结果，并在组装完整结果时跳过适当数量的文档。
当与limit()一起使用时，mongos会将limit加上skip()的值传递给shard，以提高这些操作的效率。
##### 读取首选项(Read Preference)和分片
对于分片集群，mongos在读取分片时应用读取首选项。所选成员由读取首选项和replication.localPingThresholdMs配置决定，并为每个操作重新评估。
##### 确认与mongos实例的连接
要检测客户端连接到的MongoDB实例是否是mongos，可以使用isMaster命令。当客户机连接到mongos时，isMaster返回一个包含字符串"isdbgrid"的msg字段的文档。如果应用程序连接到的是mongod，则返回的文档中不包含"isdbgrid"字符串。

##### 定向操作和广播操作
通常，分片环境中最快的查询是mongos使用片键和来自配置服务器的集群元数据路由到单个分片的查询。这些定向操作使用片键值来定位满足查询文档的分片或分片子集。
对于不包含片键的查询，mongos必须查询所有分片、等待它们的响应、然后将结果返回给应用。这些“分散/聚集”查询可能是长时间运行的操作。
* 广播操作
  mongos实例向集合的所有分片广播查询，除非mongos可以确定哪个分片或分片的子集存储这些数据。
  ![image-20220112142040694](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142040694.png)
  当monogos收到来自所有分片的响应后，它将合并数据并返回结果文档。广播操作的性能取决于集群的总体负载、网络延迟、单个分片负载和每个分片返回的文档数量等因素。在可能的情况，支持导致定向操作的操作，而不是导致广播操作的操作。
  multi-update操作总是广播操作。
  updateMany()和deleteMany()方法是广播操作，除非查询文档完全指定了片键。

* 定向操作
  mongos可以将包含片键或复合片键前缀的查询路由到特定的分片或分片集。mongos使用片键来定位其范围包括片键值的chunk，并将查询执行包含该chunk的分片。

  ![image-20220112142148117](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142148117.png)

  例如，如果片键是
```json
{ a: 1, b: 1, c: 1 }
```
mongos程序可以将包含完整片键或下面任一个片键前缀的的查询路由到特定的分片或分片集。
```json
{ a: 1 }
{ a: 1, b: 1 }
```
所有insertOne()操作都定位到一个分片。insertMany()数组中的每个文档都定位到单个分片，但不保证数组中的所有文档都插入到单个分片中
所有updateOne()、replaceOne()和deleteOne()操作都必须在查询文档中包含片键或_id。如果使用这些方法时没有片键或_id，MongoDB将返回错误。
根据集群中数据的分布和查询的选择性，mongos仍然可能执行广播操作来完成这些查询。

##### 索引使用
如果查询不包含片键，mongos必须将查询作为“散布/聚集”操作发送到所有分片。每个分片将依次使用片键索引或另一个更有效的索引来完成查询。
如果查询包含多个子表达式，这些子表达式引用被分片键和辅助索引索引的字段，则mongos可以将查询路由到特定分片，并且分片将允许其使用高效的索引执行。

### 集群的安全机制
mongodb中经常用到的安全机制有两个：
* 基于角色的访问控制（Role-Based Access Control,RBAC）
* 基于keyfiles（秘钥文件）或x.509证书的的内部/成员身份认证

前者用来管理用户对实例的访问，后者用于集群内部各个节点之间相互验证身份。
#### 基于角色的访问控制
* Mongodb RBAC权限管理机制的核心是给每个用户赋予一定的权限，用户连接mongodb前需先验证，验证通过后即拥有用户的权限，权限决定了用户在某一组资源（如某个DB、某个特定集合）上可以执行哪些操作（比如增删改查、建索引）。
* RBAC的详细内容可以查阅[官方文档](https://docs.mongodb.com/manual/core/authorization/)。
* 默认情况下，MongoDB实例启动运行时是没有启用用户访问权限控制的，也就是说，在实例本机服务器上都可以随意登录实例进行各种操作，MongoDB不会对连接客户端进行用户验证，可以想象这是非常危险的。为了强制开启用户访问控制(用户验证)，则需要在MongoDB实例启动时使用选项--auth或在指定启动配置文件中添加选项auth=true。
* 分片集群中的所有节点都要开启用户访问控制。
* 集群中创建的用户分为两类：
  * 分片本地用户（Shard Local Users）
    * 有些维护操作需要直接连接到分片集群中的特定分片。要执行这些操作，必须直接连接到分片并作为分片本地管理用户进行身份认证。
    * 每个分片都有自己的分片本地用户。这些用户不能在其他分片上使用，也不能通过mongos连接到集群。
    * 要创建分片本地用户，直接连接到分片进行创建。MongoDB将分片本地用户存储在分片本身的admin数据库中。
    * 这些分片本地用户完全独立于通过mongos添加到分片集群的用户。分片本地用户是分片的本地用户，mongos无法访问。
    * 与分片的直接相连只能用于特定于分片的维护和配置。一般来说，客户端应该通过mongos连接到分片集群。
  * 分片集群用户（Shard Cluster Users）
    * 要为分片集群创建用户，需连接到mongos实例并添加用户。然后，客户端通过mongos实例对这些用户进行身份验证。在分片集群中，MongoDB将用户配置数据存储在配置服务器的admin数据库中。
    * 每个集群都有自己的集群用户。这些用户不能用于访问单个分片。
> 下文的分片集群搭建部分也会有具体的演示
#### 内部/成员身份认证
* 集群内部各个节点成员之间使用内部身份认证，可以使用秘钥文件或x.509证书。正式环境中应使用x.509证书。
* 内部身份认证的原理是集群中每一个实例彼此连接的时候都检验彼此使用的证书的内容是否相同。只有证书相同的实例彼此才可以访问，这样可以防止非集群节点加入集群。
* 内部身份认证是mongo实例级别的。
> 下文的分片集群搭建部分有内部身份认证的具体使用演示。

### 集群中的数据分布
该部分主要涉及两方面的内容：chunk（块）和zone（区）
#### Chunk
##### Chunk简介
* Chunk是分片内片键值的一段范围。启用分片之后，数据会以chunk为单位（默认64MB，大小是1~1024MB之间）根据片键(shardKey，按照自己指定的字段作为片键)均匀的分散到后端1或多个分片上。
* 每个database会有一个主分片(primary shard)，在数据库创建时分配。database下启用分片（即调用shardCollection命令）的集合，刚开始会生成一个[minKey, maxKey]的chunk，该chunk初始会存储在primary shard上，然后随着数据的写入，不断的发生chunk分裂及迁移。初始的minkey和maxkey是无穷大和无穷小。
* database 下没有启用分片的集合，其所有数据都会存储到primary shard。
![image-20220112142205534](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142205534.png)

##### Chunk的分裂
* mongos上有个sharding.autoSplit 的配置项，可用于控制是否自动触发 chunk分裂，默认是开启的。强烈建议不要关闭autoSplit。
* mongoDB的自动chunk分裂只会发生在mongos写入数据时，当写入的数据超过一定量时，就会触发chunk的分裂（chunk的默认大小是64MB，达到这个阈值就会发生分裂）。
* 分裂不能被取消。
* 如果一个chunk只包含一个分片键值，mongodb就不会分裂这个chunk,即使这个chunk超过了chunk需要分裂时的大小。
比如我们使用日期（精确到日） 作为分片键，当某一天的数据非常多时，这个分片键值（比如2021/04/19）对应的chunk会非常大，超过64M，但是这个chunk是不可分割的。这会造成数据在各个分片中不平衡，出现性能问题。

> 修改Chunk的大小
> 1.连接mongos
> 2.use config
> 3.db.settings.save({_id:“chunksize”,value:64}) //单位是MB

##### Chunk的迁移
* 默认情况下，MongoDB会开启均衡器（balancer），它会周期性的检查分片间是否存在不均衡，如果存在就会开始在各个分片间迁移chunk来让各个分片间负载均衡。用户也可以手动的调用moveChunk命令在分片之间迁移数据。
* Balancer在工作时，会根据shard tag、集合的chunk数量、分片间chunk数量差值来决定是否需要迁移。
  * 根据shard tag迁移
  MongoBD分片支持shard tag特性，用户可以给分片打上标签，然后给集合的某个range打上标签，mongoDB会通过balancer的数据迁移来保证拥有tag的range会分配到具有相同tag 的分片上。
  * 根据分片间chunk数量迁移
  针对所有启用分片的集合，如果 「拥有最多数量 chunk 的 shard」 与 「拥有最少数量 chunk 的 shard」 的差值超过某个阈值，就会触发 chunk 迁移； 有了这个机制，当用户调用 addShard 添加新的分片，或者各个分片上数据写入不均衡时，balancer 就会自动来均衡数据。
  * 还有一种情况会触发迁移，当用户调用 removeShard 命令从集群里移除shard时，Balancer 也会自动将这个分片负责的 chunk 迁移到其他节点。

##### chunkSize 对分裂及迁移的影响
MongoDB 默认的 chunkSize 为64MB，如无特殊需求，建议保持默认值；chunkSize 会直接影响到 chunk 分裂、迁移的行为。
  * chunkSize 越小，chunk 分裂及迁移越多，数据分布越均衡；反之，chunkSize 越大，chunk 分裂及迁移会更少，但可能导致数据分布不均。
  * chunkSize 太小，容易出现 jumbo chunk（即shardKey 的某个取值出现频率很高，这些文档只能放到一个 chunk 里，无法再分裂）而无法迁移；chunkSize 越大，则可能出现 chunk 内文档数太多（chunk 内文档数不能超过 250000 ）而无法迁移。
  * chunk 自动分裂只会在数据写入时触发，所以如果将 chunkSize 改小，系统需要一定的时间来将 chunk 分裂到指定的大小。
  * chunk 只会分裂，不会合并，所以即使将 chunkSize 改大，现有的 chunk 数量不会减少，但 chunk 大小会随着写入不断增长，直到达到目标大小。

总的来说，调节Chunk大小的影响可以用下表描述 

| chuank size条件 | splitting次数（分片数） | 数据跨分片数目 | 数据均匀 | 网络传输次数 | 迁移次数 | 单词迁移传输量 | 查询速度 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 变大 | 减少 | 变少 | 不太均匀 | 变少 | 变少 | 变大 | 变快 |
| 变小 | 增的 | 变多 | 更均匀 | 变多 | 变多 | 变小 | 变慢 |

##### 如何减小分裂及迁移的影响？
mongoDB sharding 运行过程中，自动的 chunk 分裂及迁移如果对服务产生了影响，可以考虑一下如下措施。
* 预分片提前分裂
在使用 shardCollection 对集合进行分片时，如果使用 hash 分片，可以对集合进行「预分片」，直接创建出指定数量的 chunk，并打散分布到后端的各个shard。详见[shardCollection](https://docs.mongodb.com/manual/reference/command/shardCollection/#mongodb-dbcommand-dbcmd.shardCollection)命令
* 合理配置 balancer
  monogDB 的 balancer 能支持非常灵活的配置策略来适应各种需求
  * Balancer 能动态的开启、关闭
  * Blancer 能针对指定的集合来开启、关闭
  * Balancer 支持配置时间窗口，只在制定的时间段内进行迁移

#### Zone
##### Zone的简介
* 一个Zone（分区）是基于特定标记（tags）集的一组分片。
* 每个Zone可以与一个或多个分片关联。
*类似的，分片集群中的每个分片可以与一个或多个Zone相关联。
* MongoDB Zone有助于基于标签跨分片分发Chunk。与Zone内文档相关的所有工作（读写）都在与该Zone匹配的分片上完成。
![image-20220112142224611](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142224611.png)
上图显示了一个有三个分片和两个区的分片集群。A区代表[0,10)，B区代表[10,20).
红色分片和蓝色分片都有A区，而蓝色分片还有B区。绿色分片没有关联分区。
> 不同的Zone不能共享范围，也不能有重叠的范围。
##### 适用场景
* 由于某些硬件配置限制而需要将数据路由到特定分片的情况。
* 如果需要将特定数据隔离到特定的分片，则分区可能很有用。例如，GDPR合规性要求企业保护欧盟内个人的数据和隐私。
* 如果应用程序在不同地理位置上使用，并且您希望查询路由到最近的分片进行读写。

## 分片集群使用指南
### 如何确定shard、mongos数量？
当您决定使用Sharded cluster时，到底应该部署多少个shard、多少个mongos？shard、mongos的数量归根结底是由应用需求决定：
* 如果您使用sharding只是解决海量数据存储问题，访问并不多。假设单个shard能存储M， 需要的存储总量是N，那么您可以按照如下公式来计算实际需要的shard、mongos数量：
  * numberOfShards = N/M/0.75 （假设容量水位线为75%）
  * numberOfMongos = 2+（对访问要求不高，至少部署2个mongos做高可用即可）
* 如果您使用sharding是解决高并发写入（或读取）数据的问题，总的数据量其实很小。您要部署的shard、mongos要满足读写性能需求，容量上则不是考量的重点。假设单个shard最大QPS为M，单个mongos最大QPS为Ms，需要总的QPS为Q。那么您可以按照如下公式来计算实际需要的shard、mongos数量：
  * numberOfShards = Q/M/0.75 （假设负载水位线为75%）
  * numberOfMongos = Q/Ms/0.75
> 说明 mongos、mongod的服务能力，需要用户根据访问特性来实测得出。

如果sharding要同时解决上述2个问题，则按需求更高的指标来预估。以上估算是基于sharded cluster里数据及请求都均匀分布的理想情况。但实际情况下，分布可能并不均衡，为了让系统的负载分布尽量均匀，就需要合理的选择shard key。

### 如何选择shard key？
MongoDB Sharded cluster支持2种分片方式：
* 范围分片，通常能很好的支持基于shard key的范围查询。
* Hash 分片，通常能将写入均衡分布到各个shard。

上述2种分片策略都无法解决以下3个问题：
* shard key取值范围太小（low cardinality），比如将数据中心作为shard key，而数据中心通常不会很多，分片的效果肯定不好。
* shard key某个值的文档特别多，这样导致单个chunk特别大（及 jumbo chunk），会影响chunk迁移及负载均衡。
* 根据非shardkey进行查询、更新操作都会变成scatter-gather查询，影响效率。

好的shard key应该拥有如下特性：
* key分布足够离散（sufficient cardinality）
* 写请求均匀分布（evenly distributed write）
* 尽量避免scatter-gather查询（targeted read）

例如某物联网应用使用MongoDB Sharded cluster存储海量设备的工作日志。假设设备数量在百万级别，设备每10s向 MongoDB汇报一次日志数据，日志包含deviceId，timestamp信息。应用最常见的查询请求是查询某个设备某个时间内的日志信息。以下四个方案中前三个不建议使用，第四个为最优方案，主要是做个对比。
* 方案1： 时间戳作为shard key，范围分片：
  * Bad。
  * 新的写入都是连续的时间戳，都会请求到同一个shard，写分布不均。
  * 根据deviceId的查询会分散到所有shard上查询，效率低。
* 方案2： 时间戳作为shard key，hash分片：
  * Bad。
  * 写入能均分到多个shard。
  * 根据deviceId的查询会分散到所有shard上查询，效率低。
* 方案3：deviceId作为shardKey，hash分片（如果ID没有明显的规则，范围分片也一样）：
  * Bad。
  * 写入能均分到多个shard。
  * 同一个deviceId对应的数据无法进一步细分，只能分散到同一个chunk，会造成jumbo chunk，根据deviceId的查询只请求到单个shard。不足的是，请求路由到单个shard后，根据时间戳的范围查询需要全表扫描并排序。
* 方案4：（deviceId，时间戳）组合起来作为shardKey，范围分片（Better）：
  * Good。
  * 写入能均分到多个shard。
  * 同一个deviceId的数据能根据时间戳进一步分散到多个chunk。
  * 根据deviceId查询时间范围的数据，能直接利用（deviceId，时间戳）复合索引来完成。

### 关于jumbo chunk及chunk size
MongoDB默认的chunk size为64MB，如果chunk超过64MB且不能分裂（比如所有文档的shard key都相同），则会被标记为jumbo chunk ，balancer不会迁移这样的chunk，从而可能导致负载不均衡，应尽量避免。

一旦出现了jumbo chunk，如果对负载均衡要求不高，并不会影响到数据的读写访问。如果一定要处理，可以尝试如下方法：

* 对jumbo chunk进行split，一旦split成功，mongos会自动清除jumbo标记。
* 对于不可再分的chunk，如果该chunk已不是jumbo chunk，可以尝试手动清除chunk的jumbo标记（注意先备份下config数据库，以免误操作导致config库损坏）。
* 调大chunk size，当chunk大小不超过chunk size时，jumbo标记最终会被清理。但是随着数据的写入仍然会再出现 jumbo chunk，根本的解决办法还是合理的规划shard key。

关于chunk size如何设置，绝大部分情况下可以直接使用默认的chunk size ，以下场景可能需要调整chunk size（取值在1-1024之间）：

* 迁移时IO负载太大，可以尝试设置更小的chunk size。
* 测试时，为了方便验证效果，设置较小的chunk size。
* 初始chunk size设置不合理，导致出现大量jumbo chunk影响负载均衡，此时可以尝试调大chunk size。
* 将未分片的集合转换为分片集合，如果集合容量太大，需要（数据量达到T级别才有可能遇到）调大chunk size才能转换成功。具体方法请参见[Sharding Existing Collection Data Size](https://docs.mongodb.com/manual/core/sharded-cluster-requirements/?spm=a2c63.p38356.879954.3.c63131f6EhmCpN)。

### Tag aware sharding
> 即前文介绍的Zone（分区）概念

Tag aware sharding是Sharded cluster很有用的一个特性，允许用户自定义一些chunk的分布规则。Tag aware sharding原理如下：
* sh.addShardTag()给shard设置标签A。
* sh.addTagRange()给集合的某个chunk范围设置标签A，最终MongoDB会保证设置标签A的chunk范围（或该范围的超集）分布设置了标签A的shard上。

Tag aware sharding可应用在如下场景
* 将部署在不同机房的shard设置机房标签，将不同chunk范围的数据分布到指定的机房。
* 将服务能力不同的shard设置服务等级标签，将更多的chunk分散到服务能力更强的shard上去。

使用Tag aware sharding需要注意：
chunk分配到对应标签的shard上无法立即完成，而是在不断insert、update后触发split、moveChunk后逐步完成的并且需要保证balancer是开启的。在设置了tag range一段时间后，写入仍然没有分布到tag相同的shard上去。

### 关于负载均衡
MongoDB Sharded cluster的自动负载均衡目前是由mongos的后台线程来做，并且每个集合同一时刻只能有一个迁移任务。负载均衡主要根据集合在各个shard上chunk的数量来决定的，相差超过一定阈值（跟chunk总数量相关）就会触发chunk迁移。

负载均衡默认是开启的，为了避免chunk迁移影响到线上业务，可以通过设置迁移执行窗口，比如只允许凌晨2:00-6:00期间进行迁移。
```
use config
db.settings.update(
{ _id: "balancer" },
{ $set: { activeWindow : { start : "02:00", stop : "06:00" } } },
{ upsert: true }
)
            
```
> 注意：在进行sharding备份时（通过mongos或者单独备份config server和所有shard），需要停止负载均衡，以免备份出来的数据出现状态不一致问题。
```
sh.stopBalancer()
```

### 分片前的注意事项
鉴于分片集群基础架构的要求和复杂性，您需要仔细地制定计划、执行和维护。
分片后，MongoDB不会提供任何方法来对分片集群进行分片。
为了确保集群的性能和效率，在选择分片键时需要仔细考虑。请参阅[选择分片键](https://docs.mongodb.com/manual/core/sharding-shard-key/#sharding-shard-key-selection)及上文“如何选择shard key？”的描述。
分片有一定的操作要求和限制。有关更多信息，请参见[分片集群中的操作限制](https://docs.mongodb.com/manual/core/sharded-cluster-requirements/)。
如果查询条件中没有包含拆分键或[复合](https://docs.mongodb.com/manual/reference/glossary/#std-term-compound-index)拆分键的前缀，mongos节点将执行[广播](https://docs.mongodb.com/manual/core/sharded-cluster-query-router/#std-label-sharding-mongos-broadcast)操作，即查询分片集群中的所有分片。这些分散/聚集查询可能会变成长耗时的操作。

## 分片集群搭建实践
### 最小集群
#### 集群组成
搭建一个由
* 1 个mongos
* 1 个config server （因要求必须是副本集，所以使用单节点副本集）
* 3 个shard （mongod单节点而不是副本集）

组成的分片集群。
> 这里只是用于演示搭建一个最简单的分片集群，所以力求简单。而实际环境中不能这样做!! 切记！
> 另外，因只是一个示例，所以分片集群的所有的节点都在一台电脑上进行。
#### 集群搭建步骤
1. 目录结构
![image-20220112142249706](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142249706.png)
```
PS D:\mongoworkdir\sharding_test> tree /F
.
├─configsvr
│  │  configsvr.conf
│  │
│  ├─data
│  └─logs
├─router
│  │  router.conf
│  │
│  └─logs
├─shard1
│  │  shard1.conf
│  │
│  ├─data
│  └─logs
├─shard2
│  │  shard2.conf
│  │
│  ├─data
│  └─logs
└─shard3
    │  shard3.conf
    │
    ├─data
    └─logs
```
其中以.conf为后缀的的配置文件，其它都是文件夹。各个配置文件内容列举如下：
* configsvr\configsvr.conf内容如下：
```
systemLog:
    destination: file
    path: D:\mongoworkdir\sharding_test\configsvr\logs\configsvr.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\sharding_test\configsvr\data
replication:
    replSetName: cfgsvr001	
sharding:
    clusterRole: configsvr
net:
    bindIp: 127.0.0.1
    port: 27300
```

* router\router.conf内容如下：
```
systemLog:
   destination: file
   path: D:\mongoworkdir\sharding_test\router\logs\router.log
   logAppend: true
sharding:
  configDB: cfgsvr001/127.0.0.1:27300
net:
  bindIp: 127.0.0.1
  port: 27400
```

* shard1\shard1.conf内容如下：
```
systemLog:
    destination: file
    path: D:\mongoworkdir\sharding_test\shard1\logs\shard1.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\sharding_test\shard1\data
    directoryPerDB: true
sharding:
    clusterRole: shardsvr
net:
    bindIp: 127.0.0.1
    port: 27201
```
* shard2\shard2.conf内容如下：
```
systemLog:
   destination: file
   path: D:\mongoworkdir\sharding_test\shard2\logs\shard2.log
   logAppend: true
storage:
   dbPath: D:\mongoworkdir\sharding_test\shard2\data
   directoryPerDB: true
sharding:
  clusterRole: shardsvr
net:
  bindIp: 127.0.0.1
  port: 27202
```
* shard3\shard3.conf内容如下：
```
systemLog:
   destination: file
   path: D:\mongoworkdir\sharding_test\shard3\logs\shard3.log
   logAppend: true
storage:
   dbPath: D:\mongoworkdir\sharding_test\shard3\data
   directoryPerDB: true
sharding:
  clusterRole: shardsvr
net:
  bindIp: 127.0.0.1
  port: 27203
```
> 注意，上面的配置文件是都是一些基本的配置项。MongoDB提供了很多的配置选择，具体可以参考官方文档。

2. 创建config server复制集

* mongos配置中sharding.configDB指定“The configuration servers for the sharded cluster.”，格式如下：
```
sharding:
  configDB: <configReplSetName>/cfg1.example.net:27019, cfg2.example.net:27019,...
```

试验过程中发现MongoDB server version: 4.0.11版本的sharding.configDB只支持复制集，不支持单个节点。

根据官方说明：
For a production deployment, deploy a config server replica set with at least three members. For testing purposes, you can create a single-member replica set.
即在测试中可以创建单节点复制集，所以为了继续搭建，创建单节点的config server复制集。

* 创建单节点的config server复制集

启动config server节点
```
PS D:\mongoworkdir\sharding_test\configsvr> mongod -f .\configsvr.conf
```
连接config server节点后，按如下方式创建
```
> use admin
> cfg = {_id: "cfgsvr001",members:[{_id: 0,host: '127.0.0.1:27300'}]};
> rs.initiate(cfg)    #使配置生效
或
> rs.initiate({_id: "cfgsvr001",members:[{_id: 0,host: '127.0.0.1:27300'}]})
```
其中27300是config server的端口。
操作日志：
```
PS D:\mongoworkdir\sharding_test> mongo --port 27300
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27300/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("764bd23f-bb0b-4e6f-ba8f-db552bb4191b") }
MongoDB server version: 4.0.11
Server has startup warnings:
2021-04-20T05:36:55.071-0700 I CONTROL  [initandlisten]
2021-04-20T05:36:55.071-0700 I CONTROL  [initandlisten] ** WARNING: Access control is not enabled for the database.
2021-04-20T05:36:55.071-0700 I CONTROL  [initandlisten] **          Read and write access to data and configuration is unrestricted.
... ...

> use admin
switched to db admin
> rs.initiate({_id: "cfgsvr001",members:[{_id: 0,host: '127.0.0.1:27300'}]})
{
        "ok" : 1,
        ... ...
}
cfgsvr001:SECONDARY>
cfgsvr001:PRIMARY>
cfgsvr001:PRIMARY> rs.status()
{
        "set" : "cfgsvr001",
        ... ...
        "members" : [
                {
                        "_id" : 0,
                        "name" : "127.0.0.1:27300",
                        "health" : 1,
                        "state" : 1,
                        ... ...
                }
        ],
        "ok" : 1,
        ... ...
}
cfgsvr001:PRIMARY>
```

3. 启动各节点，创建超管账号
* 启动各个节点
```
//启动分片节点
mongod.exe -f .\shard1.conf
mongod.exe -f .\shard2.conf
mongod.exe -f .\shard3.conf

//启动mongos节点
mongos -f .\router.conf
```
其中，启动mongos节点时，会有如下警告信息
```
2021-04-20T20:46:02.760+0800 W SHARDING [main] Running a sharded cluster with fewer than 3 config servers should only be done for testing purposes and is not recommended for production.
```
"Running a sharded cluster with fewer than 3 config servers should only be done for testing purposes and is not recommended for production."
这是因为config server复制集只有一个单节点，所以这里警告说生产环境下这样做是不推荐的。

* 分别连接config server节点、所有shard节点，创建节点超管账号(Shard Local Users)
```
//创建管理员账户
use admin
db.createUser({user:"root",pwd:"root",roles:["root"]})
db.createUser({user:"admin",pwd:"admin",roles:[{ role: "userAdminAnyDatabase", db: "admin" }]})
```

* 连接mongos节点，创建集群超管账号（Shard Cluster Users）
```
//创建管理员账户
use admin
db.createUser({user:"cluserroot",pwd:"cluserroot",roles:["root"]})
db.createUser({user:"cluseradmin",pwd:"cluseradmin",roles:[{ role: "userAdminAnyDatabase", db: "admin" }]})
```

4. 增加认证、鉴权配置；
* 首先，关闭各节点。
* 增加认证、鉴权配置主要两点：一是key文件；另一个是配置中增加security项。
  * 生成mongo.keyfile，并拷贝到各节点目录下
    ```
    PS D:\mongoworkdir\sharding_test> tree /F
    .
    ├─configsvr
    │  │  configsvr.conf
    │  │  mongo.keyfile
    │  │
    │  ├─data
    │  │  ...
    │  │
    │  └─logs
    │          configsvr.log
    │
    ├─router
    │  │  mongo.keyfile
    │  │  router.conf
    │  │
    │  └─logs
    │      │  router.log
    |         ...
    │
    ├─shard1
    │  │  mongo.keyfile
    │  │  shard1.conf
    │  │
    │  ├─data
    │  │  ...
    │  │
    │  └─logs
    │          shard1.log
    │
    ├─shard2
    │  │  mongo.keyfile
    │  │  shard2.conf
    │  │
    │  ├─data
    │  │  ...
    │  │
    │  └─logs
    │          shard2.log
    │
    └─shard3
        │  mongo.keyfile
        │  shard3.conf
        │
        ├─data
        │  ...
        │
        └─logs
                shard3.log
    ```
  * 在各节点配置中增加security项
    configsvr.conf：
    ```
    systemLog:
        destination: file
        path: D:\mongoworkdir\sharding_test\configsvr\logs\configsvr.log
        logAppend: true
    storage:
        dbPath: D:\mongoworkdir\sharding_test\configsvr\data
    replication:
        replSetName: cfgsvr001	
    sharding:
        clusterRole: configsvr
    net:
        bindIp: 127.0.0.1
        port: 27300
    security:
        keyFile: D:\mongoworkdir\sharding_test\configsvr\mongo.keyfile
        authorization: "enabled"
    ```
    在shard1.conf、shard2.conf、shard3.conf配置文件中也按上述方式增加,注意keyFile填写正确路径
    ```
    security:
        keyFile: path\to\the\mongo.keyfile
        authorization: "enabled"
    ```
    而*router.conf*比较特殊，增加如下：
    ```
    security:
        keyFile: D:\mongoworkdir\sharding_test\router\mongo.keyfile
    ```
    这是因为，mongos数据路由的配置中security不能包含authorization字段，官方文档中有说明：“The security.authorization setting is available only for mongod.”

* 认证、鉴权配置完成后，启动各节点。

5. 向集群添加分片
* 登录到mongos（27400），添加Shard节点（Add Shards to the Cluster）
```
PS D:\mongoworkdir\sharding_test> mongo.exe --port 27400
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27400/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("1462a2aa-57a8-4f06-b3c7-ba8bd72be6a3") }
MongoDB server version: 4.0.11
mongos> use admin
switched to db admin
mongos> db.auth("cluseradmin","cluseradmin")
1
mongos>
mongos> db.runCommand({ addshard:"127.0.0.1:27201" })
{
        "ok" : 0,
        "errmsg" : "not authorized on admin to execute command { addshard: \"127.0.0.1:27201\", lsid: { id: UUID(\"1462a2aa-57a8-4f06-b3c7-ba8bd72be6a3\") }, $clusterTime: { clusterTime: Timestamp(1618925900, 1), signature: { hash: BinData(0, C72DFCC97BF2415890C3F2F4196DEBA13E257F5D), keyId: 6953218371919806482 } }, $db: \"admin\" }",
        "code" : 13,
        "codeName" : "Unauthorized",
        "operationTime" : Timestamp(1618925960, 1),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618925960, 1),
                "signature" : {
                        "hash" : BinData(0,"CYkB8/+7eIqq6uQ2FFB5aWl3YcI="),
                        "keyId" : NumberLong("6953218371919806482")
                }
        }
}
mongos>
```
注意，提示没有权限，所以要先增加权限
```
db.grantRolesToUser( "cluseradmin" , [{ "role": "clusterAdmin", "db": "admin" }])
```
然后才有添加shard节点的权限。
操作日志：
```
mongos> db.grantRolesToUser( "cluseradmin" , [{ "role": "clusterAdmin", "db": "admin" }])
mongos>
mongos>
mongos> db.runCommand({ addshard:"127.0.0.1:27201" })
{
        "shardAdded" : "shard0000",
        "ok" : 1,
        "operationTime" : Timestamp(1618926116, 4),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618926116, 4),
                "signature" : {
                        "hash" : BinData(0,"Hf1mlUYAurTss8+085LHN0vsVoA="),
                        "keyId" : NumberLong("6953218371919806482")
                }
        }
}
mongos> db.runCommand({ addshard:"127.0.0.1:27202" })
{
        "shardAdded" : "shard0001",
        "ok" : 1,
        "operationTime" : Timestamp(1618926128, 2),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618926128, 2),
                "signature" : {
                        "hash" : BinData(0,"WHp59Zvc3Q9vlbq6GUAfa+tHH40="),
                        "keyId" : NumberLong("6953218371919806482")
                }
        }
}
mongos> db.runCommand({ addshard:"127.0.0.1:27203" })
{
        "shardAdded" : "shard0002",
        "ok" : 1,
        "operationTime" : Timestamp(1618926133, 2),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618926133, 2),
                "signature" : {
                        "hash" : BinData(0,"2fg/T0chwtbk57xMcJVVzpW/Awk="),
                        "keyId" : NumberLong("6953218371919806482")
                }
        }
}
mongos>
```
查看列表，如果是3个分片。就OK了。
```
mongos> db.runCommand({listshards:1})
{
        "shards" : [
                {
                        "_id" : "shard0000",
                        "host" : "127.0.0.1:27201",
                        "state" : 1
                },
                {
                        "_id" : "shard0001",
                        "host" : "127.0.0.1:27202",
                        "state" : 1
                },
                {
                        "_id" : "shard0002",
                        "host" : "127.0.0.1:27203",
                        "state" : 1
                }
        ],
        "ok" : 1,
        "operationTime" : Timestamp(1618926146, 1),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618926146, 1),
                "signature" : {
                        "hash" : BinData(0,"skRfpR6zgyqt7En9Ah8B53qzcEg="),
                        "keyId" : NumberLong("6953218371919806482")
                }
        }
}
mongos>
```
向集群添加分片成功。

6. 打开数据库分片使能，分片集合
* 允许数据库分片（Enable Sharding for a Database）
Before you can shard a collection, you must enable sharding for the collection's database. Enabling sharding for a database does not redistribute data but make it possible to shard the collections in that database.
From the mongo shell connected to the mongos, use the sh.enableSharding() method to enable sharding on the target database. Enabling sharding on a database makes it possible to shard collections within a database.

执行如下命令(假设要使能分片存储的数据为 shardtest ):
```
mongos> db.runCommand({ enablesharding:"shardtest" }) #设置分片存储的数据库
```
操作日志：
```
mongos>
mongos> db.runCommand({ enablesharding:"shardtest" })
{
        "ok" : 1,
        "operationTime" : Timestamp(1618926506, 2),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618926506, 2),
                "signature" : {
                        "hash" : BinData(0,"ZixzoU65rKK4h/FU+qpjKRwh2PM="),
                        "keyId" : NumberLong("6953218371919806482")
                }
        }
}
mongos>
```

* 分片集合（Shard a Collection）
Before you can shard a collection you must first enable sharding for the database where the collection resides.
To shard a collection, connect to the mongos from the mongo shell and use the sh.shardCollection() method.

执行如下命令(假设要分片的集合为 user )：
```
mongos> db.runCommand({ shardcollection: 'shardtest.user', key: {name: 1}})
```
操作日志：
```
mongos> db.runCommand({ shardcollection: 'shardtest.user', key: {name: 1}})
{
        "collectionsharded" : "shardtest.user",
        "collectionUUID" : UUID("c5153b2c-430d-40dd-934c-383c75106513"),
        "ok" : 1,
        "operationTime" : Timestamp(1618926750, 8),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618926750, 8),
                "signature" : {
                        "hash" : BinData(0,"Bb06sDB+fF9XzKBGSdahcQn3mAI="),
                        "keyId" : NumberLong("6953218371919806482")
                }
        }
}
mongos>
```
* 使用测试
  * 创建普通用户
  ```
  # 创建普通用户
  use shardtest
  db.createUser({user:"shardUser",pwd:"shardUser", roles:[{ role: "readWrite", db: "shardtest" }]})
  ```
  * 插入记录
  ```
  PS D:\mongoworkdir\sharding_test> mongo.exe --port 27400
  MongoDB shell version v4.0.11
  connecting to: mongodb://127.0.0.1:27400/?gssapiServiceName=mongodb
  Implicit session: session { "id" : UUID("5a06cdd1-7aff-481b-b96d-f0b89327d90c") }
  MongoDB server version: 4.0.11
  mongos>
  mongos> use shardtest
  switched to db shardtest
  mongos> db.auth("shardUser","shardUser")
  1
  mongos> db.user.insert({name:"Jack", age: 19, gender: "male", telno: 1233, addr:"Hangzhou"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Lily", age: 20, gender: "female", telno: 1234, addr:"Hangzhou"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Lucy", age: 20, gender: "female", telno: 1234, addr:"Hangzhou"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Mary", age: 20, gender: "female", telno: 1234, addr:"Hangzhou"})
  WriteResult({ "nInserted" : 1 })
  mongos>  db.user.insert({name:"HanMeimei", age: 20, gender: "female", telno: 1234, addr:"Hangzhou"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Tom", age: 21, gender: "male", telno: 1234, addr:"Shanghai"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Peter", age: 30, gender: "male", telno: 1234, addr:"Hunan"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Daniel", age: 29, gender: "male", telno: 1234, addr:"Shanddong"})
  WriteResult({ "nInserted" : 1 })
  mongos>  db.user.insert({name:"Martin", age: 30, gender: "male", telno: 1234, addr:"Anhui"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Clerk", age: 20, gender: "male", telno: 1234, addr:"Fujian"})
  WriteResult({ "nInserted" : 1 })
  mongos> db.user.insert({name:"Luice", age: 25, gender: "male", telno: 1234, addr:"Fuyang"})
  WriteResult({ "nInserted" : 1 })
  mongos>
  mongos> db.user.find()
  { "_id" : ObjectId("607ede0cfef5b4377e917440"), "name" : "Jack", "age" : 19, "gender" : "male", "telno" : 1233, "addr" : "Hangzhou" }
  { "_id" : ObjectId("607ede13fef5b4377e917441"), "name" : "Lily", "age" : 20, "gender" : "female", "telno" : 1234, "addr" : "Hangzhou" }
  { "_id" : ObjectId("607ede19fef5b4377e917442"), "name" : "Lucy", "age" : 20, "gender" : "female", "telno" : 1234, "addr" : "Hangzhou" }
  { "_id" : ObjectId("607ede1ffef5b4377e917443"), "name" : "Mary", "age" : 20, "gender" : "female", "telno" : 1234, "addr" : "Hangzhou" }
  { "_id" : ObjectId("607ede26fef5b4377e917444"), "name" : "HanMeimei", "age" : 20, "gender" : "female", "telno" : 1234, "addr" : "Hangzhou" }
  { "_id" : ObjectId("607ede2dfef5b4377e917445"), "name" : "Tom", "age" : 21, "gender" : "male", "telno" : 1234, "addr" : "Shanghai" }
  { "_id" : ObjectId("607ede33fef5b4377e917446"), "name" : "Peter", "age" : 30, "gender" : "male", "telno" : 1234, "addr" : "Hunan" }
  { "_id" : ObjectId("607ede39fef5b4377e917447"), "name" : "Daniel", "age" : 29, "gender" : "male", "telno" : 1234, "addr" : "Shanddong" }
  { "_id" : ObjectId("607ede3ffef5b4377e917448"), "name" : "Martin", "age" : 30, "gender" : "male", "telno" : 1234, "addr" : "Anhui" }
  { "_id" : ObjectId("607ede45fef5b4377e917449"), "name" : "Clerk", "age" : 20, "gender" : "male", "telno" : 1234, "addr" : "Fujian" }
  { "_id" : ObjectId("607ede4afef5b4377e91744a"), "name" : "Luice", "age" : 25, "gender" : "male", "telno" : 1234, "addr" : "Fuyang" }
  mongos>
  ```
  * 查看集合分片存储状态
  ```
  mongos> db.user.stats()
  {
          "sharded" : true,
          "capped" : false,
          ... ...
          "ns" : "shardtest.user",
          "count" : 11,
          "size" : 1121,
          "storageSize" : 32768,
          "totalIndexSize" : 65536,
          "indexSizes" : {
                  "_id_" : 32768,
                  "name_1" : 32768
          },
          "avgObjSize" : 101,
          "maxSize" : NumberLong(0),
          "nindexes" : 2,
          "nchunks" : 1,
          "shards" : {
                  "shard0001" : {
                          "ns" : "shardtest.user",
                          "size" : 1121,
                          "count" : 11,
                          "avgObjSize" : 101,
                          "storageSize" : 32768,
                          "capped" : false,
                          ... ...
                          "nindexes" : 2,
                          "totalIndexSize" : 65536,
                          "indexSizes" : {
                                  "_id_" : 32768,
                                  "name_1" : 32768
                          },
                          "ok" : 1,
                          "$configServerState" : {
                                  "opTime" : {
                                          "ts" : Timestamp(1618927270, 1),
                                          "t" : NumberLong(2)
                                  }
                          }
                  }
          },
          "ok" : 1,
          ... ...
  }
  mongos>
  ```
  可见，插入的11个文档全部存储在了shard0001分片上。这可能是片键选择不合理所致。但已经说明分片集群已经能正常工作，下面要对分片功能做进一步的测试。

#### 分片功能测试
上一章节已经完成了集群搭建并初步验证了通过mongos能正常存取数据，但还不足以验证分片功能，接下来进行进一步的验证。
* 测试方法
通过向分片集群写入10000条测试数据，采用hash分片，理论上这些数据会比较均匀的分布在三个分片上。

* 测试步骤
  * 开启集合分片
  ```
  db.runCommand({ shardcollection: 'shardtest.students', key:  { "_id": "hashed" }})
  ```
  操作日志：
  ```
  PS D:\mongoworkdir\sharding_test> mongo.exe --port 27400
  MongoDB shell version v4.0.11
  connecting to: mongodb://127.0.0.1:27400/?gssapiServiceName=mongodb
  Implicit session: session { "id" : UUID("f30a5b4c-b888-4c59-b2fb-463e96061bf9") }
  MongoDB server version: 4.0.11
  mongos> use admin
  switched to db admin
  mongos> db.auth("cluseradmin","cluseradmin")
  1
  mongos> db.runCommand({listshards:1})
  {
          "shards" : [
                  {
                          "_id" : "shard0000",
                          "host" : "127.0.0.1:27201",
                          "state" : 1
                  },
                  {
                          "_id" : "shard0001",
                          "host" : "127.0.0.1:27202",
                          "state" : 1
                  },
                  {
                          "_id" : "shard0002",
                          "host" : "127.0.0.1:27203",
                          "state" : 1
                  }
          ],
          "ok" : 1,
          "operationTime" : Timestamp(1618967751, 1),
          "$clusterTime" : {
                  "clusterTime" : Timestamp(1618967751, 1),
                  "signature" : {
                          "hash" : BinData(0,"DHl/UyUUkY44aRTW11/88R2Jq7c="),
                          "keyId" : NumberLong("6953218371919806482")
                  }
          }
  }
  mongos> db.runCommand({ shardcollection: 'shardtest.students', key:  { "_id": "hashed" }})
  {
          "collectionsharded" : "shardtest.students",
          "collectionUUID" : UUID("b6276a79-a024-4867-9412-ccc66c000515"),
          "ok" : 1,
          "operationTime" : Timestamp(1618967766, 3),
          "$clusterTime" : {
                  "clusterTime" : Timestamp(1618967766, 3),
                  "signature" : {
                          "hash" : BinData(0,"tAy5/mD1lsDYhkHWl17dXlLm1Jk="),
                          "keyId" : NumberLong("6953218371919806482")
                  }
          }
  }
  mongos>
  ```
  * 编写一个MongoDB客户端（基于golang）

  向已开启分片的students集合写入10000条记录，测试代码如下：
  ```
  package main
  
  import (
    "context"
    "errors"
    "flag"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
  )
  
  type student struct {
    ID  int    `bson:"_id"`
    Name string
    Gender string
    SNO int
  }
  
  var (
    host	string
    port	string
    uname	string
    pwd		string
    dbname	string	//数据库名
  )
  
  func parseStartParams() {
    flag.StringVar(&host, "h", "127.0.0.1", "server ip")
    flag.StringVar(&port, "P", "27017", "server port")
    flag.StringVar(&uname, "u", "isv", "database user")
    flag.StringVar(&pwd, "p", "isv", "user's password")
    flag.StringVar(&dbname, "d", "isv", "database name")
  
    flag.Parse()
  }
  
  func main() {
    parseStartParams()
  
    client, err := connectMongo()
    if err != nil {
      return
    }
    defer client.Disconnect(context.TODO())
  
    db, err := getDatabase(client, dbname)
    if err != nil {
      return
    }
  
    colName := "students"
    col := getCollection(db, colName)
  
    //插入测试数据
    for i := 0; i < 10000; i++ {
      s := student{
        ID: 100000 + i,
        Name: fmt.Sprintf("stu%05d", i),
        SNO: i,
      }
      if i%2 == 0 {
        s.Gender = "male"
      } else {
        s.Gender = "female"
      }
  
      _, err := col.InsertOne(context.TODO(), &s)
      if err != nil {
        fmt.Println("InsertOne error: ", err)
        break
      }
    }
  }
  
  func getCollection(db *mongo.Database, colNam string) (*mongo.Collection) {
    return db.Collection(colNam)
  }
  
  func getDatabase(client *mongo.Client, dbnam string) (*mongo.Database, error) {
    //查看要操作的数据库是否存在
    dbs, err := client.ListDatabaseNames(context.TODO(), bson.D{})
    if err != nil {
      fmt.Println("ListDatabaseNames error:", err)
      return nil, err
    }
    //fmt.Println("dbs: ", dbs)
  
    var exist bool
    for _, db := range dbs {
      if db == dbnam {
        exist = true
        break
      }
    }
    if !exist {
      fmt.Printf("database %s not exist!\n", dbnam)
      return nil, errors.New("database not exist")
    }
  
    return client.Database(dbnam), nil
  }
  
  func connectMongo() (*mongo.Client, error) {
    host := fmt.Sprintf("mongodb://%s:%s@%s:%s/shardtest?retryWrites=false", uname, pwd, host, port)
  
    // 设置客户端连接配置
    clientOptions := options.Client().ApplyURI(host)
  
    // 连接到MongoDB
    var err error
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
      fmt.Println("mongo.Connect error:", err)
      return nil, err
    }
  
    // 检查连接
    err = client.Ping(context.TODO(), nil)
    if err != nil {
      fmt.Println("mongo client.ping error:", err)
      return nil, err
    }
  
    fmt.Println("Connected to MongoDB!\n")
  
    return client, nil
  }
  ```
  编译、运行（连接的是mongos的端口27400）:
  ```
  main.exe -P 27400 -u shardUser -p shardUser -d shardtest
  
  ```

  * 插入10000调记录后，查看分片情况

  可以使用db.stats()或db.students.stats()，前者查看整个数据库的分片情况，后者查看具体集合的分片情况
  操作记录：
  ```
  PS D:\mongoworkdir\sharding_test> mongo.exe --port 27400
  MongoDB shell version v4.0.11
  connecting to: mongodb://127.0.0.1:27400/?gssapiServiceName=mongodb
  Implicit session: session { "id" : UUID("508d19fa-1e80-4fe8-8484-0120aa614f54") }
  MongoDB server version: 4.0.11
  mongos> use shardtest
  switched to db shardtest
  mongos> db.auth("shardUser","shardUser")
  1
  mongos> show collections
  students
  user
  mongos> db.user.drop()  //删除之前创建的user集合后，数据库只有students一个集合
  true
  mongos>
  mongos> db.stats()
  {
          "raw" : {
                  "127.0.0.1:27203" : {
                          "db" : "shardtest",
                          "collections" : 1,
                          "views" : 0,
                          "objects" : 3355,
                          "avgObjSize" : 60.00327868852459,
                          "dataSize" : 201311,
                          "storageSize" : 77824,
                          "numExtents" : 0,
                          "indexes" : 2,
                          "indexSize" : 122880,
                          "fsUsedSize" : 70920318976,
                          "fsTotalSize" : 500364734464,
                          "ok" : 1
                  },
                  "127.0.0.1:27201" : {
                          "db" : "shardtest",
                          "collections" : 1,
                          "views" : 0,
                          "objects" : 3368,
                          "avgObjSize" : 59.98159144893112,
                          "dataSize" : 202018,
                          "storageSize" : 102400,
                          "numExtents" : 0,
                          "indexes" : 2,
                          "indexSize" : 208896,
                          "fsUsedSize" : 70920318976,
                          "fsTotalSize" : 500364734464,
                          "ok" : 1
                  },
                  "127.0.0.1:27202" : {
                          "db" : "shardtest",
                          "collections" : 1,
                          "views" : 0,
                          "objects" : 3277,
                          "avgObjSize" : 60.015563014952704,
                          "dataSize" : 196671,
                          "storageSize" : 106496,
                          "numExtents" : 0,
                          "indexes" : 2,
                          "indexSize" : 208896,
                          "fsUsedSize" : 70920318976,
                          "fsTotalSize" : 500364734464,
                          "ok" : 1
                  }
          },
          "objects" : 10000,
          "avgObjSize" : 59.6632,
          "dataSize" : 600000,
          "storageSize" : 286720,
          "numExtents" : 0,
          "indexes" : 6,
          "indexSize" : 540672,
          "fileSize" : 0,
          "extentFreeList" : {
                  "num" : 0,
                  "totalSize" : 0
          },
          "ok" : 1,
          "operationTime" : Timestamp(1618969341, 1),
          "$clusterTime" : {
                  "clusterTime" : Timestamp(1618969341, 1),
                  "signature" : {
                          "hash" : BinData(0,"Ol7HyjGlHCL4I0hIvjx/Mj8zXCc="),
                          "keyId" : NumberLong("6953218371919806482")
                  }
          }
  }
  mongos>
  mongos> db.students.stats()
  {
          "sharded" : true,
          "capped" : false,
          ... ...
          "ns" : "shardtest.students",
          "count" : 10000,
          "size" : 600000,
          "storageSize" : 286720,
          "totalIndexSize" : 540672,
          "indexSizes" : {
                  "_id_" : 200704,
                  "_id_hashed" : 339968
          },
          "avgObjSize" : 59.6632,
          "maxSize" : NumberLong(0),
          "nindexes" : 2,
          "nchunks" : 6,
          "shards" : {
                  "shard0002" : {
                          "ns" : "shardtest.students",
                          "size" : 201311,
                          "count" : 3355,
                          "avgObjSize" : 60,
                          "storageSize" : 77824,
                          "capped" : false,
                          ... ...
                          "nindexes" : 2,
                          "totalIndexSize" : 122880,
                          "indexSizes" : {
                                  "_id_" : 49152,
                                  "_id_hashed" : 73728
                          },
                          "ok" : 1,
                          ... ...
                  },
                  "shard0000" : {
                          "ns" : "shardtest.students",
                          "size" : 202018,
                          "count" : 3368,
                          "avgObjSize" : 59,
                          "storageSize" : 102400,
                          "capped" : false,
                          ... ...
                          "nindexes" : 2,
                          "totalIndexSize" : 208896,
                          "indexSizes" : {
                                  "_id_" : 73728,
                                  "_id_hashed" : 135168
                          },
                          "ok" : 1,
                          ... ...
                  },
                  "shard0001" : {
                          "ns" : "shardtest.students",
                          "size" : 196671,
                          "count" : 3277,
                          "avgObjSize" : 60,
                          "storageSize" : 106496,
                          "capped" : false,
                          ... ...
                          "nindexes" : 2,
                          "totalIndexSize" : 208896,
                          "indexSizes" : {
                                  "_id_" : 77824,
                                  "_id_hashed" : 131072
                          },
                          "ok" : 1,
                          ... ...
                  }
          },
          "ok" : 1,
          ... ...
  }
  mongos>
  ```
  分片shard0000：3368
  分片shard0001：3277
  分片shard0002：3355
  共计10000条，且大致均匀的分布在三个分片上。
  * 查询测试
  ```
  mongos> db.students.find()
  { "_id" : 100005, "name" : "stu00005", "gender" : "female", "sno" : 5 }
  { "_id" : 100009, "name" : "stu00009", "gender" : "female", "sno" : 9 }
  { "_id" : 100018, "name" : "stu00018", "gender" : "male", "sno" : 18 }
  { "_id" : 100023, "name" : "stu00023", "gender" : "female", "sno" : 23 }
  { "_id" : 100024, "name" : "stu00024", "gender" : "male", "sno" : 24 }
  { "_id" : 100025, "name" : "stu00025", "gender" : "female", "sno" : 25 }
  { "_id" : 100027, "name" : "stu00027", "gender" : "female", "sno" : 27 }
  { "_id" : 100028, "name" : "stu00028", "gender" : "male", "sno" : 28 }
  { "_id" : 100030, "name" : "stu00030", "gender" : "male", "sno" : 30 }
  { "_id" : 100034, "name" : "stu00034", "gender" : "male", "sno" : 34 }
  { "_id" : 100036, "name" : "stu00036", "gender" : "male", "sno" : 36 }
  { "_id" : 100037, "name" : "stu00037", "gender" : "female", "sno" : 37 }
  { "_id" : 100040, "name" : "stu00040", "gender" : "male", "sno" : 40 }
  { "_id" : 100043, "name" : "stu00043", "gender" : "female", "sno" : 43 }
  { "_id" : 100054, "name" : "stu00054", "gender" : "male", "sno" : 54 }
  { "_id" : 100055, "name" : "stu00055", "gender" : "female", "sno" : 55 }
  { "_id" : 100058, "name" : "stu00058", "gender" : "male", "sno" : 58 }
  { "_id" : 100059, "name" : "stu00059", "gender" : "female", "sno" : 59 }
  { "_id" : 100062, "name" : "stu00062", "gender" : "male", "sno" : 62 }
  { "_id" : 100064, "name" : "stu00064", "gender" : "male", "sno" : 64 }
  Type "it" for more
  mongos> db.students.find({sno:{$lt:10}})
  { "_id" : 100002, "name" : "stu00002", "gender" : "male", "sno" : 2 }
  { "_id" : 100004, "name" : "stu00004", "gender" : "male", "sno" : 4 }
  { "_id" : 100000, "name" : "stu00000", "gender" : "male", "sno" : 0 }
  { "_id" : 100001, "name" : "stu00001", "gender" : "female", "sno" : 1 }
  { "_id" : 100003, "name" : "stu00003", "gender" : "female", "sno" : 3 }
  { "_id" : 100006, "name" : "stu00006", "gender" : "male", "sno" : 6 }
  { "_id" : 100007, "name" : "stu00007", "gender" : "female", "sno" : 7 }
  { "_id" : 100008, "name" : "stu00008", "gender" : "male", "sno" : 8 }
  { "_id" : 100005, "name" : "stu00005", "gender" : "female", "sno" : 5 }
  { "_id" : 100009, "name" : "stu00009", "gender" : "female", "sno" : 9 }
  mongos>
  ```

* 经过上面的测试，显示分片集群的相关功能正常。

### 完备集群
#### 集群说明
上面的最小集群虽能正常工作，但在高可用上是不足的。所谓完备，如下图所示应该是满足下面的三个条件：
![image-20220112142314921](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142314921.png)

* 分片应该是副本集的形式
* 配置服务器也应该是副本集的形式
* 部署多个mongos

只有满足上述条件，才能避免单点故障、满足高可用的要求。

#### 部署副本集
在另一篇文档中已经专门对副本集做了详细说明，这里不再赘述。

#### 部署多个mongos
在上面最小集群的基础上再增加一个mongos，如下所示，增加一个router2目录
```
  PS D:\mongoworkdir\sharding_test> tree /F
  .
  ├─configsvr
  │  │  configsvr.conf
  │  │  mongo.keyfile
  │  │
  │  ├─data
  │  │  ...
  │  │
  │  └─logs
  │          configsvr.log
  │
  ├─router
  │  │  mongo.keyfile
  │  │  router.conf
  │  │
  │  └─logs
  │      │  router.log
  |         ...
  ├─router2
  │  │  mongo.keyfile
  │  │  router2.conf
  │  │
  │  └─logs
  │
  ├─shard1
  │  │  mongo.keyfile
  │  │  shard1.conf
  │  │
  │  ├─data
  │  │  ...
  │  │
  │  └─logs
  │          shard1.log
  │
  ├─shard2
  │  │  mongo.keyfile
  │  │  shard2.conf
  │  │
  │  ├─data
  │  │  ...
  │  │
  │  └─logs
  │          shard2.log
  │
  └─shard3
      │  mongo.keyfile
      │  shard3.conf
      │
      ├─data
      │  ...
      │
      └─logs
              shard3.log
```
其中，routers.conf配置如下：
```
systemLog:
   destination: file
   path: D:\mongoworkdir\sharding_test\router2\logs\router2.log
   logAppend: true
sharding:
  configDB: cfgsvr001/127.0.0.1:27300
net:
  bindIp: 127.0.0.1
  port: 27500
security:
    keyFile: D:\mongoworkdir\sharding_test\router2\mongo.keyfile
```
与另一个mongos的配置除了部分文件路径不同，其它完全一样。
启动即可：
```
PS D:\mongoworkdir\sharding_test\router2> mongos.exe -f .\router2.conf
```
使用测试：
* 连接mongos-27400
```
PS C:\Users\YUANTINGZHONG> mongo.exe --port 27400
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27400/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("be92b3b9-e4dd-479c-921a-2e500fcd5382") }
MongoDB server version: 4.0.11
mongos> use shardtest
switched to db shardtest
mongos> db.auth("shardUser","shardUser")
1
mongos> db.students.find({sno:{$lt:10}})
{ "_id" : 100002, "name" : "stu00002", "gender" : "male", "sno" : 2 }
{ "_id" : 100004, "name" : "stu00004", "gender" : "male", "sno" : 4 }
{ "_id" : 100005, "name" : "stu00005", "gender" : "female", "sno" : 5 }
{ "_id" : 100009, "name" : "stu00009", "gender" : "female", "sno" : 9 }
{ "_id" : 100000, "name" : "stu00000", "gender" : "male", "sno" : 0 }
{ "_id" : 100001, "name" : "stu00001", "gender" : "female", "sno" : 1 }
{ "_id" : 100003, "name" : "stu00003", "gender" : "female", "sno" : 3 }
{ "_id" : 100006, "name" : "stu00006", "gender" : "male", "sno" : 6 }
{ "_id" : 100007, "name" : "stu00007", "gender" : "female", "sno" : 7 }
{ "_id" : 100008, "name" : "stu00008", "gender" : "male", "sno" : 8 }
mongos>
```
* 连接mongos-27500
```
PS C:\Users\YUANTINGZHONG> mongo.exe --port 27500
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27500/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("794e167b-d578-4aba-a73a-c59e018e7dce") }
MongoDB server version: 4.0.11
mongos> use shardtest
switched to db shardtest
mongos> db.auth("shardUser","shardUser")
1
mongos> db.students.find({sno:{$lt:10}})
{ "_id" : 100000, "name" : "stu00000", "gender" : "male", "sno" : 0 }
{ "_id" : 100001, "name" : "stu00001", "gender" : "female", "sno" : 1 }
{ "_id" : 100003, "name" : "stu00003", "gender" : "female", "sno" : 3 }
{ "_id" : 100006, "name" : "stu00006", "gender" : "male", "sno" : 6 }
{ "_id" : 100007, "name" : "stu00007", "gender" : "female", "sno" : 7 }
{ "_id" : 100008, "name" : "stu00008", "gender" : "male", "sno" : 8 }
{ "_id" : 100005, "name" : "stu00005", "gender" : "female", "sno" : 5 }
{ "_id" : 100009, "name" : "stu00009", "gender" : "female", "sno" : 9 }
{ "_id" : 100002, "name" : "stu00002", "gender" : "male", "sno" : 2 }
{ "_id" : 100004, "name" : "stu00004", "gender" : "male", "sno" : 4 }
mongos>
```
有了多个mongos之后，还可以考虑负载均衡，通常这个可以通过搭建HaProxy或者LVS的方式实现。
这里不再进一步说明。

## 分片集群实验及使用实践

### 分区使用实验
> 以下部分都是基于上文“分片集群搭建实践”中搭建的最小集群进行测试的。下文中提到的集群如无特殊说明都是指该最小集群。

#### 实验说明
* 实验目的
  演示分区的使用
* 实验方法
  设想一个场景，一个学校（数据库school）的学生（集合students）擅长不同的学科，我们现在只关注擅长数学和生物的，而且因为某些因素，必须将擅长数学和擅长生物的学生各存储到一个单独的分片中，比如擅长数学的只能存储到分片1中，而擅长生物的只能存储到分片2中。擅长其它学科的学生可以存储到任意分片中。
  考虑下这种需求如何实现？
  这就是分区的用武之地。
> 这只是一个例子，重在演示，不必较真其是否合理。

#### 基于分区的存储方案
* 集群中共有三个分片：shard0000，shard0001和shard0002。
* 针对上面描述的存储需求，设计存储方案如下：
  * 设置一个分区"MATH"存储擅长数学的学生，并把分片shard0000添加到该区；
  * 设置一个分区"BIOLOGY"存储擅长生物的学生，并把分片shard0001添加到该区；

#### 具体分区设置步骤
* 连接mongos，查看当前集群分区情况
```
PS C:\Users\YUANTINGZHONG> mongo.exe --port 27400
... ...
mongos> use admin
switched to db admin
mongos> db.auth("cluseradmin","cluseradmin")
1
mongos> db.runCommand({listshards:1})
{
        "shards" : [
                {
                        "_id" : "shard0000",
                        "host" : "127.0.0.1:27201",
                        "state" : 1
                },
                {
                        "_id" : "shard0001",
                        "host" : "127.0.0.1:27202",
                        "state" : 1
                },
                {
                        "_id" : "shard0002",
                        "host" : "127.0.0.1:27203",
                        "state" : 1
                }
        ],
        "ok" : 1,
        ... ...
}
mongos>
```
集群各分片正常。

* 创建数据库school并为其创建用户
```
use school
db.createUser({user:"zonetester",pwd:"zonetester", roles:[{ role: "readWrite", db: "school" }])
```

* 打开数据库分片使能，分片集合
```
db.runCommand({ enablesharding:"school" })
sh.shardCollection("school.students", {subject: 1, sId: 1});    //注意片键
```

* 分区设置
  * 关闭均衡器（balancer）
  ```
  sh.disableBalancing("school.students") //分区设置期间不允许Chunk分裂迁移
  ```
  * 向分区添加分片
  ```
  sh.addShardTag( "shard0000" , "MATHS");
  sh.addShardTag( "shard0001" , "BIOLOGY");
  ```
  * 查看分片状态
  ```
  mongos> sh.status()
  --- Sharding Status ---
  ... ...
  shards:
        {  "_id" : "shard0000",  "host" : "127.0.0.1:27201",  "state" : 1,  "tags" : [ "MATHS" ] }
        {  "_id" : "shard0001",  "host" : "127.0.0.1:27202",  "state" : 1,  "tags" : [ "BIOLOGY" ] }
        {  "_id" : "shard0002",  "host" : "127.0.0.1:27203",  "state" : 1 }
  ... ... 
  ```
  从tags可以看到添加分片到分区成功。
  * 为每个分区定义范围
  ```
  sh.addTagRange(
        "school.students",
        { "subject" : "maths", "sId" : MinKey},
        { "subject" : "maths", "sId" : MaxKey},
        "MATHS"
        )
  
  sh.addTagRange(
        "school.students",
        { "subject" : "biology", "sId" : MinKey},
        { "subject" : "biology", "sId" : MaxKey},
        "BIOLOGY"
        )
  ```
  即MATHS分区的范围定义为"subject"字段值是"maths","sId"字段值任意的所有记录；
  而BIOLOGY分区的范围是"subject"字段值为"biology","sId"字段值任意的所有记录。
  * 开启均衡器
  ```
  sh.enableBalancing("school.students")
  ```

至此，分区设置完毕。

#### 分区测试
> 使用go语言开发的demo进行数据写入测试，上文已经列出源代码，略加修改即可，此处不再赘述；

**测试一**
向分片集群写入10000条记录，这10000条记录的"subject"都是“math”
查看此时分片数据存储情况
命令：
```
db.students.stats()
```
数据分布情况：
shard0000:  "count" : 10000
shard0001:  "count" : 0
shard0002:  "count" : 0
可见写入的10000条“subject”为“maths”的记录全部写入了shard0000分片，shard0000分片属于"MATHS"分区，与预期一致。

**测试二**
再次写入10000条记录，这10000条记录的"subject"都是“biology”
再次查看此时分片数据存储情况:
shard0000:  "count" : 10000
shard0001:  "count" : 10000
shard0002:  "count" : 0
即再次写入的10000条“subject”为“biology”的记录全部写入了shard0001分片,shard0001分片属于"BIOLOGY"分区，与预期一致。

测试显示，通过分区达到了存储需求。

### Chunk分裂迁移实验
#### 实验目的
* 接着上面的“分区使用实验”，考虑下面的问题：
MATHS分区和BIOLOGY分区目前都只包含一个分片，随着数据的不断写入，可能会不断逼近该分片的存储能力，这时我们就要考虑进行扩容。
假如现在BIOLOGY分区数据量已经快快超过其对应分片shard0001的存储能力了，可以在集群中再增加一个分片，并且将该分片划入BIOLOGY分区，这样BIOLOGY分区又多了一个分片，其存储空间得以扩展。
那么问题来了：对BIOLOGY分区增加分片后，新增的分片和原来的分片关系如何？原来的分片有很多数据，而新增分片没有数据，会发生数据迁移吗？

* MongoDB有Chunk分裂和迁移的功能，理论上讲，新增分片后，均衡器会负责把数据从原来的分片迁移到新增分片以达到各分片数据的平衡。
现在我们通过实验来进一步验证MongoDB的迁移功能，回答上面的问题。

#### 实验步骤
* 在之前的基础上，再写入4000000条“subject”为“biology”的记录，再次查看分片数据分布情况
shard0000:  "count" : 10000
shard0001:  "count" : 4010000
shard0002:  "count" : 0
即再次写入的4000000条“subject”为“biology”的记录全部写入了shard0001分片，shard0001分片属于"BIOLOGY"分区，与预期一致

* 新增一个Mongod节点，作为新增分片shard4，其配置如下：
```
systemLog:
   destination: file
   path: D:\mongoworkdir\sharding_test\shard4\logs\shard4.log
   logAppend: true
storage:
   dbPath: D:\mongoworkdir\sharding_test\shard4\data
   directoryPerDB: true
sharding:
  clusterRole: shardsvr
net:
  bindIp: 127.0.0.1
  port: 27204
# 在创建超管账号前，下面的语句要注释掉
#security:      
#    keyFile: D:\mongoworkdir\sharding_test\shard4\mongo.keyfile
#    authorization: "enabled"
```

* 然后按照如下步骤进行
  * 启动新节点
  ```
  mongod.exe -f .\shard4.conf
  ```
  * 给新节点创建超管账号
  ```
  use admin
  db.createUser({user:"root",pwd:"root",roles:["root"]})
  db.createUser({user:"admin",pwd:"admin",roles:[{ role: "userAdminAnyDatabase", db: "admin" }]})
  ```
  * 重新启动节点，开启认证鉴权
  ```
  security:
    keyFile: D:\mongoworkdir\sharding_test\shard4\mongo.keyfile
    authorization: "enabled"
  ```
  * 将节点shard4加入分片集群
  ```
  mongo.exe --port 27400
  db.runCommand({ addshard:"127.0.0.1:27204" })
  ```
  操作记录
  ```
  PS C:\Users\YUANTINGZHONG> mongo.exe --port 27400
  ... ... 
  mongos> use admin
  switched to db admin
  mongos> db.auth("cluseradmin","cluseradmin")
  1
  mongos> db.runCommand({ addshard:"127.0.0.1:27204" })
  {
        "shardAdded" : "shard0003",
        "ok" : 1,
        ... ...
  }
  mongos>
  ```
  * 查看分片信息
  ```
  mongos> db.runCommand({listshards:1})
  {
        "shards" : [
                {
                        "_id" : "shard0000",
                        "host" : "127.0.0.1:27201",
                        "tags" : [
                                "MATHS"
                        ],
                        "state" : 1
                },
                {
                        "_id" : "shard0001",
                        "host" : "127.0.0.1:27202",
                        "tags" : [
                                "BIOLOGY"
                        ],
                        "state" : 1
                },
                {
                        "_id" : "shard0002",
                        "host" : "127.0.0.1:27203",
                        "state" : 1
                },
                {
                        "_id" : "shard0003",
                        "host" : "127.0.0.1:27204",
                        "state" : 1
                }
        ],
        "ok" : 1,
        ... ...
  }
  ```
  可见，现在集群中已经有4个分片了，新增的分片是shard0003。
  * 将新增的分片加入“BIOLOGY”分区
  ```
  sh.addShardTag( "shard0003" , "BIOLOGY");
  ```
  操作记录
  ```
  mongos> sh.addShardTag( "shard0003" , "BIOLOGY");
  {
        "ok" : 1,
        ... ...
  }
  mongos>
  ```
  * 查看分片状态
  操作记录
  ```
  mongos> sh.status()
  --- Sharding Status ---
  ... ...
  shards:
        {  "_id" : "shard0000",  "host" : "127.0.0.1:27201",  "state" : 1,  "tags" : [ "MATHS" ] }
        {  "_id" : "shard0001",  "host" : "127.0.0.1:27202",  "state" : 1,  "tags" : [ "BIOLOGY" ] }
        {  "_id" : "shard0002",  "host" : "127.0.0.1:27203",  "state" : 1 }
        {  "_id" : "shard0003",  "host" : "127.0.0.1:27204",  "state" : 1,  "tags" : [ "BIOLOGY" ] }
  active mongoses:
        "4.0.11" : 2
  autosplit:
        Currently enabled: yes
  balancer:
        Currently enabled:  yes
        Currently running:  no
        Failed balancer rounds in last 5 attempts:  0
        Migration Results for the last 24 hours:
                10 : Success
  ... ...
  mongos>
  ```
  可见，现在对应BIOLOGY分区已经有两个分片了。
  另外，注意上面的几个输出：
  autosplit用于控制是否自动触发chunk分裂，且当前是开启状态；
  balancer是均衡器，会把数据在分片间迁移以保持各分片数据平衡，且当前也是开启状态。
  其它字段含义参考[官方文档](https://docs.mongodb.com/manual/reference/method/sh.status/)。
  至此，增加一个分片，并且将该分片划入BIOLOGY分区的相关设置已经完成。下面，在不对集群进行数据写入操作的情况下查看分片数据是否会出现迁移现象。

* 查看分片数据分布状态
  * 查看分片数据操作记录
    ```
    PS C:\Users\YUANTINGZHONG>  mongo.exe --port 27400
    ... ...
    mongos> use school
    switched to db school
    mongos> db.auth("zonetester","zonetester")
    1
    mongos> db.students.stats()
    {
        "sharded" : true,
        ... ...
        "ns" : "school.students",
        "count" : 5770003,
        ... ...
        "nchunks" : 19,
        "shards" : {
                "shard0001" : {
                        "ns" : "school.students",
                        "size" : 367800000,
                        "count" : 4010000,
                        ... ...
                        "ok" : 1,
                        ... ...
                },
                "shard0003" : {
                        "ns" : "school.students",
                        "size" : 159880275,
                        "count" : 1750003,
                        ... ...
                        "ok" : 1,
                        ... ...
                },
                "shard0000" : {
                        "ns" : "school.students",
                        "size" : 880000,
                        "count" : 10000,
                        ... ...
                        "ok" : 1,
                        ... ...
                },
                "shard0002" : {
                        "ns" : "school.students",
                        "size" : 0,
                        "count" : 0,
                        ... ...
                        "ok" : 1,
                        ... ...
                }
        },
        "ok" : 1,
        ... ...
    }
    mongos>
    ```
    可见数据分布情况：
    shard0000:  "count" : 10000
    shard0001:  "count" : 4010000
    shard0002:  "count" : 0
    shard0003:  "count" : 1750003
    这里有两个情况值得注意：
    一是shard0003是新增分片，而且自其加入集群后，集群没有进行过数据写入操作，所以shard0003上的1750003条记录应该是shard0001中的Chunk分裂迁移而来；
    二是此时shard0001上的数据量并没有因为1750003条记录的迁移而减少，仍然是4010000，所以导致整体数据量超出了1750003条。
    这很奇怪，貌似有问题，但不用担心，这只是Chunk分裂迁移以达到数据平衡的一个过渡状态而已，并不是真的出了问题。
  * 过一段时间，再次查看分片数据分布情况
  shard0000:  "count" : 10000
  shard0001:  "count" : 2259997
  shard0002:  "count" : 0
  shard0003:  "count" : 1750003
  shard0001 & shard0003 ：2259997 + 1750003 = 4010000
  与上面的数据分布比较，可以看到shard0001上的数据已经变为2259997，减少的数据量恰是迁移到shard0003上的1750003条数据。
  由上面的实验现象可以看到，Chunk分裂、迁移发生后的一段时间内，原来的分片数据量可能会保持不变，一定时间后就会恢复正常，也就是说存在延迟。在实际使用查看时要注意这一点，不要误以为是出了问题。

* 上面的实验显示的是新增分片时已经存在的数据会发生迁移现象，那么对于新写入的数据会怎么样呢？
  * 再写入1000000条“subject”为“biology”的记录，然后查看分片数据分布情况
  shard0000:  "count" : 10000
  shard0001:  "count" : 3259997 （+1000000）
  shard0002:  "count" : 0
  shard0003:  "count" : 2250004 （+500001）
  又有奇怪的事情发生了， 增加了1000000+500001条，显然是不对的。但别急，等下再查看下
  * 等一段时间，再次查看分片数据分布情况
  shard0000:  "count" : 10000
  shard0001:  "count" : 2759996 （+499999）
  shard0002:  "count" : 0
  shard0003:  "count" : 2250004 （+500001）
  shard0001和shard0003增加的数据量为499999+500001=1000000，即写入的正确的数据量。
  总之，因为有Chunk分裂与迁移的存在，所以可能会出现一些中间态，导致集合总记录数与写入的记录数不一致的情况，待迁移完成即可。

## 分片集群部分问题补充说明

### 关于分片的索引问题
* 如下图所示
![image-20220112142347077](E:\study\studygo\Golang学习笔记\mongodb分片.assets\image-20220112142347077.png)
由上图可见，分片集群中的索引文件是分开存放的，每个分片存储、管理自己数据的索引。
