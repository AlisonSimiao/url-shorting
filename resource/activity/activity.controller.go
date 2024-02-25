package activity

import "github.com/gin-gonic/gin"

type ActivityController struct {}

func NewActivityController() *ActivityController {
	return &ActivityController{}
}

func (ac *ActivityController) create(ctx *gin.Context) {}

func (ac *ActivityController) paginate(ctx *gin.Context) {}

func (ac *ActivityController) findOne(ctx *gin.Context) {}

func (ac *ActivityController) update(ctx *gin.Context) {}
