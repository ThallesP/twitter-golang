package entity

import (
	"time"
)

type User struct {
	Id           string    `json:"id"`
	Email        string    `json:"email" pg:",notnull"`
	PasswordHash string    `json:"-" pg:",notnull"`
	CreatedAt    time.Time `json:"createdAt" pg:",notnull"`
}
