
### 目的

这里摘录的是非常简易版的 gin 
https://github.com/ddl-killer/gin-read/tree/main/simplified_gin/lorgin

原始的 gin 框架代码是比较繁琐的，这里首先整理调试一个简易版的 gin 框架，方便学习核心思路。

### 测试命令

```azure
   curl http://localhost:8080/email/narvar
   curl http://localhost:8080/email/hhh
   curl http://localhost:8080/email/shopify
   curl http://localhost:8080/panic
```

### 问题
engine.RouterGroup.engine = engine  gin框架中这种设计有什么好处，为什么成员是自己本身呢

在 gin 框架中，engine.RouterGroup.engine = engine 这种设计主要是为了实现路由分组的功能。  
在 gin 中，Engine 结构体既是一个 HTTP 的路由处理器，也是一个路由分组 RouterGroup。
Engine 结构体中的 RouterGroup 是作为根路由分组存在的，它的 engine 字段指向 Engine 自身，这样做的好处是可以在 RouterGroup 中访问到 Engine 的方法和字段。  
当我们创建一个新的路由分组时，新的 RouterGroup 的 engine 字段也会指向这个 Engine，这样无论我们在哪个路由分组中，都可以通过 engine 字段访问到 Engine 的方法和字段，实现了路由分组的功能。  这种设计方式使得 Engine 和 RouterGroup 之间的关系变得紧密，使得在处理 HTTP 请求时，可以方便地在不同的路由分组之间切换，同时也可以在路由分组中方便地使用 Engine 的功能