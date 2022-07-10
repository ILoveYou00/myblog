package jwt

import (
	"errors"
	"github.com/ILoveYou00/myblog/config"
	"github.com/dgrijalva/jwt-go"
)

type MyClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//GenToken 生成jwt
func GenToken(Username string, Password string) (string, error) {
	//实例化结构体对象
	c := &MyClaim{
		Username,
		Password,
		jwt.StandardClaims{
			Issuer: "my-project",
		},
	}
	//使用指定的方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获取完整的编码后的字符串token
	return token.SignedString([]byte(config.JwtSecret))

}

//ParseToken 解析token
func ParseToken(tokenString string) (*MyClaim, error) {
	//解析token
	var mc = new(MyClaim)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	//校验token
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
