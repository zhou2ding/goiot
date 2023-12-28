## 简介

## 是什么

- javascript是运行在客户端的脚本语言
- 不需要编译，运行过程中由js解释器逐行解释并执行
- 也可以Node.js进行服务器端编程
- 是解释型、动态类型语言

## 作用

- 表单动态校验、密码强度检测
- 网页特效
- 服务端开发（Node.js）
- 桌面程序（Electron）
- App（Cordova）
- 控制硬件-物联网（Ruff）
- 游戏开发（cocos2d-js）

## 浏览器执行js

- 浏览器分成两部分：渲染引擎和JS引擎
- js引擎也称为js解释器，用来读取网页的js代码并处理，如chrome的V8引擎
- js解释器逐行解释源码，转换成机器语言，由计算机执行

## 组成

- ECMAScript，基础语法
- DOM，页面文档对象模型
- BOM，浏览器对象模型

## 写法

- 行内式：直接在元素内写
- 内嵌式：写在<script></script>标签内
- 外部式：在<script></script>内通过src属性引入

## 输入输出

- prompt，浏览器弹出输入框，输入后，函数返回的结果都是string类型
- alert，浏览器弹出警示
- console.log控制台打印输出

## 报错

js是逐行解释，一行报错时，就不会继续执行后面的

# 变量

## 使用

- var age; age = 18;
- var age = 18;
- var age = prompt('输入年龄');
- age = 20; // 重新赋值，不需要var
  - 变量值以最后一次的赋值为准
- 集体声明多个变量
  - var age = 18, address = '西安', salary = 2000;

- var a = b = c = 9; 相当于var a = 9; b = 9; c = 9; // 区别于集体声明，b、c是直接赋值，没有声明
- 特殊情况
  - 只声明不赋值，结果是undefined
  - 不声明不赋值，直接使用变量 ，会报错
  - 不声明，直接赋值，可以使用（会变成全局变量，不提倡）
  - name这个变量名有特殊含义

## 命名规范

- 字母、数字、下划线、$组成
- 严格区分大小写
- 不能以数字开头
- 不能是关键字、保留字
- 必须见名知意
- 遵循小驼峰命名
- 尽量不使用name作为变量名

## 交换变量

只能使用临时变量

# 数据类型

## 简介

javascript是弱类型和动态类型语言，变量的数据类型在运行过程中根据值来确定，变量的数据类型可以变化。

## 简单数据类型

| 数据类型  | 说明                                                         | 默认值    | 控制台中的颜色 |
| --------- | ------------------------------------------------------------ | --------- | -------------- |
| number    | 数字型，包含整型和浮点型值，如21、0.21<br>数字以0开头是八进制，以0x开头是十六进制<br>Number.MAX_VALUE和Number.MIN_VALUE是js中数字的最大值和最小值<br>Infinity和-Infinity是js中的无穷大和负无穷大<br>NaN，not a number，非数字<br>isNaN()函数用来判断是否为非数字，是数字返回false，非数字返回true | 0         | 蓝色           |
| boolean   | 布尔型，进行算术运算时，true是1，false是0                    | false     | 深蓝色         |
| string    | 字符串型，只要是引号括起来的，就是字符串，单、双引号都可以<br>出现引号嵌套时，外双内单，外单内双，不用转义<br>转义字符：\n（换行），\\\，\\'，\\"，\\t（缩进），\\b（空格）<br>length属性可以获取字符串的长度<br>字符串用+可以拼接，+两边只要有一个字符串，那必然是拼接；如果是其他符号，则进行数学运算 | ""        | 黑色           |
| undefined | 未定义，var a; 声明了变量但未赋值，此时a=undifined，也可以直接var a = undefined;<br>undefined和字符串+时，是进行拼接；和数字+时，结果是NaN | undifined | 浅灰色         |
| null      | 空值，var a = null; 声明了变量a并赋值成空值null<br/>null和字符串+时，是进行拼接；和数字+时，结果是数字本身 | null      | 浅灰色         |

