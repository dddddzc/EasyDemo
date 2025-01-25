package main

import (
	"easydemo/common"
	"easydemo/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 读取配置
	InitConfig()

	// 初始化数据库
	common.InitDB()

	// 初始化gin框架
	r := gin.Default()

	// 注册路由
	r = router.CollectRoute(r)

	// 修改监听端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	// 没有指定端口则默认为8080
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("failed to read config file")
	}
}
