# Cron定时任务

## Cron表达式

Cron表达式是一个字符串，字符串以5或6个空格隔开，分为6或7个域，每一个域代表一个含义，Cron有如下两种语法格式：

- *Seconds Minutes Hours DayofMonth Month DayofWeek Year*
- *Seconds Minutes Hours DayofMonth Month DayofWeek*

| 字段                     | 允许值                                 | 允许的特殊字符             |
| ------------------------ | -------------------------------------- | -------------------------- |
| 秒（Seconds）            | 0~59的整数                             | , - * /   四个字符         |
| 分（*Minutes*）          | 0~59的整数                             | , - * /   四个字符         |
| 小时（*Hours*）          | 0~23的整数                             | , - * /   四个字符         |
| 日期（*DayofMonth*）     | 1~31的整数（但是你需要考虑每月的天数） | ,- * ? / L W C   八个字符  |
| 月份（*Month*）          | 1~12的整数或者 JAN-DEC                 | , - * /   四个字符         |
| 星期（*DayofWeek*）      | 1~7的整数或者 SUN-SAT （1=SUN）        | , - * ? / L C #   八个字符 |
| 年(可选，留空)（*Year*） | 1970~2099                              | , - * /   四个字符         |

## 注意项

```c
　　每一个域都使用数字，但还可以出现如下特殊字符，它们的含义是：

　　（1）*：表示匹配该域的任意值。假如在Minutes域使用*, 即表示每分钟都会触发事件。

　　（2）?：只能用在DayofMonth和DayofWeek两个域。它也匹配域的任意值，但实际不会。因为DayofMonth和DayofWeek会相互影响。例如想在每月的20日触发调度，不管20日到底是星期几，则只能使用如下写法： 13 13 15 20 * ?, 其中最后一位只能用？，而不能使用*，如果使用*表示不管星期几都会触发，实际上并不是这样。

　　（3）-：表示范围。例如在Minutes域使用5-20，表示从5分到20分钟每分钟触发一次 

　　（4）/：表示起始时间开始触发，然后每隔固定时间触发一次。例如在Minutes域使用5/20,则意味着5分钟触发一次，而25，45等分别触发一次. 

　　（5）,：表示列出枚举值。例如：在Minutes域使用5,20，则意味着在5和20分每分钟触发一次。 

　　（6）L：表示最后，只能出现在DayofWeek和DayofMonth域。如果在DayofWeek域使用5L,意味着在最后的一个星期四触发。 

　　（7）W:表示有效工作日(周一到周五),只能出现在DayofMonth域，系统将在离指定日期的最近的有效工作日触发事件。例如：在 DayofMonth使用5W，如果5日是星期六，则将在最近的工作日：星期五，即4日触发。如果5日是星期天，则在6日(周一)触发；如果5日在星期一到星期五中的一天，则就在5日触发。另外一点，W的最近寻找不会跨过月份 。

　　（8）LW:这两个字符可以连用，表示在某个月最后一个工作日，即最后一个星期五。 

　　（9）#:用于确定每个月第几个星期几，只能出现在DayofMonth域。例如在4#2，表示某月的第二个星期三。

　　三、常用表达式例子

　　（1）0 0 2 1 * ? *   表示在每月的1日的凌晨2点调整任务

　　（2）0 15 10 ? * MON-FRI   表示周一到周五每天上午10:15执行作业

　　（3）0 15 10 ? 6L 2002-2006   表示2002-2006年的每个月的最后一个星期五上午10:15执行作

　　（4）0 0 10,14,16 * * ?   每天上午10点，下午2点，4点 

　　（5）0 0/30 9-17 * * ?   朝九晚五工作时间内每半小时 

　　（6）0 0 12 ? * WED    表示每个星期三中午12点 

　　（7）0 0 12 * * ?   每天中午12点触发 

　　（8）0 15 10 ? * *    每天上午10:15触发 

　　（9）0 15 10 * * ?     每天上午10:15触发 

　　（10）0 15 10 * * ? *    每天上午10:15触发 

　　（11）0 15 10 * * ? 2005    2005年的每天上午10:15触发 

　　（12）0 * 14 * * ?     在每天下午2点到下午2:59期间的每1分钟触发 

　　（13）0 0/5 14 * * ?    在每天下午2点到下午2:55期间的每5分钟触发 

　　（14）0 0/5 14,18 * * ?     在每天下午2点到2:55期间和下午6点到6:55期间的每5分钟触发 

　　（15）0 0-5 14 * * ?    在每天下午2点到下午2:05期间的每1分钟触发 

　　（16）0 10,44 14 ? 3 WED    每年三月的星期三的下午2:10和2:44触发 

　　（17）0 15 10 ? * MON-FRI    周一至周五的上午10:15触发 

　　（18）0 15 10 15 * ?    每月15日上午10:15触发 

　　（19）0 15 10 L * ?    每月最后一日的上午10:15触发 

　　（20）0 15 10 ? * 6L    每月的最后一个星期五上午10:15触发 

　　（21）0 15 10 ? * 6L 2002-2005   2002年至2005年的每月的最后一个星期五上午10:15触发 

　　（22）0 15 10 ? * 6#3   每月的第三个星期五上午10:15触发
```

