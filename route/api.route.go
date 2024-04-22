package routes

import (
	API "url-shorting/api"
	"url-shorting/middleware"
	"url-shorting/resource/link"
	"url-shorting/resource/user"

	"github.com/gin-gonic/gin"
)

func Routes() {
	api := API.New().GetInstance()

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"version": "1.0.0"})
	})
	user.CreateRoutes()

	api.Use(middleware.Auth)
	user.CreatePrivateRoutes()
	link.CreateRoutes()

}
