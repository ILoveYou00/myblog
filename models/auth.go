package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func CheckAuth(p *ParamsAuth) error {
	return DB.Select("id").Where(Auth{Username: p.Username, Password: p.Password}).First(&Auth{}).Error
}
