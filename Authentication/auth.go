package auth

import (
	create "api/Create"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
)

var _, _, ctx, err = create.NewFireStore()

func Auth(App *fiber.App, appfire *firestore.Client) {

	App.Post("/login", func(c *fiber.Ctx) error {

		var login struct { // Parse request body
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&login); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		documentPath := fmt.Sprint("Users/", login.Email)

		x, err := appfire.Doc(documentPath).Get(ctx)

		emailCr := x.Data()

		if err != nil {
			return err
		}

		incomePassword := login.Password
		truePassword := emailCr["password"]

		if incomePassword == truePassword {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Acces permitted",
			})
		} else {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Acces denied",
			})
		}

	})
}
