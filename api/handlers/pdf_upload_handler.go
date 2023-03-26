package handlers

import (
	"context"
	"io"
	"log"

	storage "firebase.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddPdf is handler/controller which upload pdf to the Database
func AddPdf(storageClient *storage.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Parse the incoming PDF file
		pdf, err := c.FormFile("pdf")
		if err != nil {
			log.Print(err, "formfile")
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())

		}

		// Read the PDF file content
		pdfBytes, err := pdf.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer pdfBytes.Close()

		// Upload the PDF file to Firebase Storage
		bucketName := "minerva-95196.appspot.com"
		bucket, err := storageClient.Bucket(bucketName)
		if err != nil {
			log.Print(err, "bucket")
			return err
		}

		// Generate a unique name for the PDF file
		pdfName := uuid.New().String() + ".pdf"
		// Upload the PDF file to Firebase Storage
		pdfRef := bucket.Object(pdfName)
		writer := pdfRef.NewWriter(context.Background())
		defer writer.Close()

		if _, err := io.Copy(writer, pdfBytes); err != nil {
			log.Print(err, "copy")
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Get the download URL of the uploaded PDF file
		attrs, err := pdfRef.Attrs(context.Background())
		if err != nil {
			log.Print(err, "attrs")
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		pdfURL := attrs.MediaLink

		// Send the PDF URL back to the client
		return c.SendString(pdfURL)

	}
}
