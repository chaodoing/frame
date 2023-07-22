package main

import (
	`os`
	`time`
	
	`github.com/chaodoing/frame/boot`
	`github.com/chaodoing/frame/models`
	`github.com/chaodoing/frame/o`
	`github.com/gookit/goutil/envutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
)

var (
	ENV     = "development"
	APP     = "app"
	VERSION = "v1.0.0"
)

func handle(ctx iris.Context, container boot.Container) {
	rdx, err := container.Redis()
	if err != nil {
		o.O(ctx, 0, "Ok", nil)
		return
	}
	_, err = rdx.Set("EMAIL", "chaodoing@live.com", 648000*time.Second).Result()
	if err != nil {
		o.O(ctx, 1, err.Error(), nil)
		return
	}
	var data models.Admin
	db, err := container.MySQL()
	if err != nil {
		o.O(ctx, 3306, err.Error(), nil)
		return
	}
	err = db.Table("admin").Find(&data).Error
	if err != nil {
		o.O(ctx, 1, err.Error(), nil)
		return
	}
	o.O(ctx, 0, "Ok", data)
}

func main() {
	envutil.SetEnvMap(map[string]string{
		"APP":     envutil.Getenv("APP", APP),
		"ENV":     envutil.Getenv("ENV", ENV),
		"DIR":     envutil.Getenv("DIR", os.ExpandEnv("${PWD}")),
		"VERSION": envutil.Getenv("VERSION", VERSION),
		"BIN":     envutil.Getenv("BIN", os.ExpandEnv("${PWD}/bin")),
		"CONFIG":  envutil.Getenv("CONFIG", os.ExpandEnv("${PWD}/config/app.xml")),
	})
	var config = os.ExpandEnv("${CONFIG}")
	bt, err := boot.New(config, false)
	if err != nil {
		panic(err)
		return
	}
	bt.Router(func(app *iris.Application) {
		app.Get(`/`, hero.Handler(handle)).Name = "测试"
	}).Strap()
}
