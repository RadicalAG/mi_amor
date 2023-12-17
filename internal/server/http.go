package server

import (
	"radical/red_letter/internal/handler"
	"radical/red_letter/internal/middleware"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine      *gin.Engine
	handlers    []handler.Handler
	middlewares []middleware.Middleware
}

func NewHttpServer() *HttpServer {
	srv := gin.Default()
	return &HttpServer{
		engine: srv,
	}
}

func (h *HttpServer) AddHandler(handlers ...handler.Handler) {
	for _, handler := range handlers {
		h.handlers = append(h.handlers, handler)
	}
}
func (h *HttpServer) AddMiddleware(middlewares ...middleware.Middleware) {
	for _, middleware := range middlewares {
		h.middlewares = append(h.middlewares, middleware)
	}
}

func (h *HttpServer) Serve() {
	for _, m := range h.middlewares {
		m.RegisterMiddleware(h.engine)
	}
	for _, handler := range h.handlers {
		handler.RegisterHandler(h.engine)
	}
	h.engine.Run()
}
