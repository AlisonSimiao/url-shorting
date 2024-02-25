package project

import (
	"strconv"
	"time-wise/repository"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ps *ProjectService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		ps: NewProjectService(),
	}
}

func (pc *ProjectController) create(c *gin.Context) {
	idUser, _ := c.Get("idUser")
	data, _ := c.Get("body")
	body := make(map[string]string)

	for key, value := range data.(map[string]interface{}) {
		if stringValue, ok := value.([]string); ok {
			body[key] = stringValue[0]
		} else {
			body[key] = value.(string)
		}
	}

	status, _ := strconv.Atoi(body["status"])
	user, rest_error := pc.ps.create(Project{
		Name:        body["name"],
		IdUser:      idUser.(int),
		Description: body["description"],
		Status:      status,
		IdPhoto: 1,
	}, c)

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)

}

func (pc *ProjectController) paginate(c *gin.Context) {
	idUser, _ := c.Get("idUser")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	pageInfo := repository.Page(limit, page)

	projects, rest_error := pc.ps.paginate(idUser.(int), pageInfo)
	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}
	c.JSON(200, projects)
}

func (pc *ProjectController) findOne(c *gin.Context) {}

func (pc *ProjectController) update(c *gin.Context) {}
