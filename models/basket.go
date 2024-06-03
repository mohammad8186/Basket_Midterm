package models

import (
	"time"
)

type Basket struct {
	ID        string    `json:"id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Data      []byte    `json:"data"`
	State     string    `json:"state"`
}
