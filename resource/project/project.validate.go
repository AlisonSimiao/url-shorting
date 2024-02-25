package project

import "github.com/thedevsaddam/govalidator"

type ProjectValidate struct {
	create govalidator.MapData
	update govalidator.MapData
}

var rulesCreate = govalidator.MapData{
	"name":     {"required", "min:1", "max:100"},
	"description":  {"required", "min:1", "max:1000"},
	"status":  {"required", "numeric", "in:1,2,3,4,5,6,7,8,9,10"},
}

var rulesUpdate = govalidator.MapData{
	"name":     {"min:1", "max:100", "alpha_num"},
	"description":  {"min:1", "max:1000", "alpha_num"},
	"status":  {"numeric", "in:1,2,3,4,5,6,7,8,9,10"},
}

var rulesLogin = govalidator.MapData{
	"email":    {"required", "email"},
	"password": {"required", "min:6", "max:100", "alpha_num"},
}

func NewProjectValidate() *ProjectValidate {
	return &ProjectValidate{
		create: rulesCreate,
		update: rulesUpdate,
	}
}
