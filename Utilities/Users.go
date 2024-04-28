package utilities

import (
	"crypto/sha256"
	"encoding/hex"
	dbstore "stockManagment/DbStore"
	models "stockManagment/DbStore/Models"
	dtos "stockManagment/Dtos"
	"time"

	"github.com/google/uuid"
)

func GetAllUsers() ([]models.User, error) {
	db, err := dbstore.Connect()

	if err != nil {
		return nil, err
	}

	defer dbstore.Disconnect()

	var users []models.User

	err = db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(userInfo *dtos.User) bool {
	db, err := dbstore.Connect()

	if err != nil {
		return false
	}

	defer dbstore.Disconnect()

	hashedPass := HashString(userInfo.Password)

	newUser := models.User{
		ID:           uuid.New().String(),
		FirstName:    userInfo.FirstName,
		LastName:     userInfo.LastName,
		Email:        userInfo.Email,
		Password:     hashedPass,
		Img:          userInfo.Img,
		CreationDate: time.Now(),
	}

	result := db.Create(&newUser)

	if result.Error != nil {
		return false
	}

	return true
}

func HashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedStr := hash.Sum(nil)

	return hex.EncodeToString(hashedStr)
}
