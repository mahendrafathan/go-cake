package models

import "time"

type Cake struct {
	Id          int        `json:"id,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      float32    `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
