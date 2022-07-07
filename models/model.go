package models

import (
	"fmt"
	"github.com/ILoveYou00/myblog/config"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

//初始化数据库
func init() {

	db, err := gorm.Open(config.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Name))

	if err != nil {
		log.Println(err)
	}

	//表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
