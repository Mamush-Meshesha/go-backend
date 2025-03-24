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

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v\n", err)
	}
	log.Println("Database migration completed")

	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)

	r := gin.Default()

	// Register routes before running the server
	routes.SetupTodoRoutes(r, todoService)

	// Add root route before starting the server
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the todo app"})
	})

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}