package user

import "github.com/thedevsaddam/govalidator"

type UserValidate struct {
	Create govalidator.MapData
	Login  govalidator.MapData
	Update govalidator.MapData
}

var rulesCreate = govalidator.MapData{
	"email":    {"required", "email", "max:255"},
	"password": {"required", "min:6", "max:32", "alpha_num"},
	"name":     {"required", "min:3", "max:100", "alpha_num"},
	"username":     {"required", "min:3", "max:16", "alpha_num"},
	"status":     {"min:3", "max:100", "alpha_num"},
	"pro":     {"min:3", "max:100", "alpha_num"},
}

var rulesUpdate = govalidator.MapData{
	"email":    {"email", "max:255"},
	"password": {"min:6", "max:32", "alpha_num"},
	"name":     {"min:3", "max:100", "alpha_num"},
	"username": {"min:3", "max:16", "alpha_num"},
	"status":   {"min:3", "max:100", "alpha_num"},
	"pro":      {"min:3", "max:100", "alpha_num"},
}

var rulesLogin = govalidator.MapData{
	"username":    {"required", "min:1", "max:100", "alpha_num"},
	"password": {"required", "min:1", "max:100", "alpha_num"},
}

func NewUserValidate() *UserValidate {
	return &UserValidate{
		Create: rulesCreate,
		Login:  rulesLogin,
		Update: rulesUpdate,
	}
}
