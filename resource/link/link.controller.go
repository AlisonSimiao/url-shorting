package link

import (
	"url-shorting/resource/photo"

	"github.com/gin-gonic/gin"
)

type LinkController struct {
}

func NewLinkController() *LinkController {
	return &LinkController{}
}

var linkService = NewLinkService()
var photoService = photo.NewPhotoService()

func (uc *LinkController) FindOne(c *gin.Context) {

	c.JSON(200, "link")
}

func (uc *LinkController) FindAll(c *gin.Context) {

	c.JSON(200, "link")
}

func (uc *LinkController) Update(c *gin.Context) {

	c.JSON(200, "body")
}

func (uc *LinkController) Delete(c *gin.Context) {

	c.JSON(200, "body")
}

func (uc *LinkController) UpdateClick(c *gin.Context) {

	c.JSON(200, "body")
}

func (uc *LinkController) Create(c *gin.Context) {

	c.JSON(200, "link")
}

func (uc *LinkController) Login(c *gin.Context) {

	c.JSON(200, "link")
}
