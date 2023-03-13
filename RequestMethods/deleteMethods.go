package methods

import (
	handlers "api/RouteHandlers"

	"github.com/gofiber/fiber/v2"
)

func DeleteMethods(app *fiber.App) {

	//Delete user Route
	//app.Delete("/user" /**/)

	//Delete data from Database Routes
	app.Delete("/topic", handlers.DeleteTopicHandler)
	app.Delete("/topic/research", handlers.DeleteResearchHandler)
}
