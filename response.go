package gofuck

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResp[T any](ctx *fiber.Ctx, code int, success bool, status string, data T) error {
	response := Response[T]{
		Code:    code,
		Success: success,
		Status:  status,
		Data:    data,
	}
	return ctx.Status(code).JSON(response)
}

func ErrorResp(ctx *fiber.Ctx, code int, status string) error {
	response := Response[interface{}]{
		Code:    code,
		Success: false,
		Status:  status,
		Data:    nil,
	}
	return ctx.Status(code).JSON(response)
}

func ParseBody[T any](ctx *fiber.Ctx, dest *T) error {
	if err := ctx.BodyParser(dest); err != nil {
		return ErrorResp(ctx, fiber.StatusBadRequest, "Body Request tidak Sesuai")
	}
	return nil
}

func IdParams(ctx *fiber.Ctx, param string) (string, error) {
	id := ctx.Query(param)
	if id == "" {
		return "", ErrorResp(ctx, fiber.StatusBadRequest, "ID tidak ditemukan dalam parameter")
	}
	return id, nil
}
