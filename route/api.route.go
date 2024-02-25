package routes

import (
	API "url-shorting/api"
	"url-shorting/middleware"
	"url-shorting/resource/link"
	"url-shorting/resource/user"
)

func Routes() {
	api := API.New().GetInstance()

	user.CreateRoutes()	

	api.Use(middleware.Auth)
	user.CreatePrivateRoutes()
	link.CreateRoutes()

}
