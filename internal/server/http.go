package server

import (
	"radical/red_letter/internal/handler"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine   *gin.Engine
	handlers []handler.Handler
}

func NewHttpServer(engine *gin.Engine) *HttpServer {
	return &HttpServer{
		engine: engine,
	}
}

func (h *HttpServer) AddHandler(handler handler.Handler) {
	h.handlers = append(h.handlers, handler)
}

func (h *HttpServer) Serve() {
	for _, handler := range h.handlers {
		handler.RegisterHandler(h.engine)
	}
	h.engine.Run()
}
