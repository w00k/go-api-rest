package main

import (
	"fmt"
	"go-api-rest/src/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	r := gin.Default() // Create a default Gin router

	// Define a route for the root path ("/")
	r.GET("/timezone", handler.TimezoneHandler)

	// Start the server on port
	fmt.Println("Start go-api-rest in port " + PORT)
	r.Run(PORT)
}
