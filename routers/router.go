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
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	//自定义404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": 404,
		})
	})

	return r
}
