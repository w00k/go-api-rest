package main

import (
	"fmt"
	"go-api-rest/src/handler"
	"log"
	"os"

	docs "go-api-rest/src/cmd/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	r := gin.Default() // Create a default Gin router
	docs.SwaggerInfo.BasePath = "/api/v1/go-api-rest-world-time"

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/go-api-rest-world-time")
		{
			// Define endpoint
			eg.GET("/timezone", handler.TimezoneHandler)
		}
	}

	// Start the server on port
	fmt.Println("Start go-api-rest in port " + PORT)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(PORT)
}
