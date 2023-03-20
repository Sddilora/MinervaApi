package routes

import "github.com/gofiber/fiber/v2"

// TopicRouter is the Router for GoFiber App
func TopicRouter(app fiber.Router, service topic.Service) {
	app.Get("/topics", handlers.GetTopics(service))
	app.Post("/topics", handlers.AddTopic(service))
	app.Put("/topics", handlers.UpdateTopic(service))
	app.Delete("/topics", handlers.RemoveTopic(service))
}
