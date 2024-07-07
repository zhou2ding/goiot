# ES6～ES13

## 新语法

### 语法

- 每行结束不需要分号了
- 变量用`let`声明
- 常量用声明`const`，重新赋值给常量，运行时会报错

### 解构

#### 数组解构

```js
let [x,y] = [1,2]	// x=1,y=2，数组解构赋值
let [,,c] = [1,2,3]	// c=3，c左边的逗号相当于占位符
let [a,...b] = [1,2,3,4]	//a=1，b=[2,3,4]，扩展运算符结合解构赋值
let [a,b] = [100]	//a=100，b是undefined
let [a,b=200] = [100]	//a=100，b=200,200是默认值

// 利用解构交换值
let x = 4
let y = 5
;[x,y] = [y,x]	//必须加分号，否则报错，分号也可以加在上一行的末尾。为了避免 JavaScript 引擎误解代码，将上一行的输出与数组解构赋值连接起来
```

#### 对象解构

```js
// 定义一个对象
const person = {
    name: 'Alice',
    age: 25,
    city: 'New York'
}

// 使用对象解构语法提取属性到独立变量
const { name, age, city } = person

console.log(name)  // 输出: Alice
console.log(age)   // 输出: 25
console.log(city)  // 输出: New York
```

#### 对象解构2

```js
// 定义一个对象
const person = {
    name: 'Alice',
    age: 25,
    city: 'New York'
}

// 重命名解构
const { name: personName, age: personAge, city: personCity } = person

console.log(personName)  // 输出: Alice
console.log(personAge)   // 输出: 25
console.log(personCity)  // 输出: New York

// 对象的解构赋值，变量名称必须在对象的属性中存在
const person = {
    name: 'Alice',
    age: 25
}
const { a } = person
console.log(a) // 输出: undefined
```

## 新数据类型

### object

```js
let obj = {
  key1: val1,
  key2: val2
}
```

- obj.key3 = val3，添加元素

- delete obj.key2，删除元素

- let has = "key4" in  obj，判断是否有该属性，必须带引号

- obj[key]通过key获取val，for里面使用时key不用带引号，其他时候使用key需要带引号

- 遍历对象：for(let key in obj)，用for in，遍历对象的可枚举属性，for of是遍历可迭代对象的，注意区分

- Object.keys(obj)，获取键组成的数组

- Object.entries(obj)，获取键值对组成的数组（二维数组），元素是键和对应的值组成的数组

  ```js
  // forEach遍历对象，
  Object.entries(obj).forEach([key,val]) => {
    // do sth
  }
  ```

- obj = {}，清空对象

### array

`let arr = []`

- arr.push(元素)，尾部添加元素，返回修改后的长度

- arr.unshift(元素)，头部添加元素，返回修改后的长度

- arr.shift()，删除头部元素，返回删除的元素

- arr.pop()，删除尾部元素，返回删除的元素

- arr.splice(idx,n)，删除指定元素，从idx开始，删除n个元素，返回删除掉的元素所组成的数组

- arr.reverse()，颠倒数组

- arr.sort()

  - 默认按首个首字母升序，数字也会转成字符串按首字母排序

  - 传参比较函数，则按比较函数的逻辑排序

    ```js
    let arr = [1,13,2,15,4]
    arr.sort((a,b)=>a-b)	// a-b<0，则a在b前面，否则在b后面
    ```

- arr.filter()，筛选数组

  ```js
  let arr = [10,11,12,13,14]
  let newArr = arr.filter((val)=> {
    return val>12	//筛选大于12的元素组成的数组
  })
  ```

- for (let item of arr)，遍历数组

- arr.forEach((value,index)=> {})，遍历数组

### map

存储键值对，typeof仍然是object

```js
let m = new Map([
  ["key1",val1],
  ["key2",val2]
])
```

