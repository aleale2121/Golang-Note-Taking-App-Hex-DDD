package postgress

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	ID     uint
	Title  string
	Detail string
}

type NoteRepository struct {
	conn *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db}
}

func (noteRepo *NoteRepository) Notes() ([]Note, []error) {
	var notes []Note
	errs := noteRepo.conn.Find(&notes).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return notes, errs
}

func (noteRepo *NoteRepository) Note(id uint32) (*Note, []error) {
	note := Note{}
	errs := noteRepo.conn.First(&note, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &note, errs
}

func (noteRepo *NoteRepository) UpdateClub(note *Note) (*Note, []error) {
	errs := noteRepo.conn.Save(note).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return note, errs
}

func (noteRepo *NoteRepository) DeleteClub(id uint32) (*Note, []error) {
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

func (noteRepo *NoteRepository) StoreClub(note *Note) (*Note, []error) {
	errs := noteRepo.conn.Create(note).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return note, errs
}
