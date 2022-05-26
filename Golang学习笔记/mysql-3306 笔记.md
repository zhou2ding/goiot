# 前言

## 连接错误解决办法

**1. 查看用户信息**

`select host,user,plugin,authentication_string from mysql.user;`

![img](https://img-blog.csdn.net/20180609145407985?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTI2MDQ3NDU=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

**备注：**host为 % 表示不限制ip  localhost表示本机使用  plugin非mysql_native_password 则需要修改密码

**2. 修改用户密码**

`ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';`

更新user为root，host为% 的密码为123456

`ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';`
更新user为root，host为localhost 的密码为123456

## 修改root密码

**方法1： 用ALTER USER命令** 

 `ALTER USER 'root'@'localhost' IDENTIFIED BY '你的新密码'; `

**方法2：如果忘记密码,强行修改:**

1. 停止Mysql服务 `sudo /usr/local/mysql/support-files/mysql.server stop`

2. 进入终端输入：`cd /usr/local/mysql/bin/ `回车后;
   登录管理员权限`sudo su `回车后;
   输入以下命令来禁止mysql验证功能 `./mysqld_safe --skip-grant-tables &  回车`后mysql会自动重启（偏好设置中mysql的状态会变成running）

3. 输入命令` ./mysql` 回车后，输入命令`FLUSH PRIVILEGES`;

   回车后，输入命令 `ALTER USER 'root'@'localhost' IDENTIFIED BY '你的新密码';`
   
4. 密码`564710`
## sql语句执行顺序

1. from
2. join
3. on
4. where
5. group by (开始使用select中的别名，后面的语句中都可以使用)
6. 分组函数
7. having
8. select as
9. distinct
10. order by
11. limit

# day01

## 命令行启停服务

`net start 服务名称`

`net stop 服务名称`

## 语句分类

DQL：select

DML：insert delete update

DDL：create drop alter

TCL：commit rollback

DCL：grant revoke

## 常用命令

`mysql -uroot -p564710` 登录

先`docker exec 容器名  -it /bin/bash`，再`mysql -uroot -p564710` 登录mysql容器

`exit` 退出

`show databases;` 显示数据库

`crate database bjpowernode；` 创建数据库

`use bjpowernode;` 使用数据库

`source D:\learn\bjpowernode.sql;` 导入数据，批量执行文件中的sql语句

`show tables;` 显示表

`select * from dept;` 查看表的全部内容

`desc dept；` 查看表结构

`select version()` 查看mysql版本

`select database()` 查看当前的数据库

`system cls` 清屏

`show create table 表名;` 查看建表语句

## 注意

- 对于SQL语句来说，是通用的；

- 不区分大小写；

- 所有的SQL语句都是以分号结尾；

- 数据库中，`NULL`和任何值数学运算结果都为`null`
- 增删改查称为`CRUD`，没人说增删改查（Create、Retrieve、Update、Delete）

## 查询(DQL)

### 简单查询

`select 字段名 from 表名;`

`select 字段1,字段2 from 表名`

`select * from 表名` -- 效率低，可读性差，在代码中不写。想查询所有字段，可以把所有字段写上。

`select 字段1 as A,字段2 as B from 表名` 起别名，只是将查询的结果的列名显示为别名，不影响原名；

1. 备注：
   - as可以省略
   - 可以用单引号或双引号把别名括起来。
   - 在所用数据库中，字符串统一用单引号括起来，在Oracle中双引号用不了

2. 字段可以用数据表达式
3. 别名是中文，需要用单引号括起来

### 条件查询

``` mysql
select
	字段1,字段2,字段3...
from
	表名
where
	条件;
```

符号：

| 等于 | 不等于 | 小于 | 大于 | 小于等于 | 大于等于 |
| :--: | :----: | :--: | :--: | :------: | :------: |
|  =   | <>或!= |  <   |  >   |    <=    |    >=    |

|           之间           |                   为null                    |  不为null   | 并且 | 或者 |        多个或        | 取非        |
| :----------------------: | :-----------------------------------------: | :---------: | :--: | :--: | :------------------: | ----------- |
|     >= ... and <=...     |                   is null                   | is not null | and  |  or  |      in(a,b,c)       | not in      |
|   between ... and ...    |                =null是不行的                |  0不是null  |      |      | 等价于or a or b or c | is not null |
| 必须左小又大，且是闭区间 | 数据库中null代表空，不是一个值，不能用=比较 |             |      |      |                      |             |

1. and和or同时出现，先执行and；不确定优先级的话，加小括号
   
2. 模糊查询：liketb := db.Table("post")
   * %代表任意多个字符
   * _代表任意一个字符
   * \可以转义特殊字符
   
### 例外

   select 字面值 from table;

   得到的结果是一列内容全是该字面值的查询结果。

## 排序

``` mysql
mysql> select
    -> ename,sal
    -> from
    -> emp
    -> order by
    -> sal desc;
```

- desc：降序

- asc：升序

- 省略：默认升序

- order by后可以跟多个字段，用`,`隔开（第一个字段主导，当第一个字段值相等时，再启用第二个字段）
- order by 必须在wher后面

## 单行处理函数（数据处理函数）

**1. 特点：**一个输入对应一个输出；

**2. 常用的单行处理函数：**

   - `lower`转换成小写
   
   - `upper`转换成大写
   
   - `substr`截取字符串（`substr（字段名，起始位置，长度)`)，起始位置从1开始，没有0
   
   - `concat`拼接字符串（`concat(字段1，字段2)`）
   
   - `length`字符串长度
   
   - `trim`去除前后空白
   
   - `rouond`四舍五入（`round(数字，保留几位小数)`，0为个位，-1为十位）
   
   - `rand`小于1的随机数
   
   - `ifnull`处理null（`ifnull(字段，值)`，字段为null时，把它当做值）
   
   - `format(数字,'格式')` 数字格式化：`select ename,format(sal,'$999,999') from emp;`
   
   - `case 字段名 when 值1 then...when 值2 then...else...end`
   
     ```mysql
     # 涨薪
     # 当job是manager时，sal*1.1；
     # 当job是salesman时，sal*1.5;
     # 其他情况sal不变
     select
     	ename,job,sal as oldsal,
     (case job
     when 'manager' then sal*1.1 when 'salesman' then sal*1.5 else sal end)
     as newsal
     from
     	emp;
     	
     +--------+-----------+---------+---------+
     | ename  | job       | oldsal  | newsal  |
     +--------+-----------+---------+---------+
     | SMITH  | CLERK     |  800.00 |  800.00 |
     | ALLEN  | SALESMAN  | 1600.00 | 2400.00 |
     | WARD   | SALESMAN  | 1250.00 | 1875.00 |
     | JONES  | MANAGER   | 2975.00 | 3272.50 |
     | MARTIN | SALESMAN  | 1250.00 | 1875.00 |
     | BLAKE  | MANAGER   | 2850.00 | 3135.00 |
     | CLARK  | MANAGER   | 2450.00 | 2695.00 |
     | SCOTT  | ANALYST   | 3000.00 | 3000.00 |
     | KING   | PRESIDENT | 5000.00 | 5000.00 |
     | TURNER | SALESMAN  | 1500.00 | 2250.00 |
     | ADAMS  | CLERK     | 1100.00 | 1100.00 |
     | JAMES  | CLERK     |  950.00 |  950.00 |
     | FORD   | ANALYST   | 3000.00 | 3000.00 |
     | MILLER | CLERK     | 1300.00 | 1300.00 |
     +--------+-----------+---------+---------+
     ```
     

