# 安装配置

- 安装后，在安装路径的data目录下新建db文件夹

- `mongod.cfg`中设置相关配置

  - 外部机器要连接的话，`bindIP`改成`0.0.0.0`
  - `mongo.keyfile`的权限只能是400才行

- 登录（不用账户密码登录的话无法使用命令）：`mongo --host 127.0.0.1 --port 27017 -u "用户名" -p "密码" --authenticationDatabase "admin"`

  - `mongod.cfg`中有如下配置的话才需要使用账户密码来登录

    ```
    security:
      authorization: enabled
    ```

  - 如果登录时没带账户密码，也可在连接后再认证

    ```json
    use admin
    db.auth("用户名","密码")
    ```

    

# 基本概念

> 基于分布式文件存储的数据库，C++编写，用来处理大量的文档。是介于关系型数据库和非关系型数据库中间的产品

- 数据库
- 集合
- 文档
- 属性：值可以是数组，可以是文档对象（BSON），称为内嵌文档

# 命令

> <>括中的内容代表可替换成其他内容，[]代表可选内容
>
> 集合命名时一般用复数

- 基本命令
  - `show dbs`：查看所有数据库（创建数据库后，若没有数据，是查不到此数据库的）
  - `show collections`：查看所有集合
  - `db`：查看当前使用的数据库
  - `use xxx`：切换到xxx数据库，没有的话自动新建一个xxx
  - `db.dropDatabaes()`删除数据库
  - `ObjectId()`：手动生成`_id`字段（或称为数据域）
    - 是MongoDB自动根据时间戳和机器码计算生成的，文档的唯一标识
    - 此属性可以自己指定，但最好不这样做
  - `typeof db.<collection>.findOne().<field>`
  - `db.<collection>.update({},{$rename:{old:"new"}},false,true)`
  - `load(e:\tsmp.sql)`或`load(e:\faults_360.js)`运行脚本文件
  
- 插入
  - `db.createCollection("xxxx")`
  
  - `db.<collection>.insert(json)`
  
    - 第二个参数`ordered: <boolean>`，是否按顺序写入，默认 true，按顺序写入
  
  - `db.<collection>.insert([json1,json2])`
  
  - `db.<collection>.insertOne()`插入单个文档，语义上更明确
  
  - `db.<collection>.insertMany()`插入多个文档，语义上更明确
  
  - 循环插入
  
    ```javascript
    # 不要循环insert，而要循环push后再insert
    var arr = []
    for (var i =1;i<=20000;i++) {
        arr.push({"number":i})
    }
    db.numbers.insert(arr)
    db.users.insert(
    	{
            name:"zhangsan",
            #hobby即是内嵌文档
            hobby:{
            	cities:["beijing","shanghai"],
            	movies:["sanguo","hero"]
        	}
        }
    )
    ```
  
- 查询
  - `db.<collection>.find()`或`db.<collection>.find({})`：查询所有文档
  
  - `db.<collection>.find({<field>:"<value>"})`
    - 查询所有符合此条件的文档，有多个条件时在大括号中用逗号隔开
    - 返回的是一个数组，可以根据数组下标获取数组中的文档
    
  - `db.<collection>.findOne({<field>:"<value>"})`
    - 查询符合此条件的第一个文档，返回的是一个文档对象
    
    - 可以直接用`.`操作符调用此文档对象的属性
    
    - 也有操作符：`$lt 、$lte、$lte、$gt 、$gte $gte、$ne 、$or 、$in 、$nin 、 $not 、$exists`
    
      ```javascript
      db.student.find({
          age:{$gt:50}
          })
      ```
    
  - `db.<collection>.find().count()`或`db.<collection>.find().length()`返回查询结果的数量，一般用`count()`
  
