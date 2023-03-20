package presenter

import (
	"minerva_api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Research is the presenter object which will be taken in the request by Handler
type Research struct {
	ID                  string `json:"id"`
	ResearchTitle       string `json:"title"`
	ResearchContent     string `json:"content"`
	ResearchCreatorID   string `json:"author"`
	ResearchContributor string `json:"contributor"`
	ResearchTopicId     string `json:"topic_id"`
}

// ResearchSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func ResearchSuccessResponse(data *entities.Research) *fiber.Map {

	newResearch := Research{
		ResearchTitle:       data.Title,
		ResearchContent:     data.Content,
		ResearchCreatorID:   data.Author,
		ResearchContributor: data.Contributor,
		ResearchTopicId:     data.TopicID,
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
