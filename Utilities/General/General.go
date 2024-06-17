package general_utilities

import (
	"net/http"
	controllers "stockManagment/Controllers"
	middleware "stockManagment/Middleware"

	"github.com/gin-gonic/gin"
)

func SetRoutes(server *gin.Engine) {
	//pages
	server.GET("/", middleware.CheckCookie, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	server.GET("/users", middleware.CheckCookie, func(c *gin.Context) {
		users := controllers.GetAllUsers(c)
		c.HTML(http.StatusOK, "users.html", gin.H{
			"users": users,
		})
	})
	server.GET("/wearhouses", middleware.CheckCookie, func(c *gin.Context) {
		wearhouses := controllers.GetAllWearhouses(c)
		c.HTML(http.StatusOK, "wearhouses.html", gin.H{
			"wearhouses": wearhouses,
		})
	})
	server.GET("/myaccount", middleware.CheckCookie, func(c *gin.Context) {
		userInfo := controllers.GetSpecificUserInfo(c)
		c.HTML(http.StatusOK, "myAccount.html", gin.H{
			"userInfo": userInfo,
		})
	})
	server.GET("/stocks", middleware.CheckCookie, func(c *gin.Context) {
		c.HTML(http.StatusOK, "stocks.html", gin.H{})
	})
	server.GET("/login", middleware.CheckNoCookie, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//end points
	server.GET("/api/generate-pass", controllers.GeneratePassword)
	server.POST("/api/create-user", controllers.CreateUser)
	server.POST("/api/delete-user", controllers.DeleteUser)
	server.GET("/api/user-info/:id", controllers.GetUserInfoById)
	server.POST("/api/edit-user", controllers.EditUser)
	server.POST("/api/login", controllers.Login)
	server.POST("/api/logout", controllers.Logout)
	server.POST("/api/create-wearhouse", controllers.CreateWearhouse)
	server.POST("/api/delete-warehouse", controllers.DeleteWarehouse)
	server.GET("/api/warehouse-name/:id", controllers.GetWarehouseName)
	server.POST("/api/edit-warehouse", controllers.EditWarehouse)
}
