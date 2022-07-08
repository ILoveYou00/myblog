package v1

import (
	"fmt"
	"github.com/ILoveYou00/myblog/config"
	"github.com/ILoveYou00/myblog/models"
	"github.com/ILoveYou00/myblog/util"
	"github.com/ILoveYou00/myblog/util/app"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	fmt.Println("=========================")
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data["lists"] = models.GetTags(util.GetPage(c), config.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	app.ResponseSuccess(c, data)
}