## golang用法

- `github.com/robfig/cron`

  ```go
  // 单个定时任务
  func main() {
      i := 0
      c := cron.New()
      spec := "*/5 * * * * ?"
      c.AddFunc(spec, func() {
          i++
          log.Println("cron running:", i)
      })
      c.Start()
      
      select{}
  }
  
  // 多个定时任务
  type TestJob struct {
  }
  
  func (this TestJob)Run() {
      fmt.Println("testJob1...")
  }
  
  type Test2Job struct {
  }
  
  func (this Test2Job)Run() {
      fmt.Println("testJob2...")
  }
  
  func main() {
      i := 0
      c := cron.New()
      
      spec := "*/5 * * * * ?"
      c.AddFunc(spec, func() {
          i++
          log.Println("cron running:", i)
      })
  
      c.AddJob(spec, TestJob{})
      c.AddJob(spec, Test2Job{})
  
      c.Start()
  
      //defer c.Stop()
      
      select{}
  }
  ```

# json

## 转义字符

> json.Marshal 默认 escapeHtml 为true，会把 <、>、&转义成\u003c、\u003e、\u0026

```go
// 解决方法
type Test struct {
    Content   string
}
func main() {
    data := Test{}
    data.Content = "http://www.baidu.com?id=123&test=1"
    buffer := bytes.NewBuffer([]byte{})
    jsonEncoder := json.NewEncoder(buffer)
    jsonEncoder.SetEscapeHTML(false)
    jsonEncoder.Encode(&data)
    fmt.Println(buffer.String())
}
```

## 待补充...

# CGO

## windows

> import "C"和#include之间不能有空格；引用stdlib.h可以使用cstr := C.CString("xxx")、C.free(unsafe.Pointer(cstr))等函数

```go
/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

func main() {
    dll, _ := syscall.LoadDLL(mgr.libPath)

    proc, _ := dll.FindProc("C函数名字")
    
    result, _, _ := proc.Call(args...)	// args是C函数的参数
}
```

## linux

- `CFLAGS`中的`-I`（大写的i） 参数表示`.h`头文件所在的路径；

- `LDFLAGS`中的`-L`(大写) 表示.so文件所在的路径 `-l`(小写的L) 表示指定该路径下的库名称，比如要使用`demo.so`，则只需用`-ldemo` 表示（省略了`libdemo.so`中的`lib`和`.so`字符，`.so`文件也可换成`.a`文件）

- 编译得到`.so`：

  ```bash
  gcc -c -fPIC -o demo.o demo.c
  gcc -shared -o libdemo.so demo.o
  ```

- 编译得到`.a`：

  ```bash
  gcc -c demo.c
  ar -rv libdemoo.a demo.o
  ```

- windows平台也可使用此方法，`.a`文件换成`.lib`即可

