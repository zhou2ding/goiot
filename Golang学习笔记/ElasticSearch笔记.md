# 前言

- ES概述

  > ELasticSearch是基于Lucene做了一些封装和增强，是一个开源的高扩展的分布式全文检索引擎，可以实时的存储、检索数据
  >
  > 通过简单的RESTful API来因此Lucene的复杂性，有全文检索、高亮、搜索推荐等功能

- ELK概述

  > ElasticSearch、Logstash、Kibana的合称，是一个日志分析架构技术栈的总称
  >
  > Logstash是ELK的中央数据流引擎，用于从不同目标（文件/数据存储/MQ）收集不同的格式数据，过滤后输出到不同目的地（文件/MQ/Redis/es/kafka等）
  >
  > Kibana将ES的数据进行页面展示和实时分析

# 使用

- 目录

  - bin，启动文件
  - config，配置文件
    - `log4j2.properties`：日志配置
    - `jvm.options`：java虚拟机相关配置（内存、cpu等配置）
    - `elasticsearch.yml`：ES的配置
      - 开启跨域访问并允许所有人访问：`http.cors.enabled: true`，`http.cors.allow-origin:"*"`
  - lib，相关jar包
  - modules，功能模块
  - logs，日志
  - plugins，插件（IK分词器等）

- 启动

  1. `elasticsearch.bat`
  2. 访问`127.0.0.1:9200`：ES默认自成集群，即使没配集群，单个也是集群
  3. 修改kibana配置`i18n.locale: "zh-CN"`和`elasticsearch.hosts: ["http://localhost:9200"]`
  4. 启动kibana，然后访问`127.0.0.1:5601`

- 概述

  > ES是面向文档的数据库，最小数据就是文档，一切皆是json
  >
  > 物理设计：ES默认在后台把每个索引划分成多个分片，每个分片可以在集群中的不同服务器间迁移（只有一个的话本身就是集群）
  >
  > 逻辑设计：一个索引中包含多个文档，查找顺序：索引--->类型--->文档ID（字符串）

  | 关系型数据库       | ES                  |
  | ------------------ | ------------------- |
  | 数据库（database） | 索引（indices）     |
  | 表（tables）       | types（未来被弃用） |
  | 行（rows）         | documents           |
  | 字段（colums）     | fields              |

  - 文档：一条记录，包含字段和对应的值（key-value），可以再包含文档对象

  - 类型：文档的逻辑容器，就像表是行的容器一样。类型对字段的定义称为映射，比如name映射为字符串类型

  - 索引：映射类型的容器，是个非常大的文档集合。索引存储了映射类型的字段和其他设置，然后它们被存储到各分片上

  - 分片（shard）

    > ES中所有数据的文件块，也是数据的最小单元块
    >
    > ES中所有数据均衡的存储在集群中各个节点的分片中，默认5个分片
    >
    > 分为主分片(primary shard)和副本分片(replica shard)。主分片负责写和读，而从分片负责读和备份。

  - 倒排索引：对要搜索的数据进行分词，然后搜索哪些索引包含了要搜索的词

    >  区别于正向索引的在文章中找单词，倒排索引是通过单词找到和这些单词有关的文章，再根据哪些文章包含的单词更多划分权重，最后通过权重找到结果

  - 数据类型

    | 字符串 | text,keyword（keyword类型的字符串不会被分词器解析）          |
    | ------ | ------------------------------------------------------------ |
    | 数值   | long,integer,short,byte,double,float,half float,scaled float |
    | 日期   | date                                                         |
    | 布尔   | boolean                                                      |
    | 二进制 | binary                                                       |
    | 数组   | []                                                           |

- RESTful风格

  > 一种软件架构风格，提供了设计原则和约束条件。用于客户端和服务器交互类的软件，基于此风格设计的软件可以更简洁，更有层次，更易于实现缓存等机制。

  | method | url地址                                         | 描述                   |
  | ------ | ----------------------------------------------- | ---------------------- |
  | PUT    | localhost:9200/索引名称/类型名称/文档id         | 创建文档（指定文档id） |
  | POST   | localhost:9200/索引名称/类型名称                | 创建文档（随机文档id） |
  | POST   | localhost:9200/索引名称/类型名称/文档id/_update | 修改文档               |
  | DELECT | localhost:9200/索引名称/类型名称/文档id         | 删除文档               |
  | GET    | localhost:9200/索引名称/类型名称/文档id         | 通过文档id查询文档     |
  | POST   | localhost:9200/索引名称/类型名称/_search        | 查询所有数据           |