## 获取变量的数据类型

typeof：不是函数，是一个关键字，typeof a，返回a的数据类型

- 注意：null的type of结果是object

## 数据类型转换

表单、prompt获取过来的数据默认是string类型，使用时有时候需要进行数据类型转换

### 转成字符串

- toString()：变量.toString()
- String()：String(变量)
- +：变量+''（隐式转换，更推荐）

### 转成数字

- parseInt(string)函数
  - 如果参数是浮点数的字符串，则取整
  - 如果参数有后缀单位，则会去掉单位
  - 如果字符串不是以数字开头，结果是NaN
- parseFloat(string)函数
  - 如果参数有后缀单位，则会去掉单位
  - 如果字符串不是以数字开头，结果是NaN
- Number(string)强制转换函数
- js隐式转换（- * /），如'12'-0，'12'*1，'12'/1

### 转成布尔型

Boolean()函数

- 代表空、否定的值会转成false，如''、0、NaN、null、undefined
- 其他都是true

# 运算符

## 算术运算符

- +-*/%运算
- 浮点数在算术运算时，有精度问题，因此==不能直接判断两个浮点数是否相等==

## 递增和递减运算符

前置和后置单独使用时，效果一样，后置更常用

- 前置：先自加1，后返回值，eg：var a = 1; ++a + 10; 结果是12
- 后置：先返回值，再自加1，eg：var a = 1; a++ + 10; 结果是11

## 比较运算符

- \>、<、>=、<=、==、!=
- ==默认会进行数据类型转换，把字符串型的数字转换成数字型
- \===全等于，!\==不全等于，全等要求两侧的值和数据类型完全一致

## 逻辑运算符

- &&、||、!
- 短路运算（逻辑中断），运算符左边的布尔值确定好时，就不再进行右边表达式的运算

## 赋值运算符

- =、+=、-=、*=、/=、%=

## 运算符优先级

| 优先级 | 运算符     | 顺序           |
| ------ | ---------- | -------------- |
| 1      | 小括号     | ()             |
| 2      | 一元运算符 | ++ -- !        |
| 3      | 算数运算符 | 先* / % 后 + - |
| 4      | 关系运算符 | > >= < <=      |
| 5      | 相等运算符 | == != === !\== |
| 6      | 逻辑运算符 | 先 && 后 \|\|  |
| 7      | 赋值运算符 | =              |
| 8      | 逗号运算符 | ,              |

# 流程控制

## 顺序流程

![image-20230727165120337](./JavaScript笔记.assets/image-20230727165120337.png)

## 分支流程

### if

```javascript
if (条件表达式) {
  // 执行语句;
}
```

```javascript
if (条件表达式) {
  // 执行语句;
} else {
  // 执行语句;
}
```

```javascript
if (条件表达式) {
  // 执行语句;
} else if (条件表达式) {
  // 执行语句;
} else {
  // 执行语句;
}
```

### switch

> case匹配的时候，必须是全等
>
> break不写的话，还会执行下一个case，不管匹不匹配（和go的fallthrough效果一样）

```javascript
switch(表达式) {
  case value1:
    // 执行语句1;
    break;
  case value2:
    // 执行语句2;
    break;
  default:
    // 执行最后的语句;
}
```

## 三元表达式

三元运算符：`? :`

```javascript
条件 ? 表达式1 : 表达式2	// 如果条件为真则返回表达式1的值，如果条件为假则返回表达式2的值 
```

## 循环流程

## for

```javascript
for (初始化变量; 条件表达式; 操作表达式) {
  循环体;
}
```

## while

```javascript
初始化变量;
while (条件表达式) {
  循环体;
  操作表达式;	// 也可在循环体前面
}
```

## do while

至少执行一次

