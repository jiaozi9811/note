cn.vue.org 
 
 指令
 v-text     更新元素的textContent。如果要更新部分的textContent ，需要使用{{ Mustache }}插值
 v-once     能执行一次性地插值，当数据改变时，插值处的内容不会更新
 v-html     更新元素的innerHTML
 v-bind     动态地绑定一个或多个特性，或一个组件 prop 到表达式
 v-show     根据表达式之真假值，切换元素的 display CSS 属性
 v-if
 v-else
 v-else-if
 v-for      <div v-for="item in items">
 v-on       绑定事件监听器
      keydown  keyup  click  dbclick  load
 v-model    在表单控件或者组件上创建双向绑定
 v-pre      跳过这个元素和它的子元素的编译过程
 
 
 选项/DOM
 el     提供一个在页面上已存在的DOM元素作为Vue实例的挂载目标。可以是CSS选择器，也可以是一个HTMLElement实例
        可以用vm.$el访问
 template   一个字符串模板作为 Vue 实例的标识使用。模板将会 替换 挂载的元素
            常用的技巧是用<script type="x-template">包含模板
render


生命周期钩子
生命周期钩子自动绑定this上下文到实例中，因此你可以访问数据，对属性和方法进行运算

beforeCreate    在实例初始化之后，数据观测 (data observer) 和 event/watcher 事件配置之前被调用
created         实例创建完成后被立即调用。在这一步，实例已完成以下的配置：数据观测(data observer)，属性和方法的运算，watch/event事件回调。然而，挂载阶段还没开始，$el 属性目前不可见
beforeMount     在挂载开始之前被调用：相关的 render 函数首次被调用
mounted         el被新创建的vm.$el替换，并挂载到实例上去之后调用该钩子
beforeUpdate    数据更新时调用，发生在虚拟 DOM 打补丁之前
updated         当这个钩子被调用时，组件 DOM 已经更新，所以你现在可以执行依赖于 DOM 的操作
activated       keep-alive 组件激活时调用
deactivated     keep-alive 组件停用时调用
beforeDestroy   实例销毁之前调用。在这一步，实例仍然完全可用
destroyed       Vue 实例销毁后调用。调用后，Vue 实例指示的所有东西都会解绑定，所有的事件监听器会被移除，所有的子实例也会被销毁


实例属性
vm.$data
vm.$props       当前组件接收到的props对象。Vue 实例代理了对其 props 对象属性的访问
vm.$el          Vue 实例使用的根 DOM 元素
vm.$options     用于当前 Vue 实例的初始化选项。需要在选项中包含自定义属性时会有用处
vm.$parent      父实例，如果当前实例有的话
vm.$root        当前组件树的根 Vue 实例。如果当前实例没有父实例，此实例将会是其自己
vm.$children    当前实例的直接子组件
vm.$refs        一个对象，持有注册过 ref 特性 的所有 DOM 元素和组件实例
vm.$isServer    当前 Vue 实例是否运行于服务器




<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>JS</title> 
<script src="https://cdnjs.cloudflare.com/ajax/libs/vue/2.0.3/vue.js"></script>
</head>
<body>
  <div id="app">
    <input type="text" placeholder="insert"
           @keyup.enter="addTodo(newTodo)" v-model="newTodo">
    <h2>list</h2>
      <ul>
        <li v-for="todo in todos">
          {{todo.content}} - <a href="#" @click.prevent="removeTodo(todo)">delete</a>
        </li>
      <ul>
  </div>   
</body>
</html>
------------------------------
new Vue({
      el: '#app',
      data: {
        todos: [],
        newTodo: ''
      },
      methods:{
        addTodo: function(todo){
          this.todos.push({content:todo,completed:false})},
        removeTodo: function(todo){
          this.todos.splice(this.todos.indexOf(todo),1);}
      },
    })

