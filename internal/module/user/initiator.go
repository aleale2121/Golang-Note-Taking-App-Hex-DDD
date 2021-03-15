package note

import (
	"fmt"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/constant/model"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/storage/postgress"
)

//var ErrDuplicate = fmt.Errorf("title already exist")

type UseCase interface {
	GetAllNotes() ([]model.Note, []error)
	FindNoteByID(uint) (*model.Note, []error)
	AddNote(note model.Note) (*model.Note, []error)
	UpdateNote(note model.Note) (*model.Note, []error)
	DeleteNote(uint) (*model.Note, []error)
}

type Service struct {
	notePostgres postgress.NoteRepository
}

func NewService(notePostgres postgress.NoteRepository) UseCase {
	return &Service{
		notePostgres: notePostgres,
	}
}

func (s Service) GetAllNotes() ([]model.Note, []error) {
	return s.notePostgres.Notes()

}

func (s Service) FindNoteByID(u uint) (*model.Note, []error) {
	return s.notePostgres.Note(uint32(u))
}

func (s Service) AddNote(note model.Note) (*model.Note, []error) {
	if note.Title == "" || note.Detail == "" {
		return nil, []error{fmt.Errorf("invalid input")}
	}
	return s.notePostgres.StoreNote(&note)
}

func (s Service) UpdateNote(note model.Note) (*model.Note, []error) {
	if note.Title == "" || note.Detail == "" {
		return nil, []error{fmt.Errorf("invalid input")}
	}
	return s.notePostgres.UpdateNote(&note)
}

func (s Service) DeleteNote(u uint) (*model.Note, []error) {
	return s.notePostgres.DeleteNote(uint32(u))
}