- 修改
  
  - `db.<collection>.update(查询条件,新对象)`
  
    - 默认情况下会使用新对象来完全替换旧对象（直接把第二个参数完全替换符合条件的文档了），所以要用操作符
  
    - 修改指定的属性，用修改操作符`$set`（修改）、`$unset`（删除）
  
    - 对于用`$set`修改的属性，如果集合中不存在，会新增；用`$unset`修改的属性，不管给的值是什么，都会删掉这个属性
  
    - 其他操作符：`$push`，在数组中新增值；`$addToSet`，在数组中新增不重复的值；`$inc`：增加已有键的值，键不存在则创建
  
    - 默认情况下只改符合条件的一个文档，还可以传第三个参数来修改多个
  
      ```javascript
      db.<collection>.update(<query>,
                             <update>,
                             {
                                 upsert:<boolean>,
                                 multi:<boolean>, # true则支持修改多个
                                 writeConcern:<document>,
                                 collation:<document>,
      						}
      )
      #示例
      db.stus.update(
          {"_id" : ObjectId("59c219689410bc1dbecc0709")},
          {$set:{
              gender:"男",
              address:"流沙河"
          }}    
      )
      ```
  
  - `db.<collection>.updateMany(查询条件,新对象)`：同时修改符合条件的多个文档
  
  - `db.<collection>.updateOne(查询条件,新对象)`：同时修改符合条件的一个文档
  
  - `db.<collection>.replaceOne(查询条件,新对象)`：替换符合条件的一个文档（没有replaceMany）
  
- 删除
  
  - `db.<collection>.drop()`：直接删除集合（数据库中没有集合时，数据库也没了）
  
  - `db.<collection>.remove(<query>)`
  
    - 根据条件删除文档，默认支持删除符合条件的所有文档
  
    - 第二个参数传true，可以只删一个
  
      ```javascript
      db.<collection>.remove(<query>,<boolean>)
      ```
  
    - remove必须传参，如果传的是`{}`，则删除所有文档，删完后集合还在（性能略差，建议用drop）
  
  - `db.<collection>.deleteOne(<query>)`：类似update
  
  - `db.<collection>.deleteMany(<query>)`：类似update
  
- 实际业务的删除

  > 一般设置一个isDel的标志位，查询时只把isDel为0的结果呈现出来

- 内嵌文档

  - 属性的值也是个文档对象，MongoDB支持通过内嵌文档来查找，但字段名必须用`''`或`""`括起来，否则报错

  - 查找时，查找的条件不一定是`=`关系，也可能是包含关系

    > 比如属性a的值是个文档对象b，b的值是个数组，查找时也可以通过条件查找出来
    
    ```java
    db.student.insert(
        {
            name:"猪八戒",
            age:20,
            locations:{
                provicens:["陕西","湖北"],
                cities:["西安","武汉"]
            }
        }
    )
    db.student.find({
        "cities.provicens":"陕西"
    })
    ```

- 分页

  > MongoDB会自动调整skip和limit调用的顺序，这两个的前后顺序不影响结果

  - `limit(n)`：只显示n条，为0的时候查所有
  - `skip(n)`：跳过前n条
  - 显示k条数据：`db.<collections>.find().skip((pagenumber-1)*pagesize).limit(pagesize)`

- 排序

  > skip、limit、sort可以以任意顺序调用

  - MongoDB默认按照`_id`属性升序
  - `sort({<field:1>})`：按照field的升序排序，降序是`-1`
  - 可以指定多个字段，和mysql一样，第一个字段相等时，再根据第二个字段排序

- 操作符

  > 使用操作符时必须用`{}`括起来
  >
  > 且的关系，用逗号连接两个条件；或的关系，用`$or`操作符，后跟数组，数组里是两个条件，分别用`{}`括起来

  - `$set`
  - `$unset`
  - `$push`
  - `$addToSet`
  - `$gt`
  - `$lt`
  - `$eq`：判断等于时一般直接find
  - `$gte`
  - `$lte`
  - `$ne`
  - `$inc:n`值自增（自减时`$inc:-n`)

# 额外参数总结

- find：第二个参数：指定显示的字段`db.student.find({},{name:1,_id:0})`
- update：第三个参数：{}括起来的update的一些属性，如multi:true则支持修改多个文档
- insert：第二个参数：是否按序插入，默认true
- remove：第二个参数：置为true的话则一次只删符合条件的一个文档

# 其他知识点

