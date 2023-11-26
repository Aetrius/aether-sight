package aetherImageCapture

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func winScreenshot() ([]string, bool, error) {
	var err error
	n := screenshot.NumActiveDisplays()
	var images []string

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("./%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		images = append(images, fileName)
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
	return images, true, err
}
