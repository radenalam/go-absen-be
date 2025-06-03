package config

import (
	"go-absen-be/internal/delivery/http/controller"
	"go-absen-be/internal/delivery/http/route"
	"go-absen-be/internal/repository"
	"go-absen-be/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)


type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repository.NewUserRepository(config.Log)
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	userController := controller.NewUserController(userUseCase, config.Log)

	routeConfig := route.RouteConfig{
		UserController:    userController,
		App:               config.App,
	}
	routeConfig.Setup()
}