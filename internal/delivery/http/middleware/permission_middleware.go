package middleware

import (
	"go-absen-be/internal/model"
	"go-absen-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func RequirePermission(permissionName string, userUseCase *usecase.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authUser := c.Locals("auth").(*model.Auth)
		hasPerm, err := userUseCase.HasPermission(c.UserContext(), authUser.ID, permissionName)
		if err != nil || !hasPerm {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden – you don't have permission: " + permissionName,
			})
		}
		return c.Next()
	}
}