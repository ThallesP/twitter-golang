package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/thallesp/twitter-golang/entity"
)

type PostgresTweetsRepository struct {
	db *pg.DB
	TweetsRepository
}

func NewPostgresTweetsRepository(db *pg.DB) *PostgresTweetsRepository {
	return &PostgresTweetsRepository{
		db: db,
	}
}

func (p *PostgresTweetsRepository) Create(tweet *entity.Tweet) error {
	_, err := p.db.Model(tweet).Relation("User").Insert()

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresTweetsRepository) Find(pagination FindInput) (*[]entity.Tweet, error) {
	tweets := []entity.Tweet{}
	err := p.db.Model(&tweets).Select()

	if err != nil {
		return nil, err
	}

	return &tweets, nil
}