- 用get()、set()、delete()来操作键值对，has()判断key是否存在，size()获取大小，clear()清空map 
- 用forEach()、keys()、values()、entries()来遍历键值对
- for let val of map遍历时，value是键值对组成的数组，也就是赋值时的一组键值对
- 解构遍历for [key,val] of map，注意顺序是key,val，和forEach是反的，forEach对map是val,key，对数组是val,index 

### set

存储无序且唯一的值的集合

```js
let s = new Set([1,2,3,'2',2])	// 字符串'2'会存进去，最后一个数字2不会存进去
let s = new Set(arr)	// 接收一个数组，将其元素去重后转换成数组
```

- s.add(元素)，添加元素
- s.delete(元素)，删除元素
- s.has(元素)，判断元素是否在集合内
- s.size()，返回集合的大小
- set转成数组，let arr = Array.form(s)，可以同样用于map
- for let item of s，遍历set
- s.forEach(valuer=>{})，遍历set，set集合的索引，还是元素本身，set不能解构遍历，因为它没有key
- s.clear()清空set

### 扩展运算符

[...可迭代对象]，用于展开可迭代对象（字符串、数组、set、map）等

```js
let str = 'abc'
let arr = [...abc]	//将字符串转换成数组，元素是字符串的每个字符

let s = new Set([1,2,3])
let arr = [...s]	//将set转换成数组，效果和let arr = Array.form(s)一样。可以同样用于map
```

### class

类名首字母都是大写，构造函数、方法

```js
class Person {
  // 显式地声明字段，可以不声明，但最好声明，提高代码可读性
  name
  age
  // 构造函数
  constructor(name,age) {
    this.name=name
    this.age=age
  }
  info() {
    alert(this.name, this.age)
  }
}
let person = new Person('zs',18)
person.info()
```

- 私有属性，在声明时和this.xxx时，在属性名左边加上#，不能直接访问，即使通过方法操作或获取，也是undefined，需要通过getter和setter获取和修改

  ```js
  class Person {
    name
    #age
    constructor(name,age) {
      this.name=name
      this.#age=age
    }
    // getter和setter存取器
    get() {
      return this.#age
    }
    set(val) {
      this.#age = val
    }
    info() {
      alert(this.name, this.#age)
    }
  }
  let person = new Person('zs',18)
  person.info()	// 没有getter的话，打印出来的是undefined
  person.age = 16	// 没有setter的话，不会修改值
  ```

- 模板字符串

  ```js
  `普通字符串${字符串变量}`
  ```

- 继承extends，继承父类的属性和方法，且能声明自己的属性和方法，构造函数需要传父类属性的形参（顺序随便），且通过super调用父类的构造函数

  ```js
  class Zhangsan extends Person {
    gender
    constructor(name,age,gender) {	// 也可以先传gender再传父类的属性
      super(name,age)	// 调用父类的构造函数
      this.gender = gender
    }
  }
  ```

  

### function

- 形参可以指定默认值

  ```js
  function getPage(page=1) {
    return page
  }
  alert(getPage())	// 1
  alert(getPage(10))	// 10
  ```

- 匿名函数，通常用于回调函数

  ```js
  let sub = function(x,y) {
    return x-y
  }
  ```

- 箭头函数，也是匿名函数

  ```js
  let plus = (a,b) => {
    return a+b
  }
  // 函数体内只有一行表达式时，可以省略花括号和return
  let plus = (a,b) => a+b
  // 只有一个函数时，还可以省略小括号，没有参数时，不能省略小括号 
  let plus = x => x+10
  ```

## Promise

