package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()

	var requestUser = model.User{}
	c.BindJSON(&requestUser)
	// get params
	name := requestUser.Name
	telephone := requestUser.PhoneNumber
	password := requestUser.Password

	//fmt.Println(name, telephone, password)

	// validate data
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "phone number must be 11")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "password must be more than 16")
		return
	}
	// if name is not passed,generate a random string
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// judge if phone number exist
	if isPhoneNumberExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "phone number already exists")
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
	response.Success(c, nil, "register successfully")
}

func Login(ctx *gin.Context) {
	db := common.GetDB()

	phoneNumber := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	var user model.User
	db.Where("phone_number=?", phoneNumber).First(&user)

	if user.ID == 0 {
		response.Response(ctx, http.StatusNotFound, 404, nil, "account doesn't exist")
		return
	}
	if !util.ValidPassword(password, user.Password) {
		//fmt.Println(password, user.Password)
		response.Response(ctx, http.StatusBadRequest, 400, nil, "account or password is wrong")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "generate token failed")
		log.Printf("token generate error")
		return
	}

	response.Success(ctx, gin.H{"token": token}, "login successfully")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
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
