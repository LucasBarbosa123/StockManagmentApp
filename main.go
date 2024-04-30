package main

import (
	general_utilities "stockManagment/Utilities/General"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.Static("/js", "./templates/js")
	server.Static("/imgs", "./templates/imgs")
	server.LoadHTMLGlob("templates/*.html")

	general_utilities.SetRoutes(server)

	server.Run(":8080")
}