- 文档之间的关系

  - 一对一

    - MongoDB中通过内嵌文档的方式来体现（一个属性的值是一个非数组的文档对象，此文档和它内嵌的文档就是一对一）
    - 实际开发中场景不多

  - 一对多/多对一：使用场景最多

    - 方式一：通过内嵌文档，内嵌一个数组文档对象

    - 方式二：“多”这个集合中设置一个属性来标识“一”这个集合，类似mysql的”一对多，两张表，多的表加外键“

      > 查找时用findOne找到“一”这个集合中满足条件的文档，再通过`.`得到此文档的标识属性，再把这个结果定义成变量，再从“多”这个集合中根据这个变量去查找

  - 多对多

    - 类似一对多的方式二，只不过把那个用来标识的属性换成数组

- 投影

  > 查询的结果中想显示的字段

  - `find({},{<field1>:1,<field2>:1})`：只想显示field1和field2的信息
  - 默认`_id`会显示出来，不想看的话强制设为0

```javascript
// 根据ID查询
db.<collections>.find({"_id":ObjectId("612f161091c1255d1ec93fdc")})

// 嵌套查询
db.<collections>.find({"user.name":"张三"})

// 分组查询

// 批量操作
db.<collections>.bulkWrite([
{updateOne :{"filter": {"_id":ObjectId("612f161091c1255d1ec93fdc")},"update":{$set:{phone:2}}}},
{updateOne :{"filter": {"_id":ObjectId("612f161091c1255d1ec93fdd")},"update":{$set:{phone:2}}}},
{updateOne :{"filter": {"_id":ObjectId("612f161091c1255d1ec93fde")},"update":{$set:{phone:2}}}},
{updateOne :{"filter": {"_id":ObjectId("612f161091c1255d1ec93fde")},"update":{$set:{phone:2}}}},
])
```


# 聚合查询

- > match：用于对数据进行筛选

  - `db.<collections>.aggregate({"$match":{"字段":"条件"}})`

- > project：将一个数据结果映射为另一个结果，过程中可以对某些数据进行修改（修改查询的结果而不是修改数据） ，控制其最终显示的结果

  - `db.<collections>.aggregate({"$project":{"要保留的字段名":1,"要去掉的字段名":0,"查询结果中新增的字段名":"表达式"}})`

  - ```bash
    表达式之数学表达式，参加运算的字段不能被隐藏
    
    {"$add":[expr1,expr2,...,exprN]} #相加
    {"$subtract":[expr1,expr2]} #第一个减第二个
    {"$multiply":[expr1,expr2,...,exprN]} #相乘
    {"$divide":[expr1,expr2]} #第一个表达式除以第二个表达式的商作为结果
    {"$mod":[expr1,expr2]} #第一个表达式除以第二个表达式得到的余数作为结果
    ```

  - ```bash
    表达式之日期表达式:$year,$month,$week,$dayOfMonth,$dayOfWeek,$dayOfYear,$hour,$minute,$second
    
    #例如查看每个员工的工作多长时间
    db.emp.aggregate(
        {"$project":{"name":1,"hire_period":{
            "$subtract":[
                {"$year":new Date()},
                {"$year":"$hire_date"}
            ]
        }}}
    )
    ```

  - ```bash
    字符串表达式
    
    {"$substr":[字符串/$值为字符串的字段名,起始位置,截取几个字节]}
    {"$concat":[expr1,expr2,...,exprN]} #指定的表达式或字符串连接在一起返回,只支持字符串拼接
    {"$toLower":expr}
    {"$toUpper":expr}
    ```

  - ```bash
    逻辑表达式
    
    $and
    $or
    $not
    ```

- > group用于分组

  - 单个字段分组：`db.<collections>.aggregate({"$group":{_id:"$字段"}})`

  - 多个字段联合分组：`db.<collections>.aggregate({"$group":{_id:{字段1:"$字段1",字段2:"$字段2"}}})`

    ```bash
    聚合函数
    $sum、$avg、$max、$min、$first、$last
    
    数组操作符
    {"$addToSet":expr}：不重复
    {"$push":expr}：重复
    ```

- > sort，skip，limit，用法和普通查询的用法一样，但是分页时顺序要严格按照sort、skip、limit的顺序

# 权限

