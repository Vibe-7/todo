package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	HttpHandlers *HttpHandlers
}

func NewHTTPServer(HttpHandlers *HttpHandlers) *HttpServer {
	return &HttpServer{
		HttpHandlers: HttpHandlers,
	}
}

func (s *HttpServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.HttpHandlers.HandlersCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.HttpHandlers.HandlersGetTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.HttpHandlers.HandlersGetAllTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.HttpHandlers.HandlersGetAllUncompletedTask)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.HttpHandlers.HandlerCompletedTask)
	router.Path("/tasks").Methods("DELETE").HandlerFunc(s.HttpHandlers.HandlerDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil

}
