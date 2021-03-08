package listing

import "github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/entity"

// Service provides Post listing operations.
type Service interface {
	GetAllNotes() ([]entity.Note, error)
	FindNoteByID(uint) (entity.Note, error)
}

type Repository interface {
	GetAllNotes() ([]entity.Note, error)
	FindNoteByID(uint) (entity.Note, error)
}

type service struct {
	tR Repository
}

func (s *service) FindNoteByID(noteId uint) (entity.Note, error) {
	return s.tR.FindNoteByID(noteId)
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAllNotes() ([]entity.Note, error) {
	return s.tR.GetAllNotes()
}
