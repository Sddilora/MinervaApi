package entities

import "time"

type Topic struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	AuthorID     string    `json:"author_id"`
	UpdReqUserID string    `json:"update_request_user_id"` //This parameter takes the user id's of the person who has a update request
	CreatedAt    time.Time `json:"createdAt" `
	UpdatedAt    time.Time `json:"updatedAt" `
}
