package main

import (
	"fmt"
	"log"
	"time"

	"context"

	"github.com/gofiber/fiber/v2"

	"minerva_api/api/routes"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func main() {

	client, appFire, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!", client /*just for I hate warnings ,remove that*/)
	//Creates new fiber app
	app := fiber.New()

	routes.ResearchRouter(app, appFire)

	defer cancel()
	//Starts the HTTP server
	log.Fatal(app.Listen(":7334"))
}

func databaseConnection() (*auth.Client, *firebase.App, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//Take file and return as opt for API to use
	opt := option.WithCredentialsFile("C:/Users/sumey/Desktop/software/Back-End/Minerva/FireBase/key.json")
	//Creates a new App from the provided config and client options.
	appFire, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}
	// Create a client to access Firebase services such as Firebase Authentication and Realtime Database
	client, err := appFire.Auth(ctx)
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}
	return client, appFire, cancel, nil

}

// //Uses middlewares
// middlewares.FiberMiddlewares(app)

// //Calls allowed post methods
// methods.PostMethods(app)

// //Calls allowed put methods
// methods.PutMethods(app)

// //Calls allowed delete methods
// methods.DeleteMethods(app)
