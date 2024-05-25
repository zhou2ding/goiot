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

- 插值语法：`{{}}`，用于标签体的内容，匹配vue实例data的属性或js表达式，如`{{key}}`、`{{Date.now()}}`

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

- `v-model:`是双向的数据绑定，数据既能从data流向页面，也能从页面流向data，即在界面上修改属性值时，也会影响vue实例属性的值，但是只能用于表单类元素（输入类元素）