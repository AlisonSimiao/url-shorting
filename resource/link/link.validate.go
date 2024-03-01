package link

import "github.com/thedevsaddam/govalidator"

type LinkValidate struct {
	Create govalidator.MapData
	Update govalidator.MapData
}

var rulesCreate = govalidator.MapData{
	"original": {"required", "url", "max:255"},
	"ative":    {"regex:^(true|false)$"},
}

var rulesUpdate = govalidator.MapData{}

func NewLinkValidate() LinkValidate {
	return LinkValidate{
		Create: rulesCreate,
		Update: rulesUpdate,
	}
}
