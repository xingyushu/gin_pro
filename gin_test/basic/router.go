package main  

import  "github.com/gin-gonic/gin"

func main() {
	r :=gin.Default()

	v1 :=r.Group("/v1")
	{
		v1.GET("/read",func (c *gin.Context) {
			c.String(200,"hello v1 reading")
		})
		v1.GET("/watch",func (c *gin.Context) {
			c.String(200,"hello v1  watching")
		})
	}

	v2 :=r.Group("/v2")
	{
	v2.GET("/login",func (c *gin.Context) {
			c.String(200,"hello v2")
		})
	}

	v3 :=r.Group("/v3")
	{
	v3.GET("/name",func (c *gin.Context) {
			c.String(200,"hello v3")
		})
	}

	r.Run()
}


// 浏览器输入http://127.0.0.1:8080/v1/read
// 输出hello v1

// 浏览器输入http://127.0.0.1:8080/v2/login
// 输出hello v2

// 浏览器输入http://127.0.0.1:8080/v3/name
// 输出hello v3