package user

import "github.com/thedevsaddam/govalidator"

type UserValidate struct {
	Create govalidator.MapData
	Login  govalidator.MapData
	Update govalidator.MapData
}

var rulesCreate = govalidator.MapData{
	"email":    {"required", "email"},
	"password": {"required", "min:6", "max:100", "alpha_num"},
	"name":     {"required", "min:3", "max:100", "alpha_num"},
}

var rulesUpdate = govalidator.MapData{
	"email":    {"email"},
	"password": {"min:6", "max:100", "alpha_num"},
	"name":     {"min:3", "max:100", "alpha_space"},
}

var rulesLogin = govalidator.MapData{
	"email":    {"required", "email"},
	"password": {"required", "min:6", "max:100", "alpha_num"},
}

func NewUserValidate() *UserValidate {
	return &UserValidate{
		Create: rulesCreate,
		Login:  rulesLogin,
		Update: rulesUpdate,
	}
}
