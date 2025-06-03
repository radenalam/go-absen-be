package middleware

import (
	"go-absen-be/internal/model"
	"go-absen-be/internal/usecase"

	"strings"

	"github.com/gofiber/fiber/v2"
)


func NewAuth(userUserCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization", "NOT_FOUND")
		request := &model.VerifyUserRequest{Token : strings.TrimPrefix(authHeader, "Bearer ")}
		userUserCase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUserCase.Log.Warnf("(middleware) Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userUserCase.Log.Debugf("user_id : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
