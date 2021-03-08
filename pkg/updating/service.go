package updating

import "fmt"

var ErrDuplicate = fmt.Errorf("title already exist")

type Service interface {
	UpdateNote(Note) error
}

type Repository interface {
	UpdateNote(Note) error
}

type service struct {
	tR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateNote(note Note) error {
	if note.Title == "" || note.Detail == "" {
		return fmt.Errorf("invalid input")
	}
	return s.tR.UpdateNote(note)
}
