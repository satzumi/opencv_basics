package main

import (
	"image"

	"gocv.io/x/gocv"
)

func main() {
	im := gocv.IMRead("dog.jpg", gocv.IMReadUnchanged)
	if true == im.Empty() {
		return
	}
	winIn := gocv.NewWindow("Input")
	defer winIn.Close()

	winOut := gocv.NewWindow("Ouput-PyrDown")
	defer winOut.Close()

	imPy := gocv.NewMat()
	defer imPy.Close()

	gocv.PyrDown(im, &imPy, image.Pt(im.Rows()/2, im.Cols()/2), gocv.BorderDefault)
	for {
		winIn.IMShow(im)
		winOut.IMShow(imPy)

		if winIn.WaitKey(1) >= 0 {
			break
		}
	}
}
