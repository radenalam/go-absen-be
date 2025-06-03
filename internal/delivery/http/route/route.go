package route

import (
	"go-absen-be/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *controller.UserController

}

func (c *RouteConfig) Setup() {
	c.App.Get("/api/users", c.UserController.List)

}