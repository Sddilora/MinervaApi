package requests

import (
	create "api/FireBase"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

var _, firebaseApp = create.NewFireStore()

var fireStoreClient, err = firebaseApp.Firestore(context.Background())

func PostTopicHandler(c *fiber.Ctx) error {

	userCol := fireStoreClient.Collection("Topic")

	var newTopic struct { // Parse request body
		Topic string `json:"topic"`
	}

	if err := c.BodyParser(&newTopic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	topic := map[string]interface{}{ // Create new Topic
		"topic": newTopic.Topic,
	}

	_, err := userCol.Doc(newTopic.Topic).Set(context.Background(), &topic) // Add Topic to Firestore
	if err != nil {
		log.Printf("Failed to add Topic to Firestore: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add Topic to Firestore",
		})
	}

	resp := fiber.Map{
		"message": "Topic created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func PostResearchHandler(c *fiber.Ctx) error {

	userCol := fireStoreClient.Collection("Research")

	var newResearch struct { // Parse request body
		ResearchHeader      string `json:"research_header"`
		ResearchContent     string `json:"research_content"`
		ResearchCreator     string `json:"research_creator"`
		ResearchContributor string `json:"research_contributor"`
	}

	if err := c.BodyParser(&newResearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	research := map[string]interface{}{ // Create new research
		"research_header":      newResearch.ResearchHeader,
		"research_content":     newResearch.ResearchContent,
		"research_creator":     newResearch.ResearchCreator,
		"research_contributor": newResearch.ResearchContributor,
	}

	_, err := userCol.Doc(newResearch.ResearchHeader).Set(context.Background(), &research) // Add user to Firestore
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add Research to Firestore",
		})
	}

	resp := fiber.Map{
		"message": "Research created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}
