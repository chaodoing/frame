package command

import (
	`os`
	
	`github.com/chaodoing/frame/boot`
	`github.com/urfave/cli`
)

var env = boot.Env{
	Station: boot.Station{
		Host:        "0.0.0.0",
		Port:        8888,
		Compression: true,
		LogLevel:    "debug",
	},
	MySQL: boot.MySQL{
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "platform",
		Username: "root",
		Password: "123.com",
		Charset:  "utf8mb4",
		LogLevel: 4,
		LogFile:  "gorm-%Y-%m-%d.log",
	},
	Redis: boot.Redis{
		Host:    "127.0.0.1",
		Port:    6379,
		DbIndex: 1,
		Auth:    "123.com",
		TTL:     648000,
	},
	Upload: boot.Path{
		Maximum: 20,
		URL:     "/upload",
		PATH:    "${DIR}/upload",
	},
	Log: boot.Log{
		Stdout: true,
		Write:  false,
		Path:   "${DIR}/logs",
	},
}
var Env = cli.Command{
	Name:        "env",
	Usage:       "生成ENV环境变量",
	Description: os.ExpandEnv("APP") + " env ### 生成 ${WORKDIR}/.env 环境变量",
	Category:    "Frame",
	Action: func(c *cli.Context) (err error) {
		return env.INI(".env")
	},
}
