package models

import (
	"time"
)

type Company struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"size:200"`
}

type User struct {
	ID           string `gorm:"primaryKey"`
	FirstName    string `gorm:"size:200"`
	LastName     string `gorm:"size:200"`
	Email        string `gorm:"size:200"`
	Password     string `gorm:"size:200"`
	Img          string `gorm:"size:200"`
	CreationDate time.Time
	CompanyID    string
	Company      Company `gorm:"foreignKey:CompanyID"`
}

type Warehouse struct {
	ID           string `gorm:"primaryKey"`
	Name         string `gorm:"size:200"`
	CreationDate time.Time
	CreatorID    string
	CompanyID    string
	Company      Company `gorm:"foreignKey:CompanyID"`
}

type Color struct {
	ID           string `gorm:"primaryKey"`
	Name         string `gorm:"size:200"`
	CreationDate time.Time
	CreatorID    string
	Color        string `gorm:"size:10"`
	CompanyID    string
	Company      Company `gorm:"foreignKey:CompanyID"`
}

type Unity struct {
	ID           string `gorm:"primaryKey"`
	Name         string `gorm:"size:200"`
	Description  string `gorm:"size:10"`
	CreationDate time.Time
	CreatorID    string
	Color        string `gorm:"size:10"`
	CompanyID    string
	Company      Company `gorm:"foreignKey:CompanyID"`
}

type Stuf struct {
	ID           string `gorm:"primaryKey"`
	Name         string `gorm:"size:200"`
	CreationDate time.Time
	CreatorID    string
	CompanyID    string
	ColorID      string
	UnityID      string
	Company      Company `gorm:"foreignKey:CompanyID"`
	Color        Color   `gorm:"foreignKey:ColorID"`
	Unity        Unity   `gorm:"foreignKey:UnityID"`
}

type StuffBatch struct {
	ID           string `gorm:"primaryKey"`
	Name         string `gorm:"size:200"`
	CreationDate time.Time
	CreatorID    string
	StufID       string
	Stuf         Stuf `gorm:"foreignKey:StufID"`
}

type WarehousesStuff struct {
	ID          string `gorm:"primaryKey"`
	WarehouseID string
	StufID      string
	Quantity    int
	BatchID     string
	Warehouse   Warehouse  `gorm:"foreignKey:WarehouseID"`
	Stuf        Stuf       `gorm:"foreignKey:StufID"`
	StuffBatch  StuffBatch `gorm:"foreignKey:BatchID"`
}
