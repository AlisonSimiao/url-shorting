package project

import (
	"time-wise/api"
	"time-wise/middleware"
)

var pv = NewProjectValidate()

func CreateRoutes() {
	api := api.New()
	pc := NewProjectController()

	api.CreateRoute("/api/projects", "POST", middleware.Validator(pv.create), pc.create)
	api.CreateRoute("/api/projects", "GET", pc.paginate)
	api.CreateRoute("/api/projects/:id", "GET", pc.findOne)
	api.CreateRoute("/api/projects/:id", "PATCH", middleware.Validator(pv.update), pc.update)
}