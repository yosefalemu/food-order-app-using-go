package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	fmt.Println("in register")
}
func LoginUser(c *gin.Context) {
	fmt.Println("login user")
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