> 在任何数据库都可创建用户，创建后此`db.getUsers()`就是查的当前数据库的用户

- 添加admin用户：`db.createUser({user:"admin",pwd:"admin",roles:[{ role: "userAdminAnyDatabase", db: "admin" }]})`

- 添加root用户：`db.createUser({user:"root",pwd:"root",roles:["root"]})`

- 给用户添加权限：`db.grantRolesToUser( "admin" , [{ "role": "clusterAdmin", "db": "admin" }])`

- 查看所有用户：`db.getUsers()`

- 权限认证：`db.auth("admin","admin")`

- 其他命令

  ![image-20211207165904316](C:\Users\ZHOUDONGBIN\AppData\Roaming\Typora\typora-user-images\image-20211207165904316.png)

  ![image-20211207165912263](C:\Users\ZHOUDONGBIN\AppData\Roaming\Typora\typora-user-images\image-20211207165912263.png)

# Go MongoDB Driver

- 连接mongo

  ```go
  func main() {
     // options一般都要加上retryWrites=false，否则会报错：this MongoDB deployment does not support retryable writes
      host := "mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]"
      clientOptions := options.Client().ApplyURI(host)
      client, err := mongo.Connect(context.TODO(), clientOptions)
      err = client.Ping(context.TODO(), nil)
      collection := client.Database("db").Collection("collection") 
  }
  ```

  

- 创建索引

  ```go
  func Start() error {
      opts := options.CreateIndexes()
      // 复合索引
      pidItIndex := mongo.IndexModel{Keys: bsonx.Doc{
          {Key: "pd", Value: bsonx.Int32(1)},
          {Key: "it", Value: bsonx.Int32(-1)},
      }}
      
      // 唯一索引
      tnIndex := mongo.IndexModel{
          Keys:    bsonx.Doc{{Key: "tn", Value: bsonx.Int32(1)}},
          Options: options.Index().SetUnique(true),
      }
  
      // 创建单个索引
      collection.Indexes().CreateOne(context.Background(), pidItIndex, opts)
      
      // 创建多个索引
     	collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{trIndex, pidItIndex}, opts)
  }
  ```

- 聚合查询

  ```go
  // todo 待更新
  ```

- 备份欢迎

  ```go
  // 使用-q的查询语句时，必须有-c指定集合名称；分片集群不能有oplog参数
  mongodump -h ip:port -u root -p 123456 -o /home/zhoudongbin -d tsmp --authenticationDatabase admin -c train_records -q {_id:123} --oplog gzip
  ```

