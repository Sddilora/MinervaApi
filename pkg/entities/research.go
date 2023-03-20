package entities

import "time"

// Research Constructs your research model under entities.
type Research struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	Contributor  string    `json:"contributor"`
	TopicID      string    `json:"topic_id"`
	UpdReqUserID string    `json:"update_request_user_id"` //This parameter takes the user id's of the person who has a update request
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}

// // DeleteRequest struct is used to parse Delete Requests for Researches
// type DeleteResearchRequest struct {
// 	ID string `json:"id"`
// }
