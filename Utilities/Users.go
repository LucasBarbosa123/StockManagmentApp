package utilities

import (
	dbstore "stockManagment/DbStore"
	models "stockManagment/DbStore/Models"
	dtos "stockManagment/Dtos"
	hashing_utilities "stockManagment/Utilities/Hashing"
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

	hashedPass := hashing_utilities.HashString(userInfo.Password)

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

func DeleteUser(id string) error {
	db, err := dbstore.Connect()

	if err != nil {
		return err
	}

	defer dbstore.Disconnect()

	var user models.User
	err = db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return err
	}

	err = db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserById(userId string) (*models.User, error) {
	db, err := dbstore.Connect()

	if err != nil {
		return nil, err
	}

	defer dbstore.Disconnect()

	var user models.User
	err = db.Where("id = ?", userId).Find(&user).Error

	return &user, err
}

func EditUser(userInfo dtos.UserInfo) error {
	db, err := dbstore.Connect()

	if err != nil {
		return err
	}

	defer dbstore.Disconnect()

	var user models.User
	err = db.Where("id = ?", userInfo.Id).Find(&user).Error
	if err != nil {
		return err
	}

	user.FirstName = userInfo.FirstName
	user.LastName = userInfo.LastName
	user.Email = userInfo.Email

	err = db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}
