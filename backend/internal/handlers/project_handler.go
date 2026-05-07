package handlers

import (
	"net/http"
	"strconv"

	"finances/internal/models"
	"finances/internal/repository"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	repo *repository.ProjectRepository
}

func NewProjectHandler(repo *repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{repo: repo}
}

func (h *ProjectHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req models.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err := h.repo.Create(int64(userID.(int)), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"project": project})
}

func (h *ProjectHandler) GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")
	projects, err := h.repo.GetAll(int64(userID.(int)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"projects": projects, "total": len(projects)})
}

func (h *ProjectHandler) GetByID(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	project, err := h.repo.GetByID(int64(userID.(int)), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}

func (h *ProjectHandler) Update(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req models.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err := h.repo.Update(int64(userID.(int)), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}

func (h *ProjectHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.repo.Delete(int64(userID.(int)), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *ProjectHandler) BulkDelete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req models.BulkDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.DeleteMultiple(int64(userID.(int)), req.IDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *ProjectHandler) GetAllFiltered(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var filter models.ProjectFilter
	// Парсим query параметры
	filter.SortBy = c.DefaultQuery("sort_by", "created_at")
	filter.Order = c.DefaultQuery("order", "desc")
	filter.Stage = c.Query("stage")
	filter.Priority = c.Query("priority")
	filter.Search = c.Query("search")
	filter.StartDate = c.Query("start_date")
	filter.EndDate = c.Query("end_date")

	projects, err := h.repo.GetAllFiltered(int64(userID.(int)), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects, "total": len(projects)})
}
