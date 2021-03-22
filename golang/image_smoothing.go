package main

import (
	"image"

	"gocv.io/x/gocv"
)

func main() {

	//imag matrix of original image
	im := gocv.IMRead("dog.jpg", -1)
	defer im.Close()

	// create windows to hold our images
	windowIn := gocv.NewWindow("In")
	windowOut := gocv.NewWindow(("Out"))
	defer windowIn.Close()
	defer windowOut.Close()

	// create image matrix to hold smooth output
	imOut := gocv.NewMat()
	defer imOut.Close()

	// do the smoothing
	gocv.GaussianBlur(im, &imOut, image.Pt(5, 5), 0, 0, gocv.BorderDefault)

	for {
		// show the images
		windowIn.IMShow(im)
		windowOut.IMShow(imOut)

		// press any key to break out of loop and terminate the program
		if windowIn.WaitKey(1) >= 0 || windowOut.WaitKey(1) >= 0 {
			break
		}

	}
}
