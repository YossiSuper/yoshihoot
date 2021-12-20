//APIの実装をします。
package api

import (
	"fmt"
	"yoshihoot_WebServer/db"
	"yoshihoot_WebServer/structs"
	"yoshihoot_WebServer/uuid"
)

//LogInAPI
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
		if len(name) <= 0 {
			status = "Error"
			message = "nameは空にできません。"
			uuid = "Error"
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
					Status: "InLobby",
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
	}
	//最終的な値を代入
	UserInfo := structs.UserAPI{
		User: structs.DBUser{
			UUID:   uuid,
			Name:   name,
			Status: "InLobby",
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

//GetUserInfo
func GetUserInfo(uuid string) structs.UserAPI {
	//使用する変数を定義
	user := []structs.DBUser{}
	outUser := structs.DBUser{}
	apiStatus := structs.APIStatus{}

	//データーベース接続
	db, err := db.SqlConnect()

	//接続できなかった場合の処理
	if err != nil {
		apiStatus.Message = "ServerError:DBに接続できませんでした。"
		apiStatus.Status = "Error"
		outUser.Status = "Error"
		outUser.UUID = "Error"
		outUser.Name = "Error"
		fmt.Println(apiStatus.Message)
	} else {
		if len(uuid) <= 0 {
			apiStatus.Message = "uuidは空にできません。"
			apiStatus.Status = "Error"
			outUser.Status = "Error"
			outUser.UUID = "Error"
			outUser.Name = "Error"
		} else {
			//データベースから情報を取得
			db.Table("users").Where("uuid = ?", uuid).First(&user)

			//プレイヤーが見つからなかった場合
			if len(user) == 0 {
				apiStatus.Message = "UserNotFound"
				apiStatus.Status = "Error"
				outUser.Status = "Error"
				outUser.UUID = "Error"
				outUser.Name = "Error"
			} else {
				apiStatus.Message = "リクエストは正常に処理されました。"
				apiStatus.Status = "ok"
				outUser.Status = user[0].Status
				outUser.UUID = user[0].UUID
				outUser.Name = user[0].Name
			}
		}
	}
	//最終的な値を代入
	UserInfo := structs.UserAPI{
		User:   outUser,
		Status: apiStatus,
	}
	//データベースを閉じる
	db.Close()
	//戻り値を渡す。
	return UserInfo
}

//CreateQuestionnaireAPI
func CreateQuestionnaire() {

}
