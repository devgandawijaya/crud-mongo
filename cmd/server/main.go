package main

import (
    "crud-mongo/internal/config"
    "crud-mongo/internal/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":" + config.AppPort)
}
