package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code MyCode      `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, m MyCode) {
	rd := &ResponseData{
		Code: m,
		Msg:  m.Msg(),
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, m MyCode, errMsg interface{}) {
	rd := &ResponseData{
		Code: m,
		Msg:  errMsg,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: SUCCESS,
		Msg:  SUCCESS.Msg(),
		Data: data,
	}
	ctx.JSON(http.StatusOK, rd)
}
