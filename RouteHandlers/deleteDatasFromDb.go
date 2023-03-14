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
		//	TopicCreatorID string `json:"topic_creator_id"` bunu dbden alcam
		DelReqUserID string `json:"delete_request_user_id"` //This parameter takes the user id's of the person who has a deletion request
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
	//Deletes the  document if delete request owner and the topic creator are same person(same ID)

	firestoreGet, err := userDoc.Get(context.Background())
	if err != nil {
		log.Print("couldn't get user document")
	}
	firestoreData := firestoreGet.Data()

	TopicCreatorID := firestoreData["topic_creator_id"]

	if TopicCreatorID == deleteTopic.DelReqUserID {
		_, err := userDoc.Delete(context.Background())
		if err != nil {
			log.Printf("Error while deleting document: %v", err)
		} else {

			resp := fiber.Map{
				"message": "Topic Deleted successfully",
			}
			return c.Status(fiber.StatusCreated).JSON(resp)
		}
	} else {
		resp := fiber.Map{
			"message": "Topics are deleted only by the owner of the Topic",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	return nil
}

// PutResearchHandler handles the request and gives the proper response
func DeleteResearchHandler(c *fiber.Ctx) error {
	// Parse request body
	var deleteResearch struct {
		ResearchTopicID string `json:"research_topic_id"`
		ResearchID      string `json:"research_id"`
		//ResearchCreatorID string `json:"research_creator_id"`
		DelReqUserID string `json:"delete_request_user_id"` //This parameter takes the user id's of the person who has a deletion request
	}
	//Body parser, Error Handler
	if err := c.BodyParser(&deleteResearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//Indicates the document's path
	docRefPath := fmt.Sprintf("Topic/%s/%s/%s", deleteResearch.ResearchTopicID, deleteResearch.ResearchID, deleteResearch.ResearchID)

	//Indicates to document
	userDoc := fireStoreClient.Doc(docRefPath)

	firestoreGet, err := userDoc.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get document: %v", err)
	}
	firestoreData := firestoreGet.Data()

	ResearchCreatorID := firestoreData["research_creator_id"]

	if ResearchCreatorID == deleteResearch.DelReqUserID {
		//Deletes the document
		_, err := userDoc.Delete(context.Background())
		if err != nil {
			log.Printf("Error deleting document: %v", err)
		} else {

			resp := fiber.Map{
				"message": "Research Deleted successfully",
			}
			return c.Status(fiber.StatusCreated).JSON(resp)
		}
	} else {
		resp := fiber.Map{
			"message": "Researches are deleted only by the owner of the Research",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	return nil
}
