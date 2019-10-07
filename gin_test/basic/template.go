package main   

import  "github.com/gin-gonic/gin"


func main() {
	r :=gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index",func(c *gin.Context) {
		c.HTML(200,"index.tmpl",gin.H{
				"title":"Main site",
			})
	})


	r.Run()
}

