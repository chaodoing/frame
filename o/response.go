package o

import (
	`bytes`
	`encoding/json`
	`encoding/xml`
	`errors`
	`html/template`
	
	`github.com/chaodoing/frame/assets`
	`github.com/kataras/iris/v12`
)

type (
	Response struct {
		ctx  iris.Context
		data interface{}
	}
	Data struct {
		XMLName xml.Name    `xml:"root" json:"-"`
		Status  int         `json:"status" xml:"status"`
		Message string      `json:"message" xml:"message"`
		Data    interface{} `json:"data" xml:"data"`
	}
)

func (r Response) html() (html string, err error) {
	var data []byte
	data, err = assets.Asset("vscode/index.html")
	if err != nil {
		return
	}
	tpl, err := template.New("json").Parse(string(data))
	if err != nil {
		return
	}
	data, err = json.Marshal(r.data)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	theme := r.ctx.URLParamDefault("theme", "vs-dark")
	err = tpl.Execute(buf, map[string]string{
		"Title": "JSON",
		"Json":  string(data),
		"Theme": theme,
	})
	return buf.String(), err
}

func (r Response) Data(status int, message string, data interface{}) Response {
	r.data = Data{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return r
}

func (r Response) Set(data interface{}) Response {
	r.data = data
	return r
}

func (r Response) Write() (err error) {
	if r.data == nil {
		return errors.New("数据内容不能为空")
	}
	html, err := r.html()
	if err != nil {
		return err
	}
	r.ctx.Negotiation().JSON(r.data).XML(r.data).HTML(html).EncodingGzip()
	_, err = r.ctx.Negotiate(nil)
	return err
}

func O(ctx iris.Context, status int, message string, data interface{}) {
	var response = Response{ctx: ctx}
	var err = response.Data(status, message, data).Write()
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}

func Send(ctx iris.Context, data interface{}) {
	var response = Response{ctx: ctx}
	var err = response.Set(data).Write()
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}