# 索引


  ```shell
  4.1 概述
  索引支持在MongoDB中高效地执行查询.如果没有索引, MongoDB必须执行全集合扫描, 即扫描集合中的每个文档, 以选择与查询语句匹配的文档.这种扫描全集合的查询效率是非常低的, 特别在处理大量的数据时, 查询可以要花费几十秒甚至几分钟, 这对网站的性能是非常致命的.
  
  如果查询存在适当的索引, MongoDB可以使用该索引限制必须检查的文档数.
  
  索引是特殊的数据结构, 它以易于遍历的形式存储集合数据集的一小部分.索引存储特定字段或一组字段的值, 按字段值排序.索引项的排序支持有效的相等匹配和基于范围的查询操作.此外, MongoDB还可以使用索引中的排序返回排序结果.
  
  MongoDB 使用的是 B Tree, MySQL 使用的是 B+ Tree
  
  # create index，默认是_id为索引
  db.<collection_name>.createIndex({ userid : 1, username : -1 })	#创建userid升序和username降序的复合索引
  
  # retrieve indexes
  db.<collection_name>.getIndexes()
  
  # remove indexes
  db.<collection_name>.dropIndex(index)
  
  # there are 2 ways to remove indexes:
  # 1. removed based on the index name
  # 2. removed based on the fields
  
  db.<collection_name>.dropIndex( "userid_1_username_-1" )
  db.<collection_name>.dropIndex({ userid : 1, username : -1 })
  
  # remove all the indexes, will only remove non_id indexes
  db.<collection_name>.dropIndexes()
  4.2 索引的类型
  4.2.1 单字段索引
  MongoDB支持在文档的单个字段上创建用户定义的升序/降序索引, 称为单字段索引 Single Field Index
  
  对于单个字段索引和排序操作, 索引键的排序顺序（即升序或降序）并不重要, 因为 MongoDB 可以在任何方向上遍历索引.
  
  
  
  4.2.2 复合索引
  MongoDB 还支持多个字段的用户定义索引, 即复合索引 Compound Index
  
  复合索引中列出的字段顺序具有重要意义.例如, 如果复合索引由 { userid: 1, score: -1 } 组成, 则索引首先按 userid 正序排序, 然后 在每个 userid 的值内, 再在按 score 倒序排序.
  
  
  
  4.2.3 其他索引
  地理空间索引 Geospatial Index
  文本索引 Text Indexes
  哈希索引 Hashed Indexes
  
  地理空间索引（Geospatial Index）
  为了支持对地理空间坐标数据的有效查询, MongoDB 提供了两种特殊的索引: 返回结果时使用平面几何的二维索引和返回结果时使用球面几何的二维球面索引.
  
  文本索引（Text Indexes）
  MongoDB 提供了一种文本索引类型, 支持在集合中搜索字符串内容.这些文本索引不存储特定于语言的停止词（例如 “the”, “a”, “or”）, 而将集合中的词作为词干, 只存储根词.
  
  哈希索引（Hashed Indexes）
  为了支持基于散列的分片, MongoDB 提供了散列索引类型, 它对字段值的散列进行索引.这些索引在其范围内的值分布更加随机, 但只支持相等匹配, 不支持基于范围的查询.
  
  4.3 索引的管理操作
  4.3.1 索引的查看
  语法
  
  db.collection.getIndexes()
  默认 _id 索引： MongoDB 在创建集合的过程中, 在 _id 字段上创建一个唯一的索引, 默认名字为 _id , 该索引可防止客户端插入两个具有相同值的文 档, 不能在 _id 字段上删除此索引.
  
  注意：该索引是唯一索引, 因此值不能重复, 即 _id 值不能重复的.
  
  在分片集群中, 通常使用 _id 作为片键.
  
  4.3.2 索引的创建
  语法
  
  db.collection.createIndex(keys, options)
  options（更多选项）列表：放在后面了
  
  举个🌰
  
  $  db.comment.createIndex({userid:1})
  
  $ db.comment.createIndex({userid:1,nickname:-1})
  ...
  
  4.3.3 索引的删除
  语法
  
  # 删除某一个索引
  $ db.collection.dropIndex(index)
  
  # 删除全部索引
  $ db.collection.dropIndexes()
  提示:
  
  _id 的字段的索引是无法删除的, 只能删除非 _id 字段的索引
  
  示例
  
  # 删除 comment 集合中 userid 字段上的升序索引
  $ db.comment.dropIndex({userid:1})
  4.4 索引使用
  4.4.1 执行计划
  分析查询性能 (Analyze Query Performance) 通常使用执行计划 (解释计划 - Explain Plan) 来查看查询的情况
  
  $ db.<collection_name>.find( query, options ).explain(options) #explain可以传参""
  比如: 查看根据 user_id 查询数据的情况
  
  未添加索引之前
  
  "stage" : "COLLSCAN", 表示全集合扫描
  
  
  添加索引之后
  
  "stage" : "IXSCAN", 基于索引的扫描
  

4.4.2 涵盖的查询
  当查询条件和查询的投影仅包含索引字段是, MongoDB 直接从索引返回结果, 而不扫描任何文档或将 带入内存, 这些覆盖的查询十分有效

  https:#docs.mongodb.com/manual/core/query-optimization/#covered-query
  ```