- Promise表示承诺在未来某个时刻可能会完成并返回结果
- 对于某些需要时间来处理结果的操作，如用户登录、读取文件等，可以使用Promise对象来执行异步操作
- Promise对象有三种状态：pending（待处理）、fullfilled（已履行）、rejected（被驳回）
- 当创建一个Promise对象时，初始状态未pending，表示异步状态还未完成
- 当异步执行成功时，手动调用resolve，把Promise对象的状态改成fullfilled，可通过then得到异步操作的result，then中也可以返回Promise对象，也可通过throw抛出错误
- 当异步执行失败时，手动调用reject，把Promise对象的状态改成rejected，可通过catch捕获错误（reject返回的）
- new Promise接收一个立即执行的函数作为参数，该函数接收resolve和reject参数

```js
// 链式执行，const promise 一般可省略
const promise = new Promise((resolve, reject) => {
    // 异步操作，例如一个定时器
    setTimeout(() => {
        const success = true; // 模拟操作结果
        if (success) {
            resolve('操作成功'); // 操作成功，Promise 状态变为 fulfilled
        } else {
            reject('操作失败'); // 操作失败，Promise 状态变为 rejected
        }
    }, 1000);
})
.then(result => {
    console.log(result); // 输出 "操作成功"
})
.catch(error => {
    console.error(error); // 输出 "操作失败"
}).finally(() => {
    console.log('所有操作完成'); // 无论成功或失败都会执行
});

// 在 then 中返回一个新的 Promise，从而实现更加复杂的异步操作链
const promise = new Promise((resolve, reject) => {
    // 异步操作，例如一个定时器
    setTimeout(() => {
        const success = true; // 模拟操作结果
        if (success) {
            resolve('操作成功'); // 操作成功，Promise 状态变为 fulfilled
        } else {
            reject('操作失败'); // 操作失败，Promise 状态变为 rejected
        }
    }, 1000);
})
.then(result => {
    console.log(result); // 输出 "操作成功"
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            const success = true;
            if (success) {
                resolve('操作2成功');
            } else {
                reject('操作2失败');
            }
        }, 1000);
    });
})
.then(result => {
    console.log(result); // 输出 "操作2成功"
})
.catch(error => {
    console.error(error); // 如果任意一个Promise失败，输出错误信息
})
.finally(() => {
    console.log('所有操作完成'); // 无论成功或失败都会执行
});
```

## fetch

- new URLSearchParams
  - 更适合在需要单独处理查询参数字符串的场景中使用。
  - 可以通过传递对象或字符串初始化。
  - 可以单独创建查询参数字符串并赋值给 URL 对象的 search 属性。

- url.searchParams.append
  - 更适合在需要动态添加单个查询参数的场景中使用。
  - 直接对 URL 对象进行操作，简洁直观。

- get，url.searchParams.append

  ```js
  const baseUrl = 'https://api.example.com/data';
  const queryParams = {
      key1: 'value1',
      key2: 'value2',
  };
  
  // 使用 URL 对象和 searchParams.append 添加查询参数
  const url = new URL(baseUrl);
  Object.keys(queryParams).forEach(key => url.searchParams.append(key, queryParams[key]));
  
  // 执行 GET 请求
  fetch(url)
      .then(response => {
          if (!response.ok) {
              throw new Error('Network response was not ok ' + response.statusText);
          }
          return response.json();
      })
      .then(data => {
          console.log('GET response data:', data);
      })
      .catch(error => {
          console.error('GET request failed:', error);
      })
      .finally(() => {
          console.log('GET request completed');
      });
  ```

- get，new URLSearchParams

  ```js
  const baseUrl = 'https://api.example.com/data';
  const queryParams = {
      key1: 'value1',
      key2: 'value2',
  };
  
  // 使用 URLSearchParams 创建查询字符串
  const url = new URL(baseUrl);
  const params = new URLSearchParams(queryParams);
  url.search = params.toString();
  
  // 执行 GET 请求
  fetch(url)
      .then(response => {
          if (!response.ok) {
              throw new Error('Network response was not ok ' + response.statusText);
          }
          return response.json();
      })
      .then(data => {
          console.log('GET response data:', data);
      })
      .catch(error => {
          console.error('GET request failed:', error);
      })
      .finally(() => {
          console.log('GET request completed');
      });
  ```

