package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rest_api_test/internal/service"
)

type Auth interface {
	Hello(c echo.Context) error
}

type Handler struct {
	Auth
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//func
