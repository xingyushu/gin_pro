package main

import  "github.com/gin-gonic/gin"

type Login struct {
	User string   `form:"user"  json:"user" binding:"required"`
	Password string   `form:"password"  json:"password" binding:"required"`
}


func main() {
	r :=gin.Default()

	r.POST("/login",func(c *gin.Context) {
		var json  Login
		if c.BindJSON(&json) == nil{
			if json.User == "manu" && json.Password == "123" {
				c.JSON(200,gin.H{  "status":"Login successfully!"})
			}else{
				c.JSON(200,gin.H{
					"status":"Login Failed!",
					})
			}
		}
	})

	r.Run()
}