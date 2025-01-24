package main

import (
	"easydemo/common"
	"easydemo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = router.CollectRoute(r)
	panic(r.Run())
}
