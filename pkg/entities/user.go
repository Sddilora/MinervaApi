package entities

import "time"

type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	PhotoUrl      string    `json:"photo_url"`
	UpdReqUserJWT string    `json:"update_request_user_jwt"` //This parameter takes the user jwt's of the person who has a update request
	CreatedAt     time.Time `json:"createdAt" `
	UpdatedAt     time.Time `json:"updatedAt"`
}