## 分组函数（多行处理函数）

**1. 特点：**多个输入对应一个输出。

**2. 注意：**

   - 分组函数在使用的时候必须先进行分组，然后才能使用。

   - 如果没有对数据分组，整张表默认为一组。
   - 分组函数自动忽略null。
   - 分组函数不能直接使用在where字句中（group by执行顺序在where之后，而分组函数执行又必须先分组）。
   - 所有的分组函数可以组合起来一起用。

**3. 常用分组函数：**

   - `count` 计数
     - `count(字段)`：统计该字段下所有不为null的元素的总数
     - `count(*)`：统计表中的总行数（因为一行中所有元素都为null的数据不存在）
   - `sum` 求和
   - `avg` 求平均
   - `max` 求最大值
   - `min` 求最小值

## 分组查询

**场景：**先对数据分组，再对每一组数据进行操作。

   ```mysql
select
...
from
...
where
...
group by
...
order by
...
   ```

**重点结论：**

在一条select语句中，如果有group by语句的话，select后面只能跟参加分组的字段，以及分组函数，其他都不能跟。

**tags：**

每个、不同 ≈ group by的字段

**having子句**：

分组之后再筛选，避免了where写在group by之后和在where之后出现分组函数的情况。

但是执行效率较低，可以先筛选（根据不同字段筛选而不是分组函数），再分组；优先where，再having，where用不了的情况下再having，比如筛选的条件是要用分组函数操作的。

eg：`select deptno, max(sal) from emp where sal > 3000 group by deptno;`

# day02-1

## 数据去重

`select distinct 字段1,字段2 from table`

- distinct只能出现在字段之前
- 出现在多个字段之前则是字段联合起来后去重

## 连接查询（跨表查询）

### 注意点

