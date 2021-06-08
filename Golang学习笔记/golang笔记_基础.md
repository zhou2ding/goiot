# day01

## 常用命令

- `go build`
- `go build -o xxx`
- `go mod init`
- `go mod init xxx（文件夹）`
- `go doc builtin.delete(函数名，builtin是内置函数的意思)`
- 待补充



## 引用类型

要在内存中开辟空间后才能赋值

- slice
- map
- channel

## 字符串

### 常用函数

`s := fmt.Sprintf("%v,%v", "a","b")`：s的值是`"a,b"`

`s := strings.Split(s, ",")`

`s := strings.contains(s, "a")`

`HasPrefix`

`HasSuffix`

`Index`

`LastIndex`

`Join([]string elem, sep string)`：把一系列字符串拼接起来，中间用sep分隔

### 长度

```go
	fmt.Println(len("a")) // 1
	fmt.Println(len("啊")) // 3
	fmt.Println(len(string('a'))) // 1
	fmt.Println(len(string('啊'))) // 3
	fmt.Println(len(string("a啊"[0]))) // 1
	fmt.Println(len(string("我啊"[0]))) // 2
	s1 := "a啊"
	s2 := "我啊"
	b1 := []rune(s1)
	b2 := []rune(s2)
	fmt.Println(len(string(b1[0]))) // 1
	fmt.Println(len(string(b2[0]))) // 3
	b3 := []byte(s1)
	b4 := []byte(s2)
	fmt.Println(len(string(b3[0]))) // 1
	fmt.Println(len(string(b4[0]))) // 2
```

## 数组

- 数组如果不初始化，则元素默认为零值

- 声明时必须指定存放元素的类型和数组的容量（长度），数组的长度是数组类型的一部分

- 初始化方式

  - ```go
    //方式1
    var a1 [3]bool
    a1 = [3]bool{true,false,true}
    ```

  - ```go
    //方式2
    a10 := [...]int{0,1,2,3,4,5,6,7} //根据初始值自动推断长度
    ```

  - ```go
    //方式3
    a3 := [5]int{1,2} //不足的默认补零值
    a4 := [5]int{0: 1,4: 2} //根据索引初始化，其他的零值
    ```

- 遍历数组

  - ```go
    //方式1：根据索引遍历
    cities := [...]string{"北京","上海","深圳"}
    for i := 0; i < len(cities); i++ {
        fmt.Println(cities[i])
    }
    ```

  - ```go
    //方式2：for range遍历
    for i, v := range cities {
        fmt.Println(i, v)
    }
    ```

- **数组是值类型**

  ```go
  b1 := [3]int{1, 2, 3} //[1 2 3]
  b2 := b1 //[1 2 3]
  b2[0] = 100 // b1是[1 2 3]，b2是[100 2 3]
  ```

- 多维数组

  - 多维数组只能放相同类型的元素

  - 只有最外层才能用自动推断长度的方式
  
  - ```go
    //多维数组初始化
    var a11 [3][2]int
    a11 = [3][2]int {
        [2]int{1, 2},
        [2]int{3, 4},
        [2]int{5, 6}
    }
    //多维数组遍历
    for _, v1 := range a11 {
        fmt.Println(v1)
        for _, v2 := range v1 {
            fmt.Println(v2)
        }
    }
    ```

## 切片

- 包含三个元素：指向底层数组第一个元素的指针，切片的长度、容量

- 切片是引用类型（切片不存值），改一个值其他的指向同一个底层数组的切片都会改

- 切片不保存具体的值，指向一个底层数组，而底层数组都是占用一块连续的内存

- 初始化化方式

  ```go
  s1 := []int{1, 3, 5} //方式1
  a2 := [...]int{1, 2, 3, 4}
  s2 := a2[1:3] //方式2
  s22 := a2[:2]
  s222 := a2[1:2:3] // 容量是3-1；1可以省略，默认为0
  s3 := make([]int, 0, 10) //方式3
  //append
  s3 = append(s3, 5, 6)
  _ = append(s3, 5, 5)
  s2 = append(s2, s3...)
  ```
  
- 切片的遍历：和数组类似

- 切片的常用函数

  - `append`扩容
    - 会自动初始化，不用手动分配内存
    - 调用append函数必须用原来的切片变量接收返回值，原因如下
    - 原来的底层数组放不下来的时候，Go就会换一个底层数组
    - 扩容策略
    ![1615373093431](D:\资料\Go\src\studygo\Golang学习笔记\golang笔记_基础.assets\1615373093431.png)
  - `copy(dst, src)`
    - 切片没有长度或内存无法copy（copy之后不改变dst的值）
    - copy之后，dst和src的底层数组不一样，所以改值不影响；append也不会改
    - 当dst长度大于src长度时，src的赋值过来后，只覆盖src长度范围内的值，dst多出来的值不受影响
  - 删掉`a[k]`元素
    - `a1 = append(a1[:k], a1[k+1:]...)`
    - 此时底层数组被修改，效果等于把`a[k+1]`及的之后的值拷贝到了从`a[k]`开始的位置上
    - 删掉多个元素时，方法类似

## 指针

Go语言不存在指针操作，只有取地址和取值

- 取地址`&`
  - 不能对字面值取地址
  - 指针要先new，对刚new完的指针取值，得到的是对应类型的零值
  - new用于基本数据类型、结构体分配内存，返回的是类型指针
  - make用于slice、map、chan分配内存，返回的是类型本身（因为这三种本身就是引用类型）
- 取值`*`（得到的是指针指向的变量）

## map

- map是引用类型，必须分配内存后才能赋值

- 如果找的key不存在，则返回值的零值

  ```go
  // map初始化方法1：
  m := make(map[string]int, 10)
  m["我是"] = 18
  
  // 判断键
  if v, ok := m["不存在的键"]
  if !ok {
      ...
  } else {
    ...
  }
  // map初始化方法2：
  m1 := map[string]int {
      "stu1": 18,
      "stu2": 26,
  }
  ```

- 遍历map

  - 只遍历key，for range可以只有key
  - 遍历key和value，则两个都要有

- 删除`delete(m,"我是")`，删除不存在的键，什么都不会发生

- 元素类型为map的切片：初始化切片后要对map元素初始化

- 值为切片类型的map：构造键值对时要对切片初始化

# day02

## 函数

### 注意点

- 返回值可以命名也可以不命名（参数必须命名），命名的返回值就是返回值变量，要用括号括起来；没命名的返回值其实是个匿名变量

- 命名参数、返回值就相当于在函数中声明变量（命名的参数、返回值可以不使用）

- 命名的返回值，可以在return后面省略（不给返回值赋值时，返回零值）

- 当参数中连续多个参数类型一致时，可以省略前面参数的类型

- 可变长参数
  - `func f (x string, y ... int)`
  - y其实是个切片
  - y必须放在所有参数的最后
  
- Go语言没有默认参数的概念，只要命名了参数就必须传参

- 多个返回值

  ```go
  func f(x, y int) (sum, sub int) {
      sum = x + y
      sub = x - y
      return
  }
  ```

