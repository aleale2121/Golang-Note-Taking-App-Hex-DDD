package listing

// Service provides Post listing operations.
type Service interface {
	GetAllNotes() ([]Note, error)
	FindNoteByID(uint) (Note, error)
}

type Repository interface {
	GetAllNotes() ([]Note, error)
	FindNoteByID(uint) (Note, error)
}

type service struct {
	tR Repository
}

func (s *service) FindNoteByID(noteId uint) (Note, error) {
	return s.tR.FindNoteByID(noteId)
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAllNotes() ([]Note, error) {
	return s.tR.GetAllNotes()
}
