package entity

import (
	"time"
)

type User struct {
	Id              string    `json:"id"`
	FullName        string    `json:"fullName" pg:",notnull"`
	Username        string    `json:"username" pg:",notnull"`
	Email           string    `json:"email" pg:",notnull"`
	PasswordHash    string    `json:"-" pg:",notnull"`
	ProfileImageURL string    `json:"profileImageURL" pg:",notnull"`
	CreatedAt       time.Time `json:"createdAt" pg:",notnull"`
}
