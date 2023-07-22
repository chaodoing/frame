package command

import (
	`bytes`
	`errors`
	`html/template`
	`os`
	`strings`
	
	`github.com/chaodoing/frame/assets`
	`github.com/gookit/goutil/fsutil`
	`github.com/manifoldco/promptui`
	`github.com/urfave/cli`
)

type FAT struct {
	Env         string
	Group       string
	Version     string
	Username    string
	Directory   string
	Description string
	Execution   string
}

var (
	// Systemd 服务脚本
	Systemd = cli.Command{
		Name:        "systemd",
		Usage:       "生成Linux服务脚本",
		Description: "生成Linux [.service] 格式服务脚本",
		Category:    "Frame",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "export",
				Usage:    "服务脚本输出位置",
				Required: false,
			},
		},
		Action: func(c *cli.Context) (err error) {
			var (
				Description = promptui.Prompt{
					Label: "[Description]请输入应用描述",
					Validate: func(input string) error {
						if strings.EqualFold(input, "") {
							return errors.New("应用描述不能为空")
						}
						return nil
					},
				}
				Username = promptui.Prompt{
					Label: "[Users]请输入运行用户",
					Validate: func(input string) error {
						if strings.EqualFold(input, "") {
							return errors.New("运行用户不能为空")
						}
						return nil
					},
				}
				Group = promptui.Prompt{
					Label: "[Group]请输入运行用户组",
					Validate: func(input string) error {
						if strings.EqualFold(input, "") {
							return errors.New("运行用户组不能为空")
						}
						return nil
					},
				}
				Directory = promptui.Prompt{
					Label:   "[DIR]请输入工作目录",
					Default: os.ExpandEnv("${DIR}"),
					Validate: func(input string) error {
						if strings.EqualFold(input, "") {
							return errors.New("工作目录不能为空")
						}
						return nil
					},
				}
				Env = promptui.Prompt{
					Label: "[Env]请输入运行环境",
					Validate: func(input string) error {
						if strings.EqualFold(input, "") {
							return errors.New("运行环境不能为空")
						}
						return nil
					},
				}
				Version = promptui.Prompt{
					Label: "[Version]请输入当前版本",
					Validate: func(input string) error {
						if strings.EqualFold(input, "") {
							return errors.New("当前版本不能为空")
						}
						return nil
					},
				}
				Execution = promptui.Prompt{
					Label:   "[Execution]请输入执行命令",
					Default: os.ExpandEnv(os.ExpandEnv("${BIN}")),
					Validate: func(input string) error {
						if strings.EqualFold(input, "") {
							return errors.New("执行命令不能为空")
						}
						return nil
					},
				}
			)
			description, err := Description.Run()
			if err != nil {
				return
			}
			username, err := Username.Run()
			if err != nil {
				return
			}
			group, err := Group.Run()
			if err != nil {
				return
			}
			directory, err := Directory.Run()
			if err != nil {
				return
			}
			env, err := Env.Run()
			if err != nil {
				return
			}
			version, err := Version.Run()
			if err != nil {
				return
			}
			execution, err := Execution.Run()
			if err != nil {
				return
			}
			var fat = FAT{
				Env:         env,
				Group:       group,
				Version:     version,
				Username:    username,
				Directory:   directory,
				Description: description,
				Execution:   execution,
			}
			text, err := assets.Asset("service/systemd.service")
			if err != nil {
				return
			}
			tpl, err := template.New("systemd").Parse(string(text))
			if err != nil {
				return
			}
			buf := new(bytes.Buffer)
			err = tpl.Execute(buf, fat)
			if err != nil {
				return err
			}
			export := c.String("export")
			if strings.EqualFold(export, "") {
				_, err = os.Stdout.Write(buf.Bytes())
				if err != nil {
					panic(err)
				}
			} else {
				_, err = fsutil.PutContents(os.ExpandEnv(export), buf.String())
				if err != nil {
					panic(err)
				}
			}
			return
		},
	}
)