- 函数的返回值是一个单独的变量

  - `return x`实际上是把x赋值给返回值
  - 若已经命名了返回值，则此变量就是返回值变量，`return 10`相当于把10赋值给已经声明的返回值变量

- 函数传参，全部传递的是值

### defer

- `defer`把跟在它后面的语句延迟到函数即将返回的时候再执行（碰到panic的错误，会先执行发生panic的语句，然后执行defer，然后返回panic）
- 使用场景：函数结束之前（return、panic等）释放资源
  - 关闭文件
  - 关闭socket连接
  - 关闭数据库连接
- 一个函数中有多个`defer`时，倒序执行`defer`（先进后出）
  - Go的`return`语句在底层不是原子操作（`return x`一步执行完成就是原子操作）
  - `defer`语句执行时机在给返回值赋值之后，真正的`RET`返回之前
  ![1615902755944](D:\资料\Go\src\studygo\Golang学习笔记\golang笔记_基础.assets\1615902755944.png)
- defer延迟调用函数时，**会先把参数中的结果确定出来**

### 作用域

- 函数中查找变量的顺序
  1. 在函数内部查找（函数内部定义的变量只能该函数内部使用）
  2. 找不到就往函数外面查找，一直找到全局变量
- 分类
  - 语句块作用域
  - 函数作用域
  - 全局作用域

### 函数类型

- 根据函数的形式，有不同的类型：`func()`类型、`func() int`、`func(int, int) int`类型。。。
- 函数类型的变量可以作为参数和返回值
- 带参数的函数类型赋值给变量时，可以通过变量名加括号来调用函数

### 匿名函数

- 没有名字的函数（一般用在函数内部）

  ```go
  func main() {
      f1 := func(x, y int) {
          fmt.Println(x + y)
      }
      f1(10, 20) // f1可以调用多次
      
      // 如果只是调用一次，可以简写成立即执行函数
      func(x, y int) {
          fmt.Println(x + y)
      }(100, 200)
  }
  ```

### 闭包

- 概念：函数和它所引用的外部变量，常用于返回值是一个函数类型的函数