#### 连接查询的分类

   - 按年代分类：
     - SQL92
     - SQL99
   - 按连接方式分类：
     - 内连接
       - 等值连接
       - 非等值连接
       - 自连接
     - 外连接
       - 左外连接（左连接）
       - 右外连接（右连接）
       - 全连接（很少用）

#### 笛卡尔积现象

   - 现象：当两张表进行连接查询，没有任何条件限制的时候，最终查询结果的条数，是两张表条数的乘积，结果错误

     `select ename,dname from emp,dept;`

   - 解决错误：连接查询时加条件，满足这个条件的记录再筛选出来

     `select ename,dname frome emp,dept where emp.deptno = dept.deptno;`

     > 如上方法匹配次数和不做限制相比并没有减少，只是解决了错误

   - 提高效率：

        - 要给表起别名后用`别名.字段`来查询（提高编写sql语句的效率）
        
       ```mysql
       #SQL92语法，已被淘汰
       select
       	e.ename,d.dname
       from
       	emp e, dept d
       where
       	e.deptno = d.deptno
     ```
     - 尽量避免表的连接次数（连接次数=表A条数 \* 表B次 * 表c次数）
     
### 内连接

#### 注意点

两张表没有主次关系

#### 等值连接

- SQL92语法

  ```mysql
  #缺点：结构不清晰，表的连接条件和后期进一步筛选的条件，都放到了where后面。
  select
  	e.ename,d.dname
  from
  	emp e, dept d
  where
  	e.deptno = d.deptno;
  ```

- SQL99语法

  ```mysql
  #优点：表连接的条件是独立的，连接之后若需进一步筛选，再往后继续添加where。
  #查询每个员工的部门名称
  select
  	e.ename,d.dname
  from
  	emp e
  inner join #inner可省略，带着inner可读性更好，一眼能看出来是内连接
  	dept d
  on
  	e.deptno = d.deptno;  #条件是等量关系，所以被称为等值连接
  ```

#### 非等值连接

```mysql
#查询每个员工的工资等级
select
	e.ename,e.sal,s.grade
from
	emp e
inner join
	salgrade s
on
	e.sal between s.losal and s.hisal; #条件不是一个等量关系，称为非等值连接
```

#### 自连接

```mysql
#查询每个员工上级领导的名字，结果少一条，因为最高领导的领导是NULL，查不出来。用左外就能查出来
select
	a.ename as '员工名',b.ename as '领导名'
from
	emp a
inner join
	emp b
on
a.mgr = b.empno; #技巧：自连接，一张表看成两张表
```

### 外连接

#### 注意点

- 两张表存在着主次关系

- 将join关键字左边 / 右边这张表看成主表，主要是为了将主表的数据全部查询出来，捎带着连接查询子表

- 任何一个右连接 / 左连接都有左连接 / 右连接的写法

- 外连接的结果条数必定 >= 内连接的结果条数

#### 右外连接

right join，右表是主表

```mysql
select
	e.ename,d.dname
from
	emp e
right outer join #outer可以省略，带着可读性强
	dept d
on
	e.deptno = d.deptno
```

#### 左外连接

  left join，左表是主表

  ```mysql
  select
  	e.ename,d.dname
  from
  	dept d
  left outer join #outer可以省略，带着可读性强
  	emp e
  on
  	e.deptno = d.deptno
  ```

### 多表连接

```mysql
select
	...
from
	a
join
	b
on
	a和b的连接条件
left join
	c
on
	a和c的连接条件
```

- 案例1

  ```mysql
  # 找出每个员工的部门名和薪资等级，要求显示员工名、部门名、薪资、薪资等级
  select
  	e.ename,e.sal,d.dname,s.grade
  from
  	emp e
  join
  	dept d
  on
  	e.deptno = d.deptno
  join
  	salgrade s
  on
  	e.sal between s.losal and hisal;
  ```

- 案例2

  ```mysql
  # 找出每个员工的部门名、薪资等级和上级领导，要求显示员工名、领导名、部门名、薪资、薪资等级
  select
  	e.ename '员工',e2.ename '领导',e.sal '工资',d.dname '部门',s.grade '工资等级'
  from
  	emp e
  join
  	dept d
  on
  	e.deptno = d.deptno
  join
  	salgrade s
  on
  	e.sal between s.losal and hisal
  left join
  	emp e2
  on
  	e.mgr = e2.empno;
  ```

## 子查询

### 注意点

- select语句中嵌套select语句，被嵌套的select语句被称为子查询
- 可以出现在select、from、whre后面

### where子句中的子查询

```mysql
#找出工资比最低工资高的员工，显示员工名和薪资
select
	ename,sal
from
	emp
where
	sal > (select min(sal) from emp);
```

### from子句中的子查询

