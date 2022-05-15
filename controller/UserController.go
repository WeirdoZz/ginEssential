package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()

	// get params
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	fmt.Println(name, telephone, password)

	// validate data
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "phone number must be 11"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "password must be more than 16"})
		return
	}
	// if name is not passed,generate a random string
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// judge if phone number exist
	if isPhoneNumberExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "phone number already exists"})
		return
	}
	// create user account
	newUser := &model.User{
		Name:        name,
		PhoneNumber: telephone,
		Password:    password,
	}
	db.Create(&newUser)
	//return results
	c.JSON(200, gin.H{
		"message": "register successfully",
	})
}

func isPhoneNumberExist(db *gorm.DB, phoneNumber string) bool {
	var user model.User
	db.Where("phone_number = ?", phoneNumber).First(&user)
	fmt.Println("user :", user)
	if user.ID != 0 {
		return true
	}
	return false
}
