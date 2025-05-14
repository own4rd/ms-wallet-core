package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router         chi.Router
	Handlers       map[string]http.HandlerFunc
	WebServeerPort string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Router:         chi.NewRouter(),
		Handlers:       make(map[string]http.HandlerFunc),
		WebServeerPort: port,
	}
}

func (ws *WebServer) AddHandlers(path string, handler http.HandlerFunc) {
	ws.Handlers[path] = handler
}

func (ws *WebServer) Start() {
	ws.Router.Use(middleware.Logger)
	for path, handler := range ws.Handlers {
		ws.Router.Post(path, handler)
	}
	http.ListenAndServe(ws.WebServeerPort, ws.Router)
}
