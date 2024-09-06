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
	server.GET("/wearhouses", middleware.CheckCookie, func(c *gin.Context) {
		wearhouses := controllers.GetAllWearhouses(c)
		c.HTML(http.StatusOK, "wearhouses.html", gin.H{
			"wearhouses": wearhouses,
		})
	})
	server.GET("/users", middleware.CheckCookie, func(c *gin.Context) {
		users := controllers.GetAllUsers(c)
		c.HTML(http.StatusOK, "users.html", gin.H{
			"users": users,
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

	//users
	server.GET("/api/generate-pass", middleware.CheckCookie, controllers.GeneratePassword)
	server.GET("/api/user-info/:id", middleware.CheckCookie, controllers.GetUserInfoById)
	server.POST("/api/create-user", middleware.CheckCookie, controllers.CreateUser)
	server.POST("/api/delete-user", middleware.CheckCookie, controllers.DeleteUser)
	server.POST("/api/edit-user", middleware.CheckCookie, controllers.EditUser)
	server.POST("/api/change-user-info/:id", middleware.CheckCookie, controllers.ChangeUserInfo)

	//login
	server.POST("/api/login", middleware.CheckNoCookie, controllers.Login)
	server.POST("/api/logout", middleware.CheckCookie, controllers.Logout)

	//whearhouse
	server.GET("/api/warehouse-name/:id", middleware.CheckCookie, controllers.GetWarehouseName)
	server.POST("/api/create-wearhouse", middleware.CheckCookie, controllers.CreateWearhouse)
	server.POST("/api/delete-warehouse", middleware.CheckCookie, controllers.DeleteWarehouse)
	server.POST("/api/edit-warehouse", middleware.CheckCookie, controllers.EditWarehouse)
}
