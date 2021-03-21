package main

import (
	"image/color"

	"gocv.io/x/gocv"
)

func main() {

	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		return
	}
	defer webcam.Close()

	// window
	window := gocv.NewWindow("Web Cam")
	defer window.Close()

	// image matrix
	im := gocv.NewMat()
	defer im.Close()

	// harrcascade
	harrcascade := "haarcascade_frontalface_default.xml"

	// classifier
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)

	// color to draw rectangle
	color := color.RGBA{0, 255, 0, 0}
	for {
		// read frame from device
		if !webcam.Read(&im) {
			continue
		}
		// detect object face in the frame
		rects := classifier.DetectMultiScale(im)

		// draw rectangle(s) in frame
		for _, r := range rects {
			gocv.Rectangle(&im, r, color, 3)
		}

		window.IMShow(im)
		window.WaitKey(1)
	}
}
