package setting

import (
	"github.com/go-ini/ini"
	"github.com/labstack/gommon/log"
	"time"
)

type App struct {
	JwtSecret       string
	JwtIssuer       string
	PageSize        int
	RuntimeRootPath string

	PrefixUrl string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	QrCodeSavePath string
	FontSavePath   string
}

type Server struct {
	RunModel     string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Driver      string
	User        string
	Password    string
	Host        string
	Port        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int           // 最大空闲连接数
	MaxActive   int           // 给定时间内，允许分配的最大连接数
	IdleTimeout time.Duration // 给定时间内将保持空闲状态，超过时间关闭链接, 为0时，无限制
}

var AppSetting = &App{}
var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var RedisSetting = &Redis{}

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("failed to load conf/app.ini: %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("MapTo AppSetting err : %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("MapTo RedisSetting err: %v", err)
	}
}
