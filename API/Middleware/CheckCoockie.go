package middleware

import (
	"net/http"
	cookies_utilities "stockManagment/Utilities/Cookies"

	"github.com/gin-gonic/gin"
)

func CheckCookie(c *gin.Context) {
	if !cookies_utilities.HasSessionCookie(c) {
		c.Redirect(http.StatusFound, "/login")
	}

	c.Next()
}

func CheckNoCookie(c *gin.Context) {
	if cookies_utilities.HasSessionCookie(c) {
		c.Redirect(http.StatusFound, "/")
	}

	c.Next()
}
