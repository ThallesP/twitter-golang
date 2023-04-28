package repository

import "github.com/thallesp/twitter-golang/entity"

type UsersRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
