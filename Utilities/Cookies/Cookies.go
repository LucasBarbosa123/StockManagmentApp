package cookies_utilities

import (
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
