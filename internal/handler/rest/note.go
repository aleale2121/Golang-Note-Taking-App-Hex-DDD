package rest

import (
	"context"
	"encoding/json"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/constant/model"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/module/user"
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
	MiddleWareValidateNote(next httprouter.Handle) httprouter.Handle
}
type noteHandler struct {
	useCase note.UseCase
}

type keyNote struct{}

func (n noteHandler) MiddleWareValidateNote(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		noteX := model.Note{}
		err := noteX.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Error reading noteX", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), keyNote{}, noteX)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}

func NewNoteHandler(useCase note.UseCase) NoteHandler {

	return &noteHandler{useCase: useCase}

}

func (n noteHandler) GetNotes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	allNotes, errs := n.useCase.GetAllNotes()

	if len(errs) > 0 {
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
	noteByID, errs := n.useCase.FindNoteByID(uint(noteID))

	if len(errs) > 0 {
		http.Error(w, "Failed to get notes", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(noteByID)
}

func (n noteHandler) AddNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	noteById := r.Context().Value(keyNote{}).(model.Note)
	if _, errs := n.useCase.AddNote(noteById); len(errs) > 0 {
		http.Error(w, "Failed to add noteById", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New noteById added.")

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
	noteById := r.Context().Value(keyNote{}).(model.Note)

	_, errs := n.useCase.FindNoteByID(uint(noteID))
	if len(errs) > 0 {
		http.Error(w, "Failed to get noteById", http.StatusBadRequest)
		return
	}
	_, errs = n.useCase.UpdateNote(noteById)
	if len(errs) > 0 {
		http.Error(w, "Failed to get noteById", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(" noteById updated.")

}
