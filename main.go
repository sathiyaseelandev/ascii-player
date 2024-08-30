package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/gosuri/uilive"
	"gocv.io/x/gocv"
)

func main() {
	video, err := gocv.VideoCaptureFile("./badapple.mp4")
	if err != nil {
		fmt.Println("Failed to read video")
		os.Exit(1)
	}
	img := gocv.NewMat()
	// window := gocv.NewWindow("Bad Apple")

	x := 480
	y := 360

	rows := 36
	cols := 96

	blockHeight := y / rows
	blockWidth := x / cols

	writer := uilive.New()
	writer.Start()

	for {
		video.Read(&img)
		asciiMat := make([][]byte, rows)
		for i := 0; i < rows; i++ {
			asciiMat[i] = make([]byte, cols)
		}

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				sum := 0
				for i := blockHeight * row; i < blockHeight*row+blockHeight-1; i++ {
					for j := blockWidth * col; j < blockWidth*col+blockWidth-1; j++ {
						sum += int(img.GetVecbAt(i, j)[0])
					}
				}
				avg := sum / (blockHeight * blockWidth)
				intensity := " .:-=+*#%@"

				intensityLen := len(intensity)
				part := 256 / intensityLen
				asciiMat[row][col] = intensity[avg / part]
			}
		}

		// window.IMShow(img)
		var str strings.Builder
		for i := 0; i < rows; i++ {
			str.WriteString(string(asciiMat[i]))
			str.WriteString("\n")
		}
		writer.Flush()
		fmt.Fprint(writer, str.String())
		// window.WaitKey(1)
	}
}
