package router

import (
	"backend/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/api")

	api.Get("/users", userHandler.GetAll)
}
