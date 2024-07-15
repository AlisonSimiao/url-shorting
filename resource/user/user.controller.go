package user

import (
	//"vagas-api/resource/photo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

var userService = NewUserService()

//var photoService = photo.NewPhotoService()

func (uc *UserController) FindOne(c *gin.Context) {
	userIDString, _ := c.Get("idUser")

	user, rest_error := userService.findOne(userIDString.(int))
	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Update(c *gin.Context) {
	_body, exist := c.Get("body")

	userUpdates := User{}

	body := _body.(map[string]interface{})

	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	if name, ok := body["name"].(string); ok && name != "" {
		userUpdates.Name = name
	}
	if email, ok := body["email"].(string); ok && email != "" {
		userUpdates.Email = email
	}
	if password, ok := body["password"].(string); ok && password != "" {
		userUpdates.Password = password
	}
	if username, ok := body["username"].(string); ok && username != "" {
		userUpdates.Username = username
	}
	if status, ok := body["status"].(bool); ok {
		userUpdates.Status = status
	}
	if pro, ok := body["pro"].(bool); ok {
		userUpdates.Pro = pro
	}

	userIDString, _ := c.Get("idUser")

	rest_error := userService.update(userIDString.(int), userUpdates)

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(http.StatusOK, body)
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
		Username: body["username"],
		Status:   true,
		Pro:      false,
	})

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) Login(c *gin.Context) {

	body, exist := c.Get("body")

	if !exist {
		c.JSON(500, "Erro no servidor")
		return
	}
	Body := body.(map[string]interface{})

	user, rest_error := userService.login(UserLogin{
		Username: Body["username"].(string),
		Password: Body["password"].(string),
	})

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}
	c.SetCookie("auth_token", user.Token, 0, "/", "", false, true)
	c.JSON(http.StatusOK, user)
}
