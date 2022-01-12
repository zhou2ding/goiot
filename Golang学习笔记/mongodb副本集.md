[TOC]
# MongoDB - 副本集
## 1. 概念
### 1.1 MongoDB副本集
MongoDB副本集是由一组mongod进程组成的，能够提供**数据冗余**和**高可用**的集群。

### 1.2 副本集组成及工作原理
* 主节点（Primary） 
  *  主节点是副本集中唯一可以接受写操作的成员。MongoDB在主节点上应用写操作，然后在主节点的oplog上记录这些操作。从节点复制此日志并将这些操作应用于其数据集。如下图所示 
  ![image-20220112140402631](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112140402631.png)
  *  复制集中的所有节点都可以接受读操作。默认情况下，应用程序将其读操作定向到主节点。
  *  复制集最多只能有一个主节点。如果当前主节点不可用，则通过[选举](https://docs.mongodb.com/manual/core/replica-set-elections/)确定新的主节点。
* 从节点（Secondaries）
  *  从节点维护着主节点数据集的副本。为了复制数据，从节点在一个异步进程中将源自主节点的[oplog](https://docs.mongodb.com/manual/core/replica-set-oplog/)应用到自己的数据集。一个副本集可以有一个或多个从节点。
  *  尽管客户端不能向从节点写数据，但可以从从节点读取数据。这与客户端的“读偏好”（[Read Perference](https://docs.mongodb.com/manual/core/read-preference/)）设置有关。
  *  从节点可以变成主节点。如果当前主节点不可用，复制集会进行选举来决定哪个从节点变成主节点。如下图所示，在有三个成员的复制集中，主节点变为不可用，这将触发一个选举以选出一个从节点变成新的主节点。
  ![image-20220112141259225](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112141259225.png)
* 仲裁节点（Arbiter）
  在某些情况下，比如已经有一个主节点和从节点，但由于成本限制而不允许再添加另一个从节点，这时可以选择向副本集中添加一个仲裁节点。仲裁节点不保存数据，不能成为主节点，但可以参与选举主节点。
  例如，在下面的副本集中，包含2个数据成员（主节点和从节点），仲裁节点允许该数据集具有奇数的票数来打破平局。
  ![image-20220112141310955](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112141310955.png)

### 1.3 副本集限制
* 副本集的最低建议配置是三个成员的副本集，其中包含三个数据承载成员：一个主节点和两个从节点。在某些情况下，可以选择包含仲裁节点。仲裁节点参与选举，但不保存数据（即不提供数据冗余）。
* 副本集最多可以有50个成员，但只有7个投票成员。

## 2. 副本集部署体系结构
副本集的体系结构影响副本集的容量和性能。本文提供副本集部署的策略，并描述常见的部署体系结构。
生成系统的标准副本集部署是由三个成员组成的副本集。这些集合提供冗余和容错。要尽可能避免复杂性，但让你的应用需求来决定体系结构。

### 2.1 策略
* 确定成员的数量
  根据这些策略来添加副本集成员
  *  投票成员的最大数目
  一个副本集最多可以有50个成员，但只能有7个投票成员。如果副本集已经有了7个投票成员，其它的成员必须是[非投票成员](https://docs.mongodb.com/manual/core/replica-set-elections/#std-label-replica-set-non-voting-members)。
  *  部署奇数个成员
  确保副本集有奇数个投票成员。一个副本集最多可以有7个投票成员。如果有偶数个投票成员，则部署另一个数据承载投票成员，或者，如果条件限制禁止另一个数据承载投票成员，则部署一个仲裁节点。
  仲裁节点不存储数据的副本，并且需要较少的资源。因此，可以在应用程序服务器或其他共享进程上运行仲裁节点。在没有数据副本的情况下，可以将仲裁节点放置到不放置副本集其它成员的环境中。
  注意，通常避免为每个副本集部署多个仲裁节点。
  *  考虑容错
     副本集的容错性是指副本集有部分成员变得不可用，但仍有足够的成员来选举出一个主节点。换句话说，这是集合中的成员数与选举出主节点所需的投票成员的多数之间的差额。如果没有主节点，副本集不能接受写操作。容错性受副本集大小的影响，但这种关系不是直接的。见下表：

     | 成员数量 | 选举出新主节点所需的多数 | 容错 |
     | --- | ---  | --- |
     | 3 | 2 | 1 |
     | 4 | 3 | 1 |
     | 5 | 3 | 2 |
     | 6 | 4 | 2 |
        向副本集中添加一个成员并不总是增加容错性。但是，在这些情况下，其它成员可以为专用功能（如备份或报告）提供支持。
* 使用隐藏成员和延迟成员实现专门功能
  添加[隐藏](https://docs.mongodb.com/manual/core/replica-set-hidden-member/#std-label-replica-set-hidden-members)或[延迟](https://docs.mongodb.com/manual/core/replica-set-delayed-member/#std-label-replica-set-delayed-members)成员以支持专用功能，例如备份或报告。
* 大量读操作环境中部署上的负载均衡
  在具有非常高的读取流量的部署中，可以通过将读取分发给从节点来提高读取吞吐量。随着部署的增长，可以将成员添加或移动到备用数据中心，以提高冗余度和可用性。
* 提前增加容量
  副本集的现有成员必须具有备用容量，以支持添加新成员。总是在当前需求饱和集合容量之前添加新成员。
* 按地理位置分布成员
  为了在数据中心发生故障时保护数据，应该在备用数据中心至少保留一个成员。如果可能，使用奇数个数据中心，并选择一个成员分布以最大限度地提高这样的可能性，即即使数据中心丢失，剩余的副本集成员也可以构成多数，或者至少提供数据的副本。
  要确保主数据中心的成员比备用数据中心的成员先被选为主成员，应该将备用数据中心的成员的优先级设置为低于主数据中心的成员的优先级。
* 使用标签集指定目标操作
使用[副本集标签](https://docs.mongodb.com/manual/tutorial/configure-replica-set-tag-sets/#std-label-replica-set-configuration-tag-sets)来指定特定成员读操作或自定义写关注来获取特定成员的确认。
* 使用Journaling日志功能防止电源故障
MongoDB默认启用了[Journaling](https://docs.mongodb.com/manual/core/journaling/)日志功能。Journaling避免了服务中断引起的数据丢失，比如电源故障或意外的重启。
* 主机名
尽可能使用逻辑DNS主机名而不是ip地址，特别是在配置副本集成员或分片群集成员时。使用逻辑DNS主机名可以避免由于ip地址更改而导致的配置更改。

### 2.2 副本集命名
如果你的应用程序连接了多个副本集，每个副本集应该拥有不同的名字。一些驱动按副本集名字来为副本集连接分组。

## 3. 副本集搭建举例
### 3.1 示例说明
#### 3.1.1 目标
![image-20220112141328593](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112141328593.png)
如上图所示，该示例将详述如何搭建一个三节点（包括一个主节点、一个从节点和一个仲裁节点)的副本集。这是一个典型的副本集，按此方法可以搭建其它包括更多从节点的副本集。

#### 3.1.2 注意事项
* 实际应用中，副本集的各个节点要分布于不同的主机上，才能达到高可用的目的。但为了方便演示，示例中将所有节点都放在同一台主机上，通过不同的端口区分。这并不影响示例中演示的搭建方法的正确性，实际中搭建时，只要将示例中的IP和端口信息修改为与实际相符即可。
* 示例是在Windows上进行的，Linux环境下的操作基本相同。
* MongoDB的安装不在示例范围内，所以此处假设MongoDB已经安装，而且对于Windows环境来说，MongoDB的安装目录下的bin目录已经添加到环境变量PATH中。对于Linux来说，也有类似假设。

### 3.2 搭建步骤
#### 3.2.1 创建目录及相关文件
```
.
├─arbiter
│  │  arbiter.conf
│  │
│  ├─data
│  └─logs
├─dnode1
│  │  dnode1.conf
│  │
│  ├─data
│  └─logs
└─dnode2
    │  dnode2.conf
    │
    ├─data
    └─logs
```
如上所示，建立三个目录：dnode1、dnode2对应数据节点，arbiter对应仲裁节点。每个目录下创建子目录data和logs，另外创建各自的配置文件.conf。其中data目录是MongoDB数据存放目录，logs是相关日志存放目录，.conf文件是配置文件。
> 警告：这里为了方便，把它们集中存放，实际中一定要根据需要合理安排目录的位置。

#### 3.2.2 各节点的配置文件
节点dnode1的配置文件dnode1.conf内容
```yaml
systemLog:
    destination: file
    path: D:\mongoworkdir\replica_demo\dnode1\logs\dnode1.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\replica_demo\dnode1\data
    directoryPerDB: true
replication:
    replSetName: rs001
net:
    bindIp: 127.0.0.1
    port: 27500
```
节点dnode2的配置文件dnode2.conf内容
```yaml
systemLog:
    destination: file
    path: D:\mongoworkdir\replica_demo\dnode2\logs\dnode2.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\replica_demo\dnode2\data
    directoryPerDB: true
replication:
    replSetName: rs001
net:
    bindIp: 127.0.0.1
    port: 27501
```
节点arbiter的配置文件arbiter.conf内容
```yaml
systemLog:
    destination: file
    path: D:\mongoworkdir\replica_demo\arbiter\logs\arbiter.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\replica_demo\arbiter\data
    directoryPerDB: true
replication:
    replSetName: rs001
net:
    bindIp: 127.0.0.1
    port: 27502
```
> 注意各节点绑定的端口。

#### 3.2.3 启动并配置副本集节点 
* 打开三个终端分别启动各节点
```
\\终端1
PS D:\mongoworkdir\replica_demo\dnode1> mongod.exe -f .\dnode1.conf

\\终端2
PS D:\mongoworkdir\replica_demo\dnode2> mongod.exe -f .\dnode2.conf

\\终端3
PS D:\mongoworkdir\replica_demo\arbiter> mongod.exe -f .\arbiter.conf
```
> 警告：上面是手动在终端窗口中直接启动Mongod进程，实际中可以通过服务等进行自动化操作。这里的重点是演示操作过程。
* 再打开一个终端，用monogo客户端连接dnode1、dnode2（数据节点）中任意一个节点mongodb，进入命令行，将三个实例关联起来
```
> use admin
> cfg = {_id: "rs001",members:[{_id: 1,host: '127.0.0.1:27500'},{_id: 2,host: '127.0.0.1:27501'},{_id: 3,host: '127.0.0.1:27502',arbiterOnly: true}]};
> rs.initiate(cfg)    #使配置生效
```
或者
```
> use admin
> cfg = {_id: "rs001",members:[{_id: 1,host: '127.0.0.1:27500',priority: 2},{_id: 2,host: '127.0.0.1:27501',priority: 1},{_id: 3,host: '127.0.0.1:27502',arbiterOnly: true}]};
> rs.initiate(cfg)    #使配置生效
```
cfg是可以任意的名字，当然最好不要是mongodb的关键字，conf，config都可以。最外层的_id表示replica set的名字，members里包含的是所有节点的地址以及优先级。优先级最高的即成为主节点。特别注意的是，对于仲裁节点，需要有个特别的配置——arbiterOnly:true。这个千万不能少了，不然主备模式就不能生效。
> 注意，前者没有使用优先级，所以主节点要通过选举决定。本示例中使用前者，即不指定优先级。
* 操作记录如下(“>” 是命令输入提示符，后面跟着命令，其它部分是命令的响应。后文相同)
```
PS C:\Users\YUANTINGZHONG> mongo.exe --port 27500
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27500/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("87166327-1ef0-422c-91da-922379cd7e9b") }
MongoDB server version: 4.0.11
Server has startup warnings:
2021-04-13T00:39:49.369-0700 I CONTROL  [initandlisten]
2021-04-13T00:39:49.369-0700 I CONTROL  [initandlisten] ** WARNING: Access control is not enabled for the database.
2021-04-13T00:39:49.369-0700 I CONTROL  [initandlisten] **          Read and write access to data and configuration is unrestricted.
2021-04-13T00:39:49.369-0700 I CONTROL  [initandlisten]
---
... ... 
---

> use admin
switched to db admin
> cfg = {_id: "rs001",members:[{_id: 1,host: '127.0.0.1:27500'},{_id: 2,host: '127.0.0.1:27501'},{_id: 3,host: '127.0.0.1:27502',arbiterOnly: true}]};
{
        "_id" : "rs001",
        "members" : [
                {
                        "_id" : 1,
                        "host" : "127.0.0.1:27500"
                },
                {
                        "_id" : 2,
                        "host" : "127.0.0.1:27501"
                },
                {
                        "_id" : 3,
                        "host" : "127.0.0.1:27502",
                        "arbiterOnly" : true
                }
        ]
}
> rs.initiate(cfg)
{
        "ok" : 1,
        "operationTime" : Timestamp(1618301167, 1),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618301167, 1),
                "signature" : {
                        "hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
                        "keyId" : NumberLong(0)
                }
        }
}
rs001:SECONDARY>
rs001:PRIMARY>
```
> 注意上面提示符的变化">" 、"rs001:SECONDARY>" 、"rs001:PRIMARY>", 这反映了选举的过程。首先，连接mongodb成功，提示">"可以输入命令；其次，调用rs.initiate后，将各个节点加入副本集，这时还没有选举出主节点，所以提示""rs001:SECONDARY>""; 然后，经过选举，确定dnode1节点是主节点，而当前连接的正是dnode1节点（27500），所以提示符再次改变为"rs001:PRIMARY>"。
* 检验
配置的生效时间根据不同的机器配置会有长有短，配置不错的话基本上十几秒内就能生效，有的配置需要一两分钟。如果生效了，执行rs.status()命令会看到相关信息。即查看副本集状态命令rs.status()，操作记录如下
```
rs001:PRIMARY> rs.status()
{
        "set" : "rs001",
        "date" : ISODate("2021-04-13T08:17:59.061Z"),
        "myState" : 1,
        "term" : NumberLong(1),
        "syncingTo" : "",
        "syncSourceHost" : "",
        "syncSourceId" : -1,
        "heartbeatIntervalMillis" : NumberLong(2000),
        "optimes" : {
                "lastCommittedOpTime" : {
                        "ts" : Timestamp(1618301869, 1),
                        "t" : NumberLong(1)
                },
                "readConcernMajorityOpTime" : {
                        "ts" : Timestamp(1618301869, 1),
                        "t" : NumberLong(1)
                },
                "appliedOpTime" : {
                        "ts" : Timestamp(1618301869, 1),
                        "t" : NumberLong(1)
                },
                "durableOpTime" : {
                        "ts" : Timestamp(1618301869, 1),
                        "t" : NumberLong(1)
                }
        },
        "lastStableCheckpointTimestamp" : Timestamp(1618301839, 1),
        "members" : [
                {
                        "_id" : 1,
                        "name" : "127.0.0.1:27500",
                        "health" : 1,
                        "state" : 1,
                        "stateStr" : "PRIMARY",
                        "uptime" : 2292,
                        "optime" : {
                                "ts" : Timestamp(1618301869, 1),
                                "t" : NumberLong(1)
                        },
                        "optimeDate" : ISODate("2021-04-13T08:17:49Z"),
                        "syncingTo" : "",
                        "syncSourceHost" : "",
                        "syncSourceId" : -1,
                        "infoMessage" : "",
                        "electionTime" : Timestamp(1618301178, 1),
                        "electionDate" : ISODate("2021-04-13T08:06:18Z"),
                        "configVersion" : 1,
                        "self" : true,
                        "lastHeartbeatMessage" : ""
                },
                {
                        "_id" : 2,
                        "name" : "127.0.0.1:27501",
                        "health" : 1,
                        "state" : 2,
                        "stateStr" : "SECONDARY",
                        "uptime" : 711,
                        "optime" : {
                                "ts" : Timestamp(1618301869, 1),
                                "t" : NumberLong(1)
                        },
                        "optimeDurable" : {
                                "ts" : Timestamp(1618301869, 1),
                                "t" : NumberLong(1)
                        },
                        "optimeDate" : ISODate("2021-04-13T08:17:49Z"),
                        "optimeDurableDate" : ISODate("2021-04-13T08:17:49Z"),
                        "lastHeartbeat" : ISODate("2021-04-13T08:17:57.070Z"),
                        "lastHeartbeatRecv" : ISODate("2021-04-13T08:17:58.669Z"),
                        "pingMs" : NumberLong(0),
                        "lastHeartbeatMessage" : "",
                        "syncingTo" : "127.0.0.1:27500",
                        "syncSourceHost" : "127.0.0.1:27500",
                        "syncSourceId" : 1,
                        "infoMessage" : "",
                        "configVersion" : 1
                },
                {
                        "_id" : 3,
                        "name" : "127.0.0.1:27502",
                        "health" : 1,
                        "state" : 7,
                        "stateStr" : "ARBITER",
                        "uptime" : 711,
                        "lastHeartbeat" : ISODate("2021-04-13T08:17:57.071Z"),
                        "lastHeartbeatRecv" : ISODate("2021-04-13T08:17:57.721Z"),
                        "pingMs" : NumberLong(0),
                        "lastHeartbeatMessage" : "",
                        "syncingTo" : "",
                        "syncSourceHost" : "",
                        "syncSourceId" : -1,
                        "infoMessage" : "",
                        "configVersion" : 1
                }
        ],
        "ok" : 1,
        "operationTime" : Timestamp(1618301869, 1),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618301869, 1),
                "signature" : {
                        "hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
                        "keyId" : NumberLong(0)
                }
        }
}
rs001:PRIMARY>
```
从上面rs.status()的输出信息可以看到，副本集的名字是“rs001”, 状态"ok"为1表示工作正常；有三个成员，每个成员的的具体信息等等。
> 连接从节点或仲裁节点，也可以执行rs.status查看副本集状态。不过，默认情况下从节点不可读，读取会报错“not master and slaveok=false”。这时可以调用下：db.getMongo().setSlaveOk()，然后再执行rs.status()即可。

#### 3.2.4 添加管理员账号
MongoDB安装完成后，默认是没有[权限验证](https://docs.mongodb.com/manual/core/authorization/)的，不需要输入用户名密码即可登录的（即任何客户端都可以使用mongo IP:27017/admin命令登录mongo服务），但是往往数据库方面我们会出于安全性的考虑而设置用户名密码。
启用访问控制前，请确保在 admin 数据库中拥有 userAdmin 或 userAdminAnyDatabase 角色的用户。该用户可以管理用户和角色，例如：创建用户，授予或撤销用户角色，以及创建或修改定义角色。
按如下方法添加超管用户：

* 连接到**主节点**（注意这里一定要是主节点），使用下面的命令创建管理员账号
```
db.createUser({user:"root",pwd:"root",roles:["root"]})
db.createUser({user:"admin",pwd:"admin",roles:[{ role: "userAdminAnyDatabase", db: "admin" }]})
```
操作记录
```
rs001:PRIMARY> db.createUser({user:"root",pwd:"root",roles:["root"]})
Successfully added user: { "user" : "root", "roles" : [ "root" ] }
rs001:PRIMARY> db.createUser({user:"admin",pwd:"admin",roles:[{ role: "userAdminAnyDatabase", db: "admin" }]})
Successfully added user: {
        "user" : "admin",
        "roles" : [
                {
                        "role" : "userAdminAnyDatabase",
                        "db" : "admin"
                }
        ]
}
rs001:PRIMARY>
```
上面的“Successfully”显示已经成功创建用户。
* 可以使用下面的命令查看刚刚创建的用户
```
rs001:PRIMARY> db.getUsers()
[
        {
                "_id" : "admin.admin",
                "userId" : UUID("b2c27506-ce7c-4a87-9454-a7f5b3825021"),
                "user" : "admin",
                "db" : "admin",
                "roles" : [
                        {
                                "role" : "userAdminAnyDatabase",
                                "db" : "admin"
                        }
                ],
                "mechanisms" : [
                        "SCRAM-SHA-1",
                        "SCRAM-SHA-256"
                ]
        },
        {
                "_id" : "admin.root",
                "userId" : UUID("568ddbbc-3a4d-4ace-bdbe-2b190df90585"),
                "user" : "root",
                "db" : "admin",
                "roles" : [
                        {
                                "role" : "root",
                                "db" : "admin"
                        }
                ],
                "mechanisms" : [
                        "SCRAM-SHA-1",
                        "SCRAM-SHA-256"
                ]
        }
]
rs001:PRIMARY>
```
*  连接到从节点，使用db.getUsers()查看用户
操作记录如下
```
PS C:\Users\YUANTINGZHONG> mongo.exe --port 27501
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27501/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("7f1a879b-36d0-4aef-957f-a6f265a2bd95") }
MongoDB server version: 4.0.11
Server has startup warnings:
2021-04-13T00:40:27.778-0700 I CONTROL  [initandlisten]
2021-04-13T00:40:27.778-0700 I CONTROL  [initandlisten] ** WARNING: Access control is not enabled for the database.
2021-04-13T00:40:27.778-0700 I CONTROL  [initandlisten] **          Read and write access to data and configuration is unrestricted.
2021-04-13T00:40:27.779-0700 I CONTROL  [initandlisten]
---
... ...
---

rs001:SECONDARY> db.getUsers()
2021-04-13T17:35:00.805+0800 E QUERY    [js] Error: not master and slaveOk=false :
_getErrorWithCode@src/mongo/shell/utils.js:25:13
DB.prototype.getUsers@src/mongo/shell/db.js:1763:1
@(shell):1:1
rs001:SECONDARY> db.getMongo().setSlaveOk()
rs001:SECONDARY>
rs001:SECONDARY> db.getUsers()
[ ]
rs001:SECONDARY> use admin
switched to db admin
rs001:SECONDARY> db.getUsers()
[
        {
                "_id" : "admin.admin",
                "userId" : UUID("b2c27506-ce7c-4a87-9454-a7f5b3825021"),
                "user" : "admin",
                "db" : "admin",
                "roles" : [
                        {
                                "role" : "userAdminAnyDatabase",
                                "db" : "admin"
                        }
                ],
                "mechanisms" : [
                        "SCRAM-SHA-1",
                        "SCRAM-SHA-256"
                ]
        },
        {
                "_id" : "admin.root",
                "userId" : UUID("568ddbbc-3a4d-4ace-bdbe-2b190df90585"),
                "user" : "root",
                "db" : "admin",
                "roles" : [
                        {
                                "role" : "root",
                                "db" : "admin"
                        }
                ],
                "mechanisms" : [
                        "SCRAM-SHA-1",
                        "SCRAM-SHA-256"
                ]
        }
]
rs001:SECONDARY>
```
从上面的在从节点的操作的错误提示中，可以看到要注意一下几点：
1) 从节点默认不可读，要设置一下，前文已经就此说明过，不再赘述；
2) 查看用户时，因为管理员账号是在admin库创建的，所以要先切换到admin库才能看到；
3) 在主节点创建的账号会自动复制到所有从节点，这与前文中论述的副本集工作原理是一致的。

#### 3.2.5 开启权限认证
* 先依次关闭仲裁节点、从节点、主节点。
* 生成key文件
  *  利用openssl生成key文件
  ```bash
  openssl.exe rand -out ./mongo.keyfile -hex 128
  #或
  openssl.exe rand -out ./mongo.keyfile -base64 90 
  ```
  *  将key文件复制到各个节点所在目录,如下图所示
  ![image-20220112141354939](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112141354939.png)
  从节点、仲裁节点也以同样方式放置即可。
> 再次强调下，本文中相关文件放置的位置，只是为了演示方便，实际中要根据实际情况和需求合理安排。
* 在各个节点配置文件中增加auth和keyFile配置
```
security:
    keyFile: path\to\keyfile
    authorization: "enabled"
```
* 最终，各节点的配置文件为
dnode1.conf:
```yaml
systemLog:
    destination: file
    path: D:\mongoworkdir\replica_demo\dnode1\logs\dnode1.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\replica_demo\dnode1\data
    directoryPerDB: true
replication:
    replSetName: rs001
net:
    bindIp: 127.0.0.1
    port: 27500
security:
    keyFile: D:\mongoworkdir\replica_demo\dnode1\mongo.keyfile
    authorization: "enabled"

```
dnode2.conf:
```yaml
systemLog:
    destination: file
    path: D:\mongoworkdir\replica_demo\dnode2\logs\dnode2.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\replica_demo\dnode2\data
    directoryPerDB: true
replication:
    replSetName: rs001
net:
    bindIp: 127.0.0.1
    port: 27501
security:
    keyFile: D:\mongoworkdir\replica_demo\dnode2\mongo.keyfile
    authorization: "enabled"
```
arbiter.conf
```yaml
systemLog:
    destination: file
    path: D:\mongoworkdir\replica_demo\arbiter\logs\arbiter.log
    logAppend: true
storage:
    dbPath: D:\mongoworkdir\replica_demo\arbiter\data
    directoryPerDB: true
replication:
    replSetName: rs001
net:
    bindIp: 127.0.0.1
    port: 27502
security:
    keyFile: D:\mongoworkdir\replica_demo\arbiter\mongo.keyfile
    authorization: "enabled"
```
* 重启各节点，检查权限校验是否成功开启
(1) 在终端中启动各个节点;
(2) 连接dnode1节点，发现dnode1已经从之前的主节点变为从节点了
  ```
  PS C:\Users\YUANTINGZHONG> mongo.exe --port 27500
  MongoDB shell version v4.0.11
  connecting to: mongodb://127.0.0.1:27500/?gssapiServiceName=mongodb
  Implicit session: session { "id" : UUID("89724b6d-442b-41b0-86c0-936ec6b26e99") }
  MongoDB server version: 4.0.11
  rs001:SECONDARY>
  rs001:SECONDARY>
  ```
(3) 再连接dnode2节点，对应的发现dnode2已经从之前的从节点变为主节点
```
PS C:\Users\YUANTINGZHONG> mongo.exe --port 27501
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27501/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("715809ba-611c-4aad-905d-6badf8a465ab") }
MongoDB server version: 4.0.11
rs001:PRIMARY>
```
(4) 校验权限并查看副本集状态
  操作记录（下面的操作在主节点和从节点都可以，现象一样）

```
rs001:PRIMARY> use admin
switched to db admin
rs001:PRIMARY> db.getUsers()
2021-04-13T19:17:46.414+0800 E QUERY    [js] Error: command usersInfo requires authentication :
_getErrorWithCode@src/mongo/shell/utils.js:25:13
DB.prototype.getUsers@src/mongo/shell/db.js:1763:1
@(shell):1:1
rs001:PRIMARY> db.auth("admin","admin")
1
rs001:PRIMARY> rs.status()
{
        "operationTime" : Timestamp(1618312721, 1),
        "ok" : 0,
        "errmsg" : "not authorized on admin to execute command { replSetGetStatus: 1.0, lsid: { id: UUID(\"715809ba-611c-4aad-905d-6badf8a465ab\") }, $clusterTime: { clusterTime: Timestamp(1618312671, 1), signature: { hash: BinData(0, 510302E241F817ADBD8B7B5A1A1ABE64172AA7B3), keyId: 6950550638883241985 } }, $db: \"admin\" }",
        "code" : 13,
        "codeName" : "Unauthorized",
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618312721, 1),
                "signature" : {
                        "hash" : BinData(0,"W9Qe7QeWlJAB2j7I9FCiIRzeqkg="),
                        "keyId" : NumberLong("6950550638883241985")
                }
        }
}
rs001:PRIMARY> db.grantRolesToUser( "admin" , [{ "role": "clusterAdmin", "db": "admin" }])
rs001:PRIMARY> rs.status()
{
        "set" : "rs001",
        "date" : ISODate("2021-04-13T11:19:33.472Z"),
        "myState" : 1,
        "term" : NumberLong(3),
        "syncingTo" : "",
        "syncSourceHost" : "",
        "syncSourceId" : -1,
        "heartbeatIntervalMillis" : NumberLong(2000),
        "optimes" : {
                "lastCommittedOpTime" : {
                        "ts" : Timestamp(1618312771, 1),
                        "t" : NumberLong(3)
                },
                "readConcernMajorityOpTime" : {
                        "ts" : Timestamp(1618312771, 1),
                        "t" : NumberLong(3)
                },
                "appliedOpTime" : {
                        "ts" : Timestamp(1618312771, 1),
                        "t" : NumberLong(3)
                },
                "durableOpTime" : {
                        "ts" : Timestamp(1618312771, 1),
                        "t" : NumberLong(3)
                }
        },
        "lastStableCheckpointTimestamp" : Timestamp(1618312761, 1),
        "members" : [
                {
                        "_id" : 1,
                        "name" : "127.0.0.1:27500",
                        "health" : 1,
                        "state" : 2,
                        "stateStr" : "SECONDARY",
                        "uptime" : 183,
                        "optime" : {
                                "ts" : Timestamp(1618312771, 1),
                                "t" : NumberLong(3)
                        },
                        "optimeDurable" : {
                                "ts" : Timestamp(1618312771, 1),
                                "t" : NumberLong(3)
                        },
                        "optimeDate" : ISODate("2021-04-13T11:19:31Z"),
                        "optimeDurableDate" : ISODate("2021-04-13T11:19:31Z"),
                        "lastHeartbeat" : ISODate("2021-04-13T11:19:32.076Z"),
                        "lastHeartbeatRecv" : ISODate("2021-04-13T11:19:32.796Z"),
                        "pingMs" : NumberLong(0),
                        "lastHeartbeatMessage" : "",
                        "syncingTo" : "127.0.0.1:27501",
                        "syncSourceHost" : "127.0.0.1:27501",
                        "syncSourceId" : 2,
                        "infoMessage" : "",
                        "configVersion" : 1
                },
                {
                        "_id" : 2,
                        "name" : "127.0.0.1:27501",
                        "health" : 1,
                        "state" : 1,
                        "stateStr" : "PRIMARY",
                        "uptime" : 188,
                        "optime" : {
                                "ts" : Timestamp(1618312771, 1),
                                "t" : NumberLong(3)
                        },
                        "optimeDate" : ISODate("2021-04-13T11:19:31Z"),
                        "syncingTo" : "",
                        "syncSourceHost" : "",
                        "syncSourceId" : -1,
                        "infoMessage" : "",
                        "electionTime" : Timestamp(1618312599, 1),
                        "electionDate" : ISODate("2021-04-13T11:16:39Z"),
                        "configVersion" : 1,
                        "self" : true,
                        "lastHeartbeatMessage" : ""
                },
                {
                        "_id" : 3,
                        "name" : "127.0.0.1:27502",
                        "health" : 1,
                        "state" : 7,
                        "stateStr" : "ARBITER",
                        "uptime" : 180,
                        "lastHeartbeat" : ISODate("2021-04-13T11:19:32.076Z"),
                        "lastHeartbeatRecv" : ISODate("2021-04-13T11:19:33.016Z"),
                        "pingMs" : NumberLong(0),
                        "lastHeartbeatMessage" : "",
                        "syncingTo" : "",
                        "syncSourceHost" : "",
                        "syncSourceId" : -1,
                        "infoMessage" : "",
                        "configVersion" : 1
                }
        ],
        "ok" : 1,
        "operationTime" : Timestamp(1618312771, 1),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618312771, 1),
                "signature" : {
                        "hash" : BinData(0,"Y+IrsNKCNr6B3I+Lv1KMcA2MABM="),
                        "keyId" : NumberLong("6950550638883241985")
                }
        }
}
rs001:PRIMARY>
```
上面的操作中：
a. 切换到admin后直接调用db.getUsers()会报错“command usersInfo requires authentication”，说明权限校验已经打开；
b. 进行权限校验db.auth("admin","admin")，返回1表示权限校验通过；
c. 使用命令rs.status()查看状态报错"Unauthorized"，这是因为admin账号只是被赋予了userAdminAnyDatabase权限，而不具有查看副本集状态的权限，所以在操作使用rs.status()时，提示没有操作权限。需要重新赋予副本集的操作权限，如下
```
db.grantRolesToUser( "admin" , [{ "role": "clusterAdmin", "db": "admin" }])
```
然后就可以rs.status()查看状态了。
经过上述操作，MongoDB的权限校验被打开，各种操作都要先进行权限校验。
* 上面的操作中还有以下几点需要注意
  * 重启各个节点会出现重新选举主节点的情况，也就是说主节点并不是固定的；
  * 重启后，副本集内各节点会自动关联起来（加入副本集），不需要手动操作；
  * 连接仲裁节点，可以看到仲裁节点还是安装后的默认状态（没有账号、未开启权限校验，可以直接查看副本集状态），如下
  ```
  PS D:\mongoworkdir\replica_demo> mongo.exe --port 27502
  MongoDB shell version v4.0.11
  connecting to: mongodb://127.0.0.1:27502/?gssapiServiceName=mongodb
  Implicit session: session { "id" : UUID("d0f4c0ea-9785-4f8f-8372-8e23f54e662d") }
  MongoDB server version: 4.0.11
  rs001:ARBITER> use admin
  switched to db admin
  rs001:ARBITER> db.auth("admin","admin")
  Error: Authentication failed.
  0
  rs001:ARBITER> rs.status()
  {
          "set" : "rs001",
          ... ... ... 
  }
  ```
  这是因为之前创建管理员账号时是在主节点创建的，从节点会同步（复制）主节点的操作，也会创建同样的账号，但仲裁节点不会同步主节点操作。另外，仲裁节点配置中开启权限校验的配置项实际上是并不发挥作用（但不能从配置中省略）。因仲裁节点不存储任何数据，所以这些操作都是合理的。
#### 3.2.6 副本集使用测试
* 主节点是副本集中唯一接受写操作的节点。
* 连接主节点，作如下操作：
  *  创建replSetTest数据库，创建一个可读写数据库的用户replUser
  ```
  use replSetTest
  db.createUser({user:"replUser",pwd:"replUser", roles:[{ role: "readWrite", db: "replSetTest" }]})
  ```
  * 重新连接主节点，切换到replSetTest数据库，完成认证
  ```
  PS C:\Users\YUANTINGZHONG> mongo.exe --port 27501
  MongoDB shell version v4.0.11
  connecting to: mongodb://127.0.0.1:27501/?gssapiServiceName=mongodb
  Implicit session: session { "id" : UUID("2b194407-2961-4b63-9d52-eb342f137758") }
  MongoDB server version: 4.0.11
  rs001:PRIMARY> use replSetTest
  switched to db replSetTest
  rs001:PRIMARY> db.auth("replUser","replUser")
  1
  rs001:PRIMARY>
  ```
  *  插入测试数据
  ```
  rs001:PRIMARY> db.orders.insert({ "_id" : 1, "item" : "abc", "price" : 10, "quantity" : 2, "date" : ISODate("2014-03-01T08:00:00Z") })
  WriteResult({ "nInserted" : 1 })
  rs001:PRIMARY> db.orders.insert({ "_id" : 2, "item" : "jkl", "price" : 20, "quantity" : 1, "date" : ISODate("2014-03-01T09:00:00Z") })
  WriteResult({ "nInserted" : 1 })
  rs001:PRIMARY> db.orders.insert({ "_id" : 3, "item" : "xyz", "price" : 5, "quantity" : 10, "date" : ISODate("2014-03-15T09:00:00Z") })
  WriteResult({ "nInserted" : 1 })
  rs001:PRIMARY> db.orders.insert({ "_id" : 4, "item" : "xyz", "price" : 5, "quantity" : 20, "date" : ISODate("2014-04-04T11:21:39.736Z") })
  WriteResult({ "nInserted" : 1 })
  rs001:PRIMARY> db.orders.insert({ "_id" : 5, "item" : "abc", "price" : 10, "quantity" : 10, "date" : ISODate("2014-04-04T21:23:13.331Z") })
  WriteResult({ "nInserted" : 1 })
  rs001:PRIMARY>
  rs001:PRIMARY> 
  rs001:PRIMARY> db.orders.find()
  { "_id" : 1, "item" : "abc", "price" : 10, "quantity" : 2, "date" : ISODate("2014-03-01T08:00:00Z") }
  { "_id" : 2, "item" : "jkl", "price" : 20, "quantity" : 1, "date" : ISODate("2014-03-01T09:00:00Z") }
  { "_id" : 3, "item" : "xyz", "price" : 5, "quantity" : 10, "date" : ISODate("2014-03-15T09:00:00Z") }
  { "_id" : 4, "item" : "xyz", "price" : 5, "quantity" : 20, "date" : ISODate("2014-04-04T11:21:39.736Z") }
  { "_id" : 5, "item" : "abc", "price" : 10, "quantity" : 10, "date" : ISODate("2014-04-04T21:23:13.331Z") }
  rs001:PRIMARY>
  ```
  数据插入成功。
* 连接从节点，用respleUser进行权限校验及查看主节点刚刚插入的文档
```
PS C:\Users\YUANTINGZHONG> mongo.exe --port 27500
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27500/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("008f3a75-9724-49fe-b61a-01f316339331") }
MongoDB server version: 4.0.11
rs001:SECONDARY> use replSetTest
switched to db replSetTest
rs001:SECONDARY> db.auth("replUser","replUser")
1
rs001:SECONDARY> show collections;
2021-04-13T20:32:57.214+0800 E QUERY    [js] Error: listCollections failed: {
        "operationTime" : Timestamp(1618317172, 1),
        "ok" : 0,
        "errmsg" : "not master and slaveOk=false",
        "code" : 13435,
        "codeName" : "NotMasterNoSlaveOk",
        "$clusterTime" : {
                "clusterTime" : Timestamp(1618317172, 1),
                "signature" : {
                        "hash" : BinData(0,"j+O/Vpxwtzk+M9RGtC9VRFp8RPs="),
                        "keyId" : NumberLong("6950550638883241985")
                }
        }
} :
_getErrorWithCode@src/mongo/shell/utils.js:25:13
DB.prototype._getCollectionInfosCommand@src/mongo/shell/db.js:943:1
DB.prototype.getCollectionInfos@src/mongo/shell/db.js:993:20
DB.prototype.getCollectionNames@src/mongo/shell/db.js:1031:16
shellHelper.show@src/mongo/shell/utils.js:869:9
shellHelper@src/mongo/shell/utils.js:766:15
@(shellhelp2):1:1
rs001:SECONDARY>
rs001:SECONDARY> db.getMongo().setSlaveOk()
rs001:SECONDARY>
rs001:SECONDARY> show collections;
orders
rs001:SECONDARY> db.orders.find()
{ "_id" : 1, "item" : "abc", "price" : 10, "quantity" : 2, "date" : ISODate("2014-03-01T08:00:00Z") }
{ "_id" : 2, "item" : "jkl", "price" : 20, "quantity" : 1, "date" : ISODate("2014-03-01T09:00:00Z") }
{ "_id" : 3, "item" : "xyz", "price" : 5, "quantity" : 10, "date" : ISODate("2014-03-15T09:00:00Z") }
{ "_id" : 4, "item" : "xyz", "price" : 5, "quantity" : 20, "date" : ISODate("2014-04-04T11:21:39.736Z") }
{ "_id" : 5, "item" : "abc", "price" : 10, "quantity" : 10, "date" : ISODate("2014-04-04T21:23:13.331Z") }
rs001:SECONDARY>
```
可见，主节点的操作（账号操作及数据操作等）都被精确的复制到了从节点，副本集工作正常。

## 4. 副本集读写分离
* 副本集中的从节点默认情况下是不可读的，所有的读写操作都是对主节点的。在某些情况下，可能会造成主节点的访
问压力较大，这个时候可以考虑读写分离，将读操作部分或全部转移到从节点上，从而减轻主节点的压力。
* 读写分离是在客户端实现的。
* 下面的演示都是针对上文中的副本集示例进行的，并且当前的副本集状态为:
```
rs001:PRIMARY> rs.status()
{
        "set" : "rs001",
        ... ... 
        "members" : [
                {
                        "_id" : 1,
                        "name" : "127.0.0.1:27500",
                        "health" : 1,
                        "state" : 2,
                        "stateStr" : "SECONDARY",
                        "uptime" : 5756,
                        "optime" : {
                                "ts" : Timestamp(1618318342, 1),
                                "t" : NumberLong(3)
                        },
                        "optimeDurable" : {
                                "ts" : Timestamp(1618318342, 1),
                                "t" : NumberLong(3)
                        },
                        "optimeDate" : ISODate("2021-04-13T12:52:22Z"),
                        "optimeDurableDate" : ISODate("2021-04-13T12:52:22Z"),
                        "lastHeartbeat" : ISODate("2021-04-13T12:52:24.597Z"),
                        "lastHeartbeatRecv" : ISODate("2021-04-13T12:52:25.533Z"),
                        "pingMs" : NumberLong(0),
                        "lastHeartbeatMessage" : "",
                        "syncingTo" : "127.0.0.1:27501",
                        "syncSourceHost" : "127.0.0.1:27501",
                        "syncSourceId" : 2,
                        "infoMessage" : "",
                        "configVersion" : 1
                },
                {
                        "_id" : 2,
                        "name" : "127.0.0.1:27501",
                        "health" : 1,
                        "state" : 1,
                        "stateStr" : "PRIMARY",
                        "uptime" : 5760,
                        "optime" : {
                                "ts" : Timestamp(1618318342, 1),
                                "t" : NumberLong(3)
                        },
                        "optimeDate" : ISODate("2021-04-13T12:52:22Z"),
                        "syncingTo" : "",
                        "syncSourceHost" : "",
                        "syncSourceId" : -1,
                        "infoMessage" : "",
                        "electionTime" : Timestamp(1618312599, 1),
                        "electionDate" : ISODate("2021-04-13T11:16:39Z"),
                        "configVersion" : 1,
                        "self" : true,
                        "lastHeartbeatMessage" : ""
                },
                {
                        "_id" : 3,
                        "name" : "127.0.0.1:27502",
                        "health" : 1,
                        "state" : 7,
                        "stateStr" : "ARBITER",
                        "uptime" : 2116,
                        "lastHeartbeat" : ISODate("2021-04-13T12:52:24.481Z"),
                        "lastHeartbeatRecv" : ISODate("2021-04-13T12:52:23.994Z"),
                        "pingMs" : NumberLong(0),
                        "lastHeartbeatMessage" : "",
                        "syncingTo" : "",
                        "syncSourceHost" : "",
                        "syncSourceId" : -1,
                        "infoMessage" : "",
                        "configVersion" : 1
                }
        ],
        "ok" : 1,
        ... ...
}
rs001:PRIMARY>
```
### 4.1 用mongo客户端实现读写分离
操作如下
```
PS C:\Users\YUANTINGZHONG> mongo mongodb://replUser:replUser@localhost:27500,localhost:27501/replSetTest?replicaSet=rs001
MongoDB shell version v4.0.11
connecting to: mongodb://localhost:27500,localhost:27501/replSetTest?gssapiServiceName=mongodb&replicaSet=rs001
2021-04-13T20:49:05.368+0800 I NETWORK  [js] Starting new replica set monitor for rs001/localhost:27500,localhost:27501
2021-04-13T20:49:05.420+0800 I NETWORK  [js] Successfully connected to localhost:27500 (1 connections now open to localhost:27500 with a 5 second timeout)
2021-04-13T20:49:05.422+0800 I NETWORK  [ReplicaSetMonitor-TaskExecutor] Successfully connected to localhost:27501 (1 connections now open to localhost:27501 with a 5 second timeout)
2021-04-13T20:49:05.423+0800 I NETWORK  [js] Successfully connected to 127.0.0.1:27501 (1 connections now open to 127.0.0.1:27501 with a 5 second timeout)
2021-04-13T20:49:05.423+0800 I NETWORK  [ReplicaSetMonitor-TaskExecutor] changing hosts to rs001/127.0.0.1:27500,127.0.0.1:27501 from rs001/localhost:27500,localhost:27501
2021-04-13T20:49:05.424+0800 I NETWORK  [ReplicaSetMonitor-TaskExecutor] Successfully connected to 127.0.0.1:27500 (1 connections now open to 127.0.0.1:27500 with a 5 second timeout)
Implicit session: session { "id" : UUID("5a8024b2-545b-4cf3-bc72-dd786e76470a") }
MongoDB server version: 4.0.11
rs001:PRIMARY>
rs001:PRIMARY> db.getMongo().setReadPref("secondary")
rs001:PRIMARY> db.orders.find()
2021-04-13T20:50:03.018+0800 I NETWORK  [js] Successfully connected to 127.0.0.1:27500 (1 connections now open to 127.0.0.1:27500 with a 0 second timeout)
{ "_id" : 1, "item" : "abc", "price" : 10, "quantity" : 2, "date" : ISODate("2014-03-01T08:00:00Z") }
{ "_id" : 2, "item" : "jkl", "price" : 20, "quantity" : 1, "date" : ISODate("2014-03-01T09:00:00Z") }
{ "_id" : 3, "item" : "xyz", "price" : 5, "quantity" : 10, "date" : ISODate("2014-03-15T09:00:00Z") }
{ "_id" : 4, "item" : "xyz", "price" : 5, "quantity" : 20, "date" : ISODate("2014-04-04T11:21:39.736Z") }
{ "_id" : 5, "item" : "abc", "price" : 10, "quantity" : 10, "date" : ISODate("2014-04-04T21:23:13.331Z") }
rs001:PRIMARY> db.orders.find()
{ "_id" : 1, "item" : "abc", "price" : 10, "quantity" : 2, "date" : ISODate("2014-03-01T08:00:00Z") }
{ "_id" : 2, "item" : "jkl", "price" : 20, "quantity" : 1, "date" : ISODate("2014-03-01T09:00:00Z") }
{ "_id" : 3, "item" : "xyz", "price" : 5, "quantity" : 10, "date" : ISODate("2014-03-15T09:00:00Z") }
{ "_id" : 4, "item" : "xyz", "price" : 5, "quantity" : 20, "date" : ISODate("2014-04-04T11:21:39.736Z") }
{ "_id" : 5, "item" : "abc", "price" : 10, "quantity" : 10, "date" : ISODate("2014-04-04T21:23:13.331Z") }
rs001:PRIMARY>
```
要点：
* 连接参数中要携带参数，表名连接一个复制集，否则只是连接一个具体的mongodb服务节点，即
```
mongodb://replUser:replUser@localhost:27500,localhost:27501/replSetTest?replicaSet=rs001
```
其中"replicaSet=rs001"不可少。
* 连接成功后，要设置读偏好
```
db.getMongo().setReadPref("secondary") //secondary表示只从从节点读取数据
```
* 注意上面的操作中，第一次查询时db.orders.find()，有一条打印
“Successfully connected to 127.0.0.1:27500 (1 connections now open to 127.0.0.1:27500 with a 0 second timeout)”
这是与从节点（127.0.0.1:27500）建立连接，进行读取操作。为了进一步证实，使用wireshark抓包，结果如下
![mongo_capture][image_5]
从抓包也可以看出，读操作确实是对从节点进行的，即实现了读写分离。

### 4.2 使用MongoDB Compass客户端实现读写分离
* 连接参数设置界面如下
![image-20220112142428746](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112142428746.png)
注意，上图红框标出的内容，Read Preference设置了Secondary。
* 使用compass查询到结果
![image-20220112142442505](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112142442505.png)
* 抓包
![image-20220112142453204](E:\study\studygo\Golang学习笔记\mongodb副本集.assets\image-20220112142453204.png)
从抓包中可以看到，确实是从从节点读取了数据，即实现了读写分离。

### 4.3 其它客户端
* 一般mongo客户端都会提供读写分离的功能，如有需要，根据使用的具体客户端，检查其实现读写分离的方法；
* 读偏好（Read Perferences）描述与副本集相关的读取操作的行为。这些参数允许你在连接字符串中基于每个连接指定读取偏好。
* MongoDB有5种ReadPreference模式：
  *  primary
  主节点，默认模式，读操作只在主节点，如果主节点不可用，报错或者抛出异常。
  *  primaryPreferred
  首选主节点，大多情况下读操作在主节点，如果主节点不可用，如故障转移，读操作在从节点。
  *  secondary
  从节点，读操作只在从节点， 如果从节点不可用，报错或者抛出异常。
  *  secondaryPreferred
  首选从节点，大多情况下读操作在从节点，特殊情况（如单主节点架构）读操作在主节点。
  *  nearest
  最邻近节点，读操作在最邻近的成员，可能是主节点或者从节点。
> 实际中要根据需要选择合理的读偏好。
