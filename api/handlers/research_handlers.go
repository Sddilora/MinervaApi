package handlers

import (
	"errors"
	"minerva_api/api/presenter"
	"minerva_api/pkg/entities"
	"minerva_api/pkg/research"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AddResearch is handler/controller which creates Researches in the Database
func AddResearch(service research.ResearchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Research

		//Body Parser,Error Handler
		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}

		if requestBody.Author == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(errors.New(
				"Please specify title and author")))
		}

		result, err := service.InsertResearch(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		return c.JSON(presenter.ResearchSuccessResponse(result))
	}
}

// UpdateResearch is handler/controller which updates data of Researches in the database
func UpdateResearch(service research.ResearchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Research
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		result, err := service.UpdateResearch(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		return c.JSON(presenter.ResearchSuccessResponse(result))
	}
}

// RemoveResearch is handler/controller which removes Researches from the Database
func RemoveResearch(service research.ResearchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Research
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		err = service.RemoveResearch(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

// GetResearches is handler/controller which lists all Researches from the database
func GetResearches(service research.ResearchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Research
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		fetched, err := service.FetchResearch(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ResearchErrorResponse(err))
		}
		return c.JSON(presenter.ResearchesSuccessResponse(fetched))
	}
}
