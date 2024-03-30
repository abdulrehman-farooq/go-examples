package main

import (
    "github.com/gin-gonic/gin"
    "github.com/abdulrehman-farooq/go-examples/go-clean-microservice/internal/interface/controller"
)
func main() {
    router := gin.Default()
    router.GET("/health", controller.Health())
    router.Run("localhost:8080")
}
