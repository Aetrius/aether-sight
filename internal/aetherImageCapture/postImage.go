package aetherImageCapture

import (
	"bytes"
	//"encoding/base64"
	"fmt"
	"net/http"
)

func postImage(base64In []byte) {
	base64String := base64In

	// Create a request body from the base64 string
	requestBody := []byte(base64String)

	// URL to forward the request to
	url := "http://localhost:8080/v1/image/"

	// Create a POST request with the base64 string as the body
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers if needed, for example:
	// req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	fmt.Println("Response Status:", resp.Status)
}
