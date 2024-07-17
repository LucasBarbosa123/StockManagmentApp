package utilities

import (
	dbstore "stockManagment/DbStore"
	models "stockManagment/DbStore/Models"

	hashing_utilities "stockManagment/Utilities/Hashing"
)

func Login(email string, pass string) (string, *models.User) {
	db, err := dbstore.Connect()

	if err != nil {
		return "Problems connection with the data base.", nil
	}

	defer dbstore.Disconnect()

	var user models.User

	err = db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return "Problems interacting with the data base.", nil
	}

	if user.ID == "" {
		return "No user found with that email.", nil
	}

	correctPass := CheckPass(&user, pass)
	if !correctPass {
		return "Incorrect password.", nil
	}

	return "", &user
}

func CheckPass(user *models.User, pass string) bool {
	hashedPass := hashing_utilities.HashString(pass)
	return (hashedPass == user.Password)
}
