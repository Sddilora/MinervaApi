package research

import (
	"minerva_api/api/presenter"
	"minerva_api/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type ResearchService interface {
	InsertResearch(research *entities.Research) (*entities.Research, error)
	FetchResearch(research *entities.Research) (*[]presenter.Research, error)
	UpdateResearch(research *entities.Research) (*entities.Research, error)
	RemoveResearch(research *entities.Research) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) ResearchService {
	return &service{
		repository: r,
	}
}

// InsertResearch is a service layer that helps insert research in Database
func (s *service) InsertResearch(book *entities.Research) (*entities.Research, error) {
	return s.repository.CreateResearch(book)
}

// FetchResearches is a service layer that helps fetch all researches in Database
func (s *service) FetchResearch(research *entities.Research) (*[]presenter.Research, error) {
	return s.repository.ReadResearch(research)
}

// UpdateBook is a service layer that helps update books in Database
func (s *service) UpdateResearch(research *entities.Research) (*entities.Research, error) {
	return s.repository.UpdateResearch(research)
}

// RemoveResearch is a service layer that helps remove researches from Database
func (s *service) RemoveResearch(research *entities.Research) error {
	return s.repository.DeleteResearch(research)
}
