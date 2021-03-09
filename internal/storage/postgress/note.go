package postgress

import (
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/constant/model"
	"github.com/jinzhu/gorm"
)

type NoteRepository struct {
	conn *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db}
}

func (noteRepo *NoteRepository) Notes() ([]model.Note, []error) {
	var notes []model.Note
	errs := noteRepo.conn.Find(&notes).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return notes, errs
}

func (noteRepo *NoteRepository) Note(id uint32) (*model.Note, []error) {
	note := model.Note{}
	errs := noteRepo.conn.First(&note, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &note, errs
}

func (noteRepo *NoteRepository) UpdateNote(note *model.Note) (*model.Note, []error) {
	errs := noteRepo.conn.Save(note).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return note, errs
}

func (noteRepo *NoteRepository) DeleteNote(id uint32) (*model.Note, []error) {
	note, errs := noteRepo.Note(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = noteRepo.conn.Delete(note, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return note, errs
}

func (noteRepo *NoteRepository) StoreNote(note *model.Note) (*model.Note, []error) {
	errs := noteRepo.conn.Create(note).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return note, errs
}
