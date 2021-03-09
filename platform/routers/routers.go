package routers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Routers interface {
	Serve()
}

type Router struct {
	Method string
	Path   string
	Handle httprouter.Handle
}
type routing struct {
	host    string
	port    string
	routers []Router
}

func NewRouting(host, port string, routers []Router) Routers {
	return &routing{
		host,
		port,
		routers,
	}
}

func (r *routing) Serve() {
	httpRouter := httprouter.New()
	for _, router := range r.routers {
		httpRouter.Handle(router.Method, router.Path, router.Handle)
	}
	addr := fmt.Sprintf("%s:%s", r.host, r.port)
	err := http.ListenAndServe(addr, httpRouter)
	if err != nil {
		panic(err)
	}
}
