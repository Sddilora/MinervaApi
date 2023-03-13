package methods

import (
	handlers "api/RouteHandlers"
	auth "api/Users"

	"github.com/gofiber/fiber/v2"
)

func PostMethods(app *fiber.App) {

	//Auth Routes
	app.Post("/login", auth.CreateUserHandler)
	app.Post("/user", auth.SigninHandler)

	//Save data to Database Routes
	app.Post("/topic", handlers.PostTopicHandler)
	app.Post("/topic/research", handlers.PostResearchHandler)

	//Retrieve data from Database Routes
	app.Post("/topics", handlers.PostTopicsHandler)
	app.Post("/topic/researches", handlers.PostResearchesHandler)

}
