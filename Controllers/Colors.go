package controllers

import (
	"net/http"
	dtos "stockManagment/Dtos"
	utilities "stockManagment/Utilities"
	cookies_utilities "stockManagment/Utilities/Cookies"

	"github.com/gin-gonic/gin"
)

func CreateColor(c *gin.Context) {
	var req dtos.ColorInfo
	err := c.ShouldBind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	userId := cookies_utilities.RetriveCurrentUserId(c)
	companyId := cookies_utilities.GetCurrentCompanyId(c)
	success := utilities.CreateColor(&req, companyId, userId)
	if !success {
		c.IndentedJSON(http.StatusBadRequest, false)
		return
	}

	c.IndentedJSON(http.StatusBadRequest, true)
	return
}