- post

  ```js
  const postUrl = 'https://api.example.com/submit';
  const postData = {
      key1: 'value1',
      key2: 'value2',
  };
  
  // 执行 POST 请求
  fetch(postUrl, {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
      },
      body: JSON.stringify(postData),
  }).then(response => {
          if (!response.ok) {
              throw new Error('Network response was not ok ' + response.statusText);
          }
          return response.json();
      })
      .then(data => {
          console.log('POST response data:', data);
      })
      .catch(error => {
          console.error('POST request failed:', error);
      })
      .finally(() => {
          console.log('POST request completed');
      });
  ```

- post，上传文件

  ```js
  const uploadUrl = 'https://api.example.com/upload';
  const formData = new FormData();
  formData.append('file', fileInput.files[0]); // 假设有一个文件输入元素 <input type="file" id="fileInput">
  
  // 执行 POST 请求
  fetch(uploadUrl, {
      method: 'POST',
      body: formData,
  }).then(response => {
          if (!response.ok) {
              throw new Error('Network response was not ok ' + response.statusText);
          }
          return response.json();
      })
      .then(data => {
          console.log('File upload response data:', data);
      })
      .catch(error => {
          console.error('File upload failed:', error);
      })
      .finally(() => {
          console.log('File upload completed');
      });
  ```

## Axios

相比fetch的原生API，用起来更方便

- **GET 请求**：axios.get(url, config)
  - url: 请求的地址。
  - config: 配置对象，包含 params 和 headers 等。

- **POST 请求**：axios.post(url, data, config)
  - url: 请求的地址。
  - data: 请求体数据。
  - config: 配置对象，包含 params 和 headers 等。

- get

  ```js
  const axios = require('axios');
  
  const baseUrl = 'https://api.example.com/data';
  const queryParams = {
      key1: 'value1',
      key2: 'value2',
  };
  
  // 执行 GET 请求
  axios.get(baseUrl, { params: queryParams })
      .then(response => {
          console.log('GET response data:', response.data);
      })
      .catch(error => {
          console.error('GET request failed:', error);
      })
      .finally(() => {
          console.log('GET request completed');
      });
  ```

- post

  ```js
  const postUrl = 'https://api.example.com/submit';
  const postData = {
      key1: 'value1',
      key2: 'value2',
  };
  
  // 执行 POST 请求，headers可省略，axios 默认会设置 Content-Type 为 application/json
  axios.post(postUrl, postData, {
      headers: {
          'Content-Type': 'application/json',
      },
  })
      .then(response => {
          console.log('POST response data:', response.data);
      })
      .catch(error => {
          console.error('POST request failed:', error);
      })
      .finally(() => {
          console.log('POST request completed');
      });
  ```

- post，上传文件

  ```js
  const uploadUrl = 'https://api.example.com/upload';
  const formData = new FormData();
  formData.append('file', fileInput.files[0]); // 假设有一个文件输入元素 <input type="file" id="fileInput">
  
  // 执行文件上传的 POST 请求，headers可省略，axios 和浏览器会自动设置适当的 Content-Type
  axios.post(uploadUrl, formData, {
      headers: {
          'Content-Type': 'multipart/form-data',
      },
  })
      .then(response => {
          console.log('File upload response data:', response.data);
      })
      .catch(error => {
          console.error('File upload failed:', error);
      })
      .finally(() => {
          console.log('File upload completed');
      });
  ```

## 模块化开发

script标签的type属性需要是module

