package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thallesp/twitter-golang/usecase"
	"github.com/thallesp/twitter-golang/utils"
)

type RegisterDTO struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8"`
}

type RegisterController struct {
	registerUseCase *usecase.RegisterUseCase
}

func NewRegisterController(registerUseCase *usecase.RegisterUseCase) *RegisterController {
	return &RegisterController{
		registerUseCase: registerUseCase,
	}
}

func (r *RegisterController) Handle(c *fiber.Ctx) error {
	registerDTO := new(RegisterDTO)

	if err := c.BodyParser(registerDTO); err != nil {
		return err
	}

	errors := utils.ValidateStruct(registerDTO)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user, err := r.registerUseCase.Execute(&usecase.RegisterInput{
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
	})

	if err != nil {
		return err
	}

	return c.Status(200).JSON(user)
}
