package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverMw "github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddlewares(app *fiber.App) {

	//Middlewares
	app.Use(logger.New())    //records the details of incoming requests when any HTTP request is made. This can be used for purposes such as debugging and performance optimization.
	app.Use(recoverMw.New()) //catches any errors that may cause the program to crash or interrupt and keep the server running.
	app.Use(cors.New())      //It helps applications bypass CORS restrictions by providing appropriate responses that allow or deny HTTP requests access to their resources.

}
