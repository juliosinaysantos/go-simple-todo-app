package middlewares

import "github.com/gofiber/fiber/v2"

func ParseIDParam(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Param id must be an unsigned integer")
	}

	if id <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Param id must be an unsigned integer")
	}
	ctx.Locals("id", id)
	return ctx.Next()
}
