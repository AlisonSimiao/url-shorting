package link

import (
	"url-shorting/api"
	"url-shorting/middleware"
)

var lv = NewLinkValidate()

func CreateRoutes() {
	api := api.New()
	lc := NewLinkController()

	api.CreateRoute("/link", "POST", middleware.Validator(lv.Create), lc.Create)
	api.CreateRoute("/link/:hash", "GET", lc.FindOne)
	api.CreateRoute("/link", "GET", lc.FindAll)
	api.CreateRoute("/link/:hash", "PATCH", lc.Update)
	api.CreateRoute("/link/:hash", "DELETE", lc.Delete)
	api.CreateRoute("/link/:hash/clicks", "PATCH", lc.UpdateClick)

}
