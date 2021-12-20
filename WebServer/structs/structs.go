package structs

type User struct {
	UUID string
	Name string
}

type DBUser struct {
	UUID   string
	Name   string
	Status string
}

type UserAPI struct {
	User   DBUser
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

type DBRoom struct {
	Count int
	Room  []Room
}

type Questionnaire struct {
	Question string
	Asnwer   string
	Owner    string
	UUID     string
}

type QuestionnaireAnswer struct {
	Index  int
	Answer string
	Count  int
}