- 应用场景

  - 把类型不匹配的函数当参数传给另一个函数

    ```go
    // 需求：把f2传参给f1
    
    func f1(f func()) {
        fmt.Println("this is f1")
        f()
    }
    
    func f2(x, y int) {
        fmt.Println("this is f2")
        fmt.Println(x + y)
    }
    // 把f2包装一下
    func f3(f func(int, int), x, y int) func() {
        tmp := func() {
            f(x, y)
        }
        return tmp
    }
    
    func main() {
        ret := f3(f2, 100, 200)
        f1(ret)
    }
    ```

  - 把调用函数的时机改一下

    ```go
    func adder(x int) func(int) int {
        return func(y int) int {
            x += y
            return x
        }
    }
    
    func main() {
        ret1 := adder(100) // ret1指向匿名函数，和其所引用的外层变量x
        ret2 := ret1(200) // ret2 = 300
        ret3 := ret1(50) //  ret3 = 350
    }
    ```

  - 网上搜的应用场景

    -  `HTTP router middleware`，比如`go-chi/chi`的[middleware](https://github.com/go-chi/chi/tree/master/middleware) 
    -  你可以使用闭包来做装饰器 
    -  函数式编程中这种闭包的使用比较常见，特别是在nodejs中，几乎成了高级语言的标配 
    -  需要临时开启goroutine时需要 go func(){}()
      在panic的recover时使用 
    -  需要调用函数, 且该函数与运行环境有关系时, 就使用闭包 

### 函数嵌套

- 在一个命名的函数中不能再声明一个命名函数
- 可以声明一个匿名函数

### 内置函数

- `len()`求长度，用于string、数组、slice、map、chan
- `make()`分配内存，用于引用类型（slice、map、chan），返回对应类型本身
- `new()`分配内存
  - 用于值类型（基本数据类型、struct），返回对应类型的指针
  - 结构体new之后没有地址，要给结构体中的元素初始化后才有地址，且地址已经固定，再给其他元素赋值或改值都不会改变地址
- `close()`关闭channel
- `append()`给数组和切片追加元素
- `panic/recover`
  - `panic()`会让程序崩溃退出
  - 必须在`panic`之前`defer `释放连接
  - `recover()`会把`panic`的信息保存起来，返回一个error，不会让程序退出（不推荐用）；但发生panic的语句之后的程序都不会执行了

## fmt包

1. 占位符
   - `%d, %02d, %T, %c, %t, %b, %o, %x, %X, %f, %F, %U, %s, %p, %q, %v, %+v, %#v, %b, %e, %E, %g, %G, %+, %-`
   - %后可带宽度和精度
2. 其他常用函数
   - `fmt.Errorf()`：格式化输出后返回一个`error`类型
   - `errors.New()`：创建一个`error`类型
3. 获取输入（应该传指针）
   - `fmt.Scan(a ...interface{})`
     - 从cmd扫描文本，读取由空格符保存的值传递给本函数的参数，换行符视为空格
     - 如果读取的数据个数比提供的参数少，则返回错误
   - `fmt.Scanf()`：`fmt.Scanf("%s, %d, %s\n", &name, &age, &email)`
   - `fmt.Scanln()`：`fmt.Scanln(&name, &age, &email)`，检测到空白符就结束
   - `fmt.Fscan(r io.Reader, a ...interface{})`系列：从`r`扫描文本，将成功读取的空白分隔的值保存进成功传递给本函数的参数
   - `fmt.Sscan(str string, a ...interface{})`系列：同上
4. `fmt.Sprint`系列：拼接字符串，返回一个字符串
5. `fmt.Fprint(w io.Writer, a ...interface{})`系列：往`w`中打印内容，如`os.StdOut`、`os.*File`等

# day03

## 结构体基础

- 结构体是值类型
- 结构体中的元素占连续的内存空间
- 空结构体不占空间

### 类型别名和自定义类型

- 自定义类型
  - `type`后面跟的是类型，`type myInt int`
  - 打印出的类型是`函数.自定义类型`
- 类型别名：`type myInt = int`
  - 只在代码编写中有效，编译完成后就不存在了
  - 打印出来的类型就是原类型

### 自定义结构体

- ```go
  type person strurc {
      name string
      age int
      gender string
      hobby []string
  }
  ```

- 先声明，后初始化

  ```go
  func main() {
  	var p person
      // 没有给元素赋值时，自动给零值
      p.name = "张三"
      p.age = 18
      p.hobby = []string{"吃","喝","玩"}
  }
  ```

- **如下两种初始化方式不能混用**

- key-value初始化（声明的同时初始化）（常用）

  ```go
  func main() {
      // 可直接用&得到结构体指针，不用new
      var p2 = person {
          name: "李四",
          age: 15,
          //gender不赋值，默认为零值
          hobby: []string{"吃","玩"},
      }
  }
  ```

- 值列表初始化（初始化值的顺序必须和结构体定义时元素的顺序一致）

  ```go
  func main() {
      p3 := person {
  		"王五",
          12,
          "男",
          []string{"睡觉","玩乐"}
      }
  }
  ```

### 匿名结构体

直接把变量声明为结构体，多用于临时场景

```go
// 方式1
var s struct {
    name string
    age int
}
s.name = "张三"
s.age = 18
// 方式2
var a = struct {
    name string
    age int
}{"李四", 25}
```

### 指针类型结构体

- 使用`new`关键字得到结构体类型的指针

  ```go
  func main() {
      p := new(person) // 不能直接&person
      p.name = "张三"
      p.age = 18
      p.hobby = []string{"吃","喝","玩"}
  }
  ```

## 结构体相关

### 构造函数

构造函数：返回一个结构体变量的函数（一般都返回指针类型）

```go
// golang是面向接口编程，不像面向对象编程的语言那样自带构造函数
type person struct {
    name string
    age int
}

// 结构体中的元素比较多时，要返回结构体指针，减少程序运行的内存开销
// 约定成俗，Go中的构造函数都是new开头
func newPerson(name string, age int) *person {
    return &person {
        name: name,
        age: age,
    }
}
```

### 方法和接收者

- 作用于特定类型的函数，即为方法

- 在函数名前指定接收者，即限定了调用此函数的特定类型

- 想改元素的值时，接收者要是指针类型（一般都使用指针接收者）

  ```go
  type person struct {
      name string
      age int
  }
  
  // p是调用该方法的具体类型的变量，多用类型名首字母小写表示
  func (p person) bark() {
      fmt.Printf("%s在叫，他今年%d岁了", p.name, p.age)
  }
  
  // 需要修改接收者元素的值时，要用指针接收者
  // 接收者是拷贝代价比较大的对象
  // 如果某个方法给了指针接收者，其他方法也要统一用指针接收者
  func (p *person) grow() {
      p.age++
  }
  ```

### 给自定义类型加方法

- 不能给基本数据类型加方法

- 不能给别的包的类型添加方法，只能给自己的包里的类型加方法

- 可以自定义类型后再添加方法

  ```go
  type myInt int
  func (m myInt) hello() {
      fmt.Println("我是一个myInt")
  }
  func main() {
      m := myInt(100)
      m.hello()
  }
  ```

### 匿名字段

- 相同类型只能写一个字段
- 场景：字段比较少也比较简单
- 匿名字段并不常用

```go
type person struct {
    string
    int
}
func main() {
    p1 := person {
        "张三"，
        18
    }
    fmt.Println(p1.string)
    fmt.Println(p1.int)
}
```

### 嵌套结构体

- 场景：多个结构体中有重复字段时
- 匿名嵌套结构体
  - 可以直接用变量获取嵌套结构体的字段（比较常用），详见如下代码
  - 当两个匿名嵌套结构体中有相同字段时，会冲突，不要这么写

```go
type address {
	city string
    province string
}
type location {
    country string
    avenue string
}
type person struct {
    name string
    age int
    addr address // 嵌套结构体
    location // 匿名嵌套结构体
}
type company struct {
    name string
    addr address
}
func main() {
    // 嵌套结构体
    p1 := person {
        name: "张三",
        age: 18,
        addr: address {
            city: "北京",
            province: "河北"，
        }，
    }
    fmt.Println(p1.addr.city)
    // 匿名嵌套结构体
    p2 := person {
        name: "李四",
        age: 15,
        location: location {
            country: "美利坚",
            avenue: "皇后区"，
        }
    }
    fmt.Println(p2.avenue)
}
```

### 继承

golang没有继承，可用结构体模拟继承（匿名嵌套结构体）

```go
type animal struct {
    name string
}
type dog strcut {
    feet int
    animal
}
func (a *animal) move() {
}
func (d *dog) bark() {
}
func main() {
    d1 := dog {
        feet: 4,
        animal: animal {name: "张三"},
    }
    d1.bank()
    d1.move() // 继承了animal的move方法
    fmt.Println(d1.name)
}
```

## JSON

- 跨语言的数据格式：JavaScript Object

- JSON是一个字符串，不同的语言会把JSON字符串转换成对应语言的对象

  - 序列化：把Go中的结构体转成JSON字符串

    ```go
    import "encoding/json"
    type person struct {
        Name string // 首字母大写后，其他包可见，才能进行json转换；但转换后的字段也是大写
        Age int `json:"age"` // 会有字段是小写的需求，要用此写法
        Id int65 `json:"id" db:"id" ini:"id"`
    }
    
    func main() {
        p1 := person {
            Name: "张三",
            Age: 18,
        }
        b, err := json.Marshal(p1) // b是一个[]byte
        if err != nil {
            // 错误处理
        }
        fmt.Println(string(b))
    }
    ```

  - 反序列化：把JSON字符串转成Go能识别的结构体变量，必须传指针

    > 1. 当对函数的返回值进行反序列化，即使返回值的类型是引用类型，也必须传指针
    > 2. 直接在main函数中对引用类型的变量反序列化，传变量本身即可
    
    ```go
    func main() {
        str := `{"Name":"张三", "age":18, "id":10010}`
        var p2 person
        // 把str取出来后放到p2中去，传指针是为了能修改p2的值
        json.Unmarshal([]byte(str), &p2)
        fmt.Println(p2)
    }
    ```

# day3

## 接口（interface）

### 概述

- 接口是一种类型，它规定了变量有哪些方法，接口就是一个==需要实现的方法列表==
- 场景

  - 不关心函数传进来的参数是什么类型，只关心这个参数能调用的方法
  - 当传进来不同类型的变量，但这些变量有统一的方法
- 接口类型的变量存储时存两部分：
  - 动态类型
  - 动态值
- 注意事项：只有当两个或以上的具体类型必须以相同的方式进行处理时才需要定义接口


  ```go
  type speaker interface {
      speak()
  }
  type cat struct{}
  type dog struct{}
  func beat(x speaker) {
      x.speak() // 只要实现了speak方法，就属于speaker类型的变量
  }
  func (c cat) speak(){}
  func (d dog) speak(){}
  func main() {
      var(
      	c cat
        d dog
      )
      beat(c)
      beat(d)
      
      var ss speaker // 此时ss的类型为nil
      ss = c // cat类型赋值给speaker类型的变量，此时ss的类型为main.cat
      ss = d
  }
  ```

### 接口的定义

用来给变量/参数/返回值等设置类型

```go
type 接口名 interface {
    方法名1()
    方法名2(参数1) (返回值1) // 参数的变量名可以省略
}
```

### 接口的实现

一个变量如果实现了接口中规定的所有方法，那么此变量就实现了这个接口，即可以称为这个接口类型的变量

### 值接收者和指针接收者

- 使用值接收者实现了接口的方法后，结构体类型、结构体指针类型的变量都能赋值给接口类型
- 使用指针接收者实现了接口的方法后，只能把结构体指针类型的变量赋值给接口类型

### 接口和类型的关系

- 多个类型可以实现同一个接口

- 一个类型可以实现多个接口

- 接口也可以嵌套

  ```go
  type animal interface {
      mover
      eater
  }
  type mover interface {
      move()
  }
  type eater interface {
      eat(string)
  }
  ```

### 空接口

- 任何类型都实现了空接口，即任意类型的变量都能保存到空接口中

- 当作为函数的参数时，传任何变量都可以

- 空接口没有必要起名字

  ```go
  interface{} // 空接口
  ```

- map的键、值类型也可以是空接口

  ````go
  func main() {
      var m1 map[string]interface{}
      m1 = make(map[string]interface{}, 16)
  }
  ````

### 类型断言

- `x.(T)`：根据传进来的类型返回是否猜对了类型
- 场景：想知道传进来的空接口的类型

```go
func assign(a interface{}) {
    // 方式1
    str, ok := a.(string) // 猜对了就把空接口a转成string，猜错了就返回传进来的类型的零值
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("猜对了，", str)
	}
    // 方式2
    switch t := a.(type) {
        case string:
        	fmt.Println("是一个字符串", t)
        case int64:
        	.....
    }
}
```

## 包（package）

### 概述

- `main包`：
  1. `main包`才能编译成`exe`
  2. `main函数`是程序的入口
  3. `main包`不一定叫`main.go`
  4. 非`main包`也可以有`main函数`，但无法`go run`，也不能编译成`exe`
- 在包中定义的标识符，首字母大写才能被其他包调用
- 一般情况下包名和文件夹名是一样的，一个文件夹就是一个包
- import包
  - 默认把目录名当成包名
  - 三种特殊情况
    - 可以起别名
    - `_`匿名导入，表示只调用init函数，不导入整包，不使用包中的任何标识符
    - `.`表示调用包中的标识符时，可以省略包名（不用此方式，容易混乱）
  - 内置包和第三方包间会自动隔一行
  - 包名的导入是从`GOPATH/src`路径后开始的，分隔符用`/`（有了go module后被弃用）
  - 禁止循环导入（A导B，B导C，C导A），禁止导入后不使用

### init函数

- 导入包语句会自动触发包内部的init函数

- 没有参数、返回值，一个包里只有一个init

- 不能手动调用，只能自动调用（导入包时调用，main函数执行前调用）

- 执行顺序

  ![1616591545897](D:\资料\Go\src\studygo\Golang学习笔记\golang笔记_基础.assets\1616591545897.png)

  

  ![1616591517045](D:\资料\Go\src\studygo\Golang学习笔记\golang笔记_基础.assets\1616591517045.png)

## 文件操作

### 打开文件

- `os.Open()`可以打开一个文件（打开后只能读），返回一个`*os.File`类型的变量和一个`error`类型的变量，函数中可以传相对路径和绝对路径
- 记得关闭文件`defer fileObj.Cloes()`，要放在`if err != nil`后面，否则会出现空指针调用引发panic

### 读文件

- `fileObj.Read()`读文件：按字节读

  - 接收一个byte切片，表示一次读多少字节，返回读到的字节数（切片的长度）和err，把读到的内容存到切片中，读到末尾时返回0和EOF错误

  - 读指定的长度

    ```go
    var tmp [128]byte
    for {
        n, err := fileObj.Read(tmp[:])
        if err != nil {
            return
        }
        if err == os.EOF {
            return
        }
        fmt.Println("读了",n,"个字节")
        fmt.Println(string(tmp[:n])) //每读128个字节就把读到的内容打印出来
        // 读完了
        if n < 128 {
            return
        }
    }
    ```

- `bufio`读文件：按行读

  - 优点：快；缺点：断电后数据会丢失

  - `ReadString()`接收一个字符，返回读到的字符串和err（按行读时就返回读到的整行）

    ```go
    for {
        reader := bufio.NewReader(fileObj) //先创建一个从文件读内容的对象
    	line, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        if err == os.EOF {
            return
        }
        fmt.Print(line)
    }
    ```

  - `ReadLine()`：是一个低水平的行数据读取原语。大多数调用者应使用`ReadBytes('\n')`或`ReadString('\n')`代替，或者使用`Scanner`

- `ioutil`读文件

  - `ioutril.ReadFile()`

    - 接收文件路径字符串，返回一个byte切片和一个error，存的是读到的内容
    - 自带打开和关闭文件

    ```go
    ret, err := ioutil.ReadFile("文件路径")
    if err != nil {
        return
    }
    fmt.Println(string(ret))
    ```

### 写文件

- 必须用`os.OpenFile()`才能写文件

- 接收三个参数

  - 文件路径的字符串
  - 标志位
    - 打开文件的方式：`os.O_xxx`（创建、只读、只写、读写、清空、追加等）
    - 标志位可以相或，底层会按照不同的bit位上的1来进行相应的操作
    - 常用：
      - `os.O_WRONLY | os.O_CREATE | os.O_APPEND`，每次打开后继续写
      - `os.O_WRONLY | os.O_CREATE | os.O_TRUNC`，每次打开后重新写
  - 文件的权限（0644，0666，777等，windows下没作用）

- `fileObj.Write()`和`fileObj.WriteString()`，分别接收一个byte切片和字符串

- `bufio.WriteString()`

  ```go
  wr := bufio.NewWriter(fileObj)
  wr.WriteString("xxx") // 写到缓存中
  wr.Flush() // 把缓存的内容刷到文件中
  ```

- `ioutil.WriteFile()`

  ```go
  str := "写入文件的内容"
  err := ioutil.WriteFile("xx.txt", []byte(str), 0666)
  ```

- `io.Copy()`：将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。

- 解决`fmt.Scan`一读到空白符就结束的方法

  ```go
  var s string
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("请输入:")
  s, _ = reader.ReadString('\n')
  fmt.Printf("你输入了:%s", s)
  ```

### seek

- 任何语言都没有办法直接在文件中插入内容，只能
  1. 创建一个临时文件
  2. 读原文件读到插入位置之前
  3. 要插入的内容写到临时文件
  4. 后续内容写入临时文件
- 移动光标的位置 
- 接收两个参数：偏移量，相对哪里的偏移量（0是文件头，2是文件尾）

## time包

### time对象

- `time.Time`表示时间类型，`time.Now()`获取当前的时间对象，为`time.Time`类型
- `time.Now().Year()`等获取年月日时分秒和`Date()`

### 时间戳

- 自1970年1月1日8点整的总秒数
- `time.Now().Unix()`或`time.Now().UnixNano()`获取时间戳，分别是秒和毫秒
- `time.Unix()`可以把时间戳转换为`time.Time`类型，接受一个时间戳，返回一个类型

### 时间间隔

- 是一些常量，包括`time.Second`等
- 利用时间间隔常量进行时间操作
  - `time.Now().Add(time.Hour)`：当前时间1小时后
  - `Sub()`：判断两个`time.Time`对象的之间的时间间隔
  - `Equal()`判断两个时间是否相等，会考虑时区的影响
  - `Before()`、`After()`，判断两个时间谁前谁后
- 数字可以直接和`time.Duration`运算，但变量不行，变量必须强转

### 定时器

- `time.Tick(time.Second)`返回一个1秒间隔的定时器（channel类型），对此定时器执行`for range`，即可每秒进行一些操作

### 时间格式化

- 把Time对象转换成字符串
  - `time.Now().Format()`：接收一个字符串，把当前时间转换成字符串
  - 格式化字符串`2006-01-02 15:04:05.000`或`2006-01-02 03:04:05 PM`
- 把字符串转成时间
  - `time.Parse()`：按照对应的格式解析字符串
  - 接收两个参数，分别是时间格式和需要解析的时间字符串，返回一个`time.Time`类型，两个字符串必须格式一致
- `time.Sleep()`：程序等待多长时间后继续往下走，接收一个`time.Dutarion`类型，可以直接传数字，要是传int等类型的变量，要进行类型转换

### 时区

- `time.Now()`：本地时间
- `time.Parese()`：第二个参数是UTC时间
- `time.LoadLocation("Asia\Shanghai")`：返回一个`*Location`类型的变量和一个 `error`
- `time.ParseInLocation()`：接受三个参数：时间格式，要解析成`Time`类型的字符串，`*Location`变量，返回一个`Time`类型变量和一个`error`

## log包

- `log.Println()/Printf()/Print()`：在终端中打印日志，比fmt打印的东西多了时间
- `log.Fatal/Fatalln/Fatalf`：打印比较严重的日志
- `log.SetOutPut()`：接收一个`io.Writer（os.StdOut/os.*File等）`，设置打印的位置

# 日志库作业

## 日志分级

- Debug
- Trace
- Info
- Warning
- Error
- Fatal

## 常用库

1. `pc, file, line, ok := runtime.Caller(n)`
   - n是调用的层数，想获取此语句所在函数的信息，n就是0
   - 想获取调用此语句所在函数的外层函数，n就是1往上
2. `funcName := runtime.FuncForPC(pc).Name()`
   - 可获取函数的名字
   - 还有`.Entry()`、`.FileLine()`
3. `path := path.Base(file)`
   - 可获取路径的最后一节
   - 还有 `.Clean()`、`.Dir()`、`.Ext()`、`.ErrBadPattern()`、`.IsAbs()`、`.Join()`、`.Match()`、`.Split()`、`.ErrBadPattern.Error`
4. `fileInfo, err := fileObj.Stat()`
   - 获取文件信息
   - 后面可以接`.Name()`、`.Size()`、`.Mode()`、`.ModTime()`、`.IsDir()`、`.Sys()`
5. `os.Rename(old, new)`

## 反射（很少写，关注原理）

### 概述

- 在程序运行期间对程序进行访问和修改
- 用于ORM、ini文件解析/xml解析/json解析等
-  程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息
- 支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们

### 通过反射读变量的值

> 任意接口值在反射中都可以理解为由`reflect.Type`和`reflect.Value`两部分组成，分别由如下两个函数获取

- `reflect.TypeOf(x)`：获取反射类型
  - `.Name()`、`.Kind()`，Kind是底层类型，是种类（对结构体变量调用这两个方法分别是结构体名和"struct"）
  - Kind包含：`reflect.Int64`等等

- `reflect.ValueOf(x)`：获取反射值
  - `.Kind()`
  - 后面再跟`.Int()`、`.Bytes()`、`.Float()`等获取反射值的原始值，没有8、32、64这些
- `reflecy.TypeOf(x).Elem().Kind()`：获取指针变量所指向的类型

### 通过反射设置变量的值

- 传x的时候要传指针，否则会panic

- 必须用`reflecy.ValueOf(x).Elem().SetInt(100)`：`Elem()`方法可以获取指针对应的值

### 其他

- `reflect.ValueOf(x).isNil()`：判断引用类型是否为空

- `reflect.ValueOf(x).isValid()`：判断是否持有一个值
- `reflect.ValueOf(b).FieldByName("abc")`，查找结构体b的字段“abc”
- `reflect.ValueOf(b).MethodByName("abc")`，查找结构体b的方法“abc”
- `reflect.ValueOf(b).MapIndex(reflect.ValueOf("abc"))`，查找map类型变量b的abc键

### 结构体反射

-  任意值通过`reflect.TypeOf()`获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（`reflect.Type`）的`NumField()`和`Field()`方法获得结构体成员的详细信息。
- 一般是用结构体指针来获取字段信息，所以是`reflect.TypeOf(x).Elem().Field()` 

|                            方法                             |                             说明                             |
| :---------------------------------------------------------: | :----------------------------------------------------------: |
|                  Field(i int) StructField                   |          根据索引，返回索引对应的结构体字段的信息。          |
|                       NumField() int                        |                   返回结构体成员字段数量。                   |
|        FieldByName(name string) (StructField, bool)         |       根据给定字符串返回字符串对应的结构体字段的信息。       |
|            FieldByIndex(index []int) StructField            | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) |              根据传入的匹配函数匹配需要的字段。              |
|                       NumMethod() int                       |                返回该类型的方法集中方法的数目                |
|                     Method(int) Method                      |                返回该类型方法集中的第i个方法                 |
|             MethodByName(string)(Method, bool)              |              根据方法名返回该类型方法集中的方法              |

