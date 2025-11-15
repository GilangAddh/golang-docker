package handler

import (
	"backend/http"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	data, err := h.userService.GetAll(ctx.Context())
	if err != nil {
		return http.ResponseApi(ctx, fiber.StatusOK, "Success get all users", data, nil)
	}
	return http.ResponseApi(ctx, fiber.StatusOK, "Success get all users", data, nil)

}
