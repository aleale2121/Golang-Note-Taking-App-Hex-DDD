package rest

import (
	"context"
	"encoding/json"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/constant/model"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/module/user"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/entity"
	"github.com/julienschmidt/httprouter"

	"net/http"
	"strconv"
)

type NoteHandler interface {
	GetNotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetNoteById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	EditNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	MiddleWareValidateNote(next http.Handler) http.Handler
}
type noteHandler struct {
	useCase note.UseCase
}

type keyNote struct{}

func (n noteHandler) MiddleWareValidateNote(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		note := entity.Note{}
		err := note.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Error reading note", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), keyNote{}, note)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}

func NewNoteHandler(useCase note.UseCase) NoteHandler {

	return &noteHandler{useCase: useCase}

}

func (n noteHandler) GetNotes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	allNotes, err := n.useCase.GetAllNotes()

	if err != nil {
		http.Error(w, "Failed to get notes", http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(allNotes)
}

func (n noteHandler) GetNoteById(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	noteID, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil || noteID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	note, errs := n.useCase.FindNoteByID(uint(noteID))

	if len(errs) > 0 {
		http.Error(w, "Failed to get notes", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (n noteHandler) AddNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	note := r.Context().Value(keyNote{}).(model.Note)
	if _, errs := n.useCase.AddNote(note); len(errs) > 0 {
		http.Error(w, "Failed to add note", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New note added.")

}

func (n noteHandler) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	noteID, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil || noteID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if _, errs := n.useCase.DeleteNote(uint(noteID)); len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("note deleted.")
}

func (n noteHandler) EditNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	noteID, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil || noteID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	note := r.Context().Value(keyNote{}).(model.Note)

	_, errs := n.useCase.FindNoteByID(uint(noteID))
	if len(errs) > 0 {
		http.Error(w, "Failed to get note", http.StatusBadRequest)
		return
	}
	_, errs = n.useCase.UpdateNote(note)
	if len(errs) > 0 {
		http.Error(w, "Failed to get note", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(" note updated.")

}
