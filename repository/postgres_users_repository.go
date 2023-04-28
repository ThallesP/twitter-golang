package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/thallesp/twitter-golang/entity"
)

type PostgresUsersRepository struct {
	db *pg.DB
}

func NewPostgresUsersRepository(db *pg.DB) *PostgresUsersRepository {
	return &PostgresUsersRepository{
		db: db,
	}
}

func (p *PostgresUsersRepository) Create(user *entity.User) error {
	_, err := p.db.Model(user).Insert()

	return err
}

func (p *PostgresUsersRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{Email: email}
	err := p.db.Model(user).Where("email = ?", email).Select()

	if err != nil {
		return nil, err
	}

	return user, nil
}
