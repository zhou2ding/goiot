# vue教程

## 基础

- 一套用于构建用户界面的渐进式JavaScript框架

- 特点

  - 采用组件化模式，提高代码复用率，一个模块/组件所需的html、css、js都在xxx.vue里

  - 声明式编码，无需直接操作DOM

  - 采用虚拟DOM+优秀的Diff算法，尽量复用DOM节点

    ![image-20240522222319505](vue教程.assets/image-20240522222319505.png)

## 基本使用

`new Vue({el:'css选择器',data:{}})`，vue实例和容器只能一一对应

### 模板语法

- 插值语法：`{{}}`，用于标签体的内容，匹配vue实例属性(如果是methods里定义的函数的话，需要加()才能执行函数)或js表达式，如`{{key}}`、`{{Date.now()}}`、`{{1+1}}`

  ```js
  <div>
      <h1>
          helllo, {{key}}, {{key1}}
      </h1>
  </div>
  new Vue({
  	el: 'css选择器',
  	data: {
  		key: val,
  		key1: val1,
          key2: {
              subkey: subval
          }
  	}
  })
  ```

- 指令语法：`v-bind:`，或简写`:`，用于标签属性的值、标签体内容、绑定事件等，匹配vue实例data的属性或js表达式，，如`:href=key1`、`:href:Date.now()`

- 当data的属性也是个对象时，用插值或指令语法匹配时，就用`key.subkey`的形式

### 数据绑定

- `v-bind:`是单向的数据绑定，数据只能从data流向页面，即只能把vue实例的属性绑定给容器的属性，改变容器的属性值，不会影响vue实例的属性值，如

  ```js
  <div id="root">
      <input type="text", :value="key">
  </div>
  new Vue({
      el: '#root',
      data: {
          key: val
      }
  })
  ```

  在界面上改变输入框的内容时，不会改变vue实例的key的值

- `v-model:`是双向的数据绑定，数据既能从data流向页面，也能从页面流向data，即在界面上修改属性值时，也会影响vue实例属性的值

  - 但是只能用于表单类元素（输入类元素）
  - `v-model:value`可以简写为`v-model`，因为它默认收集的就是表单的value值

### el和data其他写法

```js
const v = new Vue()
// el的其他写法
v.$mount('css选择器')

// data的其他写法：函数式（日常开发基本只用这种方式），不能是箭头函数
const v = new Vue({
  data: function() {
    return {
      key: val
    }
  }
})
const v = new Vue({
  data() {
    return {
      key: val
    }
  }
})
```

### MVVM模型

- M：model，模型，对应data中的数据

- V：view，视图，对应Vue模板

- VM：ViewModel，视图模型，对应Vue实例对象

  ![image-20240525163803391](./vue教程.assets/image-20240525163803391.png)

  > 此图就是双向数据绑定，model的数据通过data bindings映射到view，而ViewModel要监听view的变化来改变model中的数据
  >
  > 所以用变量接收vue实例时，变量名称一般都是vm

- vue模板匹配的值，除了前面说的那些，vm对象里的所有属性都能得到

### 数据代理

通过一个对象代理另一个对象中属性的操作（读/写）

![image-20240525193234825](./vue教程.assets/image-20240525193234825.png)

### 事件处理

`v-on:`，和vue对象的`methods`属性匹配

- 简写形式：`@`
- 只能传一个参数，要么是普通形参，要么是event
- 如果既需要普通形参和event，就写成`$event`，且普通形参只能有一个
- 如果只需要event参数，则调用的地方可以不用传参，定义的地方需要写event形参
- 事件绑定的函数名，可以直接写函数体，且不需要{{}}就能直接得到vm中的属性，但是表达式只能是vm中的属性（只有函数体是一行时再这么写）

```html
<button v-on:click="clickFunc1">点击</button>
<button @click="clickFunc2(1)">点击</button>	<-- 简写形式，传普通参数 -->
<button @click="clickFunc3(1,$event)">点击</button>	<-- 传参，既需要普通参数，也需要event，event的位置随便 -->
<button @click="clickFunc4">点击</button> <-- 不传参，也能有event -->
<script>
  new Vue({
  ...
  methods: {
    clickFunc1() {	// 不需要function关键字，可以接受event（事件对象）形参
      
    },
    clickFunc2(num) {
      
    },
    clickFunc3(num, event) {
      
    },
    clickFunc4(event) {
      
    },
  }
})
</script>
```

