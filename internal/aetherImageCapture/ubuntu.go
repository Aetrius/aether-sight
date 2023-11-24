package aetherImageCapture

import (
	"os/exec"
)

func linuxScreenshot() (bool, error) {
	var err error
	cmd := exec.Command("gnome-screenshot", "-f", "screen_capture.png")

	if err = cmd.Run(); err != nil {
		return false, err
	}

	return true, err
}
