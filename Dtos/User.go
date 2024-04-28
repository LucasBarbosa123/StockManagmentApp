package dtos

type User struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Img       string `json:"img"`
}