```mysql
#from后面的子查询，可以将子查询的结果当做一张临时表
#找出每个岗位的平均工资的薪资等级
select
	t.*,s.grade
from
	(select job,avg(sal) as avgSal from emp group by job) t
join
	salgrade s
on t.avgSal between s.losal and s.hisal;
```

### select子句中的子查询（了解，不掌握）

```mysql
#找出每个员工的部门名称，要求显示员工名，部门名
#缺陷：对于select后面的子查询，这个子查询只能一次返回一条结果，多于一条及报错
select
	e.ename,(select d.dname from dept d where e.deptno = d.deptno) as dname
from
	emp;
```
## 查询收尾

### union

- 效率高一些：
  - join：对于表连接来说，每连接一次新表，则匹配的次数满足笛卡尔积现象，成倍的翻；
  - union：但是union可以减少匹配的次数。在减少次数的情况下，还能完成两个结果集的拼接。
  - eg：a 连接 b 连接 c，a、b、c均10条记录
    - 连接查询的匹配次数是1000次
    - a连接b，100次，a连接c100次，union之后200次

- 注意事项：

  - 合并的时候，被合并结果集的列数必须相同
  - 结果集的列的数据类型也必须相同（mysql不报错，Oracle会报错）

  ```mysql
  #简单案例
  select ename,job from emp where job = 'manager'
  union
  select ename,job from emp where job = 'salesman'
  ```


### limit

- 概述：将查询结果的一部分取出来，通常使用在分页查询当中。
- 使用：
  - 完整用法：`limit startIndex, lengh`，startIndex是起始下标，从0开始；length是长度
  - 缺省用法：`limit n`，取前n条
- 注意：limit在order by之后执行

### 分页

- 每页显示pageSize条记录，第pageNo页：`limit (pagaNo - 1) * pageSize, pageSize`

# day02-2

## 表的创建(DDL)

### 概述

- 语法格式：`crate table 表名(字段名1 数据类型 约束条件 默认值, 字段名2 数据类型, 字段名3 数据类型);`
  - 表名：以`t_`开始或`tbl_`开始
  - 字段名：见名知意
  - 建表的时候可以指定默认值，在数据类型后跟`default xxx`
  - 可带`if not exists`
- 数据类型：
  - varchar：可变长度字符串，会根据实际的数据长度动态地分配空间，但速度慢（最长255）
  - char：定长字符串，分配固定长度的空间去存储数据，速度快，但可能会造成空间浪费（255）
  - int：整形（最长11位）
  - bigint：长整形
  - double：双精度浮点型
  - float：单精度浮点型
  - date：短日期
  - datetime：长日期
  - clob：字符大对象，最多可以存4G的字符串（超过255个字符的都要采用存储clob）
  - blob：二进制大对象
    - 专门用来存储图片、声音、视频等流媒体数据
    - 往blog类型的字段上插入数据需要使用IO流

```mysql
#案例
create table t_student(
	no int,
    name varchar(32),
    sex char(1) default 'm',
    age int(3),
    email varchar(255)
);
```

## 插入数据（DML）

### 插入日期

- 语法格式：`insert into 表名(字段1,字段2,字段3) values(值1,值2,值3)`

- 注意
  - 字段名要和值一一对应（数量、数据类型要对应）
  - insert语句但凡执行成功，必然会多一条记录，没有给其他字段指定值的话，默认值是NULL或建表时给的默认值
  - **省略字段名的话等于都写上了，所以值都得写上**（主键的值可以直接用字段名代替）
  
- 操作日期的单行处理函数
  - `str_to_date('日期字符串','日期格式')`
    - 将字符串varchar转换成date，通常在插入时使用
    - 如果提供的字符串是`'%Y-%m-%d'`，则可以不用此函数
  - `date_format(date类型数据,'日期格式')`
    - 将date转换成具有特定格式的varchar，通常在查询时使用
    - 不使用时，mysql会默认进行格式化，采用默认格式（`%Y-%m-%d`）
  - 日期格式
    - %Y 年
    - %m 月
    - %d 日
    - %h 时
    - %i 分
    - %s 秒
  - 日期默认格式
    - 短日期：`%Y-%m-%d`，按长日期格式插入时，会省略时分秒
    - 长日期：`%Y-%m-%d %h:i:%s`
  
- 实际业务

  ```mysql
  create table if not exists test(
  	id int primary key auto_increment,
      name varchar(32),
      time datetime default now(),
      ....
  );
  insert into test(name) values('张三');
  ```

### 插入多条数据

`insert into 表名(字段1,字段2,字段3) values(值1,值2,值3),(值1,值2,值3),(值1,值2,值3)`

## 修改数据

- 语法格式：`update 表名 set 字段1=值1,字段2=值2 where 条件`
- 注意：没有条件会使所有数据都被改了

## 删除数据

