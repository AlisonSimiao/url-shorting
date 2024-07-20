package routes

import (
	API "vagas-api/api"
	"vagas-api/middleware"
	"vagas-api/resource/user"
	_ "vagas-api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title TODO APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /api/v1
// @schemes htpt

func Routes() {
	api := API.New().GetInstance()

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"version": "1.0.0"})
	})
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	user.CreateRoutes()
	api.Use(middleware.Auth)
	user.CreatePrivateRoutes()
}
