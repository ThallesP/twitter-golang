package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thallesp/twitter-golang/usecase"
)

type DeleteTweetController struct {
	deleteTweetUseCase *usecase.DeleteTweetUseCase
}

func NewDeleteTweetController(deleteTweetUseCase *usecase.DeleteTweetUseCase) *DeleteTweetController {
	return &DeleteTweetController{
		deleteTweetUseCase: deleteTweetUseCase,
	}
}

func (d *DeleteTweetController) Handle(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)
	tweetId := c.Params("tweetId")

	return d.deleteTweetUseCase.Execute(&usecase.DeleteTweetInput{
		TweetId: tweetId,
		UserId:  sub,
	})
}
