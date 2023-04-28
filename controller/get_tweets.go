package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thallesp/twitter-golang/entity"
	"github.com/thallesp/twitter-golang/usecase"
)

type GetTweetsController struct {
	getTweetsUsecase *usecase.GetTweetsUseCase
}

type GetTweetsResponse []struct {
	*entity.Tweet
}

func NewGetTweetsController(getTweetsUseCase *usecase.GetTweetsUseCase) *GetTweetsController {
	return &GetTweetsController{
		getTweetsUsecase: getTweetsUseCase,
	}
}

func (g *GetTweetsController) Handle(c *fiber.Ctx) error {
	tweets, err := g.getTweetsUsecase.Execute()

	if err != nil {
		return err
	}

	fmt.Println(tweets)
	return c.Status(200).JSON(tweets)
}
