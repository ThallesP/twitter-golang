package usecase

import "github.com/thallesp/twitter-golang/repository"

type DeleteTweetUseCase struct {
	tweetsRepository repository.TweetsRepository
}

func NewDeleteTweetUseCase(tweetsRepository repository.TweetsRepository) *DeleteTweetUseCase {
	return &DeleteTweetUseCase{
		tweetsRepository: tweetsRepository,
	}
}

type DeleteTweetInput struct {
	TweetId string
	UserId  string
}

func (d *DeleteTweetUseCase) Execute(deleteTweetInput *DeleteTweetInput) (err error) {
	tweet, err := d.tweetsRepository.FindByID(deleteTweetInput.TweetId)

	if err != nil {
		return NewException("tweet not found", 404, "tweet_not_found")
	}

	if tweet.UserId != deleteTweetInput.UserId {
		return NewException("You can't delete other users' tweets", 403, "cant_delete_other_tweets")
	}

	err = d.tweetsRepository.Delete(deleteTweetInput.TweetId)

	return
}
