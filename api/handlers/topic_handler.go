package handlers

import (
	"context"
	"errors"
	"fmt"
	"minerva_api/api/presenter"
	"minerva_api/pkg/entities"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

func GetTopics(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody *entities.Topic
		//Body Parser,Error Handler
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TopicErrorResponse(err))
		}

		if requestBody.AuthorID == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TopicErrorResponse(errors.New(
				"Please specify title and author")))
		}

		var topics []presenter.Topic
		// Creates a reference to a collection to Topic path.
		userCol := Client.Collection("Topic")

		//Documents returns an iterator over the query's resulting documents.
		query := userCol.Documents(context.Background())

		for {
			//Next returns the next result. Its second return value is iterator.Done if there are no more results
			doc, err := query.Next()
			if err == iterator.Done {
				break
			}
			var topic presenter.Topic
			doc.DataTo(&topic)
			topics = append(topics, topic)
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   &topics,
			"err":    nil,
		})
	}
}

func AddTopic(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody *entities.Topic

		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TopicErrorResponse(err))
		}

		if requestBody.AuthorID == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TopicErrorResponse(errors.New(
				"Please specify title and author")))
		}

		// Creates a reference to a collection to Topic path.
		userCol := Client.Collection("Topic")
		//Creates unıq id for document
		docRefUID := userCol.NewDoc()
		// Add Topic to Firestore
		_, err := docRefUID.Set(context.Background(), &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TopicErrorResponse(err))
		}

		return c.JSON(presenter.TopicSuccessResponse(requestBody))
	}
}
func UpdateTopic(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody *entities.Topic

		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TopicErrorResponse(err))
		}

		if requestBody.AuthorID == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TopicErrorResponse(errors.New(
				"Please specify title and author")))
		}

		//Indicates the document's path
		docRefPath := fmt.Sprintf("Topic/%s", requestBody.ID)
		//Indicates to document
		userDoc := Client.Doc(docRefPath)

		//Updates the given parameters at the Document
		_, err := userDoc.Update(context.Background(), []firestore.Update{
			{Path: "topic_name", Value: requestBody.Title},
		})

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TopicErrorResponse(err))
		}

		return c.JSON(presenter.TopicSuccessResponse(requestBody))
	}
}
func RemoveTopic(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody *entities.Topic

		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TopicErrorResponse(err))
		}

		if requestBody.AuthorID == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TopicErrorResponse(errors.New(
				"Please specify title and author")))
		}

		//Indicates the document's path
		docRefPath := fmt.Sprintf("Topic/%s", requestBody.ID)
		//Indicates to document
		userDoc := Client.Doc(docRefPath)
		_, err := userDoc.Delete(context.Background())
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TopicErrorResponse(err))
		}

		return c.JSON(presenter.TopicSuccessResponse(requestBody))
	}
}
