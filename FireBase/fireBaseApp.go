package create

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFireStore() (*auth.Client, *firebase.App) {

	//Take file and return as opt for API to use
	opt := option.WithCredentialsFile("C:/Users/sumey/Desktop/software/Back-End/Minerva/FireBase/key.json")
	//Creates a new App from the provided config and client options.
	appFire, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil { //Error check
		log.Fatalf("Failed while creating firebaseApp: %v", err)
	}

	// Create a client to access Firebase services such as Firebase Authentication and Realtime Database
	client, err := appFire.Auth(context.Background())
	if err != nil { //Error check
		log.Fatalf("Something went wrong while creating client: %v", err)
	}

	return client, appFire
}
