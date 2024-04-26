package link

import (
	"url-shorting/api"
	"url-shorting/middleware"
)

var lv = NewLinkValidate()

func CreateRoutes() {
	api := api.New()
	lc := NewLinkController()

	api.CreateRoute("/links", "POST", middleware.Validator(lv.Create), lc.Create)
	api.CreateRoute("/links/:hash", "GET", lc.FindOne)
	api.CreateRoute("/links", "GET", lc.FindAll)
	api.CreateRoute("/links/:hash", "PATCH", middleware.Validator(lv.Update), lc.Update)
	api.CreateRoute("/links/:hash", "DELETE", lc.Delete)
	api.CreateRoute("/links/:hash/clicks", "PATCH", lc.UpdateClick)

}
