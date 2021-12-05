// ---------------------------------------------------------------------
/*

   server.go

                       Jun/11/2018
*/
// ---------------------------------------------------------------------
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

// ---------------------------------------------------------------------
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// ---------------------------------------------------------------------
// サイトで共通情報
type ServiceInfo struct {
	Title string
}

var serviceInfo = ServiceInfo{
	"サイトのタイトル",
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

func GenerateUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return u.String()
}

// ---------------------------------------------------------------------
func main() {
	_, err := sqlConnect()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	fmt.Println(GenerateUUID())

	e := echo.New()

	e.Renderer = t

	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		data := struct {
			CountOfPlayers int
			CountOfRooms   int
			UserID         string
		}{
			CountOfPlayers: 20,
			CountOfRooms:   2,
			UserID:         GenerateUUID(),
		}
		return c.Render(http.StatusOK, "index", data)
	})

	e.GET("/page1", func(c echo.Context) error {
		// テンプレートに渡す値

		data := struct {
			ServiceInfo
			Content_a string
			Content_b string
			Content_c string
			Content_d string
		}{
			ServiceInfo: serviceInfo,
			Content_a:   "雨が降っています。",
			Content_b:   "明日も雨でしょうか。",
			Content_c:   "台風が近づいています。",
			Content_d:   "Jun/11/2018",
		}
		return c.Render(http.StatusOK, "page1", data)
	})

	e.Logger.Fatal(e.Start(":80"))
}

// ---------------------------------------------------------------------
