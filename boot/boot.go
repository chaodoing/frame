package boot

import (
	`fmt`
	`net/http`
	`strings`
	`time`
	
	`github.com/gookit/goutil/envutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`github.com/kataras/iris/v12/middleware/logger`
	`github.com/kataras/iris/v12/middleware/recover`
)

type Boot struct {
	app    *iris.Application
	config iris.Configuration
	env    Env
}

// New 初始化
func New(file string, isJson bool) (boot Boot, err error) {
	var env Env
	if isJson {
		env, err = JSON(file)
		if err != nil {
			return
		}
	} else {
		env, err = XML(file)
		if err != nil {
			return
		}
	}
	env = env.parse()
	var container = Container{env: env}
	container.db, err = container.MySQL()
	if err != nil {
		panic(err)
	}
	container.rdx, err = container.Redis()
	if err != nil {
		panic(err)
	}
	hero.Register(container)
	var app = iris.New()
	w, err := container.logWrite("iris-%Y-%m-%d.log")
	if err != nil {
		return
	}
	app.Logger().SetOutput(w).SetLevel(env.Station.LogLevel)
	return Boot{
		app: app,
		env: env,
	}, nil
}

func (b Boot) IrisConfiguration(config iris.Configuration) Boot {
	b.config = config
	return b
}

func (b Boot) Handle(handles ...Handle) Boot {
	for _, handle := range handles {
		handle(b.app, b.env)
	}
	return b
}

func (b Boot) Router(routes ...Route) Boot {
	for _, route := range routes {
		route(b.app)
	}
	return b
}

func (b Boot) Strap() {
	if b.env.Station.Compression {
		b.app.Use(iris.Compression)
	}
	var env = envutil.Getenv("ENV", "development")
	b.app.UseRouter(recover.New())
	b.app.UseRouter(logger.New(logger.Config{
		Status:           true,
		IP:               true,
		Method:           true,
		Path:             true,
		PathAfterHandler: true,
		Query:            true,
		TraceRoute:       true,
		LogFunc:          nil,
		LogFuncCtx:       nil,
		Skippers:         nil,
	}))
	
	b.config.PostMaxMemory = b.env.Upload.Maximum
	b.config.DisableStartupLog = !strings.EqualFold(env, "development")
	b.config.Other = map[string]interface{}{
		"routes": b.app.GetRoutes(),
	}
	err := b.app.Run(iris.Server(&http.Server{
		Addr:              fmt.Sprintf("%s:%d", b.env.Station.Host, b.env.Station.Port),
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Second * 6,
		WriteTimeout:      time.Second * 30,
	}), iris.WithConfiguration(b.config))
	if err != nil {
		panic(err)
	}
}
