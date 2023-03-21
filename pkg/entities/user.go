package entities

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	PhotoUrl  string    `json:"photo_url"`
	CreatedAt time.Time `json:"createdAt" `
	UpdatedAt time.Time `json:"updatedAt"`
}
