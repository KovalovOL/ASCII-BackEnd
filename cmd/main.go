package main

import (
	"net/http"
	"app/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		},
	)
	r.POST("/pixelart", routers.PixelArtHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}