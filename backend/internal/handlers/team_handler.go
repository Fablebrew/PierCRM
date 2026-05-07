package handlers

import (
	"net/http"
	"strconv"

	"finances/internal/models"
	"finances/internal/repository"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	repo *repository.TeamRepository
}

func NewTeamHandler(repo *repository.TeamRepository) *TeamHandler {
	return &TeamHandler{repo: repo}
}

// CreateTeamMember - POST /api/team
func (h *TeamHandler) CreateTeamMember(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req models.CreateTeamMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	member, err := h.repo.Create(int64(userID.(int)), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"member": member})
}

// GetTeamMembers - GET /api/team
func (h *TeamHandler) GetTeamMembers(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	members, err := h.repo.GetAll(int64(userID.(int)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"members": members,
		"total":   len(members),
	})
}

// GetTeamMember - GET /api/team/:id
func (h *TeamHandler) GetTeamMember(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	memberID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid member ID"})
		return
	}

	member, err := h.repo.GetByID(int64(userID.(int)), memberID)
	if err != nil {
		if err.Error() == "team member not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "team member not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"member": member})
}

// UpdateTeamMember - PUT /api/team/:id
func (h *TeamHandler) UpdateTeamMember(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	memberID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid member ID"})
		return
	}

	var req models.UpdateTeamMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	member, err := h.repo.Update(int64(userID.(int)), memberID, &req)
	if err != nil {
		if err.Error() == "team member not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "team member not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"member": member})
}

// DeleteTeamMember - DELETE /api/team/:id
func (h *TeamHandler) DeleteTeamMember(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	memberID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid member ID"})
		return
	}

	if err := h.repo.Delete(int64(userID.(int)), memberID); err != nil {
		if err.Error() == "team member not found or not authorized" {
			c.JSON(http.StatusNotFound, gin.H{"error": "team member not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "team member deleted successfully"})
}

// DeleteMultipleTeamMembers - POST /api/team/bulk-delete
func (h *TeamHandler) DeleteMultipleTeamMembers(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req models.BulkDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	if err := h.repo.DeleteMultiple(int64(userID.(int)), req.IDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "team members deleted successfully"})
}
