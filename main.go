package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var (
	ctx     context.Context
	appfire *firestore.Client
	userCol *firestore.CollectionRef
)

func main() {

	jsonData := Key()

	// Create a new Firestore client using the Google Application Credentials path
	ctx = context.Background()
	var err error
	appfire, err = firestore.NewClient(ctx, "minerva-95196", option.WithCredentialsJSON(jsonData))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer appfire.Close()

	// Firestore client
	userCol = appfire.Collection("users")

	setup := Setup()
	log.Fatal(setup.Listen(":7334"))
}

func Key() []byte {

	// Load the service account key file
	keyFile, err := os.Open("key.json")
	if err != nil {
		log.Fatalf("Failed to open service account key file: %v", err)
	}
	defer keyFile.Close()

	// Parse the service account key JSON data
	var keyData map[string]interface{}
	if err := json.NewDecoder(keyFile).Decode(&keyData); err != nil {
		log.Fatalf("Failed to parse service account key file: %v", err)
	}

	// Convert keyData to JSON format
	jsonData, err := json.Marshal(keyData)
	if err != nil {
		log.Fatalf("Failed to marshal credentials: %v", err)
	}
	return jsonData
}

func Setup() *fiber.App {

	app := fiber.New()

	//We will add login part later
	// app.Post("/login", func(c *fiber.Ctx) error {

	// })

	app.Post("/users", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Users")

		// Parse request body
		var newUser struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// Create new user
		user := map[string]interface{}{
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		}

		// Add user to Firestore
		_, err := userCol.Doc(newUser.Email).Set(ctx, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add user to Firestore",
			})
		}

		resp := fiber.Map{
			"message": "User created successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	})

	///////////////////////////////TOPIC//////////////////////////////////////////////

	app.Post("/topic", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Topic")

		// Parse request body
		var newTopic struct {
			Topic string `json:"topic"`
		}

		if err := c.BodyParser(&newTopic); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// Create new Topic
		topic := map[string]interface{}{
			"topic": newTopic.Topic,
		}

		// Add Topic to Firestore
		_, err := userCol.Doc(newTopic.Topic).Set(ctx, topic)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add Topic to Firestore",
			})
		}

		resp := fiber.Map{
			"message": "Topic created successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	})

	//////////////////////////////RESEARCH//////////////////////////////////////////////////

	app.Post("/topic/research", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Research")

		// Parse request body
		var newResearch struct {
			ResearchHeader      string `json:"research_header"`
			ResearchContent     string `json:"research_content"`
			ResearchCreator     string `json:"research_creator"`
			ResearchContributor string `json:"research_contributor`
		}

		if err := c.BodyParser(&newResearch); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// Create new user
		research := map[string]interface{}{
			"research_header":      newResearch.ResearchHeader,
			"research_content":     newResearch.ResearchContent,
			"research_creator":     newResearch.ResearchCreator,
			"research_contributor": newResearch.ResearchCreator,
		}

		// Add user to Firestore
		_, err := userCol.Doc(newResearch.ResearchHeader).Set(ctx, research)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add Research to Firestore",
			})
		}

		resp := fiber.Map{
			"message": "Research created successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	})

	return app
}
