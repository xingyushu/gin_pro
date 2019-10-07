package main  

import (
	 "github.com/gin-gonic/gin"
	 "net/http"
)

func main() {
	r :=gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index",func(c *gin.Context) {
		c.HTML(200,"posts/index.tmpl",gin.H{
				"title":"Posts",
			})
		
	})

	r.GET("/users/index",func(c *gin.Context) {
	c.HTML(200,"users/index.tmpl",gin.H{
				"title":"Users",
			})
		
	})

	//重定向网址
	r.GET("/Redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

    //路径重定向
	r.GET("/test", func(c *gin.Context) {
    	c.Request.URL.Path = "/test2"
    	r.HandleContext(c)
	})

	r.GET("/test2", func(c *gin.Context) {
    	c.JSON(200, gin.H{"hello": "world"})
	})

	r.Run()
}

