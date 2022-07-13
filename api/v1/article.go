package v1

import (
	"fmt"
	"github.com/ILoveYou00/myblog/config"
	"github.com/ILoveYou00/myblog/global"
	"github.com/ILoveYou00/myblog/models"
	"github.com/ILoveYou00/myblog/util"
	"github.com/ILoveYou00/myblog/util/app"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/unknwon/com"
)

//GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	//1.解析参数
	id := com.StrTo(c.Param("id")).MustInt()
	//2.业务处理
	//判断id是否存在
	if models.ExistArticleByID(id) {
		data, err := models.GetArticle(id)
		if err != nil {
			app.ResponseError(c, app.INVALID_PARAMS)
			return
		}
		app.ResponseSuccess(c, data)
	} else {
		app.ResponseError(c, app.ERROR_NOT_EXIST_ARTICLE)
	}

}

//GetArticles 获取多个文章
func GetArticles(c *gin.Context) {
	//2.业务处理
	var err error
	var p models.ParamsGetArticle
	p.Article, err = models.GetArticles(util.GetPage(c), config.AppSetting.PageSize)
	p.Total, err = models.GetArticleTotal()
	if err != nil {
		app.ResponseError(c, app.INVALID_PARAMS)
	}
	app.ResponseSuccess(c, p)
}

//AddArticle 新增文章
func AddArticle(c *gin.Context) {
	//1.解析参数
	var p models.ParamsCreateArticle
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			app.ResponseError(c, app.INVALID_PARAMS)
			return
		}
		app.ResponseErrorWithMsg(c, app.INVALID_PARAMS, errs.Translate(global.Trans))
		return
	}
	fmt.Println(p)
	p.State = com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	//2.业务处理
	article := &models.Article{
		TagID:     p.TagId,
		Title:     p.Title,
		Desc:      p.Desc,
		Content:   p.Content,
		CreatedBy: p.CreatedBy,
		State:     p.State,
	}
	err := models.CreateArticle(article)
	//3.返回响应
	if err != nil {
		app.ResponseError(c, app.INVALID_PARAMS)
		return
	}
	app.ResponseSuccess(c, nil)
}

//EditArticle  修改文章
func EditArticle(c *gin.Context) {
	//1.解析参数
	var p models.ParamsUpdateArticle
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
	err := models.EditArticle(&p)
	if err != nil {
		app.ResponseError(c, app.INVALID_PARAMS)
		return
	}
	app.ResponseSuccess(c, nil)
}

//DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	//解析参数
	id := com.StrTo(c.Param("id")).MustInt()
	err := models.DeleteArticle(id)
	if err != nil {
		app.ResponseError(c, app.INVALID_PARAMS)
	}
	app.ResponseSuccess(c, nil)
}
