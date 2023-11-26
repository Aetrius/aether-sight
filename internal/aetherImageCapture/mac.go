package aetherImageCapture

import (
	"fmt"
	"os/exec"
)

func macScreenshot() ([]string, bool, error) {
	var err error
	var images []string

	// macs require you to specify a value as an X coordinate followed by the screensize.
	// cmd := exec.Command("screencapture", "-x", "-R1921,0,3840,1080", "second_screen_capture.png")
	cmd := exec.Command("screencapture", "-x", "screen_capture.png")

	images = append(images, "./screen_capture.png")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to capture screen:", err)
		return images, false, err
	}

	return images, true, err
}