以上方法返回的`StructField`类型是一个结构体类型

```go
type StructField struct {
    // Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
    // 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
    Name    string
    PkgPath string
    Type      Type      // 字段的类型
    Tag       StructTag // 字段的标签
    Offset    uintptr   // 字段在结构体中的字节偏移量
    Index     []int     // 用于Type.FieldByIndex时的索引切片
    Anonymous bool      // 是否匿名字段
}
```

`StructField.Tag`有`.Get("json")`方法，可以在tag中查找对应的字符串

# day04

## strconv包

**前言**

- 用`string()`强转int类型变量时，得到的是对应的ASCII值
- int类型变量不能强转成string

**用法1**

- `strconv.ParseInt(str, 10, 64)`，接受的三个参数分别是
  - 待转换的字符串
  - 进制
  - 位数
- 返回得到的`int64`变量和一个`error`
- 当传进来的位数不是64时，虽然函数返回的是`int64`类型，但此时把返回值强转成低位的类型，不会丢失
- 传的进制是0时，对应的位数是`int`，可传0、8、16、32、64

**用法2**

- `strconv.Atoi()`，仅适用于`string`转换成`int`
- 接收一个字符串，返回一个`int`类型变量和`error`，更简单
- `strconv.Itoa()`，仅适用于`int`转换成`string`

