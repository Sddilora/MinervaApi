package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SigninHandler(c *fiber.Ctx) error {

	var UserSignin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&UserSignin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	k := authenticationHandler(UserSignin.Email, UserSignin.Password)
	log.Print("bakalım", k)

	resp := fiber.Map{
		"message": "Login access permitted ",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)

}

func authenticationHandler(email string, password string) map[string]interface{} {
	// Create Http request
	apiKey, err := os.ReadFile("Users/apikey.txt")
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", string(apiKey))
	requestData := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}
	//Encode(json) requestData as requestBody
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	//Set the header entries
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request and process response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		panic(err)
	}

	return responseBody
}
