package main

import (
	"log"
	"time"

	"gocv.io/x/gocv"
)

func main() {

	window := gocv.NewWindow("2-11")
	defer window.Close()

	//winLog := gocv.NewWindow("Log Polar")

	vc, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		return
	}
	defer vc.Close()
	time.Sleep(5000)
	//fps := vc.Get(gocv.VideoCaptureFPS)
	//frameWidth := gocv.VideoCaptureFrameWidth
	//frameHeight := gocv.VideoCaptureFrameHeight

	vwfile, err := gocv.VideoWriterFile("cam_backup.avi",
		"XVID", 20, 640, 480, true)

	if err != nil {
		log.Fatalln("vwfile:" + err.Error())
	}
	defer vwfile.Close()

	im := gocv.NewMat()
	defer im.Close()

	for {
		if !vc.Read(&im) {
			continue
		}

		window.IMShow(im)
		if vwfile.Write(im) != nil {
			log.Println("error writing to the file.")
		}

		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
