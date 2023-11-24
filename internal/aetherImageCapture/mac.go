package aetherImageCapture

import (
	"fmt"
	"os/exec"
)

func macScreenshot() (bool, error) {
	var err error
	cmd := exec.Command("screencapture", "-x", "screen_capture.png")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to capture screen:", err)
		return false, err
	}
	
	return true, err
}
