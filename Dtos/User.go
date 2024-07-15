package dtos

type User struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Img       string `json:"img"`
}

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"email"`
}

type UserChangerInfo struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
