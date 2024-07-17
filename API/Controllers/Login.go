package controllers

import (
	"net/http"
	dtos "stockManagment/Dtos"
	utilities "stockManagment/Utilities"
	utilities_controllers "stockManagment/Utilities/Controllers"
	cookies_utilities "stockManagment/Utilities/Cookies"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	validate := utilities_controllers.InitValidator()
	var req dtos.LoginInfo
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

	res, user := utilities.Login(req.Email, req.Password)

	if res != "" {
		c.IndentedJSON(http.StatusBadRequest, res)
		return
	}

	cookies_utilities.SetSessionCookie(c, user.ID)

	c.IndentedJSON(http.StatusOK, true)
}

func Logout(c *gin.Context) {
	cookies_utilities.KillSessionCoockie(c)
	c.IndentedJSON(http.StatusOK, true)
}
