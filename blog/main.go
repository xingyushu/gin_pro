package main 


import (
    _ "github.com/gin-gonic/gin"
    "fmt"
    "net/http"
    "blog/pkg/setting"
    "blog/routers"
 )


func main() {

	router :=routers.InitRouter()
	// router :=gin.Default()
	// router.GET("/",func(c *gin.Context) {
	// 	c.JSON(200,gin.H{
	// 		 "message":"test",
	// 		})
	// }) 

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

	s.ListenAndServe()
}