- 语法格式：`delete from 表名 where 条件`
- 注意：没有条件，整张表的数据都会被删除

## 表的删除

建议命令：`drop table if exists 表名`

## DDL顺序

- 删除表：先子后父
- 创建表：先父后子
- 删除数据：先子后父
- 插入数据：先父后子

## 快速创建表（表的复制）（了解）

- 整张表复制 `create table 表名2 as select * from 表名1`
- 部分表复制`create table 表名2 as select 语句 from 表名1`
- `insert into 表名2 select * from 表名1` 把查询结果插入表中（查询结果列数、数据类型必须和表2一样）

## 快速删除表的数据

- `delete from 表名` 
  - 速度慢，但可以回滚
  - 删除数据后，存储数据的空间还在
- `truncate table 表名`
  - 效率高，表呗一次截断，物理删除
  - 不支持回滚

## 表结构的增删改（实际开发很少出现）

- 添加字段、删除字段、修改字段
- 使用alter
- 了解，一般用工具操作

# day02-3

## 约束constraint

### 概述

- 创建表的时候给字段加约束来保证数据的完整性、有效性
- 添加在列后面的叫`列级约束`
- 没有添加在列后面的叫`表级约束`，希望多个字段联合起来添加某个约束时才使用`表级约束`
- 对于一张表，只要是主键或加了unique约束的字段会自动创建索引

### 分类

- 非空约束：not null
- 唯一性约束：unique
- 主键约束：primary key（PK）
- 检查约束：check（mysql不支持，Orcale支持）

#### 非空约束 not null

- not null约束的字段不能为null

- ```mysql
  create table t_vip(
  	id int,
      name varchar(255) not null
  );
  insert into t_vip(id) values(1) #会报错
  ```
  
- **只有列级约束，没有表级约束**

#### 唯一性约束 unique

- unique约束的字段的值不能重复，可以为null（null不是重复值）

- ```mysql
  create table t_vip(
  	id int,
      name varchar(255) unique, #列级约束
      email varchar(255)
  );
  insert into t_vip(id,name,email) values(1,'zhangsan',"zhangsan@123.com")
  insert into t_vip(id,name,email) values(2,'zhangsan',"lisi@123.com") #会报错
  ```

- **两个字段联合起来后具有唯一性**

  - 如果在两个字段后分别添加`unique`约束，则是两个字段各自唯一，不符合此需求

  - 应这样做：`unique(字段1,字段2)`

    ```mysql
    drop table if exists t_vip;
    create table t_vip(
    	id int,
        name varchar(255),
        emai varchar(255),
        unique(name,email) #表级约束
    );
    ```

#### 主键约束 primary key

1. 相关术语

   - 主键约束：一种约束
   - 主键字段：该字段上添加了主键约束
   - 主键值：主键字段中的每个值

2. 概念

   - 主键值是每一行记录的**唯一标识**（每一行的身份证号）
   - 任何一张表都应该有主键，没有主键的表是无效的
   - **主键的特征：not null + unique（主键值不能为null，也不能重复）**
   - **一张表，主键约束只能有1个**
   - 主键值建议使用int，bigint，char等类型，一般是数字、定长，不建议用varchar

3. 分类一

   - 单一主键：一个字段作主键

     ```mysql
     drop table if exists t_vip;
     create table t_vip(
     	#id int primary key, #列级约束
         primary key(id)， 表级约束
         name varchar(255),
         sex int(1)
     );
     insert into t_vip(id,name,sex) values(1,'zs',0),(2,'ls',1);
     insert into t_vip(id,name,sex) values(2,'ww',0); #报错，主键值重复
     insert into t_vip(name,sex) values('zl',0); #报错，主键值为null
     ```

   - 复合主键：多个字段联合起来作主键（实际开发不用）

     ```mysql
     #实际开发中不建议使用复合主键，因为主键的意义就是这行记录的身份证号
     drop table if exists t_vip;
     create table t_vip(
     	id int,
         name varchar(255),
         sex int(1),
         primary key(id,name) #两个字段联合起来作主键，复合主键
     );
     ```

4. 分类二

   - 自然主键：主键值是一个自然数，和业务没关系
   - 业务主键：主键值和业务紧密关联，例如银行卡号
   - 实际开发中，自然主键用的多，因为主键的作用就是这行记录的唯一标识，不需要有意义；一旦和业务挂钩，当业务变动的时候可能会影响主键值

5. mysql自动维护主键值

   ```mysql
   drop table if exists t_vip;
   create table t_vip(
   	id int primary key auto_increment, #自增，从1开始以1递增
       name varchar(255)
   );
   insert into t_vip(name) values('zhangsan');
   insert into t_vip(name) values('zhangsan');
   ```


#### 外键约束 foreign key

