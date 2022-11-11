package handler

import (
	"github.com/gin-gonic/gin"
	"rest_api_test/internal/service"
)

type Creating interface {
	CreateUser(c *gin.Context)
}

type Handler struct {
	services *service.Service
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/create_user", h.CreateUser)
	return router
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
