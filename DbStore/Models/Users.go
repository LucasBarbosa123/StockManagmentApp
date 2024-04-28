package models

import "time"

type User struct {
	ID           string    `gorm:"primaryKey"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	Email        string    `gorm:"column:email"`
	Img          string    `gorm:"column:img"`
	Password     string    `gorm:"column:password"`
	CreationDate time.Time `gorm:"column:creation_date"`
}
