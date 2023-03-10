package main

import (
	requests "api/RouteHandlers"

	auth "api/Users"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverMw "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New()

	//Middlewares
	app.Use(logger.New())    //records the details of incoming requests when any HTTP request is made. This can be used for purposes such as debugging and performance optimization.
	app.Use(recoverMw.New()) //catches any errors that may cause the program to crash or interrupt and keep the server running.
	app.Use(cors.New())      //It helps applications bypass CORS restrictions by providing appropriate responses that allow or deny HTTP requests access to their resources.

	//_, _, fireApp := create.NewFireStore()

	app.Post("/signup", auth.CreateUserHandler)
	app.Post("/signin", auth.SigninHandler)

	// Authenticaed Routes.
	app.Post("/topic", requests.PostTopicHandler)
	app.Post("/topic/research", requests.PostResearchHandler)
	// app.Get("/topics" /*,func will be added */)           // Ignore the auth by IgnoreUrls config
	// app.Get("/topic/researches" /*,func will be added */) // Ignore the auth by IgnoreUrls config

	log.Fatal(app.Listen(":7334"))
}
