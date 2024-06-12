package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"time"
)

type Job struct {
	Image   image.Image
	OutPath string
}

var imagePaths = []string{
	"images/image1.jpeg",
	"images/image2.jpeg",
	"images/image3.jpeg",
	"images/image4.jpeg",
	"images/image5.jpeg",
	"images/image6.jpeg",
	"images/image7.jpeg",
	"images/image8.jpeg",
}

func withoutConcurrentPipeline() {
	for _, p := range imagePaths {
		Image := imageprocessing.ReadImage(p)

		resizedImage := imageprocessing.Resize(Image)

		greyscaledImage := imageprocessing.Grayscale(resizedImage)

		imageprocessing.WriteImage("output/"+p, greyscaledImage)
	}

}

func main() {
	nowTime := time.Now()
	withoutConcurrentPipeline()
	fmt.Printf("Without concurrent done in %s\n", time.Since(nowTime))
}