```js
初始化变量;
do {
  循环体;
  操作表达式;	// 也可在循环体前面
} while (条件表达式)
```

## continue和break

用法和其他语言一样。

# 数组

一组数据的集合，存储在单个变量下，每个数据称作元素，js的数组可以放任何类型的元素。

## 创建

- new：var 数组名 = new Array();
- 数组字面量
  - 空数组：var 数组名 = [];
  - 有元素的数组：var 数组名 = [1, 5.3, 'abc', true];

## 操作

### 访问元素

- 索引：数组名[索引号]，==越界后得到的是undefined==

### 遍历

```js
for (var i = 0; i < arr.length; i++) {
  arr[i]
}
```

### 新增元素

- 修改数组的长度（此时多出来的空元素是undefined）；给空元素赋值。
- 直接用索引塞进去

```js
// 从一个数组筛选元素塞进新数组
for (var i = 0; i < arr.length; i++) {
  if (筛选条件) {
    newArr[newArr.length] = arr[i]
  }
}
```

# 函数

## 声明

参数和返回值不需要写类型

```js
function demo(para1, para2) {
  // do something
}
demo()
```

js中也有匿名函数，可以赋值给变量

```js
var fun = function() {
  // do something
}
fun()
```

## 参数

- 如果实参的个数和形参的个数相等，则正常输出结果
- 如果实参的个数多于形参的个数，则多出的实参忽略
- 如果实参的个数少于形参的个数，则缺的参数是undefined

## 返回值

- 不需要返回变量名，直接return，return之后的代码都不会执行
- ==return只能返回一个值，有多个值的话用逗号隔开，则只返回最右边的值==
- ==需要返回多个值时，返回数组==
- 函数没有return时，返回的是undefined

## arguments

不确定有多少个参数传递时，用arguments来获取。JS中，arguments是当前函数的一个内置对象，所有函数都内置了一个arguments对象，**存储了传递的所有实参**。

arguments展示形式是一个伪数组，有以下特点

- 有length属性
- 按索引方式存储数据，因此可以遍历
- 不具有数组的push、pop等方法

## 调用

函数可以互相调用，可以调用后面声明的函数

# 作用域

## 全局作用域

- script标签
- 单独的js文件

## 局部作用域

在函数内部就是局部作用域，也叫函数作用域

## 全局变量和局部变量

- 全局变量
  - 全局作用域下的变量，全局都能使用
  - 没有声明直接赋值的变量，也是全局变量

- 局部变量
  - 局部作用域下的变量，只能函数内部使用
  - 函数的形参也是局部变量
  - 在别的地方调用局部变量，会报错Uncaught ReferenceError
- 全局变量只有浏览器关闭时才销毁，比较占内存；局部变量在程序执行完毕就会销毁

## 作用域链

内部函数访问外部函数的变量，采取的是链式查找的方式来决定用哪个值（就近原则），称为作用域链

# 预解析

js引擎运行js过程：

1. 预解析：js引擎会把js里面所有的var和function声明提升到当前作用域的最前面
2. 执行代码：按照代码顺序从上往下执行

## 变量预解析

把所有的**变量声明**提升到当前作用域的最前面，**不提升赋值**

## 函数预解析

把所有的函数声明提升到当前作用域的最前面，不提升调用

# 对象

## 概念

js中，对象是一组无序的相关属性和方法的集合

- 属性：事物的特征
- 方法：事物的行为

## 创建方式

### 字面量

使用对象字面量{}创建对象，属性或方法采取键值对的形式，之间用逗号隔开

```js
var obj = {};	// 空对象
var obj2 = {
  uname: 'zs',
  age: 10,
  sex: true,
  sayHi: function() {
    // do something
  }
}
```

### new Object()

```js
var obj = new Object();
obj.uname = 'ls';
obj.age = 11;
obj.sex = false;
obj.sayHi = function() {
  // do something
}
```

## 调用方式

