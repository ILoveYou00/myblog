package routers

import (
	v1 "github.com/ILoveYou00/myblog/api/v1"
	"github.com/ILoveYou00/myblog/config"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(config.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
	}

	//自定义404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": 404,
		})
	})

	return r
}
