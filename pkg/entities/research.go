package entities

import "time"

// Research Constructs your research model under entities.
type Research struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	AuthorID      string    `json:"author_id"`
	Contributor   string    `json:"contributor"`
	TopicID       string    `json:"topic_id"`
	UpdReqUserJWT string    `json:"update_request_user_jwt"` //This parameter takes the user jwt's of the person who has a update request
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
