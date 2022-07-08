package app

type MyCode int

const (
	SUCCESS        MyCode = 200
	ERROR          MyCode = 500
	INVALID_PARAMS MyCode = 400 //请求参数错误

	ERROR_EXIST_TAG         MyCode = 10001 //已存在该标签名称
	ERROR_NOT_EXIST_TAG     MyCode = 10002 //该标签不存在
	ERROR_NOT_EXIST_ARTICLE MyCode = 10003 //该文章不存在

	ERROR_AUTH_CHECK_TOKEN_FAIL    MyCode = 20001 //Token鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT MyCode = 20002 //Token已超时
	ERROR_AUTH_TOKEN               MyCode = 20003 //Token生成失败
	ERROR_AUTH                     MyCode = 20004 //Token错误
)

var MsgFlags = map[MyCode]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func (code MyCode) Msg() string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
