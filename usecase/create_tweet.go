package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/thallesp/twitter-golang/entity"
	"github.com/thallesp/twitter-golang/repository"
)

type CreateTweetInput struct {
	Content string
	UserId  string
}

type CreateTweetUseCase struct {
	tweetsRepository repository.TweetsRepository
}

func NewCreateTweetUseCase(tweetsRepository repository.TweetsRepository) (c *CreateTweetUseCase) {
	return &CreateTweetUseCase{
		tweetsRepository: tweetsRepository,
	}
}

func (c *CreateTweetUseCase) Execute(createTweetInput CreateTweetInput) (*entity.Tweet, error) {
	tweet := &entity.Tweet{
		Id:        uuid.New().String(),
		Content:   createTweetInput.Content,
		CreatedAt: time.Now(),
		UserId:    createTweetInput.UserId,
	}

	err := c.tweetsRepository.Create(tweet)

	if err != nil {
		return nil, err
	}

	return tweet, nil
}
