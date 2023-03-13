package methods

import (
	handlers "api/RouteHandlers"

	"github.com/gofiber/fiber/v2"
)

func PutMethods(app *fiber.App) {

	// //Update user Route
	// app.Put("/user" /**/)

	//Update datas from Database Routes
	app.Put("/topic", handlers.PutTopicHandler)
	app.Put("/topic/research", handlers.PutResearchHandler)
}
