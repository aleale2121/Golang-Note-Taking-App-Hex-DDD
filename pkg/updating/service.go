package updating

import (
	"fmt"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/entity"
)

var ErrDuplicate = fmt.Errorf("title already exist")

type Service interface {
	UpdateNote(note entity.Note) error
}

type Repository interface {
	UpdateNote(entity.Note) error
}

type service struct {
	tR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateNote(note entity.Note) error {
	if note.Title == "" || note.Detail == "" {
		return fmt.Errorf("invalid input")
	}
	return s.tR.UpdateNote(note)
}
