package api

import (
	"yoshihoot_WebServer/structs"
	"yoshihoot_WebServer/uuid"
)

//ログインAPI
func LogIn(name string) structs.UserAPI {
	UserInfo := structs.UserAPI{
		User: structs.User{
			UUID: uuid.GenerateUUID(),
			Name: name,
		},
		Status: structs.APIStatus{
			Status:  "OK",
			Message: "None",
		},
	}
	return UserInfo
}
