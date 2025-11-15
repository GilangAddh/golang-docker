package router

import (
	"backend/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/api")

	api.Get("/users", userHandler.GetAll)
	api.Get("/users/:id", userHandler.GetByID)
	api.Post("/users", userHandler.Create)
	api.Put("/users/:id", userHandler.Update)
	api.Delete("/users/:id", userHandler.Delete)
}
