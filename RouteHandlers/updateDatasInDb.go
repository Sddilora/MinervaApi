package handlers

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
)

// PutTopicHandler handles the request and gives the proper response
func PutTopicHandler(c *fiber.Ctx) error {

	// Parse request body
	var updatedTopic struct {
		TopicID      string `json:"topic_id"`
		TopicName    string `json:"topic_name"`
		UpdReqUserID string `json:"update_request_user_id"` //This parameter takes the user id's of the person who has a update request

	}
	//Body parser, Error Handler

	if err := c.BodyParser(&updatedTopic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//Indicates the document's path
	docRefPath := fmt.Sprintf("Topic/%s", updatedTopic.TopicID)
	//Indicates to document
	userDoc := fireStoreClient.Doc(docRefPath)

	dataGet, err := userDoc.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get document: %v", err)
	}

	//read all document
	firestoreData := dataGet.Data()

	TopicCreatorID := firestoreData["topic_creator_id"]

	if TopicCreatorID == updatedTopic.UpdReqUserID {
		//Updates the given parameters at the Document
		_, err = userDoc.Update(context.Background(), []firestore.Update{
			{Path: "topic_name", Value: updatedTopic.TopicName},
		})
		if err != nil {
			log.Printf("Error updating document: %v", err)
		} else {

			resp := fiber.Map{
				"message": "Topic Updated successfully",
			}
			return c.Status(fiber.StatusCreated).JSON(resp)
		}
	} else {
		resp := fiber.Map{
			"message": " Topic can only updated by topic creator",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	return nil
}

// PutResearchHandler handles the request and gives the proper response
func PutResearchHandler(c *fiber.Ctx) error {

	// Parse request body
	var updatedResearch struct {
		ResearchTopicID string `json:"research_topic_id"`
		ResearchHeader  string `json:"research_header"`
		ResearchContent string `json:"research_content"`
		//ResearchContributor string `json:"research_contributor"`
		ResearchID   string `json:"research_id"`
		UpdReqUserID string `json:"update_request_user_id"` //This parameter takes the user id's of the person who has a update request

	}
	//Body parser, Error Handler
	if err := c.BodyParser(&updatedResearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//Indicates the document's path
	docRefPath := fmt.Sprintf("Topic/%s/%s/%s", updatedResearch.ResearchTopicID, updatedResearch.ResearchID, updatedResearch.ResearchID)
	//Indicates to document
	userDoc := fireStoreClient.Doc(docRefPath)

	dataGet, err := userDoc.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get document: %v", err)
	}

	//read all document
	firestoreData := dataGet.Data()

	ResearchCreatorID := firestoreData["research_creator_id"]

	if ResearchCreatorID == updatedResearch.UpdReqUserID {
		//Updates the given parameters at the Document
		_, err = userDoc.Update(context.Background(), []firestore.Update{
			{Path: "research_header", Value: updatedResearch.ResearchHeader},
			{Path: "research_content", Value: updatedResearch.ResearchContent},
			//{Path: "research_contributor", Value: updatedResearch.ResearchContributor},
		})
		if err != nil {
			log.Printf("Error updating document: %v", err)
		} else {

			resp := fiber.Map{
				"message": "Topic Updated successfully",
			}
			return c.Status(fiber.StatusCreated).JSON(resp)
		}
	} else {
		resp := fiber.Map{
			"message": " Research can only updated by research creator", //ilerleyen zamanlarda buraya contributor için de izin ekleyebiliriz
		}
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	return nil
}
