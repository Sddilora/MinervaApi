package presenter

import (
	"minerva_api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Research is the presenter object which will be taken in the request by Handler
type Research struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	TopicId string `json:"topic_id"`
}

// ResearchSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func ResearchSuccessResponse(data *entities.Research) *fiber.Map {

	newResearch := Research{
		ID:      data.ID,
		Title:   data.Title,
		Content: data.Content,
		TopicId: data.TopicID,
	}
	return &fiber.Map{
		"status": true,
		"data":   newResearch,
		"error":  nil,
	}
}

// ResearchesSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func ResearchesSuccessResponse(data *[]Research) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// ResearchErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ResearchErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
