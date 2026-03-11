package handlers

import (
	"finances/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type registerRequest struct {
	Email          string `json:"email" binding:"required,email"`
	Name           string `json:"name" binding:"required"`
	BusinessForm   string `json:"business_form" binding:"required"`
	NalogSystem    string `json:"nalog_system" binding:"required"`
	Employees      string `json:"employees" binding:"required"`
	BusinessSphere string `json:"business_sphere" binding:"required"`
	Password       string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(c.Request.Context(), service.RegisterInput{
		Email:          req.Email,
		Name:           req.Name,
		BusinessForm:   req.BusinessForm,
		NalogSystem:    req.NalogSystem,
		Employees:      req.Employees,
		BusinessSphere: req.BusinessSphere,
		Password:       req.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
