package methods

import (
	requests "api/RouteHandlers"
	auth "api/Users"

	"github.com/gofiber/fiber/v2"
)

func PostMethods(app *fiber.App) {

	//Auth Routes
	app.Post("/signup", auth.CreateUserHandler)
	app.Post("/signin", auth.SigninHandler)

	//Save data to Database Routes
	app.Post("/topic", requests.PostTopicHandler)
	app.Post("/topic/research", requests.PostResearchHandler)

	//Retrieve data from Database Routes
	app.Post("/topics", requests.PostTopicsHandler)
	app.Post("/topic/researches", requests.PostResearchesHandler)

}