- 导出

  - 单独导出

    ```js
    // math.js
    export function add(a, b) {
        return a + b;
    }
    
    export function subtract(a, b) {
        return a - b;
    }
    
    export const PI = 3.14159;
    ```

  - 默认导出，用于导出一个默认值，一个模块只能有一个默认导出

    ```js
    // calculator.js
    export default class Calculator {
        add(a, b) {
            return a + b;
        }
    
        subtract(a, b) {
            return a - b;
        }
    }
    ```

  - 统一导出

    ```js
    // index.js
    function add(a, b) {
        return a + b;
    }
    
    function subtract(a, b) {
        return a - b;
    }
    
    const PI = 3.14159;
    
    class Calculator {
        add(a, b) {
            return a + b;
        }
    
        subtract(a, b) {
            return a - b;
        }
    }
    
    export { add, subtract, PI };
    export default Calculator;	//导入时，import Calculator, {add, subtract, PI} from './math'
    
    // 也可以
    export default { add, subtract, PI };	//导入时，import math, {Calculator} from './math'
    export Calculator;
    ```

- 导入

  - 直接导入

    ```js
    // main.js
    import { add, subtract, PI } from './math';
    
    console.log(add(2, 3));  // 输出 5
    console.log(subtract(5, 2));  // 输出 3
    console.log(PI);  // 输出 3.14159
    ```

  - 重命名导入

    ```js
    // main.js
    import { add as addition, subtract as subtraction } from './math';
    
    console.log(addition(2, 3));  // 输出 5
    console.log(subtraction(5, 2));  // 输出 3
    ```

  - 导入整个模块

    ```js
    // main.js
    import * as math from './math';
    
    console.log(math.add(2, 3));  // 输出 5
    console.log(math.subtract(5, 2));  // 输出 3
    console.log(math.PI);  // 输出 3.14159
    ```

  - 默认导入，不需要使用大括号

    > 可以和普通的导出同时用：默认导入不带大括号，普通导入带大括号

    ```js
    // main.js
    import Calculator from './calculator';
    
    const calc = new Calculator();
    console.log(calc.add(2, 3));  // 输出 5
    console.log(calc.subtract(5, 2));  // 输出 3
    ```

## async和await

- 当一个函数被标记为async后，该函数会返回一个Promise对象
- 只能在async标记的函数内部使用，加上await关键字后，会在执行到这一行时暂停函数的剩余部分，等待网络请求完成后，继续执行并获取到请求返回的数据
- 使用async和await可以用同步的方式编写异步的代码，避免回调地狱

```js
const getData() = async () => {
  try {
    const resp = await axios.get(url1)
  }
}
```

# Vue3

## 响应式对象

响应式对象是 Vue 3 中的一种数据结构，它能够自动追踪数据的变化，并在数据发生改变时自动更新视图，创建的响应式对象，用{{}}模板语法就可以访问到

- 创建一个vue应用

```html
<script type="module">
  import {createApp, reactive, ref} from "./vue.esm-browser.js"
	// 创建一个Vue3应用
	createApp({
    // setup选项，用于设置响应式数据和方法等
    setup() {
      return {
        key1: val1,
        key2: val2
      }
    }
  }).mount(选择器)	// mount用于绑定到对象上，和vue2的el一样
</script>
```

- ref：用于存储单个基本数据类型的数据，使用ref创建的响应式对象，要用.value来访问和修改值
  - 也可以用来创建对象或数组

```js
// 创建一个响应式的基本数据类型
const count = ref(0);

// 访问和修改值
console.log(count.value); // 0
count.value = 1;
console.log(count.value); // 1

// 创建一个响应式对象
const user = ref({
  name: 'John',
  age: 30
});

// 访问和修改对象的属性
console.log(user.value.name); // John
user.value.age = 31;
console.log(user.value.age); // 31
```

- reactive：用于存储复杂数据类型，使用reactive创建的响应式对象，可直接用属性名或索引来访问和修改值

```js
// 创建一个响应式对象
const user = reactive({
  name: 'Jane',
  age: 25,
  address: {
    city: 'New York',
    zip: '10001'
  }
});

// 访问和修改对象的属性
console.log(user.name); // Jane
user.age = 26;
console.log(user.age); // 26

// 嵌套对象也是响应式的
console.log(user.address.city); // New York
user.address.city = 'Los Angeles';
console.log(user.address.city); // Los Angeles
```

