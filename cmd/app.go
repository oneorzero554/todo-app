package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "todo-app/internal/handlers"
    "todo-app/internal/repositories"
    "todo-app/internal/services"
    "todo-app/pkg/db"
)

func main() {

    dbc, err := db.NewPostgresClient(db.Config{})

    if err != nil {
        log.Fatal(err)
    }

    repos := repositories.InitRepostitories(dbc)
    svcs := services.InitServices(repos)
    hdls := handlers.InitHandlers(svcs)

    router := gin.Default()

    api := router.Group("api")
    {
        list := api.Group("list")
        {
            list.GET("/", hdls.IndexList)
            list.GET("/:id", hdls.ShowList)
            list.POST("/", hdls.CreateList)
            list.PUT("/:id", hdls.UpdateList)
            list.DELETE("/:id", hdls.DeleteList)

            items := list.Group(":id/items")
            {
                items.GET("/", hdls.IndexTask)
                items.POST("/", hdls.CreateTask)
            }
        }

        items := api.Group("items")
        {
            items.GET("/:id", hdls.ShowTask)
            items.PUT("/:id", hdls.UpdateTask)
            items.DELETE("/:id", hdls.DeleteTask)
        }
    }

    if err := router.Run(); err != nil {
        log.Fatal(err)
    }
}