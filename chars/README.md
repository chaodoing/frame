# 字符串处理函数

```go
import (
    "github.com/chaodoing/frame/chars"
)

func main() {
    fmt.Println(chars.Snake("developersMenu", '-')) // 输出 developers-menu
    fmt.Println(chars.Camel("create-at", '-')) // 输出 CreateAt
}
```