package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"minerva_api/api/presenter"
	"minerva_api/pkg/entities"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

func AddUser(appFire *firebase.App) fiber.Handler {
	//Creates auth.Client for CreateUser func
	AuthClient, _ := appFire.Auth(context.Background())
	return func(c *fiber.Ctx) error {

		var requestBody *entities.User
		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		requestBody.CreatedAt = time.Now()

		params := (&auth.UserToCreate{}).
			Email(requestBody.Email).
			Password(requestBody.Password).
			DisplayName(requestBody.Name).
			PhotoURL(requestBody.PhotoUrl)

			//Creates a new user with the specified properties and returns an userRecord
		_, err := AuthClient.CreateUser(context.Background(), params)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(requestBody))

	}
}

func Signin(appFire *firebase.App) fiber.Handler {
	//Creates auth.Client for CreateUser func
	AuthClient, _ := appFire.Auth(context.Background())
	return func(c *fiber.Ctx) error {
		var requestBody *entities.User
		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		authUserRecord, err := AuthClient.GetUserByEmail(context.Background(), requestBody.Email)
		if err != nil {
			log.Println(err)
		}

		signedinUsersId := authUserRecord.UID

		confirmation := authenticationHandler(requestBody.Email, requestBody.Password)

		JWT := confirmation["idToken"].(string)

		_, ok := confirmation["registered"].(bool)
		if ok {
			resp := fiber.Map{
				"message":   "Login access permitted",
				"author_id": signedinUsersId,
				"jwt":       JWT,
			}
			return c.Status(fiber.StatusUnauthorized).JSON(resp)
		} else {
			resp := fiber.Map{
				"message": "Login access denied",
			}
			return c.Status(fiber.StatusOK).JSON(resp)
		}

	}
}

func UpdateUser(appFire *firebase.App) fiber.Handler {
	//Creates auth.Client for CreateUser func
	AuthClient, _ := appFire.Auth(context.Background())
	return func(c *fiber.Ctx) error {
		var requestBody *entities.User
		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		authUserRecord, err := AuthClient.GetUser(context.Background(), requestBody.ID)
		if err != nil {
			log.Println(err)
		}
		params := (&auth.UserToUpdate{}).
			Email(requestBody.Email).
			Password(requestBody.Password).
			DisplayName(requestBody.Name).
			PhotoURL(requestBody.PhotoUrl)

		uid := authUserRecord.UID
		_, err = AuthClient.UpdateUser(context.Background(), uid, params)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status":  true,
			"message": "User Updated Successfully",
			"error":   nil,
		})
	}
}

func RemoveUser(appFire *firebase.App) fiber.Handler {
	//Creates auth.Client for CreateUser func
	AuthClient, _ := appFire.Auth(context.Background())
	return func(c *fiber.Ctx) error {
		var requestBody *entities.User
		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		authUserRecord, err := AuthClient.GetUser(context.Background(), requestBody.ID)
		if err != nil {
			log.Println(err)
		}

		uid := authUserRecord.UID
		err = AuthClient.DeleteUser(context.Background(), uid)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status":  true,
			"message": "User Removed Successfully",
			"error":   nil,
		})
	}
}

func authenticationHandler(email string, password string) map[string]interface{} {
	// Create Http request
	apiKey, err := os.ReadFile("api/apikey.txt")
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

// func GetUsers