**其他**

- `strconv.ParseBool()`接收一个字符串，返回`bool`和`error`
- `strconv.ParseFloat()`接收一个字符串和位数，返回`float64`和`error`

## rand包

**常用函数**

- `rand.Int()`返回一个`int64`的随的数，一般都特别大；`rand.Int63()`、`rand.Int31()`分别返回int64、int32类型的随机数
- `rand.Intn(n)`返回`[0,n)`的随机数；int64、32位随机数的获取和上面类似

**随机数种子**

- 不加种子的话，每次执行程序，得到的随机数都一样
- `rand.Seed()`添加种子，接收一个`int64`变量，一般把`time.Now.UnixNano()`传进去
- `Seed`不能和其他`rand`函数并发执行

# 并发

## 概述

### 相关概念

- Go天生为并发而生，其他语言是程序调用OS的线程接口（内核态），需要程序员维护一个线程池，自己包装任务、维护上下文切换
- 而Go是自己模拟出线程（用户态），更轻量级
- `goroutine`是`runtime`调度的，而OS线程是OS自己调度

### 并发并行

- 并发：同一时间段内，同时执行多个任务
- 并行：同一时刻，同时执行多个任务

## goroutine

### 概述

- 一个`goroutine`必定对应一个函数
- 程序启动之后会自动创建一个`main goroutine`去执行，是和main函数中启动的goroutine并发执行的
- `main`函数结束后，在`main`函数中启动的`goroutine`也都结束了

