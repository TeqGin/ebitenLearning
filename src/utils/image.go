package utils

import (
	"bytes"
	"ebitenLearning/src/resource"
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nfnt/resize"
)

func ResizeImageFromFile(path string, scalar float64) image.Image {
	// Open the original image file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Decode the original image
	originalImg, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil
	}
	return ResizeImage(originalImg, scalar)
}

func ResizeImageFromReader(path string, scalar float64) image.Image {
	b, _ := resource.Asset(path)
	// Decode the original image
	originalImg, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		fmt.Println(path)
		fmt.Println("Error decoding image:", err)
		return nil
	}
	return ResizeImage(originalImg, scalar)
}

func ResizeImage(originalImg image.Image, scalar float64) image.Image {
	size := originalImg.Bounds().Size()
	// Resize the image to the new dimensions
	resizedImg := resize.Resize(uint(float64(size.X)*scalar), uint(float64(size.Y)*scalar), originalImg, resize.Lanczos3)
	return resizedImg
}

func NewEbitenImangeFromFile(path string) *ebiten.Image {
	b, _ := resource.Asset(path)
	img, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(b))
	return img
}