- 使用方法

  - `elasticsearch-plugin list`查看已安装插件

  - ik分词器，有`ik_smart`和`ik_max-word`两种分词算法，分别为最少切分和最细粒度切分（穷尽字典词库的可能）

    - json命令

      ```json
      GET _analyze
      {
        "analyzer": "ik_max_word",	//或ik_smart，keyword，standard
        "text": "早上好，苦逼的打工人"
      }
      ```

    - 增加自己的字典

      1. 新建`my.dic`字典，添加自己的词语
      2. `IKAnalyzer.cfg.xml`中添加字典`<entry key="ext_dict>my.dic</entry>"`
      3. 重启ES和kibana

  - CRUD

    ```json
    //创建文档，没有提前创建规则的话，会默认指定类型
    PUT /test1/type1/1	//重复put的话，会更新这个文档的版本，而不是生成新的数据；POST随机id的话不会这样
    {
      "name":"张三",
      "age":3,
      "birth":"1997-04-17"
    }
    //创建规则，创建结果是一个空的索引，指定了文档的数据类型
    PUT /test2
    {
      "mappings": {
        "properties": {
          "name":{
            "type": "text"
          },
          "age":{
            "type": "long"
          },
          "birth":{
            "type": "date"
          }
        }
      }
    }
    ```

    ```json
    GET /test1/type1/1							//不能只查到type，要么只到索引，要么就到id
    GET /test2									//查看索引的基本信息
    
    POST /test1/type1/_search					//查看所有文档
    GET /test1/type1/_search?q=name:zhangsan	//精确匹配查找文档，此方法不常用，详见复杂查询
    
    GET _cat/health								//查看健康状态
    GET _cat/indices?v							//查看索引信息
    ```

    ```json
    POST /test1/type1/1/_update	//PUT也能更新，但是现在都推荐此种方法
    {
      "doc":{					//"doc"是固定的
        "name":"王五"				//字段可以少，用PUT更新的话，少字段时就为空了
      }
    }
    ```

    ```json
    DELETE /test1			//删除索引
    DELETE /test1/type1/1	//删除文档
    ```

  - 复杂查询

    ```json
    //模糊匹配
    GET /test1/type1/_search
    {
      "query":{
        "match": {					//match先用分词器解析，再用分析的结果去查询；效率比精确查询的term低
          "name": "张"				//单个字段是数组时，如果有多个查询条件时用空格隔开，只要满足其中一个									//就会被查出来，再根据分数排序/筛选
        }
      },
      _source:["name","age"],	//结果过滤，指定只显示哪些字段
      "sort": [					//排序，给了固定的排序后，score就为null了
          {
              "age":{
                  "order":"desc"
              }
          }
      ]
      "from":0,
      "size":2					//分页，from和size就是mysql中limit的两个参数
    }
    ```

    > 查询结果
    >
    > ![image-20210506164404598](D:\资料\Go\src\studygo\Golang学习笔记\ElasticSearch笔记.assets\image-20210506164404598.png)

    > must，mysql中的and（用两个match会报错，必须用must）
    >
    > should，mysql中的or
    >
    > must_not，即not，反向操作
    >
    > ![image-20210506200756924](D:\资料\Go\src\studygo\Golang学习笔记\ElasticSearch笔记.assets\image-20210506200756924.png)

    > 筛选，gt gte lt lte eq，类似MongoDB
    >
    > ![image-20210506201238919](D:\资料\Go\src\studygo\Golang学习笔记\ElasticSearch笔记.assets\image-20210506201238919.png)

    > 精确查询：term查询，是通过倒排索引进行精确查找的
    >
    > 精确查询是把查询的值和分词的结果精确匹配的，详见试验结论

  - 高亮查询

    ![image-20210506215025513](D:\资料\Go\src\studygo\Golang学习笔记\ElasticSearch笔记.assets\image-20210506215025513.png)

    ![image-20210506215236279](D:\资料\Go\src\studygo\Golang学习笔记\ElasticSearch笔记.assets\image-20210506215236279.png)

# 试验结论

- match
  - text类型的随便一个字或单词都能匹配到
  - keyword类型的，完全一致才能匹配到
- term
  - text类型的，只有单个字或单个单词才能匹配到
  - keyword类型的，完全一致才能匹配到