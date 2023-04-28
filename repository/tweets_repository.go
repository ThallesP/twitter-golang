package repository

import "github.com/thallesp/twitter-golang/entity"

type FindInput struct {
	Page int
}

func NewFindInput(page int) *FindInput {
	return &FindInput{
		Page: 1,
	}
}

type TweetsRepository interface {
	Create(tweet *entity.Tweet) error
	Find(pagination FindInput) (*[]entity.Tweet, error)
}
