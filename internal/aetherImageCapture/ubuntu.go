package aetherImageCapture

import (
	"fmt"
	"os/exec"
)

// needs to implement multiple screens if possible
func linuxScreenshot() ([]string, bool, error) {
	var err error
	var images []string // parse through multiple screens if required

	cmd := exec.Command("/usr/bin/gnome-screenshot", "-f", "screen_capture.png")

	if err = cmd.Run(); err != nil {
		fmt.Println("you need to make sure you have gnome-screenshot installed...")
		return images, false, err
	}

	return images, true, err
}
