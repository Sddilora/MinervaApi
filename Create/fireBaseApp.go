package create

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFireStore() (*auth.Client, *firestore.CollectionRef, *firebase.App) {

	var UserCol *firestore.CollectionRef

	opt := option.WithCredentialsFile("C:/Users/sumey/Desktop/software/Back-End/Minerva/Create/key.json")
	appFire, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed while creating firebaseApp: %v", err)
	}

	// Create a client to access Firebase services such as Firebase Authentication and Realtime Database
	client, err := appFire.Auth(context.Background())
	if err != nil {
		log.Fatalf("Something went wrong while creating client: %v", err)
	}

	return client, UserCol, appFire
}
