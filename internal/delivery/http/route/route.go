package route

import (
	"go-absen-be/internal/delivery/http/controller"
	"go-absen-be/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *controller.UserController
	AuthMiddleware    fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/auth/register", c.UserController.Register)
	c.App.Post("/api/auth/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)

	user := c.App.Group("/api/users")
	user.Post("/create",middleware.RequirePermission("create_user", c.UserController.UseCase), c.UserController.Create)
	user.Get("/list", middleware.RequirePermission("read_user", c.UserController.UseCase), c.UserController.List)
	user.Get("/:id", middleware.RequirePermission("read_user", c.UserController.UseCase), c.UserController.GetByID)
}