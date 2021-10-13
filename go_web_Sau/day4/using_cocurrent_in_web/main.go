package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	imgfile, err := os.Open("cat_m.jpg")

	if err != nil {
		fmt.Println("img.png file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()

	img, _, err := image.Decode(imgfile)

//	fmt.Println(img.At(10, 10))

	bounds := img.Bounds()

	//r := image.Rectangle{image.Point{0,0}, image.Point{100,100}}

	fmt.Println(bounds)

	fmt.Println("Min X:", bounds.Min.X)
	fmt.Println("Min Y:", bounds.Min.Y)
	fmt.Println("Max X:", bounds.Max.X)
	fmt.Println("Max Y:", bounds.Max.Y)

	img1 := resize(img, 20)
	bounds = img1.Bounds()

	//r := image.Rectangle{image.Point{0,0}, image.Point{100,100}}
	fmt.Println("=========")
	fmt.Println(bounds)

	fmt.Println("Min X:", bounds.Min.X)
	fmt.Println("Min Y:", bounds.Min.Y)
	fmt.Println("Max X:", bounds.Max.X)
	fmt.Println("Max Y:", bounds.Max.Y)
	//canvas := image.NewNRGBA(bounds) // <-- here
	//
	//stride := canvas.Stride
	//
	//fmt.Println("Stride : ", stride)
}