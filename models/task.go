package models

import (
	"time"
)

type Task struct {
	ID	string	`json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}


type Tasks []Task