package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_test/internal/service"
)

type Auth interface {
	Hello(c gin.Context)
}

type Handler struct {
	services *service.Service
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", h.hello)
	router.POST("/create_user", h.createUser)
	return router
}

func (h Handler) hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
