package main

import (
	"ginEssential/common"
	"ginEssential/router"
	"github.com/gin-gonic/gin"
)

func main() {
	err := InitConfig()
	common.InitDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r = router.CollectRouter(r)
	panic(r.Run(":" + common.ServerSetting.Port))
}

func InitConfig() error {
	setting, err := common.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &common.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("DBSource", &common.DBSetting)
	if err != nil {
		return err
	}
	return nil
}
