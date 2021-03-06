package application

import (
	router "seifwu/config/routers"
	global "seifwu/global"

	"github.com/spf13/viper"
)

// Run 运行服务
func Run() {
	r := router.Routers()
	defer global.DB.Close()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 默认监听并在 0.0.0.0:8080 上启动服务
}