```go
#cgo CFLAGS: -I../../lib
#cgo LDFLAGS: -L../../lib -lhi
#include "demo.h"
#include <stdlib.h>
*/
import "C"

func main() {
    C.f() 	// 直接通过C.来调用头文件中的函数
}
```

## 类型转换

- 基本类型转换

  ```c++
  char -->  C.char -->  byte
  signed char -->  C.schar -->  int8
  unsigned char -->  C.uchar -->  uint8
  short int -->  C.short -->  int16
  short unsigned int -->  C.ushort -->  uint16
  int -->  C.int -->  int
  unsigned int -->  C.uint -->  uint32
  long int -->  C.long -->  int32 or int64
  long unsigned int -->  C.ulong -->  uint32 or uint64
  long long int -->  C.longlong -->  int64
  long long unsigned int -->  C.ulonglong -->  uint64
  float -->  C.float -->  float32
  double -->  C.double -->  float64
  wchar_t -->  C.wchar_t  --> 
  void * -> unsafe.Pointer
  
  如果需要传递指针类型，需要先进行unsafe.Pointer的转化，再进行强转
  ```

- 将[]byte作为char*传递

  > 由于slice自身是一个结构体，我们需要传递的只是其中的data部分，所以需要取数据部分的首地址。这里还有个地方需要注意，由于我们传递的是数据段的首地址，所以这段数据并没有结束符，所以需要我们手动添加结束符，或者同时传递slice的真实长度，避免内存越界。

  ```go
  #将[]byte作为char*传递
  str := []byte("hello world")
  tmp := (*C.char)(unsafe.Pointer(&str[0]))
  ```

# pprof

## 通用方法
* 在代码库中引入net/http/pprof包，然后在main入口中起个端口监听
  ```go
    // pprof 的init函数会将pprof里的一些handler注册到http.DefaultServeMux上
    import(
        _ "net/http/pprof" // 必须，引入 pprof 模块
    )
  
    //如果程序已有Http服务且使用了http.DefaultServeMux，则下面的代码就不需要了
    //否则，自己启用一个使用了http.DefaultServeMux的Http服务，就需要下面的代码
    go func() {
        http.ListenAndServe("0.0.0.0:8080", nil)
    }()
  ```
## 使用github.com/gorilla/mux包情况下的pprof使用
* 方法一，即上面的通用方法，另开启一个http服务
    ```go
    package main
    
    import (
        "fmt"
        "log"
        "net/http"
        _ "net/http/pprof"
    
        "github.com/gorilla/mux"
    )
    
    func SayHello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello!")
    }
    
    func main() {
        r := mux.NewRouter()
        r.HandleFunc("/hello", SayHello)
    
        //Added another native http server on a different port and it just works out of the box
        go func() {
            log.Fatal(http.ListenAndServe(":6060", http.DefaultServeMux))
        }()
    
        http.ListenAndServe(":18080", r)
    }
    ```

* 方法二（无需新增监听端口）
    ```go
    package main
    
    import (
        "net/http"
        _ "net/http/pprof"
    
        "github.com/gorilla/mux"
    )
    
    func main() {
        // ... ...
    
        //Let net/http/pprof register itself to http.DefaultServeMux, and then pass all requests starting with /debug/pprof/ along
        router := mux.NewRouter()
        router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
        if err := http.ListenAndServe(":服务本身的端口", router); err != nil {
            panic(err)
        }
    }
    ```

# 条件编译

> // +build linux
> // +build windows
>  限制此源文件只能在 linux或者windows平台下编译

参考资料

- [golang之条件编译](https://blog.csdn.net/phantom_111/article/details/79451635)

# Go Mod

- 下载的依赖包都在`$GOPATH$/pkg/mod`下面，下载之后如果把依赖包的文件夹删除了，再次编译的时候，就不会下载，而是`go: extracting github.com/xxx/xxx v1.0.0`
- `replace`：`go mod edit -replace [old git package]@[version]=[new git package]@[version]`
- 待补充...

# Websocket