1. 相关术语

   - 外键约束：一种约束
   - 外键字段：该字段上添加了外键约束
   - 外键值：外键字段中的每个值，该字段的值必须来自于主表的关联列的值

2. ```mysql
   drop table if exists t_student; #先删子表
   drop table if exists t_class; #后删父表
   create table t_class(
   	classno int primary key,
       classname varchar(255)
   );
   create table t_student(
   	id int primary key auto_increment,
       name varchar(255),
       cno int,
       foreign key(cno) references t_class(classno)
   );
   insert into t_class(classno,classname) values(100,'高三一班'),(101,'高三二班');
   ```

3. 概念

   - 父表：被引用为外键的字段所在的表；子表：引用某其他表的字段作为外键的表
   - 子表的外键值可以为空
   - **被引用为外键的父表中的字段，不一定是主键，但至少具有`unique`约束**

#### 两个约束联合

- `create table v_vip(id int, name varchar(255) not null unique);`
- 一个字段被`not null`和`unique`联合约束后，该字段自动变成主键字段（Oracle不会这样）

# day02-4

## 存储引擎（了解）

### 概念

- mysql中特有的术语，Oracle中不叫这个，其他数据库没有
- 存储引擎是一个表存储/组织数据的方式
- 给表指定存储引擎
  - 在`create table 表名()`后跟`ENGINE=`和`CHARSET=`来指定存储引擎和字符集
  - 默认为`InnoDB`和`utf8`
  - `engine`和`charset`之间一般跟`default`

### 分类

- `show engines \G`：查看支持哪些存储引擎，`\G`可省略

- mysql支持九大存储引擎；版本不同，支持情况不同；常用以下三个
  - MyISAM
    - 使用三个文件表示每个表
      - 格式文件：存储表结构的定义`mytable.frm`
      - 数据文件：存储表行的内容`mytable.MYD`
      - 索引文件：存储表上的索引（相当于目录）`mytabl.MI`
    - 特点：可以转换为压缩、只读表来节省空间
  - InnoDB
    - mysql默认的存储引擎，也是个重量级的引擎，非常安全，但效率不是很高，不能压缩，不能转换为只读
    - 支持事务，支持数据库崩溃后自动恢复机制（提供了一组用来记录事务性活动的日志文件）
    - 在数据库目录内，每个表以表`.frm`文件表示
    - InnoDB表空间`tablespace`用于存储表的数据和索引
  - MEMORY
    - 数据和索引存储在内存中，且行的长度固定，非常快，但不安全、关机后数据消失
    - 在数据库目录内，每个表以表`.frm`文件表示
    - 表级锁机制
    - 不能包含`TEXT`和`BLOB`字段
  
# day03

## 事务-transaction

### 概述

- 一个事务就是一个完整的业务逻辑
- 只有DML语句和事务有关
- 一个事务本质上是多条DML语句，要么同时成功，要么同时失败！！！
- InnoDB提供了一组用来`记录事务性活动的日志文件`

```mysql
事务开启了
insert
insert
delete
update
事务结束了
```

### 提交事务

- 清空`记录事务性活动的日志文件`，将数据全部持久化到数据库表中
- 提交事务标志着事务的结束，是一种全部成功的结束
- `commit;`mysql默认是支持自动提交事务的，每执行一条DML语句，就提交一次！
- `start transaction;`用来关闭msql的事务自动提交机制

### 回滚事务

- 将之前的所有DML操作全部撤销，并且清空`记录事务性活动的日志文件`
- 回滚事务标志着事务的结束，是一种全部失败的结束
- `rollback;`只能回滚到上一次的提交点！

### 事务操作

```mysql
#提交事务
start transaction;
insert into dept_bak values(10, 'sales', 'beijing')
insert into dept_bak values(10, 'sales', 'beijing')
commit;
rollback; #此时回滚没有任何作用，因为事务提交了，已经结束了！

#回滚事务
start transaction;
insert into dept_bak values(10, 'sales', 'beijing')
insert into dept_bak values(10, 'sales', 'beijing')
rollback;
```

### 事务的特性

`A`：原子性

- 事务是最小的工作单元，不可再分

`C`：一致性

- 在同一个事务当中，所有操作必须同时成功，或同时失败，以保证数据的一致性

`I`：隔离性

- A事务和B事务之间具有一定的隔离

`D`：持久性

- 事务最终结束的一个保障。事务提交就相当于没有保存到硬盘上的数据保存进硬盘

### 事务的隔离级别

#### 等级

- 读未提交：`read uncommited`（最低的隔离级别）《没提交就读到了》

  > 事务A可以读到事务B还没提交的数据，是理论上的，存在脏读现象（Dirty Read），称为读到了脏数据

