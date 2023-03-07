package main

import (
	auth "api/Authentication"
	create "api/Create"
	requests "api/Requests"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverMw "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	start := run() //var start is err

	if start != nil {
		panic(start)
	}
}

func run() error {

	app := create.NewFiber()
	appfire, _, _, _ := create.NewFireStore()
	defer appfire.Close()

	// Add middleware
	app.Use(logger.New())
	app.Use(recoverMw.New())
	app.Use(cors.New())

	requests.PostUser(app, appfire)
	requests.PostTopic(app, appfire)
	requests.Research(app, appfire)

	auth.Auth(app, appfire)

	log.Fatal(app.Listen(":7334"))

	return nil
}