| Parameter              | Type          | Description                                                  |
| :--------------------- | :------------ | :----------------------------------------------------------- |
| background             | Boolean       | 建索引过程会阻塞其它数据库操作，background可指定以后台方式创建索引，即增加 "background" 可选参数。 "background" 默认值为**false**。 |
| **unique**（比较重要） | **Boolean**   | **建立的索引是否唯一。指定为true创建唯一索引。默认值为false.** |
| name                   | string        | 索引的名称。如果未指定，MongoDB的通过连接索引的字段名和排序顺序生成一个索引名称。 |
| sparse                 | Boolean       | 对文档中不存在的字段数据不启用索引；这个参数需要特别注意，如果设置为true的话，在索引字段中不会查询出不包含对应字段的文档.。默认值为 **false**. |
| expireAfterSeconds     | integer       | 指定一个以秒为单位的数值，完成 TTL设定，设定集合的生存时间。 |
| v                      | index version | 索引的版本号。默认的索引版本取决于mongod创建索引时运行的版本。 |
| weights                | document      | 索引权重值，数值在 1 到 99,999 之间，表示该索引相对于其他索引字段的得分权重。 |
| default_language       | string        | 对于文本索引，该参数决定了停用词及词干和词器的规则的列表。 默认为英语 |
| language_override      | string        | 对于文本索引，该参数指定了包含在文档中的字段名，语言覆盖默认的language，默认值为 language. |

# 优化

## explain

> explain有三种模式，分别是：queryPlanner、executionStats、allPlansExecution。现实开发中，常用的是executionStats模式

- `db.getCollection('person').find({"age":{"$lte":2000}}).explain("executionStats")`

  ```bash
  对queryPlanner分析
  
      queryPlanner: queryPlanner的返回
  
      queryPlanner.namespace:该值返回的是该query所查询的表
  
      queryPlanner.indexFilterSet:针对该query是否有indexfilter
  
      queryPlanner.winningPlan:查询优化器针对该query所返回的最优执行计划的详细内容。
  
      queryPlanner.winningPlan.stage:最优执行计划的stage，这里返回是FETCH，可以理解为通过返回的index位置去检索具体的文档（stage有数个模式，将在后文中进行详解）。
  
      queryPlanner.winningPlan.inputStage:用来描述子stage，并且为其父stage提供文档和索引关键字。
  
      queryPlanner.winningPlan.stage的child stage，此处是IXSCAN，表示进行的是index scanning。
  
      queryPlanner.winningPlan.keyPattern:所扫描的index内容，此处是did:1,status:1,modify_time: -1与scid : 1
  
      queryPlanner.winningPlan.indexName：winning plan所选用的index。
  
      queryPlanner.winningPlan.isMultiKey是否是Multikey，此处返回是false，如果索引建立在array上，此处将是true。
  
      queryPlanner.winningPlan.direction：此query的查询顺序，此处是forward，如果用了.sort({modify_time:-1})将显示backward。
  
      queryPlanner.winningPlan.indexBounds:winningplan所扫描的索引范围,如果没有制定范围就是[MaxKey, MinKey]，这主要是直接定位到mongodb的chunck中去查找数据，加快数据读取。
  
      queryPlanner.rejectedPlans：其他执行计划（非最优而被查询优化器reject的）的详细返回，其中具体信息与winningPlan的返回中意义相同，故不在此赘述。
  
  对executionStats返回逐层分析
  
      第一层，executionTimeMillis
  
      最为直观explain返回值是executionTimeMillis值，指的是我们这条语句的执行时间，这个值当然是希望越少越好。
  
      其中有3个executionTimeMillis，分别是：
  
      executionStats.executionTimeMillis
  
      该query的整体查询时间。
  
      executionStats.executionStages.executionTimeMillisEstimate
  
      该查询根据index去检索document获得2001条数据的时间。
  
      executionStats.executionStages.inputStage.executionTimeMillisEstimate
  
      该查询扫描2001行index所用时间。
  
      第二层，index与document扫描数与查询返回条目数
  
      这个主要讨论3个返回项，nReturned、totalKeysExamined、totalDocsExamined，分别代表该条查询返回的条目、索引扫描条目、文档扫描条目。
  
      这些都是直观地影响到executionTimeMillis，我们需要扫描的越少速度越快。
  
      对于一个查询，我们最理想的状态是：
      nReturned=totalKeysExamined 且 totalDocsExamined=0（cover index，仅仅使用到了index，无需文档扫描，这是最理想状态）
  
      nReturned=totalKeysExamined=totalDocsExamined（正常index利用，无多余index扫描与文档扫描）
      如果有sort的时候，为了使得sort不在内存中进行，我们可以在保证nReturned=totalDocsExamined 的基础上，totalKeysExamined可以大于totalDocsExamined与nReturned，因为量级较大的时候内存排序非常消耗性能。
      
      第三层，stage状态分析
  
      那么又是什么影响到了totalKeysExamined和totalDocsExamined？是stage的类型。类型列举如下：
  
      COLLSCAN：全表扫描
  
      IXSCAN：索引扫描
  
      FETCH：根据索引去检索指定document
  
      SHARD_MERGE：将各个分片返回数据进行merge
  
      SORT：表明在内存中进行了排序
  
      LIMIT：使用limit限制返回数
  
      SKIP：使用skip进行跳过
  
      IDHACK：针对_id进行查询
  
      SHARDING_FILTER：通过mongos对分片数据进行查询
  
      COUNT：利用db.coll.explain().count()之类进行count运算
  
      COUNTSCAN：count不使用Index进行count时的stage返回
  
      COUNT_SCAN：count使用了Index进行count时的stage返回
  
      SUBPLA：未使用到索引的$or查询的stage返回
  
      TEXT：使用全文索引进行查询时候的stage返回
  
      PROJECTION：限定返回字段时候stage的返回
  
      对于普通查询，我希望看到stage的组合(查询的时候尽可能用上索引)：
  
      Fetch+IDHACK
  
      Fetch+ixscan
  
      Limit+（Fetch+ixscan）
  
      PROJECTION+ixscan
  
      SHARDING_FITER+ixscan
  
      COUNT_SCAN
  
      不希望看到包含如下的stage：
  
      COLLSCAN(全表扫描),SORT(使用sort但是无index),不合理的SKIP,SUBPLA(未用到index的$or),COUNTSCAN(不使用index进行count)
      
      executionStages.Stage为Sort，在内存中进行排序了，这个在生产环境中尤其是在数据量较大的时候，是非常消耗性能的，这个千万不能忽视了，我们需要改进这个点。
  ```
  

