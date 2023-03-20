package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"minerva_api/api/presenter"
	"minerva_api/pkg/entities"
	"net/http"
	"os/exec"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

// AddResearch is handler/controller which creates Researches in the Database
func AddResearch(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody *entities.Research

		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}

		if requestBody.AuthorID == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(errors.New(
				"Please specify title and author")))
		}

		collectionName := UUID()

		// Creates a reference to a collection group to Research path.
		colPath := fmt.Sprintf("Topic/%s/%s", requestBody.TopicID, collectionName)

		collection := Client.Collection(colPath)

		// Set the ID field of the research to a new ID.
		requestBody.ID = collectionName

		// Set the created and updated timestamps for the research.
		requestBody.CreatedAt = time.Now()
		requestBody.UpdatedAt = time.Now()

		_, err := collection.Doc(collectionName).Set(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}

		return c.JSON(presenter.ResearchSuccessResponse(requestBody))
	}
}

// UpdateResearch is handler/controller which updates data of Researches in the database
func UpdateResearch(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody *entities.Research
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		requestBody.UpdatedAt = time.Now()
		//Indicates the document's path
		docRefPath := fmt.Sprintf("Topic/%s/%s/%s", requestBody.TopicID, requestBody.ID, requestBody.ID)
		//Indicates to document
		userDoc := Client.Doc(docRefPath)

		//Updates the given parameters at the Document
		_, err = userDoc.Update(context.Background(), []firestore.Update{
			{Path: "research_header", Value: requestBody.Title},
			{Path: "research_content", Value: requestBody.Content},
			//{Path: "research_contributor", Value: updatedResearch.ResearchContributor},
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		return c.JSON(presenter.ResearchSuccessResponse(requestBody))
	}
}

// RemoveResearch is handler/controller which removes Researches from the Database
func RemoveResearch(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {
		var requestBody *entities.Research
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		//Indicates the document's path
		docRefPath := fmt.Sprintf("Topic/%s/%s/%s", requestBody.TopicID, requestBody.ID, requestBody.ID)

		//Indicates to document
		userDoc := Client.Doc(docRefPath)

		//Deletes the document
		_, err = userDoc.Delete(context.Background())

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "removed successfully",
			"err":    nil,
		})
	}
}

// GetResearches is handler/controller which lists all Researches from the database
func GetResearches(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody entities.Research
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}

		var researches []presenter.Research
		//Indicates the document's path
		docPath := fmt.Sprintf("Topic/%s", requestBody.TopicID)
		//Indicates to document
		userDoc := Client.Doc(docPath)

		//Documents returns an iterator over the query's resulting documents.
		query := userDoc.Collections(context.Background())

		for {
			//Next returns the next result. Its second return value is iterator.Done if there are no more results
			col, err := query.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			//Next returns the next result. Its second return value is iterator.Done if there are no more results
			doc, err := (col.Documents(context.Background())).Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			var research presenter.Research
			doc.DataTo(&research)
			researches = append(researches, research)
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   &researches,
			"err":    nil,
		})

	}
}

// UUID generates a uniq id
func UUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Println(err)
	}
	return string(newUUID)
}
