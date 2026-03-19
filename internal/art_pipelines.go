package internal

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"app/internal/utils"
	"strings"
)

func CreatePixelArt(src image.Image, charLine string) string {
	b := src.Bounds()
	bCoef := 256 / len(charLine)

	var sb strings.Builder


	for y := b.Min.Y; y < b.Max.Y; y += 2 {
		for x := b.Min.X; x < b.Max.X; x++ {

			brightness := utils.Brightness(src.At(x, y))
			charI := int(brightness) / bCoef

			if charI >= len(charLine) {
				charI = len(charLine) - 1
			}

			sb.WriteByte(charLine[charI])
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}