## 分片

> 详细参加《mongodb分片.md》

概述：

- 多个Router(mongos)，一个Config Server（配成副本集），多个Shard（每个Shard都是副本集）
  - Router(mongos)：区别于mongod（mongodb的守护进程）。应用层读写数据的唯一接口，把从Config Server读来的元数据缓存后，将请求路由到对应的分片。mongos要么是向所有分片广播查询，要么将包含片键或复核片键前缀的查询路由到特定的分片
  - Config Server：保存分片的元数据、鉴权，管理分布式锁。鉴权除了用户名和密码，还要使用秘钥文件或x.509证书进行集群内部各个节点成员之间的认证（openssh生成的keyfile）。
  - Shard：一个Shard包含多个Chunk。Chunk大小1~1024MB，默认64MB。初始的时候数据库只有一个primary shard，数据量超过chunk大小后发生chunk分裂。chunk间数据量不均衡时（最多的和最少的差值超过阈值），发生chunk迁移（从primary shard迁移到其他shard）
- 分片策略：范围分片，哈希分片
- Zone：是基于特定tag集的一组分片。将特定数据隔离到特定的分片，或应用程序在不同地理位置上使用且希望查询路由到最近的分片进行读写，则分区可能很有用

知识点

- piority越大，优先级越高；为0时不参与primary的竞争；不填的话默认为1
- add时如果新节点的priority大于已有节点的priority，则新节点自动成为primary，否则自动成为secondary
- rs.add()时，vote不填的话默认为1，可手动设为0即无投票权
- 生产环境添加节点时，建议将priority及votes设为0，添加完成后用rs.conf()查看配置，再用rs.recofig()命令进行相应调整
- 通过rs.add()添加低优先级节点时，不会改变高优先级节点的地位

| 方法名                                     | 描述                                     |
| ------------------------------------------ | ---------------------------------------- |
| sh.addShard(host)                          | #将分片添加到分片集群中,host（ip：端口） |
| sh.status()                                | #分片群集的状态                          |
| sh.enableSharding(““)                      | #配置需要分片的数据库                    |
| sh.shardCollection(“.“, shard-key-pattern) | #shards the collection                   |