![image-20240526162725508](./vue教程.assets/image-20240526162725508.png)

### 事件修饰符

`@click.prevent="clickFunc"`，直接在绑定事件的地方调用修饰符，vue有6种修饰符，常用前三个，可以链式调用

- prevent，阻止默认事件，相当于`e.preventDefault()`
- stop，阻止事件冒泡
- once，事件只触发一次
- capture，使用事件的捕获模式
- self，只有event.targe是当前操作的元素时才触发
- passive，事件的默认行为立即执行，无需等待事件回调即可执行完毕（移动端可能会用）
- number，控制输入的内容强转成数字
- lazy，输入框失去焦点后，绑定是属性才会进行数据代理
- trim，去除输入内容两边的空格

### 键盘事件对象的别名

vue常用的键盘事件的别名，也是在绑定事件的地方调用，比如：`@keydown.enter="keyFunc"`，相当于`if e.keyCode === 回车的ASCII码`时再触发事件，其他类似

- 回车，enter
- 删除，delete（捕获删除和退格键）
- 退出，esc
- 空格，space
- 换行，tab
- 上下左右，up、down、left、right

vue未提供的别名，可以使用按下键位的名称（即e.key的值）去绑定，但要转换为短横线连接小写单词的格式，比如`caps-lock`

几个特殊的键：系统修饰键，tab、ctrl、shift、alt、win/commadn键，需要用keydown绑定

Vue.config.keyCodes.自定义键名 = 键码（keycode），可以自定义别名

### 计算属性

`computed`，和data类似，只不过getter和setter（如果需要修改数据的话）需要显式地写出来（data的getter和setter是自带的 ）

![image-20240530212046520](vue教程.assets/image-20240530212046520.png)

```js
new Vue({
    ...
    data: {
        key1: val1,
        key2: val2
    },
    computed: {
        key3: {
            get() {
                retutn this.key1 + this.val2 
            }
        }
    }
})
```

- computed中属性的getter只有两种情况会被调用：
  - 该属性**首次**被调用时
  - getter所依赖的数据发生变化时
  - 注意：==data和computed中属性的setter都是只能接受一个参数==

#### 简写

如果只读取计算属性而无需修改，则getter可以去掉，把该属性作为get函数

```js
new Vue({
    ...
    data: {
        key1: val1,
        key2: val2
    },
    computed: {
        key3() {
            retutn this.key1 + this.val2 
        }
    }
})
```

### 监视属性

`watch`，监视属性（data和计算属性都可以监视）是否发生变化，如果发生变化，则调用`handler`，`handler`有两个形参，分别是变化后的值和变化前的值

![image-20240530215031944](vue教程.assets/image-20240530215031944.png)

```js
new Vue({
    ...
    data: {
        key1: val1,
        key2: val2
    },
    computed: {
        key3: {
            get() {
                retutn this.key1 + this.val2 
            }
        }
    },
    watch: {
      	immediate: true,	//默认false，true的话则是初始化时记忆调用一下handler
        key1: {	// 监视key1是否发生变化，也可以监视计算属性key3
            handler(newValue,oldValue) {
                // do something
            }
        }
    }
})
```

#### 另外一种写法

```js
const vm = new Vue({
    ...
    data: {
        key1: val1,
        key2: val2
    },
    computed: {
        key3: {
            get() {
                retutn this.key1 + this.val2 
            }
        }
    }
})
// key1外面的引号必须有
vm.$watch('key1',{
  			immediate: true,
        handler(newValue,oldValue) {
            // do something
        }
    })
```

#### 深度监视

![image-20240604190942145](./vue教程.assets/image-20240604190942145.png)

如果监视的属性是对象中的属性，则需要deep配置

```js
new Vue({
    ...
    data: {
        key1: val1,
        key2: val2,
      	keyObj1: {
          subKey1: subVal1,
          subKey2: subVal2
        }
    },
    computed: {
        key3: {
            get() {
                retutn this.key1 + this.val2 
            }
        }
    },
    watch: {
        key1: {	// 监视key1是否发生变化，也可以监视计算属性key3
            handler(newValue,oldValue) {
                // do something
            }
        },
      	keyObj1: {
          	deep: true,
          	handler() {
              // do something，keyObj1中的任何属性发生变化时，都会触发此handler
            }
        }
    }
})
```

