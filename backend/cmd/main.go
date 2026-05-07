package main

import (
	"log"
	"net/http"

	"finances/internal/database"
	"finances/internal/handlers"
	"finances/internal/middleware"
	"finances/internal/repository"
	"finances/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	teamRepo := repository.NewTeamRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo, projectRepo)
	projectHandler := handlers.NewProjectHandler(projectRepo)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, "supersecretkey")

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	teamHandler := handlers.NewTeamHandler(teamRepo)

	authMiddleware := middleware.AuthMiddleware("supersecretkey")

	// Создаём роутер БЕЗ дефолтных middleware
	r := gin.New()

	// ВАЖНО: настройки ДО регистрации маршрутов
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Публичные маршруты (без /api префикса)
	r.POST("/register", userHandler.Register)
	r.POST("/login", authHandler.Login)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is running"})
	})

	// Защищенные маршруты
	protected := r.Group("/api")
	protected.Use(authMiddleware)

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(200, gin.H{"user_id": userID})
	})

	// Команда - используем ПОЛНЫЕ ПУТИ без Group
	protected.POST("/team", teamHandler.CreateTeamMember)
	protected.GET("/team", teamHandler.GetTeamMembers)
	protected.GET("/team/:id", teamHandler.GetTeamMember)
	protected.PUT("/team/:id", teamHandler.UpdateTeamMember)
	protected.DELETE("/team/:id", teamHandler.DeleteTeamMember)
	protected.POST("/team/bulk-delete", teamHandler.DeleteMultipleTeamMembers)

	protected.POST("/transactions", transactionHandler.Create)
	protected.GET("/transactions/:projectID", transactionHandler.GetByProject)

	// Проекты
	protected.POST("/projects", projectHandler.Create)
	// protected.GET("/projects", projectHandler.GetAll)
	protected.GET("/projects", projectHandler.GetAllFiltered)
	protected.GET("/projects/:id", projectHandler.GetByID)
	protected.PUT("/projects/:id", projectHandler.Update)
	protected.DELETE("/projects/:id", projectHandler.Delete)
	protected.POST("/projects/bulk-delete", projectHandler.BulkDelete)

	log.Println("Server starting on :8080")

	// Выведем реальные маршруты
	for _, route := range r.Routes() {
		log.Printf("  %s %s", route.Method, route.Path)
	}

	r.Run(":8080")
}
