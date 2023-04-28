package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	u "github.com/thallesp/twitter-golang/usecase"
	"github.com/thallesp/twitter-golang/utils"
)

type createTweetDTO struct {
	Content string `json:"content" validate:"required,min=1,max=280"`
}

type CreateTweetController struct {
	createTweetUseCase *u.CreateTweetUseCase
}

func NewCreateTweetController(createTweetUseCase *u.CreateTweetUseCase) *CreateTweetController {
	return &CreateTweetController{
		createTweetUseCase: createTweetUseCase,
	}
}

func (c *CreateTweetController) Handle(fc *fiber.Ctx) error {
	token := fc.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)

	createTweetDTO := new(createTweetDTO)

	if err := fc.BodyParser(createTweetDTO); err != nil {
		return err
	}

	errors := utils.ValidateStruct(createTweetDTO)

	if errors != nil {
		return fc.Status(fiber.StatusBadRequest).JSON(errors)
	}

	tweet, err := c.createTweetUseCase.Execute(u.CreateTweetInput{
		Content: createTweetDTO.Content,
		UserId:  sub,
	})

	if err != nil {
		return err
	}

	return fc.Status(200).JSON(tweet)
}
