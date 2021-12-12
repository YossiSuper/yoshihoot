package api

import (
	"yoshihoot_WebServer/structs"
	"yoshihoot_WebServer/uuid"
)

func LogIn(name string) structs.UserAPI {

	//userInfo := new(structs.UserInfo)
	//userInfo.Name = name
	//userInfo.UUID = uuid.GenerateUUID()
	//return *userInfo

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