- 读已提交：`read commited`（大多数事务从这级别起步）《提交之后才能读到》

  > 事务A只能读到事务B提交之后的数据，解决了脏读，但不可重复读取数据
  >
  > 即事务A结束之前，如果有多个事务修改了数据，则事务A每次从表中读取到的数据不一致
  >
  > 读到的数据比较真实，Oracle的默认级别

- 可重复读：`repeatable read`《提交之后也读不到，每次都是读取事务刚开启时的数据》

  > 事务A开启之后，不管多久，每次A从表中读到的数据都是一致的，即使事务B对表中数据已经修改并提交，A读到的数据仍不变
  >
  > 可能存在幻影读，即读到的数据不一定真实，mysql的默认级别
  >
  > 比如：银行查询总账，开始了事务A进行select，执行的时间比较长；期间有人存取款，每次存取款都是一次事务，存取款的事务提交后，事务A查询到的总账不变

- 序列号/串行化：`serializable`（最高的隔离级别）

  > 效率最低，表示事务排队，不能并发

#### 事务隔离级别SQL语句

- 查看当前隔离级别

  - `select @@transaction_isolation;`
  - `select @@global.transaction_isolation;`

- 修改隔离级别，重启mysql后生效

  - `set transaction_isolation READ-UNCOMMITTED;`
  - `set session transaction_isolation READ-UNCOMMITTED;`
  - `set global transaction_isolation READ-UNCOMMITTED`

- 验证`read uncommited`：先修改全局隔离级别，再重启mysql

  | 事务A                | 事务B                                 | 结果                             |
  | -------------------- | ------------------------------------- | -------------------------------- |
  | start transaction    | start transaction                     | 都ok                             |
  | select * from t_user |                                       | 查询结果为空                     |
  |                      | insert into t_user values('zhangsan') | 插入ok                           |
  | select * from t_user |                                       | 事务B未提交，但事务A能查询到数据 |
  |                      | rollback                              | 回滚事务ok                       |
  | select * from t_user |                                       | 查询结果为空                     |

- 验证`read commited`：先修改全局隔离级别，再重启mysql

  | 事务A                | 事务B                                       | 结果                         |
  | -------------------- | ------------------------------------------- | ---------------------------- |
  | start transaction    | start transaction                           | 都ok                         |
  | select * from t_user |                                             | 查询结果为空                 |
  |                      | insert into t_user values('zhangsan')       | 插入ok                       |
  | select * from t_user |                                             | 事务B未提交，事务A查不到数据 |
  |                      | commit                                      | 提交事务ok                   |
  | select * from t_user |                                             | 此时事务A才能查到            |
  |                      | insert into t_user values('lisi')<br>commit | 插入ok，提交ok               |
  | select * from t_user |                                             | 事务A能查到事务B新插入的数据 |

- 验证`repeatable read`：先修改全局隔离级别，再重启mysql

  | 事务A                | 事务B                                 | 结果                                       |
  | -------------------- | ------------------------------------- | ------------------------------------------ |
  | start transaction    | start transaction                     | 都ok                                       |
  | select * from t_user |                                       | 查询结果为空                               |
  |                      | insert into t_user values('zhangsan') | 插入ok                                     |
  |                      | insert into t_user values('lisi')     | 插入ok                                     |
  | select * from t_user |                                       | 查询结果为空                               |
  |                      | commit                                | 提交事务ok                                 |
  | select * from t_user |                                       | 仍为空，查到的始终是事务A开启时表中的数据  |
  |                      | delete from t_user                    | 删除数据成功                               |
  | select * from t_user |                                       | 只要A不提交，查到的永远是事务A开启时的数据 |

- 验证`serializable`：先修改全局隔离级别，再重启mysql

  | 事务A                                 | 事务B                | 结果                             |
  | ------------------------------------- | -------------------- | -------------------------------- |
  | start transaction                     | start transaction    | 都ok                             |
  | select * from t_user                  |                      | 查询结果为空                     |
  | insert into t_user values('zhangsan') |                      | 插入ok                           |
  |                                       | select * from t_user | 查询会卡住（排队中）             |
  | commit                                |                      | 事务A提交ok，事务B的查询自动完成 |

## 索引（理解）

- 概述

  - 是在表的字段上添加的，相当于目录，为了缩小扫描范围，提高查询效率而存在
  - 可以单个字段添加索引，也可以多个字段联合起来添加索引
  - 索引是需要排序的，是`B+Tree`的数据结构，底层是一个自平衡的二叉树，遵循左小右大的原则存放，采用中序遍历方式存取数据
  - mysql在查询方面主要两种方式：全表扫描（有where语句时只扫描where的字段），根据索引检索

