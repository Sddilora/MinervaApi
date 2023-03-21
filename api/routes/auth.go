package routes

import (
	"minerva_api/api/handlers"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

// TopicRouter is the Router for GoFiber App
func AuthRouter(app *fiber.App, appFire *firebase.App) {
	app.Post("/user", handlers.AddUser(appFire))
	app.Post("/signin", handlers.Signin(appFire))
	app.Put("/user", handlers.UpdateUser(appFire))
	app.Delete("/user", handlers.RemoveUser(appFire))
}
