package activity

import "github.com/thedevsaddam/govalidator"

type ActivityValidate struct {
	create govalidator.MapData
	update govalidator.MapData
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

func NewActivityValidate() *ActivityValidate {
	return &ActivityValidate{
		create: rulesCreate,
		update: rulesUpdate,
	}
}
