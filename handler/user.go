package handler

import (
	"fmt"
	"golang-bwa/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(s user.Service) *userHandler {
	return &userHandler{s}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	fmt.Println("input first ", input)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	fmt.Println("input after ", input)

	user, err := h.service.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, user)
}