- 注意点

  - 任何数据库的主键字段都会自动添加索引对象；mysql的`unique`约束字段也会自动添加索引对象
  - 任何数据库，任何一张表的任何一条记录都在硬盘上有一个物理存储编号
  - mysql中索引是一个单独的对象，`MyISAM`引擎把索引存储在`.MYI`文件中，`InnoDB`存储在`tablespace`中，`MEMORY`存储在内存中。但不管索引存储在哪，都是一个树的形式存在（自平衡二叉树：`B+Tree`）

- 实现原理

  > 根据左小右大的存储原则，搜索时按这个原则搜，找到后直接根据物理编号拿到整条记录

  ![image-20210428121414649](D:\资料\Go\src\studygo\Golang学习笔记\mysql-3306 笔记.assets\image-20210428121414649.png)


- 索引的使用

  - 使用条件：数据量庞大，具体还要看硬件环境，需要测试；该字段经常出现在`where`后面，即该字段总是被扫描；该字段有很少的DML操作（因为DML之后索引需要重新排序）
  - 不要随便添加索引，索引多了系统性能反而会降低；建议通过主键或unique约束的字段来查询
  - `show index from emp;`查看所有的索引
  - 方式一：
    - `create index emp_ename_index on emp(ename);`在ename字段上创建索引
    - `create unique index emp_ename_index emp(ename);`在ename字段上创建唯一索引
    - `drop index emp_ename_index on emp;`删除ename对象上的索引
  - 方式二：
    - `ALTER TABLE table_name ADD INDEX index_name (column_list)`
    - `ALTER TABLE table_name ADD UNIQUE (column_list)`
    - `ALTER TABLE table_name ADD PRIMARY KEY (column_list)`
  - `explain select * from emp where ename='KING';`查看一个SQL语句是否使用了索引进行检索，如果输出的结果中`type`是`ALL`的话，就没有使用索引；是`ref`的话就是用了索引

- 索引的失效

  - 模糊查询的时候以`%`开头
  - 使用`or`，除非`or`两边的字段都有索引
  - 使用`in`
  - 使用复合索引（多个字段联合起来添加索引）的时候，没有使用左侧的列查询
  - 在where当中索引列参加了数学运算或使用了函数或使用了类型转换
  
- 索引的分类

  > 单一索引、复合索引、主键索引、唯一性索引

## mysql优化

  - 使用索引（优先考虑）
  - 尽量少用or，可以用union
  - 模糊查询尽量不要以%开头

## 视图

> view：站在不同的角度去看待同一份数据，也是以文件的形式存在的
>
> 只有DQL语句才能以view的形式创建（as后面必须是DQL）

- `create view emp_view as select * from emp;`
- `drop view emp_view`
- 面向视图对象进行增删改查，对视图对象的增删改查会导致原表被操作
- 连接查询的结果也可以用来创建视图对象
- 视图是用来简化SQL语句的（特别长的查询语句，且会在不同的位置上反复使用，可以创建一个视图对象，后续直接使用此视图对象）

## DBA常用命令

- `create user 'xxx'@'localhost' identified by 'password'`
- `grant/revoke`：授权/撤销权限
- 重点掌握：数据的导入和导出
  - 导出，在CMD中而非登入mysql：`mysqldump 数据库名>D:\xxx.sql -uroot -p123`
    - 数据库名之后跟`空格和表名`，可以导出指定表
  - 导入，先创建数据库并使用数据库之后，再导入：`source D:\xxx.sql`

## 数据库设计三范式（面试常问）

> 数据库表的设计依据，教你怎么进行设计数据库的表
>
> “多对多，三张表，关系表两个外键”
>
> “一对多，两张表，多的表加外键”
>
> “一对一，外键唯一”（子表的外键字段要额外添加unique约束）（如登录信息和用户信息拆开来，用户信息表添加一个登录id字段作为外键）

- 第一范式：任何表必须有主键，每一个字段都是原子性不可再分

  > 比如联系方式这个字段，用邮箱和电话来作为值，则不符合此范式，拆成两个字段后才是原子性不可再分了

- 第二范式：在第一范式的基础上，所有非主键字段必须完全依赖主键，不能部分依赖（即尽量不要有复合主键）

  > 字段间有多对多关系的，三张表，关系表两个外键（第三张表专门用来存关系）

- 第三范式：在第二范式的基础上，所有非主键字段必须直接依赖主键，不能传递依赖

  > 字段间有一对多关系的，两张表；一对多的关系中，多所属的表加外键

- 面试问到了除了回答三范式，还要说：只是理论，最终是为了满足客户的需求，有时会拿冗余换执行速度，因为表的连接次数越多，效率越低（笛卡尔积），对于开发人员来说，编写SQL语句的难度也会降低。![image-20210609124316442](mysql-3306 笔记.assets/image-20210609124316442.png)







