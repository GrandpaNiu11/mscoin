package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type Routers struct {
	server *rest.Server
	//中间件
	middleware []rest.Middleware
}

func NewRouters(server *rest.Server) *Routers {
	return &Routers{
		server: server,
	}
}

func (r *Routers) Get(path string, handler http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(r.middleware,
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: handler,
			},
		),
	)
}
func (r *Routers) Post(path string, handler http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(r.middleware,
			rest.Route{
				Method:  http.MethodPost,
				Path:    path,
				Handler: handler,
			},
		),
	)
}
func (r *Routers) Group() *Routers {
	return &Routers{
		server: r.server,
	}

}
