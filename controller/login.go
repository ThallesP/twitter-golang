package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thallesp/twitter-golang/usecase"
	"github.com/thallesp/twitter-golang/utils"
)

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginController struct {
	loginUseCase *usecase.LoginUseCase
}

func NewLoginController(loginUseCase *usecase.LoginUseCase) *LoginController {
	return &LoginController{
		loginUseCase: loginUseCase,
	}
}

func (l *LoginController) Handle(c *fiber.Ctx) error {
	loginDTO := new(LoginDTO)

	if err := c.BodyParser(loginDTO); err != nil {
		return err
	}

	errors := utils.ValidateStruct(loginDTO)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	login, err := l.loginUseCase.Execute(usecase.LoginInput{
		Email:    loginDTO.Email,
		Password: loginDTO.Password,
	})

	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    login.Token,
		HTTPOnly: true,
	})

	return c.Status(200).JSON(login)
}
