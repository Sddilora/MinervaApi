package presenter

import (
	"minerva_api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Topic is the presenter object which will be taken in the request by Handler
type Topic struct {
	ID             string `json:"id"`
	TopicTitle     string `json:"title"`
	TopicCreatorID string `json:"author_id"`
}

// TopicSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func TopicSuccessResponse(data *entities.Topic) *fiber.Map {

	newTopic := Topic{
		ID:             data.ID,
		TopicTitle:     data.Title,
		TopicCreatorID: data.AuthorID,
	}
	return &fiber.Map{
		"status": true,
		"data":   newTopic,
		"error":  nil,
	}
}

// TopicsSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func TopicssSuccessResponse(data *[]Topic) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// TopicErrorResponse is the ErrorResponse that will be passed in the response by Handler
func TopicErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
