package structs

type User struct {
	UUID string
	Name string
}

type UserAPI struct {
	User   User
	Status APIStatus
}

type APIStatus struct {
	Status  string
	Message string
}

type Room struct {
	UUID string
	Name string
	Quiz string
}