- 属性
  - 对象名.属性名
  - 对象名['属性名']
- 方法
  - 对象.方法名()

## 构造函数

把对象里面一些相同的属性和方法抽象出来封装到函数里。

```js
function 构造函数名() {
  this.属性 = 值;
  this.方法 = function() {}
}
new 构造函数名();
```

- 构造函数名的首字母大写
- 构造函数不需要return就能返回结果，返回的是一个对象 

## 构造函数和对象

- 构造函数是抽象的公共部分，泛指某一大类
- 对象是特指某一个，用new创建对象的过程称为对象的实例化

## new执行过程

1. new 构造函数()，在内存中创建一个空对象
2. this指向创建的这个空对象
3. 执行构造函数里的代码，给空对象添加属性和方法
4. 返回这个对象

## 遍历对象

### for...in 循环：

`for...in` 循环可以遍历对象的可枚举属性，包括对象的原型链上的属性。在遍历时，需要使用 `hasOwnProperty` 方法来确保只获取对象自身的属性，而不包括继承的属性。

```js
javascriptCopy code
var person = {
  name: 'Alice',
  age: 25,
  job: 'Engineer'
};

for (var key in person) {
  if (person.hasOwnProperty(key)) {
    console.log(key + ': ' + person[key]);
  }
}
```

### Object.keys 方法：

`Object.keys` 方法返回一个包含对象自身可枚举属性名称的数组。可以遍历这个数组来访问对象的属性。

```js
javascriptCopy code
var person = {
  name: 'Alice',
  age: 25,
  job: 'Engineer'
};

var keys = Object.keys(person);

for (var i = 0; i < keys.length; i++) {
  var key = keys[i];
  console.log(key + ': ' + person[key]);
}
```

### Object.values 方法：

`Object.values` 方法返回一个包含对象自身可枚举属性值的数组。

```js
javascriptCopy code
var person = {
  name: 'Alice',
  age: 25,
  job: 'Engineer'
};

var values = Object.values(person);

for (var i = 0; i < values.length; i++) {
  console.log(values[i]);
}
```

### Object.entries 方法：

`Object.entries` 方法返回一个包含对象自身可枚举属性键值对的数组。

```js
javascriptCopy code
var person = {
  name: 'Alice',
  age: 25,
  job: 'Engineer'
};

var entries = Object.entries(person);

for (var i = 0; i < entries.length; i++) {
  var key = entries[i][0];
  var value = entries[i][1];
  console.log(key + ': ' + value);
}
```

### forEach 方法：

对于数组形式的对象，可以使用 `forEach` 方法来遍历。

```js
javascriptCopy code 
var person = {
  name: 'Alice',
  age: 25,
  job: 'Engineer'
};

Object.keys(person).forEach(function(key) {
  console.log(key + ': ' + person[key]);
});
```

# 内置对象

> js中的对象分3种：自定义对象，内置对象，浏览器对象

## Math对象

不是一个构造函数，不需要new来调用，直接使用它的属性和方法即可

- Math.PI，圆周率

- Math.floor()，向下取整

- Math.ceil()，向上取整

- Math.round()，四舍五入

- Math.abs()，绝对值

- Math.max()/Math.min()，最大值/最小值

- 随机数

  - Math.random()，返回[0,1)之间的随机小数

  - 获取其他范围的随机数：自己用Math的其他函数结合Math.random()去实现

    - 获取两个数之间的随机整数，并包含这两个数

      ```js
      function getRandomInt(min, max) {
        return Math.floor(Math.random() * (max - min +1)) + min;
      }
      ```

## 日期对象

- 只能通过Date构造函数来实例化日期对象：var date = new Date()

  - 没有传参，则返回当前时间

  - 参数可以
    - 数字型：var date = new Date(2023, 10, 1)，2023年10月1日
    - 字符串型（最常用）：var date = new Date('2023-5-1 8:8:8')

