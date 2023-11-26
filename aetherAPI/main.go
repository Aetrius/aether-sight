package aetherAPI

import (
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

var REDIS_SERVER_IP string
var REDIS_SERVER_PORT string

func StartServer() {
	r := gin.Default()
	REDIS_SERVER_IP = os.Getenv("REDIS_SERVER_IP")
	REDIS_SERVER_PORT = os.Getenv("REDIS_SERVER_PORT")

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://localhost:8080", "localhost:8081", "localhost:8080", "*"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Host"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}))

	v1 := r.Group("/v1")

	setupImageRoutes(v1.Group("/image"))

	r.Run(":8080")
}