#### 简写

当只有handler时，可以用简写形式，直接把简写的属性当成handler函数

```js
// 写法1
new Vue({
    ...
    data: {
        key1: val1,
        key2: val2
    },
    computed: {
        key3: {
            get() {
                retutn this.key1 + this.val2 
            }
        }
    },
    watch: {
        key1(newVal,oldVal): {
           // do something
        }
    }
})

// 写法2
const vm = new Vue({
    ...
    data: {
        key1: val1,
        key2: val2
    },
    computed: {
        key3: {
            get() {
                retutn this.key1 + this.val2 
            }
        }
    }
})
vm.$watch('key1', function() {
  // do something
})
```

### 计算属性和监视属性比较

![image-20240604193555936](./vue教程.assets/image-20240604193555936.png)

### 绑定样式

#### 绑定class样式

用`v-bind`去绑定类名，vue会自动把绑定的类名附加到已有的类名后面

```js
<div class='class1' :class="bindClass" @click="clickFunc">{{demo}}</div>
<script type="text/javascript">
  new Vue({
  	//...
  	data:{
      bindClass: 'class2'
    }
})
</script>
```

当需要绑定的样式的名称和数量不确定时，可以绑定一个数组类型的属性，然后用push、shift等方式修改数组 

当需要绑定的样式的名称和数量确定，但需要动态决定用不用，可以绑定一个对象类型的属性，然后对象的属性是类名，值为布尔值（true的时候就会应用此样式）

#### 绑定style样式（不常用）

思路和绑定class类似，只不过style对象的属性必须是css中存在的属性，且使用vue要求的小驼峰命名

![image-20240604200535176](./vue教程.assets/image-20240604200535176.png)

### 条件渲染

![image-20240605200411381](./vue教程.assets/image-20240605200411381.png)

条件渲染的值可以是布尔值、表达式或vm的属性（常用vm的属性）

- `v-show="布尔值"`，底层通过控制`dislpay`属性来控制元素的显示和隐藏
- `v-if="布尔值"`，通过删除元素来控制元素的隐藏
- `v-else-if="布尔值"`，和`v-if`一起使用，逻辑和`if...else if`的逻辑一样
- `v-else`，和`v-if`一起使用，逻辑和`if...else`的逻辑一样
- ==v-if、v-else-if、v-else之间不能有其他元素，必须是连贯的==
- `<template></template>`标签，不影响css渲染的结构，只能和`v-if`结合使用，用于包裹几个元素，然后在`template`标签中写`v-if`同时让这几个元素进行显示隐藏，避免在每个元素中都写`v-if`或`v-show`

### 列表渲染

- `v-for` 

  ![image-20240605204134625](./vue教程.assets/image-20240605204134625.png)

- `v-for`中的key

![image-20240605204110972](./vue教程.assets/image-20240605204110972.png)

### 列表过滤

用计算属性和监视属性都可以实现，优先用计算属性实现

![image-20240610161504422](./vue教程.assets/image-20240610161504422.png)

### 列表排序

也用计算属性

![image-20240610162211576](./vue教程.assets/image-20240610162211576.png)

### Vue.set

![image-20240610173331505](./vue教程.assets/image-20240610173331505.png)

给vue对象添加属性：`Vue.set(target,key,val)`，也可以用`vm.$set(target,key,val)`

- target是要添加新属性的属性，==key可以是对象的属性，也可以是数组的索引==
- vm是vue实例对象，在new Vue的代码里面，就是this
- ==不能给vue对象或vue对象的根对象(data)添加属性，只能给data中是对象类型的属性添加属性==

### 收集表单属性

![image-20240611220525345](./vue教程.assets/image-20240611220525345.png)

### 过滤器

`filters`，对数据进行单向的格式化，不改变原数据

```js
<div>{{demo | filterName}}</div>
<script type="text/javascript">
  new Vue({
  	//...
    data: {
      demo: 'demoValue'
    }
  	filters:{
      filterName(value) {
        return 处理后的value
      }
    }
})
</script>
```

