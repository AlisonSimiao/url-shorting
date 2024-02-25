package middleware

import (
	"mime/multipart"
	"net/url"
	message "url-shorting/utils"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

var Messages = govalidator.MapData{
	"email": []string{
		"required:" + message.Required("email"),
		"email:" + message.Email(),
	},
	"password": []string{
		"required:" + message.Required("password"),
		"min:" + message.MinLength("password", "6"),
		"max:" + message.MaxLength("password", "20"),
		"alpha_num:" + message.AlphaNum("password"),
	},
	"name": []string{
		"required:" + message.Required("name"),
		"min" + message.MinLength("name", "3"),
		"max:" + message.MaxLength("name", "20"),
		"alpha:" + message.Alpha("name"),
		"alpha_num:" + message.AlphaNum("name"),
	},
	"description": []string{
		"required:" + message.Required("description"),
		"min" + message.MinLength("description", "3"),
		"max:" + message.MaxLength("description", "100"),
		"alpha_num:" + message.AlphaNum("description"),
	},
	"status": []string{
		"required:" + message.Required("status"),
		"in:" + message.IsIn("status", "1, 2"),
		"numeric:" + message.Numeric("status"),
	},
}

type FormData struct {
	Fields map[string]interface{}
	Files  map[string]*multipart.FileHeader
}

func parseMultipartForm(c *gin.Context) (*FormData, error) {
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB de tamanho máximo
	if err != nil {
		return nil, err
	}

	formData := &FormData{
		Fields: make(map[string]interface{}),
		Files:  make(map[string]*multipart.FileHeader),
	}

	for key, values := range c.Request.MultipartForm.Value {
		formData.Fields[key] = values
	}

	for key, files := range c.Request.MultipartForm.File {
		formData.Files[key] = files[0]
	}

	return formData, nil
}

func Validator(rules govalidator.MapData) func(*gin.Context) {

	return func(c *gin.Context) {
		var body map[string]interface{}

		opcs := govalidator.Options{
			Request:  c.Request,
			Data:     &body,
			Rules:    rules,
			Messages: Messages,
		}

		validate := govalidator.New(opcs)

		var unprocessableErros url.Values
		if c.ContentType() == "multipart/form-data" {
			unprocessableErros = validate.Validate()
		} else {
			unprocessableErros = validate.ValidateJSON()
		}

		if len(unprocessableErros) > 0 {
			var errors []string
			for _, err := range unprocessableErros {
				errors = append(errors, err...)
			}

			c.AbortWithStatusJSON(422, errors)
			return
		}
		if c.ContentType() == "multipart/form-data" {
			formData, err := parseMultipartForm(c)
			if err != nil {
				c.AbortWithStatusJSON(500, err)
				return
			}

			body = formData.Fields
			c.Set("files", formData.Files)
		}

		c.Set("body", body)
		c.Next()
	}
}
