package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mframadann/gourl/domain/auth/dto"
	"github.com/mframadann/gourl/domain/auth/services"
	"gorm.io/gorm"
)

type AuthController struct {
	AuthService services.AuthService
}

func (controller AuthController) Register(c echo.Context) error {
	user := new(dto.Register)

	if err := c.Bind(user); err != nil {
		return err
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	res := controller.AuthService.Register(*user)
	return c.JSON(res.StatusCode, res)
}

func (controller AuthController) SignIn(c echo.Context) error {
	user := new(dto.SignIn)

	if err := c.Bind(user); err != nil {
		return err
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	res := controller.AuthService.SignIn(*user)
	return c.JSON(res.StatusCode, res)
}

func NewAuthController(db *gorm.DB) AuthController {
	service := services.NewAuthService(db)
	controller := AuthController{
		AuthService: service,
	}

	return controller
}
