package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DigiConvent/testd9t/version"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	fmt.Println("Version: ", version.Version)

	router.NoRoute(handleFrontend())
	router.Run(":8080")
}

func handleFrontend() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/favicon.ico" {
			path = "/favicon.ico"
		} else if strings.HasPrefix(path, "/assets/") {

		} else {
			path = "/index.html"
		}
		c.File(os.Getenv("STATIC_FILES") + "frontend/" + path)
	}
}
