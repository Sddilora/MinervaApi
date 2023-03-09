package main

import (
	create "api/Create"
	requests "api/Requests"

	auth "api/Users"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverMw "github.com/gofiber/fiber/v2/middleware/recover"
	gofiberfirebaseauth "github.com/sacsand/gofiber-firebaseauth"
)

func main() {

	app := fiber.New()

	//Middlewares
	app.Use(logger.New())    //records the details of incoming requests when any HTTP request is made. This can be used for purposes such as debugging and performance optimization.
	app.Use(recoverMw.New()) //catches any errors that may cause the program to crash or interrupt and keep the server running.
	app.Use(cors.New())      //It helps applications bypass CORS restrictions by providing appropriate responses that allow or deny HTTP requests access to their resources.

	_, _, fireApp := create.NewFireStore()

	app.Post("/signup", auth.CreateUserHandler)
	app.Post("/signin", auth.SigninHandler)

	// app.Use("/users", auth.CreateUserHandler)
	app.Use(gofiberfirebaseauth.New(gofiberfirebaseauth.Config{
		// Firebase Authentication App Object
		// Mandatory
		FirebaseApp: fireApp,
		// Ignore urls array.
		// Optional. These url will ignore by middleware
		IgnoreUrls: []string{"GET::/topics", "GET::/topic/researches"},
	}))

	// Authenticaed Routes.
	app.Post("/topic", requests.PostTopic)
	app.Post("/topic/research", requests.Research)
	// app.Get("/topics" /*,func will be added */)           // Ignore the auth by IgnoreUrls config
	// app.Get("/topic/researches" /*,func will be added */) // Ignore the auth by IgnoreUrls config

	log.Fatal(app.Listen(":7334"))
}
