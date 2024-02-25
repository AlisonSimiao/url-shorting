package activity

import (
	"time-wise/api"
	"time-wise/middleware"
)

func CreateRoutes() {
	api := api.New()
	ac  := NewActivityController()
	av  := NewActivityValidate()

	api.CreateRoute("/projects/:idProject/activities", "POST", middleware.Validator(av.create), ac.create)
	api.CreateRoute("/projects/:idProject/activities", "GET", ac.paginate)
	api.CreateRoute("/projects/:idProject/activities/:id", "GET", ac.findOne)
	api.CreateRoute("/projects/:idProject/activities/:id", "PATCH", middleware.Validator(av.update), ac.update)
}
