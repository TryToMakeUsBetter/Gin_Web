package settings

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error

	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':'%v'", err)
	}

	HTTPPort = sec.Key("HTTP_Port").MustInt(80)
	ReadTimeout = time.Duration(sec.Key("Read_Timeout").MustInt(60) * int(time.Second))
	WriteTimeout = time.Duration(sec.Key("Write_Timeout").MustInt(60) * int(time.Second))
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app':'%v'", err)

	}

	JwtSecret = sec.Key("JWT_Secret").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_Size").MustInt(10)
}
