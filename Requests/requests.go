package requests

import (
	create "api/Create"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
)

var _, userCol, ctx, _ = create.NewFireStore()

func PostUser(App *fiber.App, appfire *firestore.Client) {

	App.Post("/user", func(c *fiber.Ctx) error {

		var newUser struct { // Parse request body
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		user := map[string]interface{}{ // Create new user
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		}

		userCol = appfire.Collection("Users")

		_, err := userCol.Doc(newUser.Email).Set(ctx, user) // Add user to Firestore
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

}

func PostTopic(App *fiber.App, appfire *firestore.Client) {
	App.Post("/topic", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Topic")

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

		_, err := userCol.Doc(newTopic.Topic).Set(ctx, topic) // Add Topic to Firestore
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
	})
}

func Research(App *fiber.App, appfire *firestore.Client) {
	App.Post("/topic/research", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Research")

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

		research := map[string]interface{}{ // Create new user
			"research_header":      newResearch.ResearchHeader,
			"research_content":     newResearch.ResearchContent,
			"research_creator":     newResearch.ResearchCreator,
			"research_contributor": newResearch.ResearchContributor,
		}

		_, err := userCol.Doc(newResearch.ResearchHeader).Set(ctx, research) // Add user to Firestore
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
}
