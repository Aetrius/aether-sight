package aetherImageCapture

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func convertScreenshot() (bool, error) {	
	var err error
	bytes, err := ioutil.ReadFile("./screen_capture.png")
	if err != nil {
		return false, err
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	//fmt.Println(base64Encoding)
	//writeBase64ToFile(toBase64(bytes))

	return true, err
}

func deleteScreenshot() (bool, error) {
	var err error
	err = os.Remove("./screen_capture.png") 

	if err != nil { 
		
		return false, err
	}

	return true, err
}

func writeBase64ToFile(base64In string) {
		base64String := base64In
		decodedData, err := base64.StdEncoding.DecodeString(base64String)
		if err != nil {
			fmt.Println("Error decoding base64:", err)
			return
		}

		filePath := "output.txt"
	
		err = ioutil.WriteFile(filePath, decodedData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
}