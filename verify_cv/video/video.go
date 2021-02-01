package main

import (
	"github.com/sudosapan/morego/verify_cv/util"
	"gocv.io/x/gocv"
)

func main() {
	camera, err := gocv.VideoCaptureDevice(0)
	util.LogFatal(err, "Can;t get device 0")

	win := gocv.NewWindow("Camera")
	img := gocv.NewMat()

	for {
		camera.Read(&img)
		win.IMShow(img)
		if win.WaitKey(1) >= 0 {
			break
		}

	}

}
