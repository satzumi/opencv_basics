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

	//classifier
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	classifier.Load("Nariz.xml")

	im := gocv.NewMat()
	defer im.Close()

	window := gocv.NewWindow("Web Cam- Nose Detect")
	color := color.RGBA{0, 255, 0, 0}

	for {
		if false == webcam.Read(&im) {
			continue
		}

		rects := classifier.DetectMultiScale(im)

		for _, r := range rects {
			gocv.Rectangle(&im, r, color, 2)
		}

		window.IMShow(im)
		window.WaitKey(3)
	}
}
