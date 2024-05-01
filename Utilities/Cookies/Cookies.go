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
