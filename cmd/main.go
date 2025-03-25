package main

import (
	"log"
	"todo/config"
	"todo/models"
	"todo/repositories"
	"todo/routes"
	"todo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Loaded config: %+v\n", cfg)

	db, err := repositories.NewDB(cfg)

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the database")

	if err := db.AutoMigrate(&models.Todo{}, &models.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v\n", err)
	}
	log.Println("Database migration completed")

	userRepo := repositories.NewUserRepository(db)
	todoRepo := repositories.NewTodoRepository(db)
	
	emailService, err := services.NewEmailService(config.LoadEmailConfig())
	if err != nil {
		log.Fatalf("Email service failed: %v", err)
	}

	services := routes.Services{
		AuthService: services.NewAuthService(userRepo, emailService),
		TodoService: services.NewTodoService(todoRepo),
		UserService: services.NewUserService(userRepo),
	}

	// Create router
	router := gin.Default()

	// Setup ALL routes in one place
	routes.SetupAllRoutes(router, services)

	

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}