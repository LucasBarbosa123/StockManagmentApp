package controllers

import (
	"crypto/rand"
	"math/big"
	"net/http"
	models "stockManagment/DbStore/Models"
	dtos "stockManagment/Dtos"
	utilities "stockManagment/Utilities"
	utilities_controllers "stockManagment/Utilities/Controllers"
	cookies_utilities "stockManagment/Utilities/Cookies"

	"github.com/gin-gonic/gin"
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
	validate := utilities_controllers.InitValidator()
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

	currentCompanyId := cookies_utilities.GetCurrentCompanyId(c)
	success := utilities.CreateUser(&req, currentCompanyId)

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

func GetUserInfoById(c *gin.Context) {
	userId := c.Param("id")

	user, err := utilities.GetUserById(userId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Problems finding the user.")
		return
	}

	userInfo := dtos.UserInfo{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	c.JSON(http.StatusOK, userInfo)
}

func EditUser(c *gin.Context) {
	validate := utilities_controllers.InitValidator()
	var req dtos.UserInfo

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

	err = utilities.EditUser(req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func ChangeUserInfo(c *gin.Context) {
	validate := utilities_controllers.InitValidator()
	userId := c.Param("id")
	var req dtos.UserChangerInfo

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

	err = utilities.ChangeUserInfo(userId, req)
	if err != nil {
		c.IndentedJSON(http.StatusTeapot, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, true)
}

func ChangePass(c *gin.Context) {
	validate := utilities_controllers.InitValidator()
	userId := c.Param("id")
	var req dtos.PassChangerInfo

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

	err = utilities.ChangeUserPassword(userId, req)
	if err != nil {
		c.IndentedJSON(http.StatusTeapot, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, true)
}

func GetAllUsers(c *gin.Context) []models.User {
	currentCompanyId := cookies_utilities.GetCurrentCompanyId(c)
	users, _ := utilities.GetAllUsers(currentCompanyId)
	return users
}

func GetSpecificUserInfo(c *gin.Context) models.User {
	userId := cookies_utilities.RetriveCurrentUserId(c)
	user, _ := utilities.GetUserById(userId)
	return *user
}
