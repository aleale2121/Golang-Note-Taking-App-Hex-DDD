package routing

import (
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/handler/rest"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/platform/routers"
	"net/http"
)

func NoteRouting(handler rest.NoteHandler) []routers.Router {
	return []routers.Router{
		{
			Method: http.MethodGet,
			Path:   "/note",
			Handle: handler.GetNotes,
		},
		{
			Method: http.MethodGet,
			Path:   "/note/:id",
			Handle: handler.GetNotes,
		},
		{
			Method: http.MethodPost,
			Path:   "/note",
			Handle: handler.AddNote,
		},
		{
			Method: http.MethodPut,
			Path:   "/note/:id",
			Handle: handler.GetNotes,
		},
		{
			Method: http.MethodDelete,
			Path:   "/note/:id",
			Handle: handler.GetNotes,
		},
	}
}
