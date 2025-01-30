package handlers

import (
	"net/http"
	"strconv"

	"github.com/codepnw/go-movie-booking/internal/models"
	"github.com/codepnw/go-movie-booking/internal/services"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service services.IUserService
}

func NewUserHandler(service services.IUserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Register(c *gin.Context) {
	var payload models.UserRegisterReq

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.Password != payload.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password not match"})
		return
	}

	// User Service
	if err := h.service.Register(c, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "register success"})
}

func (h *userHandler) Login(c *gin.Context) {
	var payload models.UserLoginReq

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Login(c, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *userHandler) GetProfile(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	user, err := h.service.GetByID(c, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
