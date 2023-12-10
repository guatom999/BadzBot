package main

import "gocv.io/x/gocv"

func main() {
	// video, err := gocv.VideoCaptureFile("../video/video.mp4")
	// if err != nil {
	// 	// return nil, err
	// 	panic(err)
	// 	// return err
	// }
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}

	// return nil
}
