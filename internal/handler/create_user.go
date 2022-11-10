package handler

import (
	"github.com/gin-gonic/gin"
	"rest_api_test"
)

func (h Handler) createUser(c *gin.Context) {
	var input rest_api_test.User

	err := c.BindJSON(&input)
	if err != nil {

	}
}