### 使用

- `go f()`：开启一个单独的`goroutine`去执行`f()`函数

- 启动匿名函数的`goroutine`

  ```go
  func main() {
      for i := 0; i < 1000; i++ {
          go func() {
              fmt.Println(i) //闭包，打印的i会有重复的
          }()
      }
      for i := 0; i < 1000; i++ {
          go func(i int) {
              fmt.Println(i) //避免闭包
          }(i)
      }
      time.Sleep(time.Second) //main函数执行太快了，不这样的话可能goroutine刚启动，main就结束了
  }
  ```

- 更高级的方式等待`goroutine`结束`waitGroup`

  ```go
  var wg sync.WaitGroup
  func f(i int) {
      defer wg.done() //函数结束后，计数器减1
      ...
  }
func main() {
      for i := 0; i < 10; i++ {
          wg.Add(1) //开启一个goroutine前，先给计数器加1
          go f(i)
      }
      wg.Wait() //等待wg的计数器减为0
  }
  ```
  

### GMP

- OS线程的栈内存是固定的，一般2MB；goroutine的栈内存是按需动态变化的，初始一般2KB，最大能到1GB
- `G`：goroutine
- `P(Processor)`：P管理着一组goroutine队列，通过`runtime.GOMAXPROCS`设定P的数量，最大256，默认为物理线程数；不设置的话，默认跑满CPU
- `M(Machine)`：Go的`runtime`对OS线程的虚拟，是真正干活的，goroutine都是放到M上运行
- P管理一组G在M上运行，发生阻塞时，runtime会新建一个M，并把阻塞的G所属的P中他其他G挂载到新建的M上，待阻塞的M完成或死掉时，回收
- P与M一一对应，M与OS线程一一对应；P的默认数量不一定是CPU核心数，一些CPU有超线程技术：双核4线程；G和OS线程是多对多，即`m:n`技术

## channel

### 概述

- 共享内存（全局变量等）在goroutine中会发生竞态问题；为了避免此，又需要加互斥锁，但此时又变成了串行
- Go提出**通过通信来共享内存**，即`CSP(Communicating Sequential Processes)`模型
- `channel`是一种类型，遵循先入先出的规则

### 使用

- `channel`是一种引用类型，需要初始化分配内存后才能使用

- 对于元素类型比较大的，一般存指针(`string`等类型)

- 声明时需要指定通道中元素的类型

  ```go
  var b chan int
  ch1 = make(chan int) // 无缓冲区的通道的初始化
  ch1 = make(chan int, 16) // 带缓冲区的通道的初始化
  ```

#### 发送

- `ch1 <- 1`

- 向无缓冲区的`channel`发送数据，会发生`deadlock`的error

- 解决方法：另启一个从无缓冲区通道中接收数据的`goroutine`

  ```go
  var wg sync.WaitGroup
  var ch1 chan int
  ch1 = make(chan int)
  wg.Add(1)
  go func(){
      defer wg.Done()
      <- ch1
  }()
  ch1 <- 10
  wg.Wait()
  ```

#### 接收

- `x := <- ch1`：接收的值存到变量x中
- `<- ch1`：丢弃接收的值

#### 关闭

- `cloese(ch1)`，不关也可以，会自动回收，但最好关了；
- 如果多个`goroutine`操作同一个通道的话，所有`goroutine`都结束之后再关
- 当通过中的值取完时，如果继续取值
  - 未关闭通道，会报`deadlock`错误
  - 关闭通道，会取到零值；两个变量接收的话就是零值和`false`

#### 从通道循环取值

`for range`可以从通道循环取值

```go
// 1. 启动一个goroutine，生成100个数发送到ch1中
// 2. 启动一个goroutine，从ch1中接收值，计算平方后发送到ch2中
// 3. 在main中，从ch2中接收值，对值进行操作
var wg sync.WaitGroup
func f1 (ch1 chan int) {
    defer wg.Done()
    for i := 0; i < 100; i++ {
        ch1 <- i
    }
    close(ch1)
}
func f2(ch1, ch2 chan int) {
    defer wg.Done()
    for {
        x, ok := <- ch1 // f1和f2是并发执行的，不这样的话可能从f1中取不到值，报deadlock错误
        if !ok { // 从ch1中读完后，返回false
            break
        }
        ch2 <- x * x
    }
    close(ch2)
}
func main() {
    a := make(chan int, 50) // 有go f2中的b通道在接收，a的缓冲小于100也没事
    b := make(chan int, 100) // 没有其他线程接收b的数据，b的缓冲必须>=100
    wg.Add(2)
    go f1(a)
    go f2(a, b) //f1和f2必须都启动线程，只有一个启动的话，属于goroutin泄露而非deadlock
    for ret := range b {
        fmt.Println(ret)
    }
    wg.Wait()
}
```

#### 单向通道

- 只写通道：`var ch1 chan<- int`

- 只读通道：`var ch2 <-chan int`
- 作用：作为函数的参数，限制只能向该参数发送 / 从该参数接收

#### 异常情况

