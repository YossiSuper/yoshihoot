package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"yoshihoot_WebServer/uuid"
)

//テンプレートエンジン用
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "yoshihoot"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

//メイン
func main() {

	fmt.Println(uuid.GenerateUUID())

	//MySQL接続
	_, err := sqlConnect()

	if err != nil {
		panic(err.Error())
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
	e.Static("/", "public")

	//トップページ
	e.GET("/", func(c echo.Context) error {
		data := struct {
			CountOfPlayers int
			CountOfRooms   int
			UserID         string
		}{
			CountOfPlayers: 20,
			CountOfRooms:   2,
			UserID:         "4",
		}
		return c.Render(http.StatusOK, "index", data)
	})

	//サーバー開始
	e.Logger.Fatal(e.Start(":80"))
}

// ---------------------------------------------------------------------
