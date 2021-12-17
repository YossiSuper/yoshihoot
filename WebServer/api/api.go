package api

import (
	"fmt"
	"yoshihoot_WebServer/db"
	"yoshihoot_WebServer/structs"
	"yoshihoot_WebServer/uuid"
)

type User struct {
	ID   int
	Name string
}

//ログインAPI
func LogIn(name string) structs.UserAPI {
	var message string
	var status string

	//データーベース接続
	db, err := db.SqlConnect()
	//エラー処理
	if err != nil {
		message = "ServerError:DBに接続できませんでした。"
		status = "Error"
		fmt.Println(message)
	} else {
		message = "リクエストは正常に処理されました。"
		status = "ok"
		fmt.Println("DB接続成功")
	}
	user := structs.User{}
	db.Where("name = ?", name).First(&user)

	fmt.Println(user)

	UserInfo := structs.UserAPI{
		User: structs.User{
			UUID: uuid.GenerateUUID(),
			Name: name,
		},
		Status: structs.APIStatus{
			Status:  status,
			Message: status,
		},
	}
	return UserInfo
}
