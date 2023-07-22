# iris+gorm框架整合


## 要定义的常量

```go
package main

import (
	`os`
)

func main() {
	_ = os.Setenv("APP", "应用名称")
	_ = os.Setenv("DIR", "应用运行目录")
	_ = os.Setenv("ENV", "运行环境")
	_ = os.Setenv("VERSION", "版本信息")
	_ = os.Setenv("CONFIG", "配置文件")
	_ = os.Setenv("BIN", "运行命令 [${DIR}/bin/${APP} [web] --config=${CONFIG}]")
}
```