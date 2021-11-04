package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	LogFilePath string
)

// init
// @Desc: 	读取初始化文件，设置全局变量
// @Notice:
func init() {
	var err error
	Cfg, err = ini.Load("with_jianyu/go_gin_example/conf/app.ini")
	if err != nil {
		log.Fatalf("fail to parse 'with_jianyu/go_gin_example/conf/app.ini': %v", err)
	}

	LoadBase()
	LoadLogFilePath()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadLogFilePath()  {
	section, err := Cfg.GetSection("runtime")
	if err != nil {
		log.Fatalf("fail to get section 'runtime': %v", err)
	}

	LogFilePath = section.Key("LogFilePath").MustString("with_jianyu/go_gin_example/runtime/logs/")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
