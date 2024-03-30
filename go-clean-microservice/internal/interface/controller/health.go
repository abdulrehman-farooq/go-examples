package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)
type HealthCheckResponse struct {
    Status string `json:"status"`
}

func Health() gin.HandlerFunc{
    return func(c *gin.Context) {
        response := HealthCheckResponse{Status: "ok"}
        c.JSON(http.StatusOK, response)
    }
}
