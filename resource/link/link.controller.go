package link

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkController struct {
}

func NewLinkController() *LinkController {
	return &LinkController{}
}

var linkService = NewLinkService()

func (lc *LinkController) FindOne(c *gin.Context) {

	c.JSON(http.StatusOK, "link")
}

func (lc *LinkController) FindAll(c *gin.Context) {

	c.JSON(http.StatusOK, "link")
}

func (lc *LinkController) Update(c *gin.Context) {

	c.JSON(http.StatusOK, "body")
}

func (lc *LinkController) Delete(c *gin.Context) {

	c.JSON(http.StatusOK, "body")
}

func (lc *LinkController) UpdateClick(c *gin.Context) {

	c.JSON(http.StatusOK, "body")
}

func (lc *LinkController) Create(c *gin.Context) {
	data, exist := c.Get("body")

	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	body := data.(map[string]interface{})

	ative := body["ative"].(bool)
	if !ative {
		ative = true
	}

	idUser := c.GetInt("idUser")

	link, rest_error := linkService.create(idUser, Link{
		Original: body["original"].(string),
		Ative:    ative,
	})

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(http.StatusCreated, link)
}

func (lc *LinkController) Login(c *gin.Context) {

	c.JSON(http.StatusOK, "link")
}
