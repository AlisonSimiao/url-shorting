package link

import (
	"net/http"
	"strconv"
	"url-shorting/repository"

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

	hash := c.Param("hash")

	link, erro := lc.ls.findOne(hash)
	if erro != nil {
		c.AbortWithStatusJSON(erro.GetStatus(), erro.GetMensagem())
		return
	}

	c.JSON(http.StatusOK, link)
}

func (lc *LinkController) FindAll(c *gin.Context) {

	//Pega Id do Usuario.
	page, erro := strconv.Atoi(c.Query("pagina"))
	if erro != nil {
		page = 0
	}

	limit, erro := strconv.Atoi(c.Query("registros"))
	if erro != nil {
		limit = 0
	}

	pageI := repository.Page(limit, page)

	idUser, exist := c.Get("idUser")
	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	//volta
	links, err := lc.ls.findAll(idUser.(int), pageI.Page, pageI.Limit)
	if err != nil {
		c.AbortWithStatusJSON(err.GetStatus(), err.GetMensagem())
		return
	}
	c.JSON(http.StatusOK, links)
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

	hash := c.Param("hash")
	if hash == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Hash não informado"})
	}
	rest_error := lc.ls.delete(hash)

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.Status(http.StatusOK)
}

func (lc *LinkController) UpdateClick(c *gin.Context) {

	hash := c.Param("hash")

	rest_error := lc.ls.updateClick(hash)

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.Status(http.StatusOK)
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
