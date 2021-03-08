package adding

import "fmt"

var ErrDuplicate = fmt.Errorf("title already exist")

type Service interface {
	AddNote(Note) error
}

type Repository interface {
	AddNote(Note) error
}

type service struct {
	tR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddNote(note Note) error {
	if note.Title == "" || note.Detail == "" {
		return fmt.Errorf("invalid input")
	}
	return s.tR.AddNote(note)
}
