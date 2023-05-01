package entity

import "time"

type Tweet struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	UserId    string    `json:"userId"`
	User      *User     `json:"user" pg:"rel:has-one"`
	CreatedAt time.Time `json:"createdAt"`
}
