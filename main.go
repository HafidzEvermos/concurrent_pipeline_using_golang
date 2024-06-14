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

func withConcurrentPipeline() {
	var (
		imageChan           = make(chan Job)
		resizedImageChan    = make(chan Job)
		greyscaledImageChan = make(chan Job)
	)

	go func() {
		for _, p := range imagePaths {
			image := imageprocessing.ReadImage(p)
			imageChan <- Job{
				Image:   image,
				OutPath: "concurrent_output/" + p,
			}
		}
		close(imageChan)
	}()

	go func() {
		for image := range imageChan {
			resizedImage := imageprocessing.Resize(image.Image)
			resizedImageChan <- Job{
				Image:   resizedImage,
				OutPath: image.OutPath,
			}
		}
		close(resizedImageChan)
	}()

	go func() {
		for image := range resizedImageChan {
			greyscaledImage := imageprocessing.Grayscale(image.Image)
			greyscaledImageChan <- Job{
				Image:   greyscaledImage,
				OutPath: image.OutPath,
			}
		}
		close(greyscaledImageChan)
	}()

	for image := range greyscaledImageChan {
		imageprocessing.WriteImage(image.OutPath, image.Image)
	}
}

func main() {
	nowTime := time.Now()
	withoutConcurrentPipeline()
	fmt.Printf("Without concurrent done in %s\n", time.Since(nowTime))

	nowTime = time.Now()
	withConcurrentPipeline()
	fmt.Printf("   With concurrent done in %s\n", time.Since(nowTime))
}
