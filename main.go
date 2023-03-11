package main

import (
	"log"

	middlewares "api/Middlewares"
	methods "api/RequestMethods"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Creates new fiber app
	app := fiber.New()

	//Uses middlewares
	middlewares.FiberMiddlewares(app)

	//Calls allowed post methods
	methods.PostMethods(app)

	//Starts the HTTP server
	log.Fatal(app.Listen(":7334"))
}
