package test

import (
	`testing`
	
	`github.com/chaodoing/frame/boot`
	`github.com/chaodoing/frame/o`
)

func TestEnv(t *testing.T) {
	var env = boot.Env{
		Station: boot.Station{
			Host:        "0.0.0.0",
			Port:        8888,
			Compression: true,
		},
		MySQL: boot.MySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DbName:   "test",
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
	var err = o.SaveXML(env, "./app.xml")
	if err != nil {
		t.Error(err)
	}
	// err := env.INI("../.env")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// t.Log("Success")
}
