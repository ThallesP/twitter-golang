package usecase

import (
	"github.com/thallesp/twitter-golang/entity"
	"github.com/thallesp/twitter-golang/repository"
)

type EditTweetUseCase struct {
	tweetsRepository repository.TweetsRepository
}

func NewEditTweetUseCase(tweetsRepository repository.TweetsRepository) *EditTweetUseCase {
	return &EditTweetUseCase{
		tweetsRepository: tweetsRepository,
	}
}

type EditTweetInput struct {
	Content string
	TweetId string
	UserId  string
}

func (e *EditTweetUseCase) Execute(editTweetInput *EditTweetInput) (*entity.Tweet, error) {
	tweet, err := e.tweetsRepository.FindByID(editTweetInput.TweetId)

	if err != nil {
		return nil, NewException("tweet not found", 404, "tweet_not_found")
	}

	if editTweetInput.UserId != tweet.UserId {
		return nil, NewException("You can't edit other users' tweets", 403, "cant_edit_others_tweets")
	}

	tweet.Content = editTweetInput.Content
	return e.tweetsRepository.Update(editTweetInput.TweetId, tweet)
}
