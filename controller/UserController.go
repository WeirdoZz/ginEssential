package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()

	// get params
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//fmt.Println(name, telephone, password)

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
		Password:    util.EncodePassword(password),
	}
	db.Create(&newUser)
	//return results
	c.JSON(200, gin.H{
		"message": "register successfully",
	})
}

func Login(ctx *gin.Context) {
	db := common.GetDB()

	phoneNumber := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	var user model.User
	db.Where("phone_number=?", phoneNumber).First(&user)

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, "account doesn't exist")
		return
	}
	if !util.ValidPassword(password, user.Password) {
		//fmt.Println(password, user.Password)
		ctx.JSON(http.StatusBadRequest, "account or password is wrong")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "generate token failed"})
		log.Printf("token generate error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login successfully",
		"data": gin.H{"token": token}})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{"user": user}})
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
