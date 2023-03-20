package routes

import (
	"minerva_api/api/handlers"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

// ResearchRouter is the Router for GoFiber App
func ResearchRouter(app fiber.Router, appFire *firebase.App) {
	app.Get("/researches", handlers.GetResearches(appFire))
	app.Post("/researches", handlers.AddResearch(appFire))
	app.Put("/researches", handlers.UpdateResearch(appFire))
	app.Delete("/researches", handlers.RemoveResearch(appFire))
}
