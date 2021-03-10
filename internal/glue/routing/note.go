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
			Path:        "/note",
			Handle:      handler.GetNotes,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/note/:id",
			Handle:      handler.GetNotes,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPost,
			Path:        "/note",
			Handle:      handler.AddNote,
			MiddleWares: []func(handle httprouter.Handle) httprouter.Handle{handler.MiddleWareValidateNote},
		},
		{
			Method:      http.MethodPut,
			Path:        "/note/:id",
			Handle:      handler.GetNotes,
			MiddleWares: []func(handle httprouter.Handle) httprouter.Handle{handler.MiddleWareValidateNote},
		},
		{
			Method:      http.MethodDelete,
			Path:        "/note/:id",
			Handle:      handler.GetNotes,
			MiddleWares: nil,
		},
	}
}
