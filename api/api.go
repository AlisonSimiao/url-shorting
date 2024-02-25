package api

import (
	"fmt"
	"os"
	"strconv"

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
		api.instance = gin.Default()
	}

	return &api
}

func (a *API) GetInstance() *gin.Engine {
	return a.instance
}

func (a *API) Start() {
	host := os.Getenv("API_HOST")
	port, _ := strconv.Atoi(os.Getenv("API_PORT"))

	if a.isRuning {
		return
	}

	if a.instance == nil {
		New()
	}

	a.host = host
	a.port = port
	a.isRuning = true
	a.instance.Run(fmt.Sprintf("%s:%d", host, port))
}

func (a *API) CreateRoute(path string, method string, handler ...gin.HandlerFunc) {
	if a.instance == nil {
		return
	}

	a.instance.Handle(method, path, handler...)
}