## v-on

- 和vue2用法差不多，`v-on:事件名="函数"`，v-on可用@代替

- 在setup中要返回v-on绑定的函数

- 键盘按下某个键：keyup.键名，如keyup.enter表示按了回车键弹起后，keydown也一样，如果是多个键的话，则继续链式起来，如keyup.ctrl.a，表示按了ctrl+a弹起后

### 事件修饰符

  `@事件名.修饰符="函数"`，直接在绑定事件的地方调用修饰符，vue有6种修饰符，常用前三个，可以链式调用

  - prevent，阻止默认事件，相当于`e.preventDefault()`
  - stop，阻止事件冒泡
  - once，事件只触发一次
  - capture，使用事件的捕获模式
  - self，只有event.targe是当前操作的元素时才触发
  - passive，事件的默认行为立即执行，无需等待事件回调即可执行完毕（移动端可能会用）

```html
<div id="app">
    <h3>{{ msg }}</h3>
    <h3>{{ web.url }}</h3>
    <h3>{{ web.user }}</h3>
    <h3>{{ sub(100, 20) }}</h3>

    <!-- v-on:click 表示在 button 元素上监听 click 事件 -->
    <button v-on:click="edit">修改</button> <br>

    <!-- @click 简写形式 -->
    <button @click="add(20, 30)">加法</button> <br>

    <!-- 
        enter space tab 按键修饰符
        keyup是在用户松开按键时才触发
        keydown是在用户按下按键时立即触发
    -->
    回车 <input type="text" @keyup.enter="add(40, 60)"> <br>
    空格 <input type="text" @keyup.space="add(20, 30)"> <br>
    Tab <input type="text" @keydown.tab="add(10, 20)"> <br>
    w <input type="text" @keyup.w="add(5, 10)"> <br>

    <!-- 组合快捷键 -->
    Ctrl + Enter <input type="text" @keyup.ctrl.enter="add(40, 60)"> <br>
    Ctrl + A <input type="text" @keyup.ctrl.a="add(20, 30)">
</div>

<script type="module">
    import { createApp, reactive } from './vue.esm-browser.js'
    
    createApp({
        setup() {
            const web = reactive({
                title: "邓瑞编程",
                url: "dengruicode.com",
                user: 0
            })

            const edit = () => {
                web.url = "www.dengruicode.com"
                //msg = "邓瑞编程" //错误示例 不能直接改变msg的值,因为msg是一个普通变量, 不是响应式数据
            }

            const add = (a, b) => {
                web.user += a + b
            }

            const sub = (a, b) => {
                return a - b
            }

            return {
                msg: "success", //普通变量, 非响应式数据, 在模板中普通变量不会自动更新
                web, //响应式数据
                edit, //方法
                add,
                sub,
            }
        }
    }).mount("#app")

</script>
```

## v-show

用v-show绑定响应式对象的值，嵌入元素的样式中，v-show为true时该元素显示，为false时，给该元素添加display:none样式

```html
<template>
  <div>
    <button @click="toggle">Toggle Visibility</button>
    <p v-show="isVisible">This paragraph is conditionally visible.</p>
    <p v-show="web.show">This paragraph is conditionally visible.</p>
  </div>
</template>

<script>
import { ref } from 'vue';
export default {
  setup() {
    const isVisible = ref(true);
    const web = reactive({
      show: true
    })
    const toggle = () => {
      isVisible.value = !isVisible.value;
      web.show = !web.show;
    };
    return {
      isVisible,
      toggle
    };
  }
};
</script>
```

## v-if

更复杂的控制元素显式的逻辑，v-if、v-else-if、v-else结合使用，满足某个分支的条件时（该分支为true），显式该分支的元素

