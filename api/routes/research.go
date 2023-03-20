package routes

import (
	"minerva_api/api/handlers"
	"minerva_api/pkg/research"

	"github.com/gofiber/fiber/v2"
)

// ResearchRouter is the Router for GoFiber App
func ResearchRouter(app fiber.Router, service research.ResearchService) {
	app.Get("/researches", handlers.GetResearches(service))
	app.Post("/researches", handlers.AddResearch(service))
	app.Put("/researches", handlers.UpdateResearch(service))
	app.Delete("/researches", handlers.RemoveResearch(service))
}
