package handlers

import (
	"finances/internal/models"
	"finances/internal/repository"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	repo        *repository.TransactionRepository
	projectRepo *repository.ProjectRepository
}

func NewTransactionHandler(repo *repository.TransactionRepository, projectRepo *repository.ProjectRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo, projectRepo: projectRepo}
}

func (h *TransactionHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req models.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаём транзакцию
	t, err := h.repo.Create(int64(userID.(int)), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Пересчитываем проект
	if err := h.recalculateProject(int64(userID.(int)), req.ProjectID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"transaction": t})
}

func (h *TransactionHandler) GetByProject(c *gin.Context) {
	userID, _ := c.Get("user_id")
	projectID, err := strconv.ParseInt(c.Param("projectID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	transactions, err := h.repo.GetByProjectID(int64(userID.(int)), projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// recalculateProject пересчитывает поля проекта после добавления транзакции
func (h *TransactionHandler) recalculateProject(userID, projectID int64) error {
	// Получаем текущий проект
	project, err := h.projectRepo.GetByID(userID, projectID)
	if err != nil {
		return err
	}

	// Получаем суммы транзакций
	totalIncome, totalExpense, err := h.repo.GetTotals(userID, projectID)
	if err != nil {
		return err
	}

	// Парсим стоимость и оплачено (используем простые функции)
	cost := parseStringToFloat(project.Cost)
	paid := parseStringToFloat(project.Paid)

	// Вычисляем новые значения
	otherIncome := totalIncome
	expense := totalExpense
	revenue := cost + otherIncome
	profit := revenue - expense
	remainder := cost - paid

	// Обновляем проект
	_, err = h.projectRepo.Update(userID, projectID, &models.UpdateProjectRequest{
		OtherIncome: formatFloatToString(otherIncome),
		Expense:     formatFloatToString(expense),
		Revenue:     formatFloatToString(revenue),
		Profit:      formatFloatToString(profit),
		Remainder:   formatFloatToString(remainder),
	})

	return err
}

// Вспомогательные функции для handler'а
func parseStringToFloat(s string) float64 {
	if s == "" {
		return 0
	}
	// Убираем " р." и заменяем "." на ""
	s = strings.ReplaceAll(s, " р.", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

func formatFloatToString(val float64) string {
	if val == 0 {
		return "0"
	}
	return fmt.Sprintf("%.0f", val)
}
