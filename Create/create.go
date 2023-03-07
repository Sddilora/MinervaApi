package create

import (
	key "api/Key"
	"context"
	"log"
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

	jsonData := key.Key() //we will make that easier

	ctx = context.Background() // Create a new Firestore client using the Google Application Credentials path
	opt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time: 2 * time.Minute,
	}))

	appfire, err := firestore.NewClient(ctx, "minerva-95196", option.WithCredentialsJSON(jsonData), opt) //we are creating new firestroe client
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	return appfire, UserCol, ctx, err
}