![channel异常总结](https://www.liwenzhou.com/images/Go/concurrence/channel01.png) 

#### water pool(goroutine池)

```go
// 从jobs接收值发送到results
func worker (id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2 //从jobs中取到值后，进行一些耗时的复杂操作
    }
}
func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    for w := 1; w <= 3; w++ { // 一般128个goroutine
        go worker(w, jobs, results) // 开启三个goroutine，去执行五个任务
    }
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    for r := 1; r <= 5; r++ {
        <- results;
    }
}
```

启动n个`goroutine`，循环操作`jobChan`中的无限个元素

### select

- 随机对某个通道进行某种通信

  ```go
  for {
      select {
          case <-ch1:
              ...
          case x := <-ch2:
              ...
          case ch1<-x:
          	...
          default:
          	...
  	}
  }
  ```

## 并发安全和锁

### sync包

>  互斥锁
>
> - 一种常用的控制共享资源访问的方法，它能够保证同时只有一个`goroutine`可以访问共享资源 
> - `lock := sync.Mutex`，是一个结构体，当参数传时要传指针，使用方法：
>    - `lock.Lock()`
>     - 操作共享的资源
>     - `lock.Unlock()`

>  读写互斥锁
>
>  - 很多场景对数据库的读写是分离的，一个写的主库一个读的从库，因为读的次数远大于写的次数
>    - 读锁：`rwlock := sync.RWMutex`：`rwlock.RLock()`，`rwlock.RUnlock()`
  >    - 写锁：`rwlock := sync.RWMutex`：`rwlock.Lock()`，`rwlock.Unlock()`
  >    - 一个`goroutine`获取读锁后，其他`goroutine`能继续读，写的话要等待
>    - 一个`goroutine`获取写锁后，其他`goroutine`读写都要等待

>  sync.Once
>
>  - 一些场景中，某个操作只需要做一次，如加载配置文件、关闭`channel`；Once是一个结构体，里面是一个标志位和一个锁，false的话加锁，true的话继续往后走
>  - `var once sync.Once`，`once.Do()`，`Do()`只接受无参、无返回值的函数类型

  ```go
  var once sync.Once
  func f1(ch1 <-chan int, ch2 chan<- int){
      once.Do(func() { close(ch2) }) // 关闭通道只执行一次
  }
  ```

>  sync.Map
>
> - Go内置的`map`不是并发安全的，最多支持20个`goroutine`操作同一个map（20以内也会有问题，只是不报错），超过20个时，会报`fatal error: concurrent map writes`
> - `m := sync.Map{}`：不用初始化，直接用
>   - `m.Store(key， value)`
>   - `value, ok := m.Load(key)`
>   - `m.LoadOrStore()`：有的话就返回取到的；没有的话先存值，再返回存入的值
>   - `m.Delete(key)`
>   - `m.Range(func(key, value interface{}) bool {...})`：参数是一个函数，函数的操作即循环中每次的操作

## 原子操作 atomic包

> 作用于基本数据类型，进行一些并发安全的操作

- 读取`Load`：Int32、64，UInt32、64，UIntPtr、Pointer，**接收的是指针**
- 写入`Store`：同上
- 修改`Add`：同上
- 交换`Swap`：同上
- 比较并交换`CompareAndSwap`：同上

# day05

## HTTP

![1618305867034](D:\资料\Go\src\studygo\Golang学习笔记\golang笔记.assets\1618305867034.png)

```go
func homeHandlerfunc(w http.ResponseWriter, r *http.Request) {
	str := `<h1 style="color:red">Hello DQ!<h1>`
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/home", homeHandlerfunc)
    // 0.0.0.0表示全网都可以访问
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		fmt.Printf("start http server failed, error:%v", err)
		return
	}
}
```

- 需要频繁发送请求，如每5秒从阿里云同步接口数据：定义一个全局的client，后续发请求的操作都使用这个全局的client
- 请求次数比较少，如一周两次请求：自定义一个client，禁用`keepAlived`连接

## 单元测试

> 测试函数覆盖率：100%
>
> 测试覆盖率：60%

### 概述

- `go test`、`go test -v`，`-v`是查看测试函数名称和运行时间
- `go test -run=Testxxx/case_x`
- `go test -cover`：检查你的代码覆盖率；`go test -cover -coverprofile=cover.out`：测试完成后将覆盖率相关的记录信息输出到一个文件；`go tool cover -html=cover.out`：打开本地的浏览器窗口生成一个HTML报告
- `go test -bench=xxx`：执行基准测试；`go test -bench=xxx -benchmem`：获得内存分配的统计数据

- 文件名`*_test.go`，`go build`不会编译这些文件

### 测试函数

- `func Testxxx(t *testing.T)`：常用Error、Fail、Fatal、Log及其格式化输出函数

  ```go
  func (c *T) Error(args ...interface{})
  func (c *T) Errorf(format string, args ...interface{})
  func (c *T) Fail()
  func (c *T) FailNow()
  func (c *T) Failed() bool
  func (c *T) Fatal(args ...interface{})
  func (c *T) Fatalf(format string, args ...interface{})
  func (c *T) Log(args ...interface{})
  func (c *T) Logf(format string, args ...interface{})
  func (c *T) Name() string
  func (t *T) Parallel()
  func (t *T) Run(name string, f func(t *T)) bool
  func (c *T) Skip(args ...interface{})
  func (c *T) SkipNow()
  func (c *T) Skipf(format strin)
  ```

- `reflect.DeepEqual()`比较字符串和切片是否相等

- 测试组，子测试`t.Run(case_name, func(t *testing.T){...})`：

```go
//搞一个testCase结构体，成员是被测函数的参数和预期结果，预期一般是string切片
//搞一个测试组，类型是map，key是子测试的名字，value是testCase
func TestSplit(t *testing.T) {
	type testCase struct {
		s    string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case_1": {s: "babcbef", sep: "b", want: []string{"", "a", "c", "ef"}},
		"case_2": {s: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case_3": {s: "abcef", sep: "bc", want: []string{"a", "ef"}},
		"case_4": {s: "沙河有沙又有河", sep: "有", want: []string{"沙河", "沙又", "河"}},
	}
	for case_name, tc := range testGroup {
		t.Run(case_name, func(t *testing.T) {
			got := Split(tc.s, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%v, but got:%v\n", tc.want, got)
			}
		})
	}
}
```

### 基准函数 

> Profiling 一般和性能测试一起使用，只有应用在负载高的情况下 Profiling 才有意义

- `Benchmarkxxx(b *testing.B)`

  ```go
  func BenchmarkSplit(b *testing.B) {
  	for i := 0; i < b.N; i++ { //b.N的值是系统根据实际情况去调整的
  		Split("a:b:c:d", ":")
  	}
  }
  ```

  - b也有t的那些Log、Fail、Fatal、Error方法

  - 命令`go test -bench=xxx`、`go test -bench=xxx -benchmem`

  - 用4个进程跑的，测试了6762436次，每次操作154纳秒，64B内存，申请了1次内存

    > goos: windows
    > goarch: amd64
    > pkg: go.mod/src/studygo/excercise/unitTest/split_string
    > BenchmarkSplit-4         6762436         154 ns/op         64 B/op         1 allocs/op
    > PASS
    > ok      go.mod/src/studygo/excercise/unitTest/split_string      1.515s

- 性能比较函数`func BenchmarkFib(b *testing.B, n int)`

  ```go
  // fib.go的Fib函数是一个计算第n个斐波那契数的函数
  func Fib(n int) int {
  	if n < 2 {
  		return n
  	}
  	return Fib(n-1) + Fib(n-2)
  }
  func BenchmarkFib(b *testing.B, n int) {
  	for i := 0; i < b.N; i++ {
  		Fib(n)
  	}
  }
  func BenchmarkFib20(b *testing.B) {
  	BenchmarkFib(b, 20)
  }
  ```

  - 命令`go test -bench=Fib2`，会把`Fib2`和`Fib20`都跑一下
  - 跑所有`go test -bench=.`
    - `go test -bench . -memprofile=mem.pprof`
    - `go test -bench . -cpuprofile=cpu.pprof`
  - 默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行 
  - 强制时间`go test -bench=Fib20 -benchtime=20s`

- 时间重置`b.ResetTimer()`

  >  放到Benchmarkxxx函数中，在这之前的处理不会算到执行时间内，也不会放到输出报告中（准备阶段的耗时可以用这个方法）

- 并行测试` go test -bench=. -cpu 1 `或代码中添加如下

  ```go
  func BenchmarkSplitParallel(b *testing.B) {
  	b.SetParallelism(1) // 设置使用的CPU数
  	b.RunParallel(func(pb *testing.PB) {
  		for pb.Next() {
  			Split("沙河有沙又有河", "沙")
  		}
  	})
  }
  ```

### 示例函数

`func Examplexxx()`

## flag包

- `os.Args`：执行程序命令和后面所跟参数，是一个切片

  > 如：flag.exe a b c
  >
  > 则os.Args是[flag.exe a b c]

- `flag.Type()`

  - 接收三个参数，标志位的名字、默认值、帮助信息，返回一个该类型的指针

  - 有 `bool`、`int`、`int64`、`uint`、`uint64`、`float` `float64`、`string`、`duration`

  - 使用命令时，可以后面跟这些参数来设置标志位的值

    ```go
    func main() {
        nameVal := flag.String("name", "张三", "请输入名字")
        ageVal := flag.Int("age", "18", "请输入年龄")
        flag.Parse()
    }
    ```

    ```cmd
    # 使用命令参数可以多种形式(布尔类型的参数必须使用等号的方式指定):
    flag.exe -name=李四 --age 1000 # 变量name被设置为"李四"，变量age被设置为1000
    ```

- `flag.TypeVar(&nameVal,"name", "张三", "请输入名字")`：和上面方法类似，也要调用`flag.Parse()`

- `flag.Args()`：返回命令行参数后的其他参数，以[]string类型；`flag.NArg()`：返回命令行参数后的其他参数个数；`flag.NFlag()`：返回使用的命令行参数个数

## pprof包

> 一般是自己写两个bool类型的flag（cpu和内存），==函数中添加对flag的if判断，if中要先os.Create(cpu/mem.pprof)==，如果执行exe时用了这些flag则开启性能分析
>
> runtime/pprof：采集工具型应用运行数据进行分析
>
> net/http/pprof：采集服务型应用运行时数据进行分析
>
> pprof开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，获取各个函数占用的CPU以及内存资源；最后通过对这些采样数据进行分析，形成一个性能分析报告。

`pprof.StartCPUProfile(w *io.Writer)`：开启CPU性能分析

`pprof.StopCPUProfile()`：停止CPU性能分析，一般defer它；应用执行结束后，就会生成一个文件，保存了我们的 CPU profiling 数据，再使用命令行分析

`pprof.WriteHeapProfile(w *io.Writer)`：开启内存性能分析（作用是记录程序的堆栈信息），无需stop

`go tool pprof cpu.pprof`：使用工具进行CPU性能分析，会进入交互界面

- `top3`：查看程序中占用CPU前3位的函数；`list 函数名`：查看具体的函数分析

>  自己写的函数占用CPU太多时，可以让线程休眠一会（如半秒），把CPU让出给其他函数使用

# day06

## Go Module

`go env`查看换变量

`set go111module=on`开启`go module`

`set goproxy=https://goproxy.cn`

`go mod init xxx`初始化项目的`go.mod`

`go get`把项目中的依赖添加到`go.mod`中

`go mod tidy`检查依赖并更新`go.mod`

`go mod download`下载依赖

## Context

> 暴露内部channel时，用函数封装起来，封成只读/只写

### 概述

- 通知子`goroutine`退出
- 处理单个请求的多个`goroutine`之间与请求域的数据、取消信号、截止时间等相关操作
- `WithCancel`、`WithDeadLine`、`WithTimeOut`、`WithValue`，deadline是绝对时间，timeout是相对时间
  - `WithDeadLine`接收`contex.Contex`和`time.Time`变量，返回类型和`WithCanel`一样，返回的context过期后，会**自动触发`ctx.Done()`**，但仍要**手动`defer canel()`**
  - `WithTImeOut`接受的时间是`time.Duration`
  - `WithValue`把context和key、value绑定起来
- `Background()`、`TODO()`生成根节点context，background用于main函数，todo是预留的，当前还不知道啥场景使用

### 使用

- `ctx, cancel := context.WithCancel(context.Background())`

  - `context.WithCancel()`接收一个`contex.Contex`类型变量(`parent context`)，可以用`BackGround()`函数生成
  - 返回`contex.Contex`和`context.CancalFun`类型的变量，`CancalFun`是一个无参无返回值的函数类型，可以直接调用

- 在子`goroutine`的`select`中，添加`case <-ctx.Done():break`

- 需要通知子`goroutine`退出时，直接`cancel()`即可，且子`goroutine`中的子`goroutine`也能退出，且能一级一级传下去

  ```go
  func f(ctx context.Context) {
  	go f2(ctx)
  	defer wg.Done()
  Loop:
  	for {
  		select {
  		case <-ctx.Done():
  			break Loop
  		default:
  		}
  		fmt.Println("my context demo")
  		time.Sleep(time.Millisecond * 500)
  	}
  }
  
  func f2(ctx context.Context) {
  	defer wg.Done()
  Loop:
  	for {
  		select {
  		case <-ctx.Done():
  			break Loop
  		default:
  		}
  		fmt.Println("my context demo")
  		time.Sleep(time.Millisecond * 500)
  	}
  }
  func main() {
  	ctx, cancel := context.WithCancel(context.Background())
  	wg.Add(1)
  	go f(ctx)
  	time.Sleep(time.Second * 5)
  	cancel()
  	wg.Wait()
  }
  ```

# tags

- vscode添加快捷函数：`ctrl+shift+p` -> `snippets`->`go.json`

  ```json
  {
  	"println":{
  		"prefix": "pln",
  		"body":"fmt.Println($0)",
  		"description": "println"
  	},
  	"printf":{
  		"prefix": "plf",
  		"body": "fmt.Printf(\"$0\")",
  		"description": "printf"
  	}
  }
  ```

- 改字符集：`设置`->`encoding`->`utf-8`

- 查看某个端口是否可访问，telnet IP 端口 或者 telnet 域名 端口，回车

- 设置中搜索：`gopls`，`settings.json`设置里面添加

  ```json
  {
      "gopls": {
          "experimentalWorkspaceModule": true
      }
  }
  ```
  
- goland快捷键：`ctrl+shift+F10`（运行光标所在的文件），`shift+F10`（运行上次运行的文件），`ctrl+shift+方向键`（移动选中的代码块）,`ctrl+shift+"-"`折叠所有代码块，`ctrl+"+"`打开折叠的代码块

![image-20210608105610382](golang笔记_基础.assets/image-20210608105610382.png)


# 面试题

```go
//阅读下面的代码，写出最后的打印结果。
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
}
```

```go
//上面代码的输出结果是？
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
//结果：
//A 1 2 3
//B 10 2 12
//BB 10 12 22
//AA 1 3 4
```

