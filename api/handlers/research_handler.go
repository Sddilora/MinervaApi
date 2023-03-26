package handlers

import (
	"context"
	"errors"
	"fmt"
	"minerva_api/api/presenter"
	"minerva_api/pkg/entities"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
				"please specify title and author")))
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
			{Path: "Title", Value: requestBody.Title},
			{Path: "Content", Value: requestBody.Content},
			{Path: "UpdatedAt", Value: requestBody.UpdatedAt},
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

// GetResearch is handler/controller which lists single Research from the database
func GetResearchByID(appFire *firebase.App) fiber.Handler {
	Client, _ := appFire.Firestore(context.Background())
	return func(c *fiber.Ctx) error {

		// Extract the research ID from the URL path
		id := c.Params("id")

		var requestBody entities.Research

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}

		//var researches []presenter.Research

		//Indicates the document's path
		docPath := fmt.Sprintf("Topic/%s/%s/%s", requestBody.TopicID, id, id)
		//Indicates to document
		userDoc := Client.Doc(docPath)

		docSnapShot, err := userDoc.Get(context.Background())
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}

		data := docSnapShot.Data()

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   &data,
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

func PostPDF(appFire *firebase.App) fiber.Handler {
	client, _ := storage.NewClient(context.Background())
	return func(c *fiber.Ctx) error {
		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form

			if token := form.Value["token"]; len(token) > 0 {
				// Get key value:
				fmt.Println(token[0])
			}

			// Get all files from "documents" key:
			files := form.File["documents"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				// Save the files to disk:
				if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
					return err
				}

				dat, _ := os.ReadFile(file.Filename)
				// Upload the files to firebase storage:
				wc := client.Bucket("bucket-name").Object(file.Filename).NewWriter(context.Background())
				wc.ContentType = file.Header["Content-Type"][0]
				// write the file to the bucket
				if _, err := wc.Write(dat); err != nil {
					return err
				}
				if err := wc.Close(); err != nil {
					return err
				}
				// print the file url
				fmt.Println(wc.Attrs().MediaLink)
			}
		}
		return c.JSON("ok")
	}
}

// UUID generates a uniq id
func UUID() string {
	newUUID := uuid.New().String()
	return newUUID
}
