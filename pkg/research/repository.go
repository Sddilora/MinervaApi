package research

import (
	"context"
	"fmt"
	"log"
	"minerva_api/api/presenter"
	"minerva_api/pkg/entities"
	"os/exec"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Repository interface allows us to access the CRUD Operations in firebase here.
type Repository interface {
	CreateResearch(research *entities.Research) (*entities.Research, error)
	ReadResearch(research *entities.Research) (*[]presenter.Research, error)
	UpdateResearch(research *entities.Research) (*entities.Research, error)
	DeleteResearch(research *entities.Research) error
}

type repository struct {
	Client *firestore.Client
}

// NewRepo is the single instance repo that is being created.
func NewRepo(client *firestore.Client) Repository {
	return &repository{
		Client: client,
	}
}

// CreateResearch is a Firebase repository that helps to create researches
func (r *repository) CreateResearch(research *entities.Research) (*entities.Research, error) {

	collectionName := UUID()

	// Creates a reference to a collection group to Research path.
	colPath := fmt.Sprintf("Topic/%s/%s", research.TopicID, collectionName)

	collection := r.Client.Collection(colPath)

	// Set the ID field of the research to a new ID.
	research.ID = collectionName

	// Set the created and updated timestamps for the research.
	research.CreatedAt = time.Now()
	research.UpdatedAt = time.Now()

	_, err := collection.Doc(collectionName).Set(context.Background(), &research)
	if err != nil {
		return nil, err
	}
	return research, nil
}

// ReadResearch is a firebase repository that helps to fetch researchs
func (r *repository) ReadResearch(research *entities.Research) (*[]presenter.Research, error) {
	var researches []presenter.Research
	//Indicates the document's path
	docPath := fmt.Sprintf("Topic/%s", research.TopicID)
	//Indicates to document
	userDoc := r.Client.Doc(docPath)

	//Documents returns an iterator over the query's resulting documents.
	query := userDoc.Collections(context.Background())

	for {
		//Next returns the next result. Its second return value is iterator.Done if there are no more results
		col, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		//Next returns the next result. Its second return value is iterator.Done if there are no more results
		doc, err := (col.Documents(context.Background())).Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var research presenter.Research
		doc.DataTo(&research)
		researches = append(researches, research)
	}
	return &researches, nil
}

// UpdateResearch is a firebase repository that helps to update researchs
func (r *repository) UpdateResearch(research *entities.Research) (*entities.Research, error) {
	research.UpdatedAt = time.Now()
	//Indicates the document's path
	docRefPath := fmt.Sprintf("Topic/%s/%s/%s", research.TopicID, research.ID, research.ID)
	//Indicates to document
	userDoc := r.Client.Doc(docRefPath)

	//Updates the given parameters at the Document
	_, err := userDoc.Update(context.Background(), []firestore.Update{
		{Path: "research_header", Value: research.Title},
		{Path: "research_content", Value: research.Content},
		//{Path: "research_contributor", Value: updatedResearch.ResearchContributor},
	})
	if err != nil {
		return nil, err
	} else {
		return research, nil
	}
}

// DeleteResearch is a firebase repository that helps to delete researches
func (r *repository) DeleteResearch(research *entities.Research) error {

	//Indicates the document's path
	docRefPath := fmt.Sprintf("Topic/%s/%s/%s", research.TopicID, research.ID, research.ID)

	//Indicates to document
	userDoc := r.Client.Doc(docRefPath)

	//Deletes the document
	_, err := userDoc.Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// UUID generates a uniq id
func UUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Println(err)
	}
	return string(newUUID)
}
