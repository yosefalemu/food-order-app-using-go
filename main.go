package main

import (
	"golang-food-store-app/intialzers"
	"golang-food-store-app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	intialzers.LoadEnvVariables()
	intialzers.ConnectToDb()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	// Call the routes
	routes.UserRouters(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
