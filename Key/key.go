package key

import (
	"encoding/json"
	"log"
	"os"
)

func Key() []byte {

	// Load the service account key file
	keyFile, err := os.Open("C:/Users/sumey/Desktop/software/Back-End/Minerva/Key/key.json")
	if err != nil {
		log.Fatalf("Failed to open service account key file: %v", err)
	}
	defer keyFile.Close()

	// Parse the service account key JSON data
	var keyData map[string]interface{}
	if err := json.NewDecoder(keyFile).Decode(&keyData); err != nil {
		log.Fatalf("Failed to parse service account key file: %v", err)
	}

	// Convert keyData to JSON format
	jsonData, err := json.Marshal(keyData)
	if err != nil {
		log.Fatalf("Failed to marshal credentials: %v", err)
	}
	return jsonData
}
