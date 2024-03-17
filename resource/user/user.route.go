package user

import (
	"url-shorting/api"
	"url-shorting/middleware"
)

var uv = NewUserValidate()

func CreateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("/users/signin", "POST", middleware.Validator(uv.Login), uc.Login)
	api.CreateRoute("/users/signup", "POST", middleware.Validator(uv.Create), uc.Create)
}

func CreatePrivateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("/users", "GET", uc.FindOne)
	api.CreateRoute("/users", "PATCH", middleware.Validator(uv.Update), uc.Update)
}
