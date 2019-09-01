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



## 　3．gin  get/post传参
　
定义方法，路由规则
对应的处理函数，包括：
　　　获取参数，定义参数展示的方式，如String,JSON等
   
```
func (*Context) Query  获取url上参数，如果没有则返回空
func (c *Context) Query(key string) string
Query returns the keyed url query value if it exists, otherwise it returns an empty string `("")`. It is shortcut for `c.Request.URL.Query().Get(key)`

    GET /path?id=1234&name=Manu&value=
	   c.Query("id") == "1234"
	   c.Query("name") == "Manu"
	   c.Query("value") == ""
	   c.Query("wtf") == ""
```


```
func (*Context) DefaultQuery　　获取url上参数，如果没有则返回默认值
func (c *Context) DefaultQuery(key, defaultValue string) string
DefaultQuery returns the keyed url query value if it exists, otherwise it returns the specified defaultValue string. See: Query() and GetQuery() for further information.

GET /?name=Manu&lastname=
c.DefaultQuery("name", "unknown") == "Manu"
c.DefaultQuery("id", "none") == "none"
c.DefaultQuery("lastname", "none") == ""
```

｀｀｀

get/post/get+post method examples:
```
package main  

import  "github.com/gin-gonic/gin"

func main() {
   r :=gin.Default()
   // r.GET("/get/:name",getting)	
   r.GET("/get",func (c *gin.Context) {
   	     firstname := c.DefaultQuery("firstname","Mike")
   	     lastname :=c.Query("lastname")
   	     c.String(200,"Hello %s %s",firstname,lastname)
   })


   r.GET("/get/:name",func(c *gin.Context){
   		name :=c.Param("name")
	    c.String(200,"Hello %s",name)
   })


   r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(200, message)
   })

   //post method
   r.POST("/post",func(c *gin.Context) {
   	   name :=c.PostForm("name")
   	   message :=c.PostForm("message")
   	   age :=c.DefaultPostForm("age","15y")
   	   
   	   c.JSON(200,gin.H{
   	   		"status":"posted",
   	   		"name":name,
   	   		"message":message,
   	   		"age":age,
   	   	})
   })

   //get+post mixed method
   r.POST("/post2",func(c *gin.Context) {
       id := c.DefaultQuery("id","001")
   	   name :=c.PostForm("name")
   	   age :=c.DefaultPostForm("age","15")   	   
   	   c.String(200,"Hello %sth user:%s %s years old",id,name,age)
   })



   r.Run()
}




// func getting(c *gin.Context) {
// 	name :=c.Param("name")
// 	c.String(200,"Hello %s",name)
	
// }
```
