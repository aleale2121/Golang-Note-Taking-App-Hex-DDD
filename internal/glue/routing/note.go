package routing

import (
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/handler/rest"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/platform/routers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NoteRouting(handler rest.NoteHandler) []routers.Router {
	return []routers.Router{
		{
			Method:      http.MethodGet,
			Path:        "/v1/note",
			Handle:      handler.GetNotes,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1/note/:id",
			Handle:      handler.GetNoteById,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPost,
			Path:        "/v1/note",
			Handle:      handler.AddNote,
			MiddleWares: []func(handle httprouter.Handle) httprouter.Handle{handler.MiddleWareValidateNote},
		},
		{
			Method:      http.MethodPut,
			Path:        "/v1/note/:id",
			Handle:      handler.EditNote,
			MiddleWares: []func(handle httprouter.Handle) httprouter.Handle{handler.MiddleWareValidateNote},
		},
		{
			Method:      http.MethodDelete,
			Path:        "/v1/note/:id",
			Handle:      handler.DeleteNote,
			MiddleWares: nil,
		},
	}
}
