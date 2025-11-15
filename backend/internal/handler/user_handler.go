package handler

import (
	"backend/http"
	"backend/internal/app"
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
		return http.NewRequestError(fiber.StatusInternalServerError, "Failed to get all users", err)
	}
	return http.ResponseApi(ctx, fiber.StatusOK, "Success get all users", data, nil)
}

func (h *UserHandler) GetByID(ctx *fiber.Ctx) error {
	dataID, err := ctx.ParamsInt("id")
	if err != nil {
		return http.NewRequestError(fiber.StatusBadRequest, "Invalid user ID", err)
	}

	user, err := h.userService.GetByID(ctx.Context(), uint(dataID))
	if err != nil {
		return http.NewRequestError(fiber.StatusInternalServerError, "Failed to get user detail", err)
	}

	return http.ResponseApi(ctx, fiber.StatusOK, "Success get user detail", user, nil)
}

func (h *UserHandler) Create(ctx *fiber.Ctx) error {
	var input app.CreateUserDTO
	if err := ctx.BodyParser(&input); err != nil {
		return http.NewRequestError(fiber.StatusBadRequest, "Invalid input", err)
	}

	result, err := h.userService.Create(ctx.Context(), input)
	if err != nil {
		return http.NewRequestError(fiber.StatusInternalServerError, "Failed to create user", err)
	}
	return http.ResponseApi(ctx, fiber.StatusOK, "Success create user", result, nil)
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	var input app.UpdateUserDTO
	if err := ctx.BodyParser(&input); err != nil {
		return http.NewRequestError(fiber.StatusBadRequest, "Invalid input", err)
	}

	dataID, err := ctx.ParamsInt("id")
	input.ID = uint(dataID)
	if err != nil {
		return http.NewRequestError(fiber.StatusBadRequest, "Invalid user ID", err)
	}

	err = h.userService.Update(ctx.Context(), uint(dataID), input)
	if err != nil {
		return http.NewRequestError(fiber.StatusInternalServerError, "Failed to update user", err)
	}

	return http.ResponseApi(ctx, fiber.StatusOK, "Success update user", nil, nil)
}

func (h *UserHandler) Delete(ctx *fiber.Ctx) error {
	dataID, err := ctx.ParamsInt("id")
	if err != nil {
		return http.NewRequestError(fiber.StatusBadRequest, "Invalid user ID", err)
	}
	err = h.userService.Delete(ctx.Context(), uint(dataID))
	if err != nil {
		return http.NewRequestError(fiber.StatusInternalServerError, "Failed to delete user", err)
	}
	return http.ResponseApi(ctx, fiber.StatusOK, "Success delete user", nil, nil)
}
