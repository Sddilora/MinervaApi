package routes

import (
	"minerva_api/api/handlers"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

// ResearchRouter is the Router for GoFiber App
func ResearchRouter(app *fiber.App, appFire *firebase.App) {
	app.Get("/topic/researches", handlers.GetResearches(appFire))
	app.Get("/topic/research/:id", handlers.GetResearchByID(appFire))
	app.Post("/topic/research", handlers.AddResearch(appFire))
	app.Put("/topic/research", handlers.UpdateResearch(appFire))
	app.Delete("/topic/research", handlers.RemoveResearch(appFire))
}
