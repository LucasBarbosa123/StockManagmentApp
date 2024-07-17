package cookies_utilities

import (
	dbstore "stockManagment/DbStore"
	models "stockManagment/DbStore/Models"

	"github.com/gin-gonic/gin"
)

func HasSessionCookie(c *gin.Context) bool {
	_, err := c.Cookie("SessionCookie")

	if err != nil {
		return false
	}

	return true
}

func SetSessionCookie(c *gin.Context, userId string) {
	maxAge := 24 * 60 * 60
	path := "/"
	domain := ""

	c.SetCookie("SessionCookie", userId, maxAge, path, domain, false, false)
}

// updates the coockie with maxAge as -1 witch deletes it
func KillSessionCoockie(c *gin.Context) {
	maxAge := -1
	path := "/"
	domain := ""

	c.SetCookie("SessionCookie", "", maxAge, path, domain, false, false)
}

func RetriveCurrentUserId(c *gin.Context) string {
	userId, err := c.Cookie("SessionCookie")

	if err != nil {
		return err.Error()
	}

	return userId
}

func GetCurrentCompanyId(c *gin.Context) string {
	db, err := dbstore.Connect()

	if err != nil {
		return ""
	}
	defer dbstore.Disconnect()

	userId := RetriveCurrentUserId(c)

	var user models.User
	err = db.Where("id = ?", userId).Find(&user).Error

	if err != nil {
		return ""
	}

	return user.CompanyID
}
