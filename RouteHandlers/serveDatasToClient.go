package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

func PostTopicsHandler(c *fiber.Ctx) error {

	// Creates a reference to a collection to Topic path.
	userCol := fireStoreClient.Collection("Topic")

	//Documents returns an iterator over the query's resulting documents.
	query := userCol.Documents(context.Background())

	for {
		//Next returns the next result. Its second return value is iterator.Done if there are no more results
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to iterate over collection: %v", err)
		}
		//Turn doc.Data into byte array
		docDataBytes, err := json.Marshal(doc.Data())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to Marshal docData",
			})
		}
		//Write appends docDataBytes into response body.
		c.Write(docDataBytes)

	}

	return nil

}

func PostResearchesHandler(c *fiber.Ctx) error {

	var wantedResearches struct {
		ResearchTopicId string `json:"research_topic_id"`
	}

	//Body Parser,Error Handler
	if err := c.BodyParser(&wantedResearches); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	docPath := fmt.Sprintf("Topic/%s", wantedResearches.ResearchTopicId)
	userDoc := fireStoreClient.Doc(docPath)

	//Documents returns an iterator over the query's resulting documents.
	query := userDoc.Collections(context.Background())

	for {
		//Next returns the next result. Its second return value is iterator.Done if there are no more results
		col, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to iterate over collection: %v", err)
		}
		//Next returns the next result. Its second return value is iterator.Done if there are no more results
		doc, err := (col.Documents(context.Background())).Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to iterate over collection: %v", err)
		}
		//Turn doc.Data into byte array
		docDataBytes, err := json.Marshal(doc.Data())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to Marshal docData",
			})
		}
		//Write appends docDataBytes into response body.
		c.Write(docDataBytes)

	}

	return nil
}
