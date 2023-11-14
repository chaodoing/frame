package boot

import (
	`fmt`
	`io`
	`log`
	`os`
	`path`
	`strings`
	`time`
	
	`github.com/go-redis/redis`
	`github.com/lestrrat-go/strftime`
	`github.com/natefinch/lumberjack`
	`gorm.io/driver/mysql`
	`gorm.io/gorm`
	`gorm.io/gorm/logger`
)

type Container struct {
	env Env
	rdx *redis.Client
	db  *gorm.DB
}

// InitContainer 初始化容器
func InitContainer(env Env) (container Container, err error) {
	container = Container{env: env}
	container.db, err = container.MySQL()
	if err != nil {
		return
	}
	container.rdx, err = container.Redis()
	if err != nil {
		return
	}
	return
}

// MySQL 连接MySQL
func (c Container) MySQL() (db *gorm.DB, err error) {
	if c.db != nil {
		return c.db, nil
	}
	schema := c.env.MySQL.schema()
	Logger, err := c.Log(c.env.MySQL.LogFile)
	if err != nil {
		return
	}
	Interface := logger.New(Logger, logger.Config{
		Colorful: false,
		LogLevel: logger.LogLevel(c.env.MySQL.LogLevel),
	})
	c.db, err = gorm.Open(mysql.Open(schema), &gorm.Config{
		Logger:                 Interface,
		SkipDefaultTransaction: true,  // SkipDefaultTransaction 跳过默认事务
		FullSaveAssociations:   true,  // FullSaveAssociations 在创建或更新时，是否更新关联数据
		PrepareStmt:            true,  // PrepareStmt 是否禁止创建 prepared statement 并将其缓存
		AllowGlobalUpdate:      false, // AllowGlobalUpdate 是否允许全局 update/delete
		QueryFields:            true,  // QueryFields 执行查询时，是否带上所有字段
	})
	return c.db, err
}

// Redis 连接Redis
func (c Container) Redis() (rdx *redis.Client, err error) {
	if c.rdx != nil {
		return c.rdx, nil
	}
	if strings.EqualFold(c.env.Redis.Auth, "") {
		c.rdx = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", c.env.Redis.Host, c.env.Redis.Port),
			DB:   c.env.Redis.DbIndex,
		})
	} else {
		c.rdx = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", c.env.Redis.Host, c.env.Redis.Port),
			Password: c.env.Redis.Auth,
			DB:       c.env.Redis.DbIndex,
		})
	}
	var pong string
	pong, err = c.rdx.Ping().Result()
	if err != nil || !strings.EqualFold(pong, "PONG") {
		return
	}
	return c.rdx, nil
}

// Log 日志文件实例化
func (c Container) Log(file string) (Log *log.Logger, err error) {
	var write io.Writer
	write, err = c.logWrite(file)
	Log = log.New(write, "", log.LstdFlags|log.Ldate|log.Ltime)
	return
}

// logWrite 日志写入实例化
func (c Container) logWrite(file string) (w io.Writer, err error) {
	var p *strftime.Strftime
	p, err = strftime.New(path.Join(c.env.Log.Path, os.ExpandEnv(file)))
	if err != nil {
		return
	}
	var drive = &lumberjack.Logger{
		Filename:   p.FormatString(time.Now()),
		MaxSize:    5,
		MaxAge:     1,
		MaxBackups: 31,
		LocalTime:  true,
		Compress:   true,
	}
	if c.env.Log.Stdout && c.env.Log.Write {
		w = io.MultiWriter(os.Stdout, drive)
	} else if c.env.Log.Stdout {
		w = os.Stdout
	} else {
		w = drive
	}
	return w, nil
}

// Attestation 用户数据认证
func (c Container) Attestation(prefix string) (a Attestation, err error) {
	var rdx *redis.Client
	rdx, err = c.Redis()
	if err != nil {
		return
	}
	prefix = prefix + ":"
	var ttl = time.Duration(c.env.Redis.TTL) * time.Second
	return Attestation{
		rdx:        rdx,
		timeToLive: ttl,
		prefix:     prefix,
	}, nil
}
