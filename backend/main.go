package main

import (
	"backend/internal/app"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://admin:admin123@127.0.0.1:5433/employees_db"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Auto migrate
	db.AutoMigrate(&app.User{})

	// Dependency Injection
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Setup Fiber
	app := fiber.New()

	// Routes
	router.SetupRoutes(app, userHandler)

	log.Println("Server running at :8080")
	app.Listen(":8080")
}
