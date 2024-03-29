package models

import (
	"database/sql"
	"fmt"
	"github.com/ILoveYou00/myblog/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on" gorm:"autoCreateTime"`
	ModifiedOn time.Time `json:"modified_on" gorm:"autoUpdateTime"`
}

//Init 初始化数据库
func Init() {

	//连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Name,
	)
	db, err := sql.Open(config.DatabaseSetting.Type, dsn)
	if err != nil {
		zap.L().Fatal("mysql connect error", zap.Error(err))
		return
	}

	//最大连接数
	db.SetMaxOpenConns(10)
	//最大空闲数
	db.SetMaxIdleConns(100)

	DB, err = Open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//设置表名为单数形式
			SingularTable: true,
			//设置表名
			TablePrefix: config.DatabaseSetting.TablePrefix,
		},
		//设置日志格式
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// Open 打开数据库连接
func Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// Close 关闭连接
func Close() {
	db, err := DB.DB()
	if err != nil {
		log.Println("db close err", err)
	}
	_ = db.Close()
}
