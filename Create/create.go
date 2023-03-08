package create

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewFiber() *fiber.App {
	App := fiber.New()

	return App
}

func NewFireStore() (*firestore.Client, *firestore.CollectionRef, context.Context, error) {

	var (
		UserCol *firestore.CollectionRef
		ctx     context.Context
	)

	fileContents, err := os.ReadFile("C:/Users/sumey/Desktop/software/Back-End/Minerva/Create/key.json") // Read and store as []byte the key file required for us to access our Firestore database
	if err != nil {
		log.Fatalf("Failed Read file: %v", err)
	}

	ctx = context.Background() // Create a new Firestore client using the Google Application Credentials path
	opt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time: 2 * time.Minute,
	}))

	appfire, err := firestore.NewClient(ctx, "minerva-95196", option.WithCredentialsJSON(fileContents), opt) //we are creating new firestroe client
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	return appfire, UserCol, ctx, err
}
