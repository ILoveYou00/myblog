package config

import (
	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"time"
)

/*var (
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

	JwtSecret string
	Header    string

	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
)*/

type App struct {
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Jwt struct {
	JwtSecret string
	Header    string
}

var JwtSetting = &Jwt{}

type Log struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

var LogSetting = &Log{}

//初始化配置文件
/*
func init() {
	file, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase(file)
	LoadApp(file)
	LoadServer(file)
	LoadDatabase(file)
	LodeJwt(file)
	LodeLog(file)

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
func LodeJwt(file *ini.File) {
	JwtSecret = file.Section("jwt").Key("JWT_SECRET").String()
	Header = file.Section("jwt").Key("HEADER").String()
}

func LodeLog(file *ini.File) {
	Level = file.Section("log").Key("LEVEL").String()
	Filename = file.Section("log").Key("FILENAME").String()
	MaxSize, _ = file.Section("log").Key("MAX_SIZE").Int()
	MaxAge, _ = file.Section("log").Key("MAX_AGE").Int()
	MaxBackups, _ = file.Section("log").Key("MAX_BACKUPS").Int()
}
*/

func Init() {
	Cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		zap.L().Fatal("Fail to parse 'conf/app.ini'")
	}
	err = Cfg.Section("app").MapTo(AppSetting)

	err = Cfg.Section("server").MapTo(ServerSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)

	err = Cfg.Section("jwt").MapTo(JwtSetting)

	err = Cfg.Section("log").MapTo(LogSetting)

	if err != nil {
		zap.L().Fatal("initialization failed")
		return
	}
}
