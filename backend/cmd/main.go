package main

import (
	"log"

	"finances/internal/database"
	"finances/internal/handlers"
	"finances/internal/middleware"
	"finances/internal/repository"
	"finances/internal/service"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	authService := service.NewAuthService(userRepo, "supersecretkey")
	authMiddleware := middleware.AuthMiddleware("supersecretkey")
	authHandler := handlers.NewAuthHandler(authService)
	// taskRepo := repository.NewTaskRepository(db)
	// taskService := service.NewTaskService(taskRepo)
	// taskHandler := handlers.NewTaskHandler(taskService)

	r := gin.Default()
	r.Use(cors.Default())

	protected := r.Group("/api")
	protected.Use(authMiddleware)

	// protected.POST("/tasks", taskHandler.Create)
	// protected.GET("/tasks", taskHandler.GetUserTasks)
	// protected.PATCH("/tasks/:id/status", taskHandler.UpdateStatus)
	// protected.DELETE("/tasks/:id", taskHandler.Delete)

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(200, gin.H{
			"user_id": userID,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})
	r.POST("/register", userHandler.Register)
	r.POST("/login", authHandler.Login)
	r.Run(":8080")
}
