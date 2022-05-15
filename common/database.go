package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
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
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
