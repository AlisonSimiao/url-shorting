package user

import (
	"url-shorting/resource/photo"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

var userService = NewUserService()
var photoService = photo.NewPhotoService()

func (uc *UserController) FindOne(c *gin.Context) {
	userIDString, _ := c.Get("idUser")

	user, rest_error := userService.findOne(userIDString.(int))
	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) Update(c *gin.Context) {
	data, exist := c.Get("body")
	body := make(map[string]string)

	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	for key, value := range data.(map[string]interface{}) {
		if stringValue, ok := value.([]string); ok {
			body[key] = stringValue[0]
		} else {
			body[key] = value.(string)
		}
	}

	userIDString, _ := c.Get("idUser")

	rest_error := userService.update(userIDString.(int), User{
		Name:     body["name"],
		Email:    body["email"],
		Password: body["password"],
	}, c)

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, body)
}

func (uc *UserController) Create(c *gin.Context) {
	data, exist := c.Get("body")
	body := make(map[string]string)

	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	for key, value := range data.(map[string]interface{}) {
		if stringValue, ok := value.([]string); ok {
			body[key] = stringValue[0]
		} else {
			body[key] = value.(string)
		}
	}

	user, rest_error := userService.create(User{
		Name:     body["name"],
		Email:    body["email"],
		Password: body["password"],
	}, c)

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) Login(c *gin.Context) {

	body, exist := c.Get("body")

	if !exist {
		c.JSON(500, "Erro no servidor")
		return
	}
	Body := body.(map[string]interface{})

	user, rest_error := userService.login(UserLogin{
		Email:    Body["email"].(string),
		Password: Body["password"].(string),
	})

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)
}
