package routes

import (
	API "time-wise/api"
	"time-wise/middleware"
	"time-wise/resource/project"
	"time-wise/resource/user"
)

func Routes() {
	api := API.New().GetInstance()
	//api.New().CreateRoute("/swagger/*any", "GET", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user.CreateRoutes()

	api.Use(middleware.Auth)
	user.CreatePrivateRoutes()
	project.CreateRoutes()
}
