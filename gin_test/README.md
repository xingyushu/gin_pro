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
