package aetherImageCapture

 import (
 	"fmt"
	"runtime"
	"errors"
 )

 func captureScreen() {
	var capture bool = false
	var err error

	if runtime.GOOS == "windows" {
		// windows specific code here...

	} else if runtime.GOOS == "darwin" {
		capture, err = macScreenshot();
		
	} else if runtime.GOOS == "linux" {
		//capture
	} else {
		err = errors.New("Unable to determine underlying OS. Supported OS Windows, Darwin, Linux/Ubuntu") 
	}

	if (err == nil && capture == true) {
		result, err := convertScreenshot();

		if (result == true && err == nil) {
			screenValidator, errDelete := deleteScreenshot();
			if (screenValidator == true && err == nil) {

			} else {
				fmt.Println(fmt.Sprintf("Screenshot failed to capture, check permissions and GOOS runtime. %s", errDelete))
			}
		} else {
			fmt.Println(fmt.Sprintf("Screenshot failed to capture, check permissions and GOOS runtime. %s", err))
		}
		
	} else {
		fmt.Println(fmt.Sprintf("Screenshot failed to capture, check permissions and GOOS runtime. %s", err))
	}
}