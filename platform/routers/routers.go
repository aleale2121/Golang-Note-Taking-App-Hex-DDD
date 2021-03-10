package routers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rileyr/middleware"
	"github.com/rileyr/middleware/wares"
	"net/http"
)

type Routers interface {
	Serve()
}

type Router struct {
	Method      string
	Path        string
	Handle      httprouter.Handle
	MiddleWares []func(handle httprouter.Handle) httprouter.Handle
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
		if router.MiddleWares == nil {
			httpRouter.Handle(router.Method, router.Path, router.Handle)
		} else {
			s := middleware.NewStack()
			for _, middle := range router.MiddleWares {
				s.Use(middle)
			}
			s.Use(wares.RequestID)
			s.Use(wares.Logging)
			httpRouter.Handle(router.Method, router.Path, s.Wrap(router.Handle))
		}
	}
	addr := fmt.Sprintf("%s:%s", r.host, r.port)
	fmt.Println(addr)
	err := http.ListenAndServe(addr, httpRouter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("server started at %s:%s", r.host, r.port)
}
