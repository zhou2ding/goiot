# day06

## 数据库

- 常见数据库`SQLlite`、`MySQL`、`SQLServer`、`Oracle`、`postgreSQL`

- 不同数据库的占位符不同

  |   数据库   |        占位符语法         |
  | :--------: | :-----------------------: |
  |   MySQL    |            `?`            |
  | PostgreSQL |       `$1`, `$2`等        |
  |   SQLite   |        `?` 和`$1`         |
  |   Oracle   | `:name`（name是字段名字） |

## MySQL

> 主流的关系型数据库

### 知识点

- SQL语句：`DQL`、`DML`、`DDL`、`DCL`、`TCL`
- 存储引擎：`InnoDB`、`MyISAM`，支持插件式的存储引擎

### database/sql包

- 原生支持连接池，是并发安全的

- 并没有具体的实现，只是列出了一些需要第三方库实现的具体内容

- 使用方法

  - `go get -u github.com/go-sql-driver/mysql`

  - `dsn := username:password@tcp(ipAddress)/database`

  - `sql.Open("msyql", dsn)`，返回一个`sql.DB`指针，不会校验用户名和密码

  - 在导入`mysql`的包时自动调用了其`init`方法，此方法向`database/sql`包中注册了`"mysql"`这个驱动

    ```go
    import (
    	"database/sql"
        
    	_ "github.com/go-sql-driver/mysql"
    )
    func main() {
    	dsn := "root:564710@tcp(localhost:3306)/bjpowernode"
    	_, err := sql.Open("mysql", dsn)
    	if err != nil {
    		fmt.Printf("open %s failed, error:%v\n", dsn, err)
    		return
    	}
    }
    ```

- `select SUBSTRING_INDEX(host,':',1) as ip , count(*) from information_schema.processlist group by ip;`获取本机连接mysql的IP地址

- `db.SetMaxOpenConns()`设置数据库连接池的最大建立连接的数量

- `db.SetMaxIdleConns()`设置数据库连接池的最大闲置连接数

#### 查询

- 单行查询

  1. 写查询单条记录的sql语句
  2. 执行查询(`QueryRow()`方法，接收一个字符串和可变长度的任意类型变量，返回`*sql.Row`对象)
  3. 拿到结果（用`Scan()`方法，且必须使用，因为此方法会释放数据库连接），不释放的话就会卡住，等待把连接归还给连接池

  ```go
  type user struct {
      id int
      name string
      age int
  }
  var u1 user
  sqlStr := `select id, name, age from user where id=?;` // ?是占位符
  db.QueryRow(sqlstr, 1).Scan(&u1.id, &u1.name, &u1.age)
  ```

- 多行查询

  1. SQL语句
  2. 执行`db.Query()`，返回`*db.Rows`对象和一个`error`
  3. `defer`关闭rows
  4. 循环取值`rows.Next()`

  ```go
  sqlStr := `select id, name, age from user where id > ?;`
  rows, _ := db.Query(sqlStr, 0)
  for rows.Next() {
      var u1 user
      _ = rows.Scan(&u1.id, &u1.name, &u1.age)
      fmt.Println(u1)
  }
  ```

#### 增删改

- `ret, err := db.Exec()`接收一个字符串和一个可变长度任意类型的变量，返回`db.Result`接口类型变量和`error`
  - 如果是插入操作，会拿到插入数据的`id`，`ret.LastInsertId()`，返回`id`和`error`
  - 如果是修改操作，会拿到受影响的行数
  - 如果是删除操作，会拿到受影响的行数

#### 预处理

>  好处

1. 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
2. 避免SQL注入问题。

>  执行过程

1. 把SQL语句分成两部分，命令部分与数据部分。
2. 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
3. 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
4. MySQL服务端执行完整的SQL语句并将结果返回给客户端。

>  使用

- `db.Prepare()`接收一个字符串，返回准备好的状态`*db.Stmt`和`error`
- `defer stmt.Cloes()`
- 调用`stmt`的`QueryRow()`、`Query()`、`Exec()`方法执行操作

#### 事务

- `db.Begin()`，无参，返回一个`db.Tx`（transaction）和一个`error`，后续执行sql语句就调动`tx`的`Query()`、`QueryRow()`、`Exec()`方法
- `db.Commit()`，无参，返回一个`error`
- `db.RollBack()`，无参，返回一个`error`

#### sqlx

> 第三方库，更方便地使用mysql
>
> 结构体的字段必须大写，因为sqlx是通过反射获取字段信息的

- `sqlx.Connect("mysql", dsn)`，`open`数据库并`ping`数据库

- `db.Get(&user,sqlStr,id)`：查询单条，不用一个字段一个字段去修改了，直接修改整个结构体变量
- `db.Select(&userlist, sqlStr)`：查询多条，`userlist`是个切片，虽然是引用类型，但也要传它的引用，因为`sqlx`只对指针类型进行了处理，其他引用类型没管

#### sql注入

``xxx or 1=1 #`

``xxx union select * from user #`

## Redis

> KV数据库

## NSQ

> Go开发的轻量级的消息队列

# day07

## Go Module