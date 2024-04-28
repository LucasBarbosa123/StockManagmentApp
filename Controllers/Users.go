package controllers

import (
	"crypto/rand"
	"math/big"
	"net/http"
	models "stockManagment/DbStore/Models"
	dtos "stockManagment/Dtos"
	utilities "stockManagment/Utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GeneratePassword(c *gin.Context) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	password := make([]byte, 8)
	for i := range password {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		password[i] = charset[idx.Int64()]
	}

	c.IndentedJSON(http.StatusOK, string(password))
}

func CreateUser(c *gin.Context) {
	initValidator()
	var req dtos.User

	err := c.ShouldBind(&req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid request body.")
		return
	}

	err = validate.Struct(&req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success := utilities.CreateUser(&req)

	if !success {
		c.IndentedJSON(http.StatusExpectationFailed, "Something is wrong creating the user.")
		return
	}

	c.IndentedJSON(http.StatusCreated, true)
}

func DeleteUser(c *gin.Context) {
	var req string

	err := c.ShouldBind(&req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err = utilities.DeleteUser(req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func GetAllUsers() []models.User {
	users, _ := utilities.GetAllUsers()
	return users
}

var validate *validator.Validate

func initValidator() {
	validate = validator.New()
}
