package handler

import (
	"golang-bwa/helper"
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

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	res := helper.ApiResponse(http.StatusOK, "success", users)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse(http.StatusConflict, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	newUser, err := h.service.RegisterUser(input)
	if err != nil {
		res := helper.ApiResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	userFormat := user.UserRegisterFormatter(newUser, "TokentokenToken")
	res := helper.ApiResponse(http.StatusOK, "success", userFormat)
	c.JSON(http.StatusOK, res)
}
