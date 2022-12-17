package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
    "todo-app/internal/handlers"
    "todo-app/internal/repositories"
    "todo-app/internal/services"
    "todo-app/pkg/db"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatal("error loading .env file")
    }

    dbc, err := db.NewPostgresClient(db.Config{
        Host: os.Getenv("DB_HOST"),
        Port: os.Getenv("DB_PORT"),
        Username: os.Getenv("DB_USERNAME"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName: os.Getenv("DB_NAME"),
        SSLMode: os.Getenv("DB_SSLMODE"),
    })

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