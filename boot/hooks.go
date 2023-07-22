package boot

import (
	`github.com/kataras/iris/v12`
)

type Hooks interface {
	// View 注册视图
	View(app *iris.Application)
	// Directory 注册目录
	Directory(app *iris.Application)
}

type Route func(app *iris.Application)
type Handle func(app *iris.Application, env Env)
