package api

import (
	"fmt"
	"yoshihoot_WebServer/db"
	"yoshihoot_WebServer/structs"
	"yoshihoot_WebServer/uuid"
)

//ログインAPI
func LogIn(name string) structs.UserAPI {

	//使用する変数を定義
	var message string
	var status string
	uuid := uuid.GenerateUUID()

	//データーベース接続
	db, err := db.SqlConnect()

	//接続できなかった場合の処理
	if err != nil {
		message = "ServerError:DBに接続できませんでした。"
		status = "Error"
		fmt.Println(message)
	} else {
		//データベースからユーザーを取得
		user := []structs.User{}
		db.Where("name = ?", name).First(&user)
		//使用されていない名前の場合
		if len(user) == 0 {
			//ユーザー情報
			newUser := structs.DBUser{
				UUID:   uuid,
				Name:   name,
				Status: "Lobby",
			}
			//データベースにユーザー情報を追加
			db.Table("users").Create(&newUser)
			message = "リクエストは正常に処理されました。"
			status = "ok"
		} else {
			//すでに使用されている名前の場合
			status = "Error"
			message = "この名前はすでに使われています。"
			uuid = "None"
		}
	}

	//最終的な値を代入
	UserInfo := structs.UserAPI{
		User: structs.User{
			UUID: uuid,
			Name: name,
		},
		Status: structs.APIStatus{
			Status:  status,
			Message: message,
		},
	}
	//データベースを閉じる
	db.Close()
	//リターン
	return UserInfo
}