- 日期格式化

  var now = new Date();

  | 方法名            | 说明                     |
  | ----------------- | ------------------------ |
  | now.getFullYear() | 获取当年                 |
  | now.getMonth()    | 获取当月（0-11）         |
  | now.getDate()     | 获取当天日期             |
  | now.getDay()      | 获取周几（周日0到周六6） |
  | now.getHours()    | 获取当前小时             |
  | now.getMinutes()  | 获取当前分钟             |
  | now.getSeconds()  | 获取当前秒钟             |

- 1970年1月1日至今的总的毫秒数（时间戳）

  - var now = new Date();

    - now.valueOf()

    - now.getTime()

  - 最常用写法：var now = +new Date();

    > +new Date()可以传参数

  - H5新增方法：Date.now()

  ![image-20231220155726510](./JavaScript笔记.assets/image-20231220155726510.png)

## 数组对象

- 创建

  - 数组字面量：var arr = [];	

  - new Array()

    - 创建一个空数组：var arr = new Array();

    - 创建长度为2且有两个空元素的数组：var arr = new Array(2);

    - 创建有元素的数组：至少传两个参数 var arr = new Array(2, 3);

      > 等价于 var arr = [2, 3];

- 检测是否为数组

  - instanceof：if (arr instanceof Array) {}

  - Array.isArray()：if (Array.isArray(arr)) {}

    > H5新增的方法 

- 添加元素

  - push：在末尾添加元素，返回push后数组的长度，var arr = []; arr.push(1, 'abc');
  - unshift：在开头添加元素，返回unshift后数组的长度

- 删除元素

  - pop：删除最后一个元素，返回被删除的元素
  - shift：删除第一个元素，返回被删除的元素

- 排序

  - reverse()：翻转数组

  - sort() 

    - 排序，默认按照元素转换为字符串的各字符的Unicode位点进行排序（和mysql的排序一样）

    - 自定义排序：

      ```js
      arr.sort(function(a, b) {
        return a - b;	// 升序排序，b-a则是降序
      })
      ```

- 索引

  - indexof()：查找给定元素的第一个索引，不存在则返回-1
  - lastIndexOf()：查找给定元素的最后一个索引，不存在则返回-1

- 转换为字符串

  - toString()：得到的字符串，数组元素间用逗号分隔
  - join(分隔符)：默认分隔符是逗号 

- 其他

  - concat()：连接两个或多个数组，返回连接后的数组，不影响原数组
  - slice(begin, end)：截取数组，返回begin到end之间的新数组
  - splice(begin, n)：从begin开始删除n个元素，返回被删除元素后的新数组，会影响原数组

## 字符串对象

> 所有字符串的方法都不会修改字符串本身

- 根据字符返回索引：indexof(target, [start])和lastIndexof(target, [start])
  - str.indexOf('a'); 默认从0开始查找，lastIndexof一样
  - str.indexOf('a', 3); 从索引3的位置开始查找，lastIndexof一样
- 根据索引返回字符
  - charAt(index)，返回index位置的字符
  - charCodeAt(index)，返回index位置的字符的ASCII码，作用：**判断用户按了哪个键**
  - str[index]，H5新增，和charAt等效
- 操作字符串
  - concat，拼接字符串，用+拼接更常用
  - substr(start,length)，从start开始，截取length长度的字符串
  - slice(start,end)，从start开始截取到end，不包括end
  - substring(start,end)，基本和slice相同，不接受负值
  - replace('被替换的字符', '替换为的字符')，替换字符，只替换第一个出现的字符
  - split，把字符串转换为数组

# 数据类型总结

## 简单类型

> 也叫基本数据类型或值类型

- string、number、boolean、undefined、null（空的object）  
- 如果有个变量以后打算存储为对象，暂时没想好放什么，就给null

## 复杂类型

> 也叫引用类型

- 用new关键字创建，Object、Array、Date等

## 堆栈

