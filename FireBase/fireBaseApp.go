package create

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFireStore() (*auth.Client, *firebase.App) {

	//var UserCol *firestore.CollectionRef

	opt := option.WithCredentialsFile("C:/Users/sumey/Desktop/software/Back-End/Minerva/FireBase/key.json")
	appFire, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed while creating firebaseApp: %v", err)
	}

	// Create a client to access Firebase services such as Firebase Authentication and Realtime Database
	client, err := appFire.Auth(context.Background())
	if err != nil {
		log.Fatalf("Something went wrong while creating client: %v", err)
	}

	// fireStoreClient, err := appFire.Firestore(context.Background())
	// if err != nil {
	// 	log.Fatalf("Failed while creating firebaseClient: %v", err)
	// }

	//UserCol = fireStoreClient.Collection("Collection")

	return client, appFire
}
