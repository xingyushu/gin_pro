//this is the conduct of gin-H example
//author:Mike
//2019--09-01

package main  

import  "github.com/gin-gonic/gin"

func main() {
	r  :=gin.Default()
	r.GET("/ping",func (c *gin.Context) {
		c.JSON(200,gin.H{  "info":gin.H{
				"status":"login successfully",
				"name":"Mike",
		    },
			"message":"Mike coding",
			})
	})
	r.Run()
}