```html
<template>
  <div>
    <button @click="setStatus('loading')">Loading</button>
    <button @click="setStatus('success')">Success</button>
    <button @click="setStatus('error')">Error</button>

    <div v-if="status === 'loading'">
      Loading...
    </div>
    <div v-else-if="status === 'success'">
      Data loaded successfully.
    </div>
    <div v-else>
      Error loading data.
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
export default {
  setup() {
    const status = ref('');
    const setStatus = (newStatus) => {
      status.value = newStatus;
    };
    return {
      status,
      setStatus
    };
  }
};
</script>
```

## v-bind

给属性绑定动态值，v-bind:value="响应式对象.属性"，简写：省略v-bind，只留冒号

v-bind 可以绑定以下内容：

- HTML 属性：绑定 HTML 元素的属性，例如 href、src、title 等。

  ```html
  <a v-bind:href="url">链接</a>
  ```

- 组件的 prop：绑定到子组件的 prop 上。

  ```html
  <my-component v-bind:some-prop="someValue"></my-component>
  ```

- class 和 style：绑定到元素的 class 和 style 属性上，可以是对象或数组形式。

  - 使用对象时，键是类名，值是布尔值。值为 true 时会添加该类，值为 false 时会移除该类。

  ```html
  <div v-bind:class="{ active: isActive, 'text-danger': hasError }"></div>
  <div v-bind:style="{ color: activeColor, fontSize: fontSize + 'px' }"></div>
  ```

  - 使用数组时，根据数组中的元素来动态地添加类名。每个元素可以是字符串、对象或者其它表达式。

  ```html
  <div v-bind:class="[activeClass, errorClass]"></div>
  ```

  - 对象和数组也可以结合起来使用

- 布尔属性：对于布尔类型的属性，如 disabled、required，当绑定的值为 true 时，属性会存在，否则属性会被移除。

  ```html
  <button v-bind:disabled="isDisabled">按钮</button>
  ```

- DOM 属性：绑定到 DOM 属性，例如 value、checked、disabled 等。

  ```html
  <input v-bind:value="inputValue" /> 
  ```

## v-for

用于基于一个数组来渲染一个列表，Vue 强烈推荐给v-for提供 key 属性。key 主要用在 Vue 的虚拟 DOM 算法中，为每个节点提供唯一标识，以便更高效地更新 DOM。

- 遍历数组和对象

```js
//待遍历的数组
data() {
  return {
    items: [
      { id: 1, text: 'Apple' },
      { id: 2, text: 'Banana' },
      { id: 3, text: 'Cherry' }
    ]
  }
}

//待遍历的对象
data() {
  return {
    object: {
      name: 'John',
      age: 30,
      city: 'New York'
    }
  }
}
```

```html
<!--遍历数组-->
<li v-for="item in items" :key="item.id">
  {{ item.text }}
</li>

<!--带index的遍历-->
<li v-for="(item, index) in items" :key="item.id">
  {{ index }} - {{ item.text }}
</li>

<!--遍历对象-->
<div v-for="(value, key) in object" :key="key">
  {{ key }}: {{ value }}
</div>
```

- 当需要循环渲染多个元素时，可以用 <template> 包裹

  > <template> 标签在 Vue.js 中并不会渲染为实际的 DOM 元素。它只是一个包含模板代码的容器，可以在其他地方被引用和使用。<template> 标签通常用于在 Vue 组件中包含结构化的 HTML 片段，或者用于包裹多个元素以进行条件渲染或循环渲染。

```html
<template v-for="item in items">
  <h1>{{ item.title }}</h1>
  <p>{{ item.description }}</p>
</template>
```

- v-for还能和v-if结合使用进行条件遍历

```html
<template v-for="item in items">
  <div v-if="item.visible" :key="item.id">{{ item.text }}</div>
</template>
```

## v-model

