package link

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkController struct {
	ls *LinkService
}

func NewLinkController() *LinkController {
	return &LinkController{
		ls: NewLinkService(),
	}
}

func (lc *LinkController) FindOne(c *gin.Context) {

	c.JSON(http.StatusOK, "link")
}

func (lc *LinkController) Update(c *gin.Context) {
	_body, exist := c.Get("body")
	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	hash := c.Param("hash")

	linkUpdate := LinkUpdate{}

	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	body := _body.(map[string]interface{})

	if ative, ok := body["ative"].(bool); ok {
		linkUpdate.Ative = ative

	}

	if original, ok := body["original"].(string); ok {
		linkUpdate.Original = original
	}

	rest_error := lc.ls.update(
		hash,
		linkUpdate,
	)

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(http.StatusOK, body)
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

	link, rest_error := lc.ls.create(idUser, Link{
		Original: body["original"].(string),
		Ative:    ative,
	})

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(http.StatusCreated, link)
}
