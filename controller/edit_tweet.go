package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thallesp/twitter-golang/usecase"
	"github.com/thallesp/twitter-golang/utils"
)

type editTweetDTO struct {
	Content string `json:"content" validate:"required,min=1,max=280"`
}

type EditTweetController struct {
	editTweetUseCase *usecase.EditTweetUseCase
}

func NewEditTweetController(editTweetUseCase *usecase.EditTweetUseCase) *EditTweetController {
	return &EditTweetController{
		editTweetUseCase: editTweetUseCase,
	}
}

func (e *EditTweetController) Handle(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)

	tweetId := c.Params("tweetId")
	editTweetDTO := new(editTweetDTO)

	if err := c.BodyParser(editTweetDTO); err != nil {
		return err
	}

	errors := utils.ValidateStruct(editTweetDTO)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	editTweet, err := e.editTweetUseCase.Execute(&usecase.EditTweetInput{
		Content: editTweetDTO.Content,
		TweetId: tweetId,
		UserId:  sub,
	})

	if err != nil {
		return err
	}

	return c.Status(200).JSON(editTweet)
}
