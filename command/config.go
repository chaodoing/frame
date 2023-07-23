package command

import (
	`os`
	
	`github.com/chaodoing/frame/o`
	`github.com/urfave/cli`
)

var Config = cli.Command{
	Name:        "config",
	Usage:       "生成配置文件",
	Description: os.Getenv("APP") + " config --export [输出文件位置]",
	Category:    "Frame",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "file,f",
			Usage:    "配置文件输出位置",
			Required: true,
			Value:    os.Getenv("CONFIG"),
		},
	},
	Action: func(c *cli.Context) (err error) {
		return o.SaveXML(env, c.String("file"))
	},
}
