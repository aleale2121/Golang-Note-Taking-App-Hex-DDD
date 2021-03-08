package deleting

import "fmt"

var ErrDuplicate = fmt.Errorf("title already exist")

type Service interface {
	DeleteNote(uint) error
}

type Repository interface {
	DeleteNote(uint) error
}

type service struct {
	tR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteNote(noteID uint) error {
	if noteID == 0 {
		return fmt.Errorf("invalid note id")
	}
	return s.tR.DeleteNote(noteID)
}
