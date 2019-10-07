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