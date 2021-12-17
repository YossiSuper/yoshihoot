package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"

	"yoshihoot_WebServer/router"
)

//テンプレートエンジン用
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//メイン
func main() {
	//テンプレートエンジン設定
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	//サーバーインスタンス作成
	e := echo.New()

	//サーバーの設定
	e.Renderer = t
	router.SetRouter(e)

	//サーバー開始
	e.Logger.Fatal(e.Start(":3000"))
}
