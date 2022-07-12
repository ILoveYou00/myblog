package routers

import (
	"github.com/ILoveYou00/myblog/api"
	v1 "github.com/ILoveYou00/myblog/api/v1"
	_ "github.com/ILoveYou00/myblog/docs"
	"github.com/ILoveYou00/myblog/middleware/jwt"
	"github.com/ILoveYou00/myblog/pkg/logging"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(logging.GinLogger())

	r.Use(logging.GinRecovery(true))

	//gin.SetMode(config.RunMode)

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	//使用身份验证中间件
	apiv1.Use(jwt.JWTAuthMiddleware())
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
	{
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	//自定义404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": 404,
		})
	})

	return r
}
