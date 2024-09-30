package dtos

type ColorInfo struct {
	Name string `validate:"required"`
	Ref  string `validate:"required"`
}
