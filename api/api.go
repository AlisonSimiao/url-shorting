package api

import (
	"os"
	"github.com/gin-gonic/gin"
)

type API struct {
	host     string
	port     int
	isRuning bool
	instance *gin.Engine
}

var api = API{}

func (a *API) IsRuning() bool {
	return a.isRuning
}

func New() *API {
	if api.instance == nil {
		gin.SetMode(gin.DebugMode)
		api.instance = gin.Default()
	}

	return &api
}

func (a *API) GetInstance() *gin.Engine {
	return a.instance
}

func (a *API) Start() {
	var port int
	host := os.Getenv("API_HOST")

	if a.isRuning {
		return
	}

	if a.instance == nil {
		New()
	}

	a.host = host
	a.port = port
	a.isRuning = true
	a.instance.Run(host)
}

func (a *API) CreateRoute(path string, method string, handler ...gin.HandlerFunc) {
	if a.instance == nil {
		return
	}

	a.instance.Handle(method, path, handler...)
}
