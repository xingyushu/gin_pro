package routers

import (
    "github.com/gin-gonic/gin"

    "blog/pkg/setting"
    "blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunMode)

    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })


    apiv1 :=r.Group("/api/v1")
    {
        //get
        apiv1.GET("/tags",v1.GetTags)

        //add
        apiv1.POST("/tags",v1.AddTag)

        //update
        apiv1.PUT("/tags/:id",v1.EditTag)

        //delete
        apiv1.DELETE("/tags/:id",v1.DeleteTag)


    }

    return r
}