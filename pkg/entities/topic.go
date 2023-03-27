package entities

import "time"

type Topic struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt" `
	UpdatedAt time.Time `json:"updatedAt" `
}