- 栈：操作系统自动分配释放内存，存放参数、局部变量等变量的值
- 堆：手动分配释放内存，复杂数据类型的值存在堆里，若不手动释放，由垃圾回收机制回收
- 简单数据类型的值存在栈里；复杂数据类型的地址存在在栈里，值存放在堆里

# Web API

## 简介

### API

Application Programming Interface，应用程序编程接口，一些预先定义的函数，目的是提供应用程序与开发人员基于软件或硬件得以访问一组例程的能力，而无需访问源码或理解内部工作机制的细节。

### Web APIs

是浏览器提供的一套操作浏览器功能和页面元素的API（BOM和DOM）

## DOM

Document Object Model，文档对象模型，通过DOM接口可以改变网页的内容、结构和样式

### DOM树

![image-20231221155833823](./JavaScript笔记.assets/image-20231221155833823.png)

- 文档：一个页面就是一个文档，DOM中用document表示
- 元素：所有的标签就是元素，DOM中有element表示
- 节点：网页中的所有内容都是节点（标签、属性、文本、注视等），DOM中用node表示
- DOM中把以上内容都看做对象

### 获取元素

> 以下函数的传参都是字符串

- 根据ID获取：document.getElementById(id)，id是大小写敏感的字符串，找到则返回DOM的element对象，否则返回null
- 根据标签名获取
  - document.getElementsByTagName(name)，找到则返回一组对象（伪数组的形式存储），否则返回null，得到的元素是动态的（原始元素中的内容发生变化，得到的就会变化）
  - element.getElementsByTagName(name)
    - document.getElementsByTagName(name)[index].getElementsByTagName(name)
    - document.getElementById(id).getElementsByTagName(name)
- 通过H5新增的方法获取
  - document.getElementsByClassName(name)
  - document.querySelector(选择器)，返回指定选择器选中的第一个元素
  - document.querySelectorAll(选择器)，返回指定选择器选中的所有元素
- 特殊元素获取
  - document.body，获取body标签
  - document.documentElement，获取html标签


### 事件

触发-响应机制，网页中的每个元素都可以产生某些可以触发js的事件，事件由三部分组成

- 事件源，触发事件的对象，html中的元素
- 事件类型，如何触发，比如鼠标点击(onclick)、鼠标经过、键盘按下等
- 事件处理程序，处理事件的函数

  ==事件函数中也可以用this，指向的是事件函数的调用者==

```js
<script>
  var btn = document.getElementById('btn');	// 获取事件源

	// 绑定/注册事件
	btn.onclick = function() {
    alert('点击了按钮');
    this.disabled = true;
  }	//  
</script>
```

| 鼠标事件    | 触发条件         |
| ----------- | ---------------- |
| onclick     | 鼠标点击左键触发 |
| onmouseover | 鼠标经过触发     |
| onmouseout  | 鼠标离开触发     |
| onfocus     | 获得鼠标焦点触发 |
| onblur      | 失去鼠标焦点触发 |
| onmouseup   | 鼠标弹起触发     |
| onmousedown | 鼠标按下触发     |

### 操作元素

> 利用DOM来改变元素的内容、属性等，可以通过事件来修改，也可以直接修改，打开或刷新页面后直接显示修改后的内容。

#### 修改或获取元素内容

- element.innerText，非标准，从起始位置到终止位置到内容，不识别标签，且去掉空格和换行
- element.innerHTML，W3C标准，从起始位置到终止位置到内容，识别标签，且会包含空格和换行

#### 常用元素的属性操作

- innerText、innerHTML，改变元素内容
- src、href
- id、alt、title

#### 表单元素的操作

- type 
- value
- checked
- selected
- disabled

样式属性操作（自动把修改后的样式设为行内样式，因此优先级比css高，会覆盖css）

- element.style，行内样式操作，如：`this.style.backgroundColor = ''`
- element.className，类名样式操作

## BOM