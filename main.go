package main

import (
	"fmt"
	"github.com/ILoveYou00/myblog/config"
	"github.com/ILoveYou00/myblog/models"
	"github.com/ILoveYou00/myblog/pkg/logging"
	"github.com/ILoveYou00/myblog/routers"
	"net/http"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termOfService http://swagger.io/terms/
// @contact.name zxp
// @contact.url http://www.swagger.io/support
// @contact.email support@seagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8889
// @BasePath
func main() {
	//初始化配置文件
	config.Init()
	//初始化日志
	logging.Init()
	//初始化数据库
	models.Init()
	//关闭数据库连接
	defer models.Close()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    config.ServerSetting.ReadTimeout,
		WriteTimeout:   config.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
