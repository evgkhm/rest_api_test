package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_test"
)

func (h Handler) CreateUser(c *gin.Context) {
	var input rest_api_test.User

	err := c.BindJSON(&input)
	if err != nil {
		//newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		c.JSON(http.StatusBadRequest, "Bad request")
	}
	h.services.CreateUser(input)
}
