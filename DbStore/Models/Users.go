package models

import "time"

type Company struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type User struct {
	ID           string    `gorm:"primaryKey"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	Email        string    `gorm:"column:email"`
	Img          string    `gorm:"column:img"`
	Password     string    `gorm:"column:password"`
	CreationDate time.Time `gorm:"column:creation_date"`
	CompanyID    string    `gorm:"column:company_id"`
	Company      Company   `gorm:"foreignKey:CompanyID"`
}

type Warehouse struct {
	ID           string    `gorm:"primaryKey"`
	Name         string    `gorm:"name"`
	CreationDate time.Time `gorm:"column:creation_date"`
	CreatorID    string    `gorm:"column:creator_id"`
	CompanyID    string    `gorm:"column:company_id"`
	Company      Company   `gorm:"foreignKey:CompanyID"`
}

type Stuff struct {
	ID           string    `gorm:"primaryKey"`
	Name         string    `gorm:"name"`
	Unit         string    `gorm:"unitary_cost"`
	CreationDate time.Time `gorm:"column:creation_date"`
	CreatorID    string    `gorm:"column:creator_id"`
	UnitaryCost  float64   `gorm:"column:unitary_cost"`
	CompanyID    string    `gorm:"column:company_id"`
	Company      Company   `gorm:"foreignKey:CompanyID"`
}

type WarehouseStuff struct {
	ID          string `gorm:"primaryKey"`
	WarehouseID string `gorm:"column:wear_house_id"`
	StuffID     string `gorm:"column:stuf_id"`
	Quantity    int    `gorm:"quantity"`
}
