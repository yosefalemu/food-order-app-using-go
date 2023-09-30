package controllers

import (
	"fmt"
	"golang-food-store-app/intialzers"
	"golang-food-store-app/models"
	"net/http"
	"time"

	"golang-food-store-app/utils"

	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context) {
	var body struct {
		First_Name  string
		Middle_Name string
		Last_Name   string
		Email       string
		Phonenumber string
		Password    string
		Avatar      string
		IsAdmin     bool
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   *time.Time
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to bind the user register"})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}
	if !utils.IsValidEmail(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is not valid"})
		return
	}
	user := models.User{
		First_Name:  body.First_Name,
		Middle_Name: body.Middle_Name,
		Last_Name:   body.Last_Name,
		Email:       body.Email,
		Phonenumber: body.Phonenumber,
		Password:    body.Password,
		Avatar:      body.Avatar,
		IsAdmin:     body.IsAdmin,
	}
	result := intialzers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)

}
func LoginUser(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to bind the user login"})
		return
	}
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}
	var user models.User
	result := intialzers.DB.Find(&user, "Email = ?", body.Email)

	fmt.Println("get all users out")
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Check if no records were found
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
func GetUsers(c *gin.Context) {
	fmt.Println("get all users")

}
func GetSingleUser(c *gin.Context) {
	fmt.Println("get single user")
}
func GetUserProfile(c *gin.Context) {
	fmt.Println("get user profile")
}
func UpdateUser(c *gin.Context) {
	fmt.Println("update user")
}
func ChangePassword(c *gin.Context) {
	fmt.Println("change password")
}
func ResetPassword(c *gin.Context) {
	fmt.Println("reset password")
}
func DeleteUser(c *gin.Context) {
	fmt.Println("delete user")
}
