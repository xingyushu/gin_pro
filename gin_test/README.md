# go gin started

 "github.com/gin-gonic/gin"
 
 https://blog.csdn.net/u013210620/article/details/82773905
 
 https://godoc.org/github.com/gin-gonic/
 
 https://github.com/gin-gonic/gin

## 　１．对gin.H的理解

// H is a shortcut for map[string]interface{}
type H map[string]interface{}


也就是说，

c.JSON(http.StatusOK, gin.H{ "status": "登录成功"})
等价于

c.JSON(http.StatusOK, map[string]interface{}{ "status": "登录成功"})
因为没用过这个框架，简单推测了一下，引入 gin.H 这个东西可以简化生成 json 的方式，如果需要嵌套 json，那么嵌套 gin.H 就可以了。例子：

```
	r.GET("/ping",func (c *gin.Context) {
		c.JSON(200,gin.H{  "info":gin.H{
				"status":"login successfully",
				"name":"Mike",
		    },
			"message":"Mike coding",
			})
	})
 ```
 
 效果如图：
 
 ```
 {"info":{"name":"Mike","status":"login successfully"},"message":"Mike coding"}
 ```
 
 ## 　2．对gin结构的理解
 ```
 package main
 import "github.com/gin-gonic/gin"
　func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
```
简单几行代码，就能实现一个web服务。使用gin的Default方法创建一个路由handler。然后通过HTTP方法绑定路由规则和路由函数。不同于net/http库的路由函数，gin进行了封装，把request和response都封装到gin.Context的上下文环境。最后是启动路由的Run方法监听端口。麻雀虽小，五脏俱全。当然，除了GET方法，gin也支持POST,PUT,DELETE,OPTION等常用的restful方法。

