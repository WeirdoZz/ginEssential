package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name        string `gorm:"varchar(20);not null"`
	PhoneNumber string `gorm:"type:varchar(11);not null;unique"`
	Password    string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		// get params
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

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
			name = RandomString(10)
		}
		fmt.Println(name, telephone, password)
		// judge if phone number exist
		if isPhoneNumberExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "phone number already exists"})
			return
		}
		// create user account
		newUser := &User{
			Name:        name,
			PhoneNumber: telephone,
			Password:    password,
		}
		db.Create(&newUser)
		//return results
		c.JSON(200, gin.H{
			"message": "register successfully",
		})
	})
	panic(r.Run(":9090"))
}

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gin_essential"
	username := "root"
	password := "123"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: driverName,
		DSN:        args,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed  to connect to database" + err.Error())
	}
	//db.AutoMigrate(&User{}).Error()
	return db
}

func isPhoneNumberExist(db *gorm.DB, phoneNumber string) bool {
	var user User
	db.Where("phone_number=?", phoneNumber).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
