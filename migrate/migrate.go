package main

import (
	"golang-food-store-app/intialzers"
	"golang-food-store-app/models"
)

func init() {
	intialzers.LoadEnvVariables()
	intialzers.ConnectToDb()

}
func main() {
	intialzers.DB.AutoMigrate(&models.User{})

}
