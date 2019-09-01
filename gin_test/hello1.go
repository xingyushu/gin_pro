//this is the conduct of gin hello-world example
//author:Mike
//2019--09-01

package main  

import (
   "github.com/gin-gonic/gin"
   "net/http"
)


func main() {
	
	router :=gin.Default()
	router.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"Hello world")
		
	})
	router.Run(":8000")
}

