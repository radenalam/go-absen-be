package controller

import (
	"go-absen-be/internal/model"
	"go-absen-be/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log: logger,
		UseCase: useCase,
	}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.LoginResponse]{Data: response})
}

func (c *UserController) List(ctx *fiber.Ctx) error {
	request := &model.SearchRequest{
		Page: ctx.QueryInt("page", 1),
		Size: ctx.QueryInt("size", 10),
	}

	response, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to list product stock : %+v ", err)
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}
	
	return ctx.JSON(model.WebResponse[[]model.UserResponse]{		
		Data:   response,
		Paging: paging,
	})
}