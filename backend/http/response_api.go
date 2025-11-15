package http

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func ResponseApi(
	c *fiber.Ctx,
	status int,
	message string,
	data interface{},
	paginations ...interface{},
) error {
	response := fiber.Map{
		"status":  status,
		"message": message,
		"payload": data,
	}

	if data != nil && reflect.TypeOf(data).Kind() == reflect.Slice && reflect.ValueOf(data).Len() == 0 {
		response["payload"] = []interface{}{}
	}

	if paginations != nil {
		response["pagination"] = paginations[0]
	}

	return c.Status(status).JSON(response)
}

type ResponsePagination struct {
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
}
