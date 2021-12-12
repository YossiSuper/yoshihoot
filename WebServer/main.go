package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"

	"yoshihoot_WebServer/db"
	"yoshihoot_WebServer/router"
	"yoshihoot_WebServer/uuid"
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

	fmt.Println(uuid.GenerateUUID())

	//MySQL接続
	_, err := db.SqlConnect()
	if err != nil {
		fmt.Println("DBに接続できませんでした。")
	} else {
		fmt.Println("DB接続成功")
	}

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
	e.Logger.Fatal(e.Start(":80"))
}

// ---------------------------------------------------------------------
