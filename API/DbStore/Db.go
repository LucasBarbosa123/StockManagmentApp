package dbstore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (*gorm.DB, error) {
	connectionString := "root@tcp(localhost:3306)/stockmanagment?charset=utf8mb4&parseTime=True&loc=Local"

	conn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db = conn

	return db, nil
}

func Disconnect() {
	conn, err := db.DB()

	if err != nil {
		return
	}

	conn.Close()
}
