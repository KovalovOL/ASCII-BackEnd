package utils

import (
	"os"
	"image"
	_ "image/jpeg"
	"image/png"
	"image/color"
	"golang.org/x/image/draw"
	"os/exec"
	"runtime"
	"fmt"
)


func GetImage(path string) (image.Image, string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return image.Decode(file)
}

func ScaleImage(src image.Image, scale float64) image.Image {
	if scale <= 0 {
		return src
	}

	bounds := src.Bounds()

	newWidth := int(float64(bounds.Dx()) * scale)
	newHeight := int(float64(bounds.Dy()) * scale)

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	draw.CatmullRom.Scale(
		dst,
		dst.Bounds(),
		src,
		bounds,
		draw.Over,
		nil,
	)

	return dst
}

func ScaleToHeight(src image.Image, targetHeight int) image.Image {
	b := src.Bounds()

	srcWidth := b.Dx()
	srcHeight := b.Dy()

	aspect := float64(srcWidth) / float64(srcHeight)

	newHeight := targetHeight
	newWidth := int(float64(targetHeight) * aspect)

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	draw.CatmullRom.Scale(
		dst,
		dst.Bounds(),
		src,
		b,
		draw.Over,
		nil,
	)

	return dst
}

func SavePNG(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func ToGrayscale(src image.Image) image.Image {
	b := src.Bounds()

	dst := image.NewRGBA(b)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {

			r, g, b, a := src.At(x, y).RGBA()
			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)

			gray := uint8(0.299*r8 + 0.587*g8 + 0.114*b8)

			dst.Set(x, y, color.RGBA{
				R: gray,
				G: gray,
				B: gray,
				A: uint8(a >> 8),
			})
		}
	}

	return dst
}

func Brightness(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	r8 := float64(r >> 8)
	g8 := float64(g >> 8)
	b8 := float64(b >> 8)

	return 0.299*r8 + 0.587*g8 + 0.114*b8
}


func OpenViewer(path string) error {
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("open", path)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to open file on macOS: %w", err)
		}
		return nil
	}

	return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
}