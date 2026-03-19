package routers

import (
	"app/internal/utils"
	"app/internal"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PixelArtHandler(c *gin.Context) {
	charLine := c.PostForm("palette")
	if charLine == "" {
		charLine = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,^`'. "
	}

	scale := c.PostForm("scale")
	if scale == "" {
		scale = "1"
	}
	scaleFloat, err := strconv.ParseFloat(scale, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "scale is not correct"})
	}	

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "image required"})
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid image"})
		return
	}
	resizedImg := utils.ScaleImage(img, scaleFloat)
	ascii := internal.CreatePixelArt(resizedImg, charLine)

	bounds := resizedImg.Bounds()
	c.JSON(http.StatusOK, gin.H{
		"ascii": ascii,
		"height": int(bounds.Dy() / 2),
		"width": bounds.Dx(),
		"palette_lenght": len(charLine),
	})
}