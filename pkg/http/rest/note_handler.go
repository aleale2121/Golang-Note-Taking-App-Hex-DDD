package rest

import (
	"context"
	"encoding/json"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/adding"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/deleting"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/entity"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/listing"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/updating"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type NoteHandler interface {
	GetNotes(w http.ResponseWriter, r *http.Request)
	GetNoteById(w http.ResponseWriter, r *http.Request)
	AddNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
	EditNote(w http.ResponseWriter, r *http.Request)
	MiddleWareValidateNote(next http.Handler) http.Handler
}
type noteHandler struct {
	l listing.Service
	a adding.Service
	d deleting.Service
	u updating.Service
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

func NewNoteHandler(l listing.Service,
	a adding.Service,
	d deleting.Service,
	u updating.Service) NoteHandler {

	return &noteHandler{l: l, a: a, d: d, u: u}

}

func (n noteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allNotes, err := n.l.GetAllNotes()

	if err != nil {
		http.Error(w, "Failed to get notes", http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(allNotes)
}

func (n noteHandler) GetNoteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	noteID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil || noteID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	note, err := n.l.FindNoteByID(uint(noteID))

	if err != nil {
		http.Error(w, "Failed to get notes", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (n noteHandler) AddNote(w http.ResponseWriter, r *http.Request) {
	note := r.Context().Value(keyNote{}).(entity.Note)
	if err := n.a.AddNote(note); err != nil {
		http.Error(w, "Failed to add note", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New note added.")

}

func (n noteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	noteID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil || noteID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err := n.d.DeleteNote(uint(noteID)); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("note deleted.")
}

func (n noteHandler) EditNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	noteID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil || noteID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	note := r.Context().Value(keyNote{}).(entity.Note)

	_, err = n.l.FindNoteByID(uint(noteID))
	if err != nil {
		http.Error(w, "Failed to get note", http.StatusBadRequest)
		return
	}
	err = n.u.UpdateNote(note)
	if err != nil {
		http.Error(w, "Failed to get note", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(" note updated.")

}
