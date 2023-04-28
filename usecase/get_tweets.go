package usecase

import (
	"github.com/thallesp/twitter-golang/entity"
	"github.com/thallesp/twitter-golang/repository"
)

type GetTweetsUseCase struct {
	tweetsRepository repository.TweetsRepository
}

func NewGetTweetsUseCase(tweetsRepository repository.TweetsRepository) *GetTweetsUseCase {
	return &GetTweetsUseCase{
		tweetsRepository: tweetsRepository,
	}
}

func (g *GetTweetsUseCase) Execute() (*[]entity.Tweet, error) {
	tweets, err := g.tweetsRepository.Find(*repository.NewFindInput(1))

	if err != nil {
		return nil, err
	}

	return tweets, nil
}
