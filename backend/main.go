package main

import (
	myhttp "backend/http"
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

	// Err Handler
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if e, ok := err.(*myhttp.RequestError); ok {
				return c.Status(e.StatusCode).JSON(fiber.Map{
					"status":  e.StatusCode,
					"message": e.Message,
					"errors":  e.Errors,
				})
			}

			// Default
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Internal Server Error",
				"errors":  err.Error(),
			})
		},
	})

	// Setup
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRoutes(app, userHandler)

	log.Println("Server running at :8080")
	app.Listen(":8080")
}
