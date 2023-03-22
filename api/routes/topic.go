package routes

import (
	"minerva_api/api/handlers"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

// TopicRouter is the Router for GoFiber App
func TopicRouter(app *fiber.App, appFire *firebase.App) {
	app.Get("/topics", handlers.GetTopics(appFire))
	app.Post("/topics", handlers.AddTopic(appFire))
	app.Put("/topics", handlers.UpdateTopic(appFire))
	app.Delete("/topics", handlers.RemoveTopic(appFire))
}
