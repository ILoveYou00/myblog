package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	RunMode string

	PageSize int

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
)

//初始化配置文件
func init() {
	file, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase(file)
	LoadApp(file)
	LoadServer(file)
	LoadDatabase(file)

}
func LoadBase(file *ini.File) {
	RunMode = file.Section("").Key("RUN_MODE").String()
}

func LoadApp(file *ini.File) {
	PageSize, _ = file.Section("app").Key("PAGE_SIZE").Int()
}

func LoadServer(file *ini.File) {
	HTTPPort, _ = file.Section("server").Key("HTTP_PORT").Int()
	ReadTimeout, _ = file.Section("server").Key("READ_TIMEOUT").Duration()
	WriteTimeout, _ = file.Section("server").Key("WRITE_TIMEOUT").Duration()
}

func LoadDatabase(file *ini.File) {
	Type = file.Section("database").Key("TYPE").String()
	User = file.Section("database").Key("USER").String()
	Password = file.Section("database").Key("PASSWORD").String()
	Host = file.Section("database").Key("HOST").String()
	Name = file.Section("database").Key("NAME").String()
	TablePrefix = file.Section("database").Key("TABLE_PREFIX").String()
}
