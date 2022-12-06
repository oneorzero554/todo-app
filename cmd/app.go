package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    api := router.Group("api")
    {
        list := api.Group("list")
        {
            list.GET("/")
            list.GET("/:id")
            list.POST("/")
            list.PUT("/:id")
            list.DELETE("/:id")

            items := list.Group(":id/items")
            {
                items.GET("/")
                items.POST("/")
            }
        }

        items := api.Group("items")
        {
            items.GET("/:id")
            items.PUT("/:id")
            items.DELETE("/:id")
        }

    }

    router.Run()
}