package general_utilities

import (
	"net/http"
	controllers "stockManagment/Controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(server *gin.Engine) {
	//pages
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	server.GET("/users", func(c *gin.Context) {
		users := controllers.GetAllUsers()
		c.HTML(http.StatusOK, "users.html", gin.H{
			"users": users,
		})
	})
	server.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//end points
	server.GET("/api/generate-pass", controllers.GeneratePassword)
	server.POST("/api/create-user", controllers.CreateUser)
	server.POST("/api/delete-user", controllers.DeleteUser)
	server.GET("/api/user-info/:id", controllers.GetUserInfoById)
	server.POST("/api/edit-user", controllers.EditUser)
	server.POST("/api/login", controllers.Login)
}
