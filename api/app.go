package main

import (
	"context"
	"fmt"
	"log"
	"minerva_api/api/routes"
	"time"

	firebase "firebase.google.com/go"
	storage "firebase.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverMw "github.com/gofiber/fiber/v2/middleware/recover"
	"google.golang.org/api/option"
)

func main() {

	appFire, cancel, storageClient, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")

	//Creates new fiber app
	app := fiber.New()

	//Middlewares
	app.Use(logger.New())    //records the details of incoming requests when any HTTP request is made. This can be used for purposes such as debugging and performance optimization.
	app.Use(recoverMw.New()) //catches any errors that may cause the program to crash or interrupt and keep the server running.
	app.Use(cors.New())      //It helps applications bypass CORS restrictions by providing appropriate responses that allow or deny HTTP requests access to their resources.

	routes.ResearchRouter(app, appFire, storageClient)
	routes.TopicRouter(app, appFire)
	routes.AuthRouter(app, appFire)

	defer cancel()
	defer app.Shutdown()
	//Starts the HTTP server
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*firebase.App, context.CancelFunc, *storage.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//Take file and return as opt for API to use
	opt := option.WithCredentialsFile("C:/Users/sumey/Desktop/software/Back-End/Minerva/api/key.json")
	//Creates a new App from the provided config and client options.
	appFire, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}

	storageClient, err := appFire.Storage(context.Background())
	if err != nil {
		log.Printf("error initializing firebase storage client: %v\n", err)
	}

	return appFire, cancel, storageClient, nil

}
