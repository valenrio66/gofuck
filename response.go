package gofuck

import (
	"github.com/gofiber/fiber/v2"
)

func WriteResponse[T any](ctx *fiber.Ctx, code int, success bool, status string, data T) error {
	response := Response[T]{
		Code:    code,
		Success: success,
		Status:  status,
		Data:    data,
	}
	return ctx.Status(code).JSON(response)
}

func WriteErrorResponse(ctx *fiber.Ctx, code int, status string) error {
	response := Response[interface{}]{
		Code:    code,
		Success: false,
		Status:  status,
		Data:    nil,
	}
	return ctx.Status(code).JSON(response)
}
