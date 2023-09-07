package boot

import (
	`encoding/json`
	`encoding/xml`
	`fmt`
	`os`
	
	`github.com/go-ini/ini`
	`github.com/gookit/goutil/fsutil`
)

type (
	// Station 站点
	Station struct {
		Host        string `json:"host" xml:"host" ini:"HOST" comment:"监听主机"`
		Port        uint16 `json:"port" xml:"port" ini:"PORT" comment:"监听端口"`
		Compression bool   `json:"compression" xml:"compression" ini:"COMPRESSION" comment:"压缩输出"`
		LogLevel    string `json:"logLevel" xml:"logLevel" ini:"LOGLEVEL" comment:"日志等级[disable fatal error warn info debug]"`
	}
	// MySQL 数据库
	MySQL struct {
		Host     string `json:"host" xml:"host" ini:"HOST" comment:"连接主机"`
		Port     uint16 `json:"port" xml:"port" ini:"PORT" comment:"连接端口"`
		DbName   string `json:"dbName" xml:"dbName" ini:"DBNAME" comment:"数据库名称"`
		Username string `json:"username" xml:"username" ini:"USERNAME" comment:"数据库用户名"`
		Password string `json:"password" xml:"password" ini:"PASSWORD" comment:"数据库密码"`
		Charset  string `json:"charset" xml:"charset" ini:"CHARSET" comment:"数据库字符集"`
		LogLevel int    `json:"log_level" xml:"logLevel" ini:"LOGLEVEL" comment:"日志等级 Silent:1 Error:2 Warn:3 Info:4"`
		LogFile  string `json:"log_file" xml:"logFile" ini:"LOGFILE" comment:"日志文件"`
	}
	// Redis 存储
	Redis struct {
		Host    string `json:"host" xml:"host" ini:"HOST" comment:"连接主机"`
		Port    uint16 `json:"port" xml:"port" ini:"PORT" comment:"连接端口"`
		DbIndex int    `json:"dbIndex" xml:"dbIndex" ini:"DBINDEX" comment:"数据库"`
		Auth    string `json:"auth" xml:"auth" ini:"AUTH" comment:"连接密码"`
		TTL     int64  `json:"ttl" xml:"ttl" ini:"TTL" comment:"默认连接时长"`
	}
	// Path 路径
	Path struct {
		Maximum int64  `json:"maximum" xml:"maximum" ini:"MAXIMUM" comment:"允许上传最大大小MB"`
		URL     string `json:"url" xml:"url" ini:"URL" comment:"访问路径"`
		PATH    string `json:"path" xml:"path" ini:"PATH" comment:"存储路径"`
	}
	// Log 日志
	Log struct {
		Stdout bool   `json:"stdout" xml:"stdout" ini:"STDOUT" comment:"输出到控制台"`
		Write  bool   `json:"write" xml:"write" ini:"WRITE" comment:"写入到文件"`
		Path   string `json:"path" xml:"path" ini:"PATH" comment:"日志存储路径"`
	}
	// Env 配置变量
	Env struct {
		XMLName xml.Name `json:"-" xml:"root" ini:"-"`
		Station Station  `json:"station" xml:"station" ini:"STATION" comment:"站点信息"`
		MySQL   MySQL    `json:"mysql" xml:"mysql" ini:"MYSQL" comment:"MySQL配置"`
		Redis   Redis    `json:"redis" xml:"redis" ini:"REDIS" comment:"Redis配置"`
		Upload  Path     `json:"upload" xml:"upload" ini:"UPLOAD" comment:"上传配置"`
		Log     Log      `json:"log" xml:"log" ini:"LOG" comment:"日志配置"`
	}
)

// XML 加载配置文件
func XML(env string) (data Env, err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(env))
	if err != nil {
		return
	}
	err = xml.Unmarshal(content, &data)
	if err != nil {
		return
	}
	data, err = data.LoadEnv()
	if err != nil {
		return data, err
	}
	data = data.parse()
	return
}

// JSON 加载配置文件
func JSON(env string) (data Env, err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(env))
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return
	}
	data, err = data.LoadEnv()
	if err != nil {
		return
	}
	data = data.parse()
	return
}

func (e Env) INI(file string) (err error) {
	var config = ini.Empty()
	err = ini.ReflectFrom(config, &e)
	if err != nil {
		fmt.Println("写入错误")
		return
	}
	err = config.SaveTo(os.ExpandEnv(file))
	return
}

// LoadEnv 加载环境变量
func (e Env) LoadEnv() (Env, error) {
	if fsutil.FileExist(".env") {
		config, err := ini.Load(".env")
		if err != nil {
			return Env{}, err
		}
		err = config.MapTo(&e)
		if err != nil {
			return Env{}, err
		}
	}
	return e, nil
}

func (e Env) parse() Env {
	e.Upload = Path{
		Maximum: e.Upload.Maximum << 20,
		URL:     os.ExpandEnv(e.Upload.URL),
		PATH:    os.ExpandEnv(e.Upload.PATH),
	}
	e.Log.Path = os.ExpandEnv(e.Log.Path)
	return e
}

func (s MySQL) schema() (data string) {
	data = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", s.Username, s.Password, s.Host, s.Port, s.DbName, s.Charset)
	return
}
