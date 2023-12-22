package utils

import (
	"fmt"
	"image"
	"os"

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

func ResizeImage(originalImg image.Image, scalar float64) image.Image {
	size := originalImg.Bounds().Size()
	// Resize the image to the new dimensions
	resizedImg := resize.Resize(uint(float64(size.X)*scalar), uint(float64(size.Y)*scalar), originalImg, resize.Lanczos3)
	return resizedImg
}
