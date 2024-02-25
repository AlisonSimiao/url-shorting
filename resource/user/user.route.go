package user

import (
	"time-wise/api"
	"time-wise/middleware"
)

var uv = NewUserValidate()

func CreateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("api/login", "POST", middleware.Validator(uv.Login), uc.Login)
	api.CreateRoute("api/signup", "POST", middleware.Validator(uv.Create), uc.Create)
}

func CreatePrivateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("api/user", "GET", uc.FindOne)
	api.CreateRoute("api/user", "PATCH", middleware.Validator(uv.Update), uc.Update)
}
