package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// PutTopicHandler handles the request and gives the proper response
func DeleteTopicHandler(c *fiber.Ctx) error {
	// Parse request body
	var deleteTopic struct {
		TopicID string `json:"topic_id"`
	}
	//Body parser, Error Handler
	if err := c.BodyParser(&deleteTopic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//Indicates the document's path
	docRefPath := fmt.Sprintf("Topic/%s", deleteTopic.TopicID)
	//Indicates to document
	userDoc := fireStoreClient.Doc(docRefPath)
	//Deletes the  document
	_, err := userDoc.Delete(context.Background())
	if err != nil {
		log.Printf("Error while deleting document: %v", err)
	} else {

		resp := fiber.Map{
			"message": "Topic Deleted successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	}
	return nil
}

// PutResearchHandler handles the request and gives the proper response
func DeleteResearchHandler(c *fiber.Ctx) error {
	// Parse request body
	var deleteResearch struct {
		ResearchTopicID string `json:"research_topic_id"`
		ResearchUID     string `json:"research_uid"`
	}
	//Body parser, Error Handler
	if err := c.BodyParser(&deleteResearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//Indicates the document's path
	docRefPath := fmt.Sprintf("Topic/%s/%s/%s", deleteResearch.ResearchTopicID, deleteResearch.ResearchUID, deleteResearch.ResearchUID)
	//Indicates to document
	userDoc := fireStoreClient.Doc(docRefPath)
	//Deletes the document
	_, err := userDoc.Delete(context.Background())
	if err != nil {
		log.Printf("Error deleting document: %v", err)
	} else {

		resp := fiber.Map{
			"message": "Topic Deleted successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	}
	return nil
}
