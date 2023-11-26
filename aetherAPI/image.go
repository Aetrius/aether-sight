package aetherAPI

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func setupImageRoutes(images *gin.RouterGroup) {
	images.PUT("/", func(c *gin.Context) {
		setImage(c)
	})

}

func setImage(c *gin.Context) {
	// Read the request body as a byte array
	debugFlag := c.DefaultQuery("debug", "false")

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	// Convert the byte array to a base64 string
	base64String := base64.StdEncoding.EncodeToString(bodyBytes)
	storeToRedis(base64String)

	if strings.ToLower(debugFlag) == "true" {
		fmt.Println("Request Body as Base64:", base64String)

	} else {
		//fmt.Println("")

	}

	c.JSON(http.StatusOK, gin.H{"message": "Set Image Complete"})
}