- 双向数据绑定，默认绑定input标签的value值
- 对于单元框，绑定单选框的选中状态和vue实例中的数据，勾选了哪个框，就把该input的value同步到绑定的数据
- 对于单个复选框，绑定选中状态和vue实例中的布尔值，也可以通过 true-value 和 false-value 属性，可以控制复选框选中和未选中时的数据值。
- 对于多个复选框，绑定选中状态vue实例中的数组，选中的值会添加到数组中

```html
<!--文本框-->
<template>
  <input type="text" v-model="message" />
   <p>message: {{ message }}</p>	<!--改输入框中的内容，这里显示的message也会变-->
</template>

<script>
export default {
  data() {
    return {
      message: ''
    };
  }
};
</script>

<!--单选框-->
<template>
  <div>
    <input type="radio" id="option1" value="Option 1" v-model="picked" />
    <label for="option1">Option 1</label>
    
    <input type="radio" id="option2" value="Option 2" v-model="picked" />
    <label for="option2">Option 2</label>

    <p>Picked: {{ picked }}</p>	<!--picked显示的是被选中的单元框的value-->
  </div>
</template>

<script>
export default {
  data() {
    return {
      picked: ''
    };
  }
};
</script>

<!--单个复选框-->
<template>
  <div>
    <input type="checkbox" id="checkbox" v-model="checked" />
    <label for="checkbox">Check me</label>
    <p>Checked: {{ checked }}</p>	<!--checked显示的是true或false，复选框被选中时显示true-->
    
    <input type="checkbox" id="checkbox" v-model="numStatus" true-value="yes" false-value="no" />
    <label for="checkbox">Check me</label>
		<p>Num Status: {{ numStatus }}</p>	<!--numStatus显示的是yes或no，复选框被选中时显示yes-->
  </div>
</template>

<script>
export default {
  data() {
    return {
      checked: false
    };
  }
};
</script>

<!--复选框-->
<template>
  <div>
    <input type="checkbox" id="option1" value="Option 1" v-model="checkedOptions" />
    <label for="option1">Option 1</label>

    <input type="checkbox" id="option2" value="Option 2" v-model="checkedOptions" />
    <label for="option2">Option 2</label>

    <p>Checked Options: {{ checkedOptions }}</p>	<!--选中一个复选框，checkedOptions数组就会多一个被选中框的value-->
  </div>
</template>

<script>
export default {
  data() {
    return {
      checkedOptions: []
    };
  }
};
</script>
```

### 修饰符

- number，自动将用户输入值转换为数字（以数字开头时，才会强转，否则显示原字符串）
- lazy，输入框失去焦点或按下回车键时才更新数据
- trim，去除输入内容两边的空格

```html
<template>
  <input v-model.lazy="message" />
  <input v-model.number="age" />
  <input v-model.trim="name" />
</template>

<script>
export default {
  data() {
    return {
      message: '',
      age: 0,
      name: ''
    };
  }
};
</script>
```

## v-text和v-html

效果和插值表达式一样，只不过v-html能解析html标签和样式（绑定的值是html标签的话，会渲染），v-text只能解析成文本

## 计算属性computed

- 需要导入computed，computed接收一个函数作为参数
- 计算属性会缓存数据，当它依赖的响应式数据发生变化时，就会重新计算

```html
<p>
  {{add()}}	<!--add是方法，需要()调用-->
  {{sub}}	<!--sub是计算属性，不用调用就能直接计算-->
</p>
<script>
import {createApp, reactive} from "./vue.js";
createApp({
  setup() {
    const data = reactive({
      x :20,
      y: 30
    })
    <!--无缓存方法-->
    let add = () => {
      return data.x + data.y
    }
    <!--计算属性-->
		const sub = computed(() => {
      return data.x - data.y
    })
    
    return {
      data,
      add,
      sub
    }
  }.mouont(筛选器)
})
</script>
```

## 侦听器watch

- 需要导入watch