package aetherImageCapture

import (
	"errors"
	"fmt"
	"runtime"
)

func captureScreen() {
	var images []string
	var capture bool = false
	var err error

	if runtime.GOOS == "windows" {
		// windows specific code here...
		images, capture, err = winScreenshot()

	} else if runtime.GOOS == "darwin" {
		images, capture, err = macScreenshot()

	} else if runtime.GOOS == "linux" {
		//capture
	} else {
		err = errors.New("unable to determine underlying OS. Supported OS Windows, Darwin, Linux/Ubuntu")
	}

	if err == nil && capture {
		result, err := convertScreenshot(images)

		if result && err == nil {
			screenValidator, errDelete := deleteScreenshot(images)
			if screenValidator && err == nil {

			} else {
				fmt.Printf("Screenshot failed to capture, check permissions and GOOS runtime. %s\n", errDelete)
			}
		} else {
			fmt.Printf("Screenshot failed to capture, check permissions and GOOS runtime. %s", err)
		}

	} else {
		fmt.Printf("Screenshot failed to capture, check permissions and GOOS runtime. %s", err)
	}
}
