package v1

import (
	"github.com/ILoveYou00/myblog/config"
	"github.com/ILoveYou00/myblog/global"
	"github.com/ILoveYou00/myblog/models"
	"github.com/ILoveYou00/myblog/util"
	"github.com/ILoveYou00/myblog/util/app"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/unknwon/com"
	"log"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data["lists"] = models.GetTags(util.GetPage(c), config.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	app.ResponseSuccess(c, data)
}

//AddTag 新增文章标签
func AddTag(c *gin.Context) {
	//1.解析参数
	var p models.ParamsCreateTag
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.INVALID_PARAMS)
			return
		}
		app.ResponseErrorWithMsg(c, app.INVALID_PARAMS, errs.Translate(global.Trans))
		return
	}
	p.State = com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	//2.业务处理
	//判断文章标签是否已经存在
	if !models.ExistTagByName(p.Name) {
		tag := &models.Tag{
			Name:      p.Name,
			CreatedBy: p.CreatedBy,
			State:     p.State,
		}
		err := models.AddTag(tag)
		if err != nil {
			log.Println(err)
		}
		app.ResponseSuccess(c, app.SUCCESS)
	} else {
		app.ResponseError(c, app.ERROR_EXIST_TAG)
	}
}

//EditTag 修改文章标签
func EditTag(c *gin.Context) {
	//1.解析参数
	var p models.ParamsUpdateTag
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.INVALID_PARAMS)
			return
		}
		app.ResponseErrorWithMsg(c, app.INVALID_PARAMS, errs.Translate(global.Trans))
		return
	}
	//2.业务处理
	//判断文章标签是否已经存在
	if !models.ExistTagByName(p.Name) {
		err := models.EditTag(&p)
		if err != nil {
			app.ResponseError(c, app.INVALID_PARAMS)
		}
		app.ResponseSuccess(c, app.SUCCESS)
	} else {
		app.ResponseError(c, app.ERROR_EXIST_TAG)
	}
}

//DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	//解析参数
	id := com.StrTo(c.Param("id")).MustInt()
	err := models.DeleteTag(id)
	if err != nil {
		app.ResponseError(c, app.INVALID_PARAMS)
	}
	app.ResponseSuccess(c, app.SUCCESS)
}
