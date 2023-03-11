package auth

import (
	create "api/FireBase"
	"context"
	"log"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

// When a POST request is sent to the /user endpoint, get the JSON object and save it to Firebase Authentication:
func CreateUserHandler(c *fiber.Ctx) error {

	_, appfire := create.NewFireStore()
	// Parse request body
	var newUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		PhotoUrl string `json:"photo_url"`
	}
	//Error Handler
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	params := (&auth.UserToCreate{}).
		Email(newUser.Email).
		EmailVerified(false). //Whether the user's primary email is verified. If not provided, the default value is false.
		Password(newUser.Password).
		DisplayName(newUser.Name).
		PhotoURL(newUser.PhotoUrl).
		Disabled(false) //Whether the user has been disabled. true for disabled; false for active . If not provided, the default value is false.

	//Returns auth.Client for CreateUser func
	client, err := appfire.Auth(context.Background())
	if err != nil {
		log.Printf("error creating user: %v\n", err)
	}
	//Creates a new user with the specified properties and returns an userRecord
	_, err = client.CreateUser(context.Background(), params)
	if err != nil {
		log.Print(err)
		resp := fiber.Map{
			"message": "user could not be created ",
		}
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		resp := fiber.Map{
			"message": "User Created Succesfully",
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

/* user update seçeneği ekleyeceğimiz zaman kullanacağız:
params := (&auth.UserToUpdate{}).
        Email("user@example.com").
        EmailVerified(true).
        PhoneNumber("+15555550100").
        Password("newPassword").
        DisplayName("John Doe").
        PhotoURL("http://www.example.com/12345678/photo.png").
        Disabled(true)
u, err := client.UpdateUser(ctx, uid, params)
if err != nil {
        log.Fatalf("error updating user: %v\n", err)
}
log.Printf("Successfully updated user: %v\n", u)*/
