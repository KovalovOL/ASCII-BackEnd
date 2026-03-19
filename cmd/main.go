package main

import (
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
	r.Run(":8080")
}