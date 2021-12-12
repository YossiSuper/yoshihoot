package api

import (
	"yoshihoot_WebServer/structs"
	"yoshihoot_WebServer/uuid"
)

func LogIn(name string) structs.UserInfo {

	userInfo := new(structs.UserInfo)
	userInfo.Name = name
	userInfo.UUID = uuid.GenerateUUID()
	return *userInfo
}
