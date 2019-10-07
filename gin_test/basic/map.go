package main  

import  "github.com/gin-gonic/gin"


func main() {
	r :=gin.Default()
	r.POST("/",func(c *gin.Context) {
		ids :=c.QueryMap("ids")
		names :=c.PostFormMap("names")

		c.JSON(200,gin.H{
			"ids":ids,
			"names":names,
		})
		c.String(200,"ids:%v names:%v",ids,names)
	})
	r.Run()
}




