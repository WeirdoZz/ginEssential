package main

import (
	"ginEssential/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = router.CollectRouter(r)
	panic(r.Run(":9090"))
}
