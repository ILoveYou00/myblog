package api

import (
	"github.com/ILoveYou00/myblog/global"
	"github.com/ILoveYou00/myblog/models"
	"github.com/ILoveYou00/myblog/pkg/jwt"
	"github.com/ILoveYou00/myblog/util/app"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// GetAuth
// @Summary Get Auth
// @Product json
// @Tags 登录
// @description 进行身份验证
// @Param p query models.ParamsAuth true "用户名和密码"
// @Failure 400 {object} app.ResponseData
// @Success 200 {object} app.ResponseData
// @Router /api/v1/auth [get]
func GetAuth(c *gin.Context) {
	//1.解析参数
	var p models.ParamsAuth
	if err := c.ShouldBindQuery(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.INVALID_PARAMS)
			return
		}
		app.ResponseErrorWithMsg(c, app.INVALID_PARAMS, errs.Translate(global.Trans))
		return
	}
	//2.处理业务
	err := models.CheckAuth(&p)
	if err != nil && err == gorm.ErrRecordNotFound {
		app.ResponseError(c, app.ERROR_AUTH)
		return
	}
	token, err := jwt.GenToken(p.Username, p.Password)
	//3.返回响应
	if err != nil {
		app.ResponseError(c, app.ERROR_AUTH_TOKEN)
		return
	}
	app.ResponseSuccess(c, token)
}
