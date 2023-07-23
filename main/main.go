package main

import (
	`fmt`
	`os`
	
	`github.com/chaodoing/frame/command`
	`github.com/gookit/goutil/envutil`
	`github.com/urfave/cli`
)

var (
	ENV     = "production"
	APP     = "app"
	VERSION = "v1.0.0"
)

func main() {
	envutil.SetEnvMap(map[string]string{
		"APP":     envutil.Getenv("APP", APP),
		"DIR":     envutil.Getenv("DIR", os.ExpandEnv("${PWD}")),
		"ENV":     envutil.Getenv("ENV", ENV),
		"VERSION": envutil.Getenv("VERSION", VERSION),
		"CONFIG":  envutil.Getenv("CONFIG", os.ExpandEnv("${PWD}/config/app.xml")),
		"BIN":     envutil.Getenv("BIN", "${DIR}/bin/${APP} web --config=${CONFIG}"),
	})
	var app = cli.NewApp()
	app.Name = envutil.Getenv("APP", APP)
	app.Version = envutil.Getenv("VERSION", VERSION)
	app.Commands = []cli.Command{
		command.Systemd,
		command.Env,
		command.Config,
	}
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
