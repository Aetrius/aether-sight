package aetherImageCapture

import (
	"bytes"
	"net/http"
	"os"

	//"encoding/base64"
	"fmt"

	"github.com/joho/godotenv"
)

func postImage(base64In []byte) {
	godotenv.Load(".env")
	//add check if the values are present
	API_IP := os.Getenv("API_IP")
	API_PORT := os.Getenv("API_PORT")
	API_VERSION := os.Getenv("API_VERSION")
	HTTPS_ENABLED := os.Getenv("HTTPS_ENABLED")
	HOST_NAME := os.Getenv("HOST_NAME")
	FULL_PATH := os.Getenv("PATH")
	var http_method string

	fmt.Printf("Full path: %s", FULL_PATH)

	base64String := base64In

	// Create a request body from the base64 string
	requestBody := []byte(base64String)

	if HTTPS_ENABLED == "false" {
		http_method = "http"
	} else {
		http_method = "https"
	}

	// URL to forward the request to
	url := fmt.Sprintf("%s://%s:%s/%s/image/host=%s", http_method, API_IP, API_PORT, API_VERSION, HOST_NAME)

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
