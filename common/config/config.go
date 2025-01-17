package config

import (
	"gopkg.in/ini.v1"
)

type config struct {
	Server server `ini:"server"`
	Db     db     `ini:"db"`
	Log    log    `ini:"log"`
}

// 服务端口配置
type server struct {
	Host  string `ini:"host"`
	Port  string `ini:"port"`
	Model string `ini:"model"`
}

// 数据库配置
type db struct {
	Driver   string `ini:"driver"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Charset  string `ini:"charset"`
	MaxIde   int    `ini:"maxIdle"`
	MaxOpen  int    `ini:"maxOpen"`
}

// log日志
type log struct {
	Path  string `ini:"path"`
	Name  string `ini:"name"`
	Model string `ini:"model"`
}

var Config config

// 配置初始化
func init() {
	// 加载配置文件
	iniFile, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	// 绑定值
	err = iniFile.MapTo(&Config)
	if err != nil {
		panic(err)
	}
}
