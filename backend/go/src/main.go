package main

import (
	"express/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Running")
	InitRouter()
}

func InitRouter() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	au := router.Group("/api/v1/auth")
	{
		au.POST("/sessions", auth.Sessions)
		au.POST("/users", auth.Users)
	}

	order := router.Group("api/v1/order")
	order.Use(auth.JWTAuth())
	{
		order.GET("/test", func(c *gin.Context) {
			claims := c.MustGet("claims").(*auth.CustomClaims)
			if claims != nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  0,
					"message": "token有效",
					"data":    claims,
				})
			}
		})
	}

	router.Run(":8080")